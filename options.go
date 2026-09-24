package gogige

import (
	"time"

	"github.com/aaronmurniadi/gogige/calib"
)

// Option configures Open / device connection.
type Option func(*openConfig)

type openConfig struct {
	logger    Logger
	timeout   time.Duration
	component Component
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

// WithComponent selects which BSCF/SFNC component GrabJPEG uses (color/depth/mono).
// Default: ComponentColor.
func WithComponent(comp Component) Option {
	return func(c *openConfig) {
		if comp != ComponentUnknown {
			c.component = comp
		}
	}
}

// GrabOption configures StartGrabber / Session.
type GrabOption func(*Session)

// GrabComponent selects which BSCF/SFNC component Session.Grab decodes.
func GrabComponent(comp Component) GrabOption {
	return func(s *Session) {
		if s != nil && comp != ComponentUnknown {
			s.component = comp
		}
	}
}

// WithOverlay enables pack-detection bounding-box projection. When enabled with
// a valid color calibration, every grabbed Sample gets its Overlay field
// populated (one pixel box per pack, matching Packs order); the consumer draws
// the green box outlines themselves. Pass false to disable (default).
func WithOverlay(enabled bool, cam calib.CamCalib) GrabOption {
	return func(s *Session) {
		if s == nil {
			return
		}
		s.overlay = camCalib{valid: enabled, color: cam}
	}
}

// Version is the library version.
const Version = "1.7.0"
