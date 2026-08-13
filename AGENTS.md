# gogige Project Rules

Pure-Go GigE Vision client (`package gogige`, module `github.com/aaronmurniadi/gogige`) targeting Huaray/Dahua 3D volume (BSCF) cameras. Primary tested device: Huaray DS5131MG30CE.

## Token saving

Be brief.

## Project intent

- **Pure Go, zero CGO.** Enforced via `mise.toml` (`CGO_ENABLED=0`) and the pre-commit hook. Do not add cgo dependencies on the hot path; `gentl/` CGO bindings are optional and off by default.
- **Zero-alloc hot path.** `gvsp.Receiver` must allocate nothing per packet during streaming. Prefer pre-allocated ring buffers over maps on the streaming path; don't regress that invariant.
- **Huaray BSCF.** Per-frame color/depth/mono components, mm-scale measurements (`WidthMm`, `HeightMm`, `LengthMm`, `PackCount`), JPEG encode. `examples/` are the executable spec / smoke surface.

## Architectural hierarchy

```
               +------------------------------------------------+
               |                  gogige API                    |
               | (Camera, Stream, Discovery — root package)     |
               +-----------------------+------------------------+
                                       |
          +----------------------------+----------------------------+
          |                                                         |
+---------v----------+                                   +----------v---------+
|     GenApi XML     |                                   |    GVSP Stream     |
|   (genapi) Parser  |                                   | Engine (UDP Data)  |
+---------+----------+                                   +----------+---------+
          |                                                         |
+---------v----------+                                   +----------v---------+
|     GVCP Client    |                                   |  Frame Buffer /    |
|  (gvcp, Control &  |                                   |  GenDC  / BSCF     |
|   Reg Map)         |                                   |  Reconstruct       |
+---------+----------+                                   +--------------------+
          |
+---------v----------+
| Physical Network / |
|   UDP Sockets      |
+--------------------+
```

## Package layout

Enforced for protocol packages (`gvcp/`, `gvsp/`, `genapi/`) and root high-level files. Prefer target names for all new code; **do not reintroduce `control/` or `vision/gvsp`** — those were removed; root-level `camera/device/stream/options/interfaces/alias` are the Phase 4 surface.

```
gogige/
├── cmd/gogige-discover/      # CLI: Discovery
├── cmd/gogige-stream/        # CLI: N-frame JPEG + BSCF measurements
├── gvcp/                     # GigE Vision Control Protocol (GenCP 1.3.1)
│   ├── client.go             # control channel, register read/write
│   ├── discovery.go          # UDP broadcast DISCOVERY_CMD + ACK ABRM parse
│   ├── heartbeat.go          # background heartbeat (HeartbeatTimeout/2)
│   ├── packet.go, pending.go # GVCP packet ser/deser, PENDING_ACK timeout
│   ├── register_map.go       # GenCP ABRM 0x0000-0x0250 + GigE ABRM/SBRM
│   ├── endianness.go         # ImplementationEndianness (0x020C) sync on TakeControl
│   └── acquisition.go        # ControlChannelPrivilege + Start/StopAcquisition RMW 0x0D04
├── gvsp/                     # GigE Vision Streaming Protocol
│   ├── receiver.go           # hot-path UDP receiver (zero-alloc)
│   ├── frame.go              # Frame + reassembly; OOOPacketRing (zero-alloc OOO store)
│   ├── buffer_pool.go        # pre-allocated frame buffers + Frame.Release
│   ├── payload.go            # BSCF + Component/Sample; IsBSCF/ParseBSCF
│   ├── resend.go             # gap detect + RESEND_CMD
│   ├── payloadtype.go        # PAYLOAD_TYPE_* constants + Frame.PayloadType
│   ├── multi_part.go, chunk_data.go, genDC_payload.go  # GenDC/MultiPart/Chunk
│   └── socket.go             # listen/transport + MTU/SCPS + SO_RCVBUF
├── genapi/                   # GenICam GenApi XML (2.1.1)
│   ├── camera_description.go # FirstURL/ManifestTable fetch + zip/deflate
│   ├── node.go, types.go     # Node interface + node parsing
│   ├── nodemap.go            # parse + feature get/set orchestration
│   ├── evaluator.go          # SwissKnife formula evaluator (int + float)
│   └── port.go               # Port node -> gvcp.Port byte-order-aware I/O
├── gentl/                    # GenTL 1.6 constants (types.go) — cti.go loader not built yet
├── internal/color/           # PFNC decode: DebayerToRGBA, DecodeHighDepth (Bayer/packed), EncodeJPEG
├── internal/genDC/           # GenDC v1.1 container/component/part/flow-table parsing
├── grab/                     # one-shot GrabJPEG convenience
├── live/                     # continuous preview loop (NewLive/WithSink/Stop, SetComponent)
└── root package files        # camera.go, device.go, stream.go, discovery.go,
                              # options.go, interfaces.go, alias.go, log.go, doc.go
```

The root package is a high-level facade over these protocol packages; `alias.go` re-exports protocol types (`Frame`, `Sample`, `Component`, `NodeMap`, `GVSPStream`, …) and constructors (`DialGVCP`, `ParseNodeMap`, `EncodeJPEG`, `FetchXML`, …) for ergonomic consumer use.

## Two consumer API paths

1. **Frame path (Phase 4):** `OpenDevice(ctx, ip)` → `*Camera`; `cam.SetInteger/SetEnum/...`; `cam.StartStream(ctx)` → `stream.Frames()` channel of pooled `*gvsp.Frame`; call `frame.Release()` to return buffers. Also `cam.StartGrabber` → `Grabber` for samples.
2. **Sample/JPEG path (Huaray BSCF):** `grab.GrabJPEG(ctx, ip)` or a `Grabber.Grab` → `gvsp.Sample` with color/depth/mono components + mm measurements. `live.NewLive(dev, live.WithSink(...))` for continuous preview driven by the app-owned `FrameSink`.

Options use functional patterns: `gige.WithLogger` (any `Logger`; default no-op), `gige.WithTimeout` (default 2s), `gige.WithComponent` / `GrabComponent` (color/depth/mono).

## Spec versions (_references/)

Authoritative standards live under `_references/` (`linguist-vendored`, excluded from archives). Map features to the pinned spec version:

| Package | Standard | Target |
| ------- | -------- | ------ |
| gvcp | GenCP (control channel over GigE UDP) | 1.3.1 |
| genapi | GenICam / GenApi | 2.1.1 |
| gvcp/gvsp | GigE Vision | 2.0 / 2.1 |
| gvsp (GenDC) | GenDC | 1.1 |
| gvsp + internal/color | PFNC pixel formats | 2.4 |
| gentl | GenTL | 1.6 (constants only today) |

Machine-readable truth for implementers: `GenDC/GenDC.h`, `GenTL/GenTL.h`, `SFNC/PFNC.h`. Track progress and open items in `ROADMAP.md`, not in prose.

## Committing (see .agents/skills/commit-changes)

The `commit-changes` skill defines the release workflow — follow it when invoked:

- Group changes by related topic; brief ≈50-char subjects (`feat:`, `fix:`, `chore:`).
- Put detail in `CHANGELOG.md` (Keep a Changelog style, dated under the next `## [X.Y.Z]`).
- Mark completed items `[ ]`/`[~]` → `[x]` in `ROADMAP.md`.
- Bump `Version` (root `options.go`) + annotated tag `vX.Y.Z` **only when code changed**. Version currently `1.4.0`.
- Don't push unless asked.
