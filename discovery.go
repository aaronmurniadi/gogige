package gige

import (
	"context"
	"time"

	"github.com/aaronmurniadi/gogige/gvcp"
)

// DeviceInfo is a camera found via GigE Vision discovery (root API).
type DeviceInfo struct {
	IP           string
	MAC          string
	Manufacturer string
	Model        string
	Serial       string
	UserName     string
}

// Discover broadcasts a GigE Vision DISCOVERY_CMD and collects acknowledgements.
// timeout bounds how long to wait for replies after the broadcast.
func Discover(ctx context.Context, timeout time.Duration) ([]DeviceInfo, error) {
	found, err := gvcp.Discover(ctx, timeout)
	if err != nil {
		return nil, err
	}
	out := make([]DeviceInfo, 0, len(found))
	for _, d := range found {
		out = append(out, DeviceInfo{
			IP:           d.IP,
			MAC:          d.MAC,
			Manufacturer: d.Manufacturer,
			Model:        d.Model,
			Serial:       d.Serial,
			UserName:     d.UserName,
		})
	}
	return out, nil
}
