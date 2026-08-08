package gogige

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/aaronmurniadi/gogige/genapi"
	"github.com/aaronmurniadi/gogige/gvcp"
)

type device struct {
	ip        string
	cam       *Camera
	log       Logger
	component Component
	mu        sync.Mutex
	closed    bool
}

// Logger returns the device logger (used by live/grab via type assertion).
func (d *device) Logger() Logger {
	if d == nil || d.log == nil {
		return NopLogger{}
	}
	return d.log
}

// Open connects to a GigE Vision camera and returns a Device.
func Open(ctx context.Context, ip string, opts ...Option) (Device, error) {
	cfg, err := resolveOpenConfig(ctx, ip, opts)
	if err != nil {
		return nil, err
	}
	cam, err := connectCamera(ip, cfg.timeout, cfg.logger)
	if err != nil {
		return nil, err
	}
	kind := cfg.component
	if kind == ComponentUnknown {
		kind = ComponentColor
	}
	return &device{ip: ip, cam: cam, log: cfg.logger, component: kind}, nil
}

// OpenDevice connects to a GigE Vision camera and returns the Camera directly
// (Phase 4 surface). Use Camera.SetInteger / SetEnum for control and
// Camera.StartStream for live frames.
func OpenDevice(ctx context.Context, ip string, opts ...Option) (*Camera, error) {
	cfg, err := resolveOpenConfig(ctx, ip, opts)
	if err != nil {
		return nil, err
	}
	cam, err := connectCamera(ip, cfg.timeout, cfg.logger)
	if err != nil {
		return nil, err
	}
	return cam, nil
}

// resolveOpenConfig validates arguments and applies options shared by Open/OpenDevice.
func resolveOpenConfig(ctx context.Context, ip string, opts []Option) (openConfig, error) {
	if ip == "" {
		return openConfig{}, errors.New("gige: ip address must not be empty")
	}
	cfg := openConfig{logger: NopLogger{}, timeout: 2 * time.Second, component: ComponentColor}
	for _, o := range opts {
		o(&cfg)
	}
	if err := ctx.Err(); err != nil {
		return openConfig{}, err
	}
	return cfg, nil
}

func (d *device) IP() string { return d.ip }

func (d *device) Features() Features { return cameraFeatures{c: d.cam} }

func (d *device) StartGrabber(ctx context.Context, opts ...GrabOption) (Grabber, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.closed {
		return nil, errors.New("gige: device closed")
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	s := NewFromCamera(d.cam)
	s.component = d.component
	for _, o := range opts {
		o(s)
	}
	if err := s.Open(d.ip); err != nil {
		return nil, err
	}
	return s, nil
}

func (d *device) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.closed {
		return nil
	}
	d.closed = true
	if d.cam != nil {
		d.cam.Close()
		d.cam = nil
	}
	return nil
}

type cameraFeatures struct{ c *Camera }

func (f cameraFeatures) SetBool(name string, v bool) error {
	return f.c.SetBooleanFeature(name, v)
}
func (f cameraFeatures) SetInt(name string, v int64) error {
	return f.c.SetIntFeature(name, v)
}
func (f cameraFeatures) SetFloat(name string, v float64) error {
	return f.c.SetFloatFeature(name, v)
}
func (f cameraFeatures) SetString(name, v string) error {
	return f.c.SetStringFeature(name, v)
}
func (f cameraFeatures) Execute(name string) error { return f.c.ExecuteCommand(name) }
func (f cameraFeatures) Has(name string) bool {
	return f.c != nil && f.c.Has(name)
}

func connectCamera(ip string, timeout time.Duration, log Logger) (*Camera, error) {
	if timeout <= 0 {
		timeout = 2 * time.Second
	}
	g, err := gvcp.DialGVCP(ip, timeout)
	if err != nil {
		return nil, err
	}
	if err := g.TakeControl(); err != nil {
		_ = g.Close()
		return nil, fmt.Errorf("gige: take control: %w", err)
	}
	xmlData, err := genapi.FetchXML(g)
	if err != nil {
		_ = g.LeaveControl()
		_ = g.Close()
		return nil, fmt.Errorf("gige: fetch XML: %w", err)
	}
	nm, err := genapi.ParseNodeMap(xmlData, g)
	if err != nil {
		_ = g.LeaveControl()
		_ = g.Close()
		return nil, err
	}
	if err := g.TakeControl(); err != nil {
		_ = g.Close()
		return nil, fmt.Errorf("gige: re-take control: %w", err)
	}
	return New(ip, g, nm, log), nil
}

var (
	_ Device                       = (*device)(nil)
	_ Features                     = cameraFeatures{}
	_ FrameSink                    = JPEGFunc(nil)
	_ Grabber                      = (*Session)(nil)
	_ interface{ Logger() Logger } = (*device)(nil)
)
