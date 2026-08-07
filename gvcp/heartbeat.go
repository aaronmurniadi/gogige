package gvcp

import (
	"sync"
	"time"
)

const (
	defaultHeartbeatTimeout = 3000 * time.Millisecond
	minHeartbeatInterval    = 100 * time.Millisecond
)

// PulseHeartbeat refreshes the GigE Vision control-channel heartbeat.
// Without periodic GVCP traffic the camera revokes CCP (~3× HeartbeatTimeout)
// and the GVSP stream stops.
func (g *GVCP) PulseHeartbeat() error {
	_, err := g.ReadReg(gvbsCCP)
	return err
}

// HeartbeatTimeout reads the device HeartbeatTimeout register (milliseconds).
// Returns the default (3000ms) when the register is unreadable or zero.
func (g *GVCP) HeartbeatTimeout() time.Duration {
	v, err := g.ReadReg(gvbsHeartbeatTO)
	if err != nil || v == 0 {
		return defaultHeartbeatTimeout
	}
	return time.Duration(v) * time.Millisecond
}

// Heartbeat is a background CCP refresh loop ticking at HeartbeatTimeout/2.
type Heartbeat struct {
	g    *GVCP
	stop chan struct{}
	done chan struct{}

	mu      sync.Mutex
	stopped bool
}

// StartHeartbeat launches a goroutine that pulses CCP every HeartbeatTimeout/2.
// Call Stop when control is no longer needed. Safe to call multiple times; each
// StartHeartbeat returns an independent handle — prefer one per GVCP session.
func (g *GVCP) StartHeartbeat() *Heartbeat {
	if g == nil {
		return nil
	}
	to := g.HeartbeatTimeout()
	interval := to / 2
	if interval < minHeartbeatInterval {
		interval = minHeartbeatInterval
	}
	h := &Heartbeat{
		g:    g,
		stop: make(chan struct{}),
		done: make(chan struct{}),
	}
	go h.loop(interval)
	return h
}

func (h *Heartbeat) loop(interval time.Duration) {
	defer close(h.done)
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-h.stop:
			return
		case <-t.C:
			_ = h.g.PulseHeartbeat()
		}
	}
}

// Stop ends the heartbeat goroutine and waits for it to exit.
func (h *Heartbeat) Stop() {
	if h == nil {
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.stopped {
		return
	}
	h.stopped = true
	close(h.stop)
	<-h.done
}
