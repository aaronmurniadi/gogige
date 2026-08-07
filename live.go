package gige

import (
	"context"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	liveGrabTimeout      = time.Second
	liveGrabErrBackoff   = 100 * time.Millisecond
	liveOpenBackoff      = time.Second
	liveGrabFailReopenAt = 8
)

// LiveOption configures NewLive.
type LiveOption func(*Live)

// WithSink attaches a FrameSink for preview JPEG delivery.
func WithSink(s FrameSink) LiveOption {
	return func(l *Live) { l.sink = s }
}

// WithOnSample registers a callback after each published sample.
func WithOnSample(fn func(Sample)) LiveOption {
	return func(l *Live) { l.onSample = fn }
}

// WithLiveComponent selects which BSCF/SFNC component Live grabs (color/depth/mono).
func WithLiveComponent(c Component) LiveOption {
	return func(l *Live) {
		if c != ComponentUnknown {
			l.component = c
		}
	}
}

// Live keeps a Grabber open and publishes Samples (optionally to a FrameSink).
type Live struct {
	dev       Device
	sink      FrameSink
	onSample  func(Sample)
	log       Logger
	component Component

	lifeMu  sync.Mutex
	running bool
	stop    chan struct{}
	done    chan struct{}
	once    sync.Once
	paused  atomic.Bool
	latest  atomic.Pointer[[]byte]
	sample  atomic.Pointer[Sample]
	grabber Grabber
}

// NewLive builds a Live preview/capture loop over an already-opened Device.
func NewLive(dev Device, opts ...LiveOption) *Live {
	l := &Live{dev: dev, log: NopLogger{}, component: ComponentColor}
	if lg, ok := dev.(hasLogger); ok {
		if log := lg.Logger(); log != nil {
			l.log = log
		}
	}
	for _, o := range opts {
		o(l)
	}
	return l
}

// Component returns the BSCF/SFNC component Live is decoding.
func (l *Live) Component() Component {
	if l == nil {
		return ComponentUnknown
	}
	l.lifeMu.Lock()
	defer l.lifeMu.Unlock()
	return l.component
}

// SetComponent switches the BSCF/SFNC component for subsequent grabs (no reconnect).
func (l *Live) SetComponent(c Component) {
	if l == nil || c == ComponentUnknown {
		return
	}
	l.lifeMu.Lock()
	defer l.lifeMu.Unlock()
	l.component = c
	if l.grabber != nil {
		l.grabber.SetComponent(c)
	}
}

// Start begins the grab loop. Safe to call again after Stop (reopens grabber).
func (l *Live) Start(ctx context.Context) {
	if l == nil {
		return
	}
	l.lifeMu.Lock()
	defer l.lifeMu.Unlock()
	l.startLocked(ctx)
}

// Restart stops and starts again.
func (l *Live) Restart(ctx context.Context) {
	if l == nil {
		return
	}
	l.lifeMu.Lock()
	defer l.lifeMu.Unlock()
	l.stopLocked()
	l.startLocked(ctx)
}

// Stop ends the grab loop and closes the Grabber owned by Live.
func (l *Live) Stop() {
	if l == nil {
		return
	}
	l.lifeMu.Lock()
	defer l.lifeMu.Unlock()
	l.stopLocked()
}

func (l *Live) startLocked(ctx context.Context) {
	if l.running {
		l.stopLocked()
	}
	if ctx == nil {
		ctx = context.Background()
	}
	stop := make(chan struct{})
	done := make(chan struct{})
	l.stop = stop
	l.done = done
	l.once = sync.Once{}
	l.running = true
	go l.loop(ctx, stop, done)
}

func (l *Live) stopLocked() {
	if !l.running {
		return
	}
	l.once.Do(func() {
		close(l.stop)
	})
	select {
	case <-l.done:
	case <-time.After(2 * time.Second):
		l.log.Warn("live stop timed out")
	}
	if l.grabber != nil {
		_ = l.grabber.Close()
		l.grabber = nil
	}
	l.running = false
	l.stop = nil
	l.done = nil
	l.once = sync.Once{}
}

// Pause freezes publishing and pauses the underlying Grabber when possible.
func (l *Live) Pause() {
	if l == nil {
		return
	}
	l.paused.Store(true)
	if l.sink != nil {
		l.sink.Freeze()
	}
}

// Resume unfreezes publishing.
func (l *Live) Resume() {
	if l == nil {
		return
	}
	l.paused.Store(false)
	if l.sink != nil {
		l.sink.Resume()
	}
}

// LatestJPEG returns a copy of the most recent JPEG, or nil.
func (l *Live) LatestJPEG() []byte {
	if l == nil {
		return nil
	}
	ptr := l.latest.Load()
	if ptr == nil {
		return nil
	}
	return append([]byte(nil), (*ptr)...)
}

// LatestSample returns a copy of the most recent Sample.
func (l *Live) LatestSample() Sample {
	if l == nil {
		return Sample{PackCount: -1}
	}
	ptr := l.sample.Load()
	if ptr == nil {
		return Sample{PackCount: -1}
	}
	s := *ptr
	if len(s.JPEG) > 0 {
		s.JPEG = append([]byte(nil), s.JPEG...)
	}
	return s
}

func (l *Live) publish(sample Sample) {
	if len(sample.JPEG) == 0 {
		return
	}
	cp := append([]byte(nil), sample.JPEG...)
	sample.JPEG = cp
	l.latest.Store(&cp)
	stored := sample
	l.sample.Store(&stored)
	if l.sink != nil {
		l.sink.SendJPEG(cp)
	}
	if l.onSample != nil {
		l.onSample(sample)
	}
}

func (l *Live) loop(parent context.Context, stop, done chan struct{}) {
	defer close(done)
	defer func() {
		l.lifeMu.Lock()
		if l.grabber != nil {
			_ = l.grabber.Close()
			l.grabber = nil
		}
		l.lifeMu.Unlock()
	}()

	fails := 0
	acqOn := false
	for {
		select {
		case <-stop:
			return
		default:
		}

		if l.paused.Load() {
			if acqOn {
				l.lifeMu.Lock()
				g := l.grabber
				l.lifeMu.Unlock()
				if g != nil {
					_ = g.Pause(context.Background())
				}
				acqOn = false
			}
			select {
			case <-stop:
				return
			case <-time.After(50 * time.Millisecond):
			}
			continue
		}

		l.lifeMu.Lock()
		g := l.grabber
		if g == nil {
			ng, err := l.dev.StartGrabber(parent, GrabComponent(l.component))
			if err != nil {
				l.lifeMu.Unlock()
				acqOn = false
				l.log.Warn("live stream open failed", "err", err)
				select {
				case <-stop:
					return
				case <-time.After(liveOpenBackoff):
				}
				continue
			}
			l.grabber = ng
			g = ng
			fails = 0
			acqOn = true
			l.log.Info("live stream opened", "camera_ip", l.dev.IP())
		} else if !acqOn {
			if err := g.Resume(context.Background()); err != nil {
				_ = g.Close()
				l.grabber = nil
				l.lifeMu.Unlock()
				acqOn = false
				l.log.Warn("live stream resume failed", "err", err)
				select {
				case <-stop:
					return
				case <-time.After(liveOpenBackoff):
				}
				continue
			}
			acqOn = true
			fails = 0
			l.log.Info("live stream resumed", "camera_ip", l.dev.IP())
		}
		l.lifeMu.Unlock()

		grabCtx, cancel := context.WithTimeout(parent, liveGrabTimeout)
		sample, err := g.Grab(grabCtx)
		cancel()
		if err != nil {
			fails++
			l.log.Debug("live grab failed", "err", err, "consecutive_fails", fails)
			if fails >= liveGrabFailReopenAt || streamDeadError(err) {
				l.lifeMu.Lock()
				if l.grabber != nil {
					_ = l.grabber.Close()
					l.grabber = nil
				}
				l.lifeMu.Unlock()
				fails = 0
				acqOn = false
				select {
				case <-stop:
					return
				case <-time.After(liveOpenBackoff):
				}
				continue
			}
			select {
			case <-stop:
				return
			case <-time.After(liveGrabErrBackoff):
			}
			continue
		}
		fails = 0
		l.publish(sample)
	}
}

func streamDeadError(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "stream not open") ||
		strings.Contains(msg, "connect") ||
		strings.Contains(msg, "AcquisitionStart") ||
		strings.Contains(msg, "take control")
}
