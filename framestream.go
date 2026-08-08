package gogige

import (
	"context"
	"sync"
)

// Stream is a live GVSP acquisition channel (Phase 4). Create one via
// Camera.StartStream, consume pooled frames from Frames(), and Release() each
// frame back to the pre-allocated buffer pool.
type Stream struct {
	s      *Session
	cancel context.CancelFunc
	done   chan struct{}
	frames chan *Frame
	once   sync.Once
}

// StartStream opens a GVSP channel on c and begins acquisition. Frames are
// delivered over Stream.Frames() until the context is cancelled or Stop is
// called. The Camera stays open; close it separately when done.
func (c *Camera) StartStream(ctx context.Context) (*Stream, error) {
	if c == nil || c.GVCP() == nil {
		return nil, errStreamNotOpen
	}
	s := NewFromCamera(c)
	if err := s.Open(c.IP); err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	child, cancel := context.WithCancel(ctx)
	st := &Stream{
		s:      s,
		cancel: cancel,
		done:   make(chan struct{}),
		frames: make(chan *Frame, 4),
	}
	go st.loop(child)
	return st, nil
}

// Frames returns the channel of pooled frames. The channel closes when the
// stream ends (context cancelled, Stop called, or the session goes away).
func (st *Stream) Frames() <-chan *Frame {
	if st == nil {
		return nil
	}
	return st.frames
}

// Stop ends acquisition, closes the stream socket, and stops the background
// goroutine. Safe to call more than once.
func (st *Stream) Stop() {
	if st == nil {
		return
	}
	st.once.Do(func() {
		st.cancel()
		<-st.done
		_ = st.s.Close()
	})
}

// Pause stops image transfer but keeps the control channel and socket open.
func (st *Stream) Pause() error {
	if st == nil {
		return nil
	}
	return st.s.PauseStreaming()
}

// Resume re-programs the stream channel and restarts acquisition.
func (st *Stream) Resume() error {
	if st == nil {
		return errStreamNotOpen
	}
	return st.s.ResumeStreaming()
}

func (st *Stream) loop(ctx context.Context) {
	defer close(st.done)
	defer close(st.frames)
	for {
		frame, err := st.s.recvFramePtr(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			if err == errStreamNotOpen {
				return
			}
			continue
		}
		select {
		case st.frames <- frame:
		case <-ctx.Done():
			frame.Release()
			return
		}
	}
}
