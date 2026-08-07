package gige

import "time"

// Option configures Open / device connection.
type Option func(*openConfig)

type openConfig struct {
	logger    Logger
	timeout   time.Duration
	imageKind ImageKind
}

// WithLogger sets the logger used by this device (default: NopLogger).
func WithLogger(l Logger) Option {
	return func(c *openConfig) {
		if l != nil {
			c.logger = l
		}
	}
}

// WithTimeout sets the GVCP dial / control timeout (default: 2s).
func WithTimeout(d time.Duration) Option {
	return func(c *openConfig) {
		if d > 0 {
			c.timeout = d
		}
	}
}

// WithImageKind selects which BSCF image block GrabJPEG uses (color/depth/mono).
// Default: ImageColor.
func WithImageKind(k ImageKind) Option {
	return func(c *openConfig) {
		if k != ImageUnknown {
			c.imageKind = k
		}
	}
}

// GrabOption configures StartGrabber / Session.
type GrabOption func(*Session)

// GrabImageKind selects which BSCF image block Session.Grab decodes.
func GrabImageKind(k ImageKind) GrabOption {
	return func(s *Session) {
		if s != nil && k != ImageUnknown {
			s.imageKind = k
		}
	}
}
