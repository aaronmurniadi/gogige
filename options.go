package gige

import "time"

// Option configures Open / device connection.
type Option func(*openConfig)

type openConfig struct {
	logger  Logger
	timeout time.Duration
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
