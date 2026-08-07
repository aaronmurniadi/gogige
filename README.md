# gogige

Pure-Go GigE Vision client (`package gige`) for Huaray/Dahua 3D volume cameras.

## Install

```bash
go get github.com/aaronmurniadi/gogige
```

## Pre-commit

Once per clone, point git at the versioned hooks (runs `gofmt` on staged `.go` files, then `go test ./...`):

```bash
git config core.hooksPath .githooks
```

## Quick start

```go
dev, err := gige.Open(ctx, "192.168.1.10",
    gige.WithLogger(logger),
    gige.WithTimeout(2*time.Second),
)
defer dev.Close()

g, err := dev.StartGrabber(ctx)
defer g.Close()
sample, err := g.Grab(ctx) // Sample: JPEG + WidthMm/HeightMm/Length/PackCount
```

One-shot:

```go
jpeg, err := gige.GrabJPEG(ctx, "192.168.1.10")
```

Discover cameras (root API):

```go
devs, err := gige.Discover(ctx, 2*time.Second)
```

Live + preview (app owns the sink):

```go
live := gige.NewLive(dev, gige.WithSink(gige.JPEGFunc(hub.Broadcast)))
live.Start(ctx)
defer live.Stop()
sample := live.LatestSample() // filter/validate in the app
```

## Package layout

| Domain | Package | Role |
|--------|---------|------|
| Root | `gige` | Discover + re-exports (`Open`, `NewLive`, `GrabJPEG`, …) |
| Control | [`control/camera`](control/camera) | Camera, Connect, Logger |
| Control | [`control/gvcp`](control/gvcp) | GVCP, CCP, FetchXML, Start/StopAcquisition |
| Control | [`control/genicam`](control/genicam) | GenICam NodeMap |
| Vision | [`vision/device`](vision/device) | Device, Open, Features, Grabber |
| Vision | [`vision/session`](vision/session) | Streaming session |
| Vision | [`vision/live`](vision/live) | Live preview loop |
| Vision | [`vision/grab`](vision/grab) | One-shot GrabJPEG |
| Vision | [`vision/gvsp`](vision/gvsp) | GVSP stream |
| Vision | [`vision/bscf`](vision/bscf) | BSCF + Sample |
| Vision | [`vision/color`](vision/color) | JPEG encode |

## Examples

Omit `-ip` to pick the first camera from GigE discovery (or pass `-ip` explicitly).

- One-shot color JPEG to disk: [`examples/grab`](examples/grab)
- Stream measurements only: [`examples/stream-measurements`](examples/stream-measurements)
- Dump GenICam XML: [`examples/dump-xml`](examples/dump-xml)
- Configure volume TCP preset: [`examples/configure-volume-tcp`](examples/configure-volume-tcp)
- Browser live stream: [`examples/websocket-stream`](examples/websocket-stream)
- CLI discover: [`cmd/gogige-discover`](cmd/gogige-discover)

Logging: `gige.WithLogger(...)` or `gige.Zerolog(z)`. Default is a no-op.
