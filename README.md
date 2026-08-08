# gogige

Pure-Go GigE Vision client (`package gige`) for Huaray/Dahua 3D volume cameras.

![Go Version](https://img.shields.io/badge/Go-1.26-blue)
![License](https://img.shields.io/badge/License-BSD--3--Clause-blue)
![Release](https://img.shields.io/badge/Release-v0.8.0-blue)
![Pure Go](https://img.shields.io/badge/Pure%20Go-Yes-green)

> Tested on:

| Manufacturer      | Model          | Description                                                                                        |
| ----------------- | -------------- | -------------------------------------------------------------------------------------------------- |
| Huaray Technology (iRAYPLE) | DS5131MG30CE   | 3D stereo industrial smart camera for machine vision, industrial automation, and precise depth measurement. |

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Examples](#examples)
- [Pre-commit](#pre-commit)
- [Contributing](#contributing)
- [License](#license)

## Features

- **GigE Vision control** — GVCP register read/write, heartbeat maintenance, and access-privilege management over a dedicated control channel.
- **Streaming** — GVSP receiver with pre-allocated buffer pool, packet resend, and multi-part/GenDC payload parsing.
- **GenICam GenApi XML** — Fetch, decompress, and build the camera node map (IntReg, Enumeration, SwissKnife, Converter, Port, …).
- **3D volume (BSCF)** — Per-frame color/depth/mono components with JPEG encoding and mm-scale measurements (`WidthMm`, `HeightMm`, `Length`, `PackCount`).
- **Live preview** — App-owned sinks (`gige.NewLive`), switch components mid-stream.
- **Zero-alloc hot path** — `gvsp.Receiver` allocates nothing per packet during streaming.

## Getting Started

### Prerequisites

- Go 1.26+
- A GigE Vision camera on the same L2 network (a simulator works for control-plane work)
- Optional: Jumbo-frames-capable NIC for high-throughput streaming

### Installation

```bash
go get github.com/aaronmurniadi/gogige
```

### Quick start

Open a device, configure nodes, and stream:

```go
cam, err := gige.Open(ctx, "192.168.1.10",
    gige.WithLogger(logger),
    gige.WithTimeout(2*time.Second),
)
defer cam.Close()

g, err := cam.StartGrabber(ctx)
defer g.Close()
sample, err := g.Grab(ctx) // Sample: JPEG + WidthMm/HeightMm/Length/PackCount
```

## Usage

**One-shot grab:**

```go
jpeg, err := gige.GrabJPEG(ctx, "192.168.1.10")
```

**Discover cameras:**

```go
devs, err := gige.Discover(ctx, 2*time.Second)
```

**Live preview (app owns the sink):**

```go
live := gige.NewLive(dev, gige.WithSink(gige.JPEGFunc(hub.Broadcast)), gige.WithLiveComponent(gige.ComponentDepth))
live.Start(ctx)
defer live.Stop()
sample := live.LatestSample() // filter/validate in the app
// live.SetComponent(gige.ComponentColor) // switch mid-stream
```

**Logging:** `gogige.WithLogger(...)` accepts any `Logger` implementation (e.g. `gogige.Slog(slog.Default())` or your own wrapper around zerolog/zap). Default is a no-op.

## Project Structure

```
gogige/
├── cmd/
│   ├── gogige-discover/      # CLI discovery utility
│   └── gogige-stream/        # CLI stream capture utility
├── control/
│   ├── camera/               # Camera, Connect, Logger
│   ├── gvcp/                 # GVCP, CCP, FetchXML, Start/StopAcquisition
│   └── genicam/              # GenICam NodeMap
├── genapi/                   # GenICam GenApi XML parser / node map
├── gvcp/                     # GigE Vision Control Protocol
├── gvsp/                     # GigE Vision Streaming Protocol
├── vision/
│   ├── device/               # Device, Open, Features, Grabber
│   ├── session/              # Streaming session
│   ├── live/                 # Live preview loop
│   ├── grab/                 # One-shot GrabJPEG
│   ├── gvsp/                 # GVSP stream
│   ├── bscf/                 # BSCF + Sample
│   └── color/                # JPEG encode
├── examples/                 # Runnable examples (smoke, grab, live, …)
├── .githooks/                # Versioned git hooks (gofmt + go test)
├── AGENTS.md                 # Project & protocol rules
├── CHANGELOG.md
├── ROADMAP.md
├── go.mod
└── LICENSE
```

## Examples

Omit `-ip` to pick the first camera from GigE discovery (or pass `-ip` explicitly).

- Full stack smoke (CCP/heartbeat/XML/stream/pause/live): [`examples/smoke`](examples/smoke)
- GenICam `Has` / `ApplyControlPair` / `Execute`: [`examples/features`](examples/features)
- One-shot color JPEG to disk: [`examples/grab`](examples/grab)
- Discover BSCF components (color/depth/mono) and grab all: [`examples/grab-components`](examples/grab-components)
- Probe available stream types (pixel formats, payload/chunk/component features, per-frame JPEG encodability): [`examples/probe-streams`](examples/probe-streams)
- Stream measurements only: [`examples/stream-measurements`](examples/stream-measurements)
- Dump GenICam XML: [`examples/dump-xml`](examples/dump-xml)
- Configure volume TCP preset: [`examples/configure-camera`](examples/configure-camera)
- Browser live stream: [`examples/websocket-stream`](examples/websocket-stream)
- CLI discover: [`cmd/gogige-discover`](cmd/gogige-discover)
- CLI N-frame capture: [`cmd/gogige-stream`](cmd/gogige-stream)

## Pre-commit

Once per clone, point git at the versioned hooks (runs `gofmt` on staged `.go` files, then `go test ./...`):

```bash
git config core.hooksPath .githooks
```

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request following these simple steps:

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the BSD 3-Clause License. See [LICENSE](LICENSE) for more information.
