package gogige

import "context"

// Device is a connected camera handle (control + ability to start streaming).
type Device interface {
	IP() string
	Features() Features
	StartGrabber(ctx context.Context, opts ...GrabOption) (Grabber, error)
	Close() error
}

// Features is GenICam feature access without exposing NodeMap.
type Features interface {
	SetBool(name string, v bool) error
	SetInt(name string, v int64) error
	SetFloat(name string, v float64) error
	SetString(name, v string) error
	Execute(name string) error
	Has(name string) bool
}

// Grabber is a live GVSP acquisition stream.
type Grabber interface {
	Grab(ctx context.Context) (Sample, error)
	GrabAll(ctx context.Context) ([]Sample, error)
	SetComponent(Component)
	Pause(ctx context.Context) error
	Resume(ctx context.Context) error
	Close() error
}

// FrameSink receives JPEG frames from Live for preview (WebSocket, MJPEG, etc.).
// Implement this in application code — the library does not ship a WebSocket server.
type FrameSink interface {
	SendJPEG(jpeg []byte)
	Freeze()
	Resume()
}

// JPEGFunc adapts a callback into a FrameSink (Freeze/Resume are no-ops).
// Handy for wiring Live to an existing broadcast function:
//
//	live := live.NewLive(dev, live.WithSink(JPEGFunc(hub.Broadcast)))
type JPEGFunc func([]byte)

func (f JPEGFunc) SendJPEG(jpeg []byte) {
	if f != nil {
		f(jpeg)
	}
}
func (JPEGFunc) Freeze() {}
func (JPEGFunc) Resume() {}

type hasLogger interface {
	Logger() Logger
}
