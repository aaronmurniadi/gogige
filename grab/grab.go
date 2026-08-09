package grab

import (
	"context"
	"fmt"
	"time"

	"github.com/aaronmurniadi/gogige"
)

const defaultGrabTimeout = 3 * time.Second

// GrabJPEG opens a short-lived device, grabs one frame, and closes.
// Convenience for one-shot capture; prefer Device + Grabber for continuous use.
func GrabJPEG(ctx context.Context, ip string, opts ...gogige.Option) ([]byte, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, defaultGrabTimeout)
		defer cancel()
	}

	start := time.Now()
	dev, err := gogige.Open(ctx, ip, opts...)
	if err != nil {
		return nil, err
	}
	defer dev.Close()

	g, err := dev.StartGrabber(ctx)
	if err != nil {
		return nil, err
	}
	defer g.Close()

	sample, err := g.Grab(ctx)
	jpeg := sample.JPEG
	if err == nil && len(jpeg) == 0 {
		err = fmt.Errorf("gige: GrabJPEG: empty frame")
	}
	if lg, ok := dev.(interface{ Logger() gogige.Logger }); ok {
		lg.Logger().Info("GrabJPEG finished",
			"camera_ip", ip,
			"elapsed", time.Since(start),
			"jpeg_bytes", len(jpeg),
			"err", err,
		)
	}
	return jpeg, err
}
