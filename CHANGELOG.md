# Changelog

## [0.11.0] - 2026-08-08



### Added

- GenDC 1.1 payload parsing (GNDC signature, container header, component/part structures)

- Multi-part payload support (GenTL v1.5) with `ParseMultiPartPayload` and `MultiPartPayload` type

- Chunk data payload support (GenTL v1.2/v1.4) with `ParseChunkPayload` and `ChunkPayload` type

- Payload type dispatcher: `ParsePayloadByType` routes by payload type ID

- YUV422_8_UYVY (0x0210001F) decode in color package alongside existing YUYV support

- Bayer format debayering (RGGB, BGGR, GBRG, GRBG patterns) with full 8-bit and packed (10p/12p/14p) formats

- Packed format decoders: `Unpack10P`, `Unpack12P`, `Unpack14P` for Mono/Bayer packed formats

- New internal modules: `internal/genDC`, `internal/color/bayer.go`, `internal/color/packed.go`

- New GVSP payload handlers: `gvsp/genDC_payload.go`, `gvsp/multi_part.go`, `gvsp/chunk_data.go`

- New tests: `TestIsGenDCPayload`, `TestMultiPartPayload`, `TestChunkPayload`

- New examples: verified all 10 examples build successfully



### Changed

- `PixelFormatYUV422` alias to `PixelFormatYUV422_8` for consistency

- GenDC and multi-part/chunk parsers integrated into `gvsp/payload.go` payload dispatch

- All tests pass, all examples build
All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.10.0] - 2026-08-08

### Added
- GenApi Phase 3 completion: constraint pointers (`pMin`, `pMax`, `pInc`) now resolved alongside static constraint values (`Min`, `Max`, `Inc`). `NodeMap.GetMin` / `GetMax` / `GetInc` methods enable parameter bounds validation. Constraint pointer nodes (typically SwissKnife formulas) are evaluated using the existing formula evaluator, falling back to static values when pointers are not present
- GenApi module refactoring: split monolithic `nodemap.go` into four focused modules per AGENTS.md architectural hierarchy: `node.go` (core Node interface + gcNode), `types.go` (XML parsing pipeline), `port.go` (gvcp.Port binding + byte order), `nodemap.go` (high-level orchestration). Zero-allocation design maintained throughout
- Test: `TestConstraintPointers` validates both pointer-based and static constraint parsing and resolution

### Fixed
- **GVCP packet size negotiation** (`gvcp.StartAcquisition`): now returns the actual negotiated packet size instead of discarding it. Previously, when device rejected the requested packet size (e.g., 9000) and fell back to 1500 via binary search, the fallback value was ignored. This caused the warning "GevSCPSPacketSize clamped by device" to be logged even when the negotiation succeeded at a lower size. Now the negotiated size is returned and used, eliminating false warnings. Actual post-acquisition downgrades (device changing the size after acquisition starts) are still warned
- `Session.ResumeStreaming` now updates `s.packetSize` with the actual negotiated value

### Changed
- `Node` interface extended with `GetConstraints()` method returning all constraint fields (pMin, pMax, pInc, Min, Max, Inc)
- `gvcp.StartAcquisition` signature changed: now returns `(actualPacketSize int, err error)` instead of `error` only

## [0.9.0] - 2026-08-08

### Added
- Phase 4 channel API: `gogige.OpenDevice` connects and returns `*Camera` directly; `Camera.SetInteger` / `Camera.SetEnum` feature setters; `Camera.StartStream` opens a GVSP channel delivering pooled frames over `Stream.Frames()` (`<-chan *gvsp.Frame`) with idempotent `Stop` plus `Pause` / `Resume`. `Stop` terminates the background loop promptly on context cancellation / Ctrl+C
- New example `examples/frames`: the Phase 4 happy path (`OpenDevice` → `StartStream` → `range Frames` → `Release`), verified to exit cleanly on SIGINT

### Changed
- Protocol alias `Stream = gvsp.Stream` renamed `GVSPStream` so the root `Stream` name is now the Phase 4 channel type
- Control examples (`configure-camera`, `features`, `probe-streams`) migrated from `gogige.Connect` to `gogige.OpenDevice` with a 3s GVCP timeout
- `Session.recvFramePtr` hands pooled frames to the channel loop without copying (zero-copy delivery)

## [0.8.1] - 2026-08-08

### Changed
- Library and all examples now declare `go 1.23` (was `1.26.3`); verified to build and pass tests with Go 1.23.12. `mise.toml` pins `go = "1.23"` so development tooling matches the module requirement

## [0.8.0] - 2026-08-08

### Added
- `gogige.Slog(*slog.Logger)` adapter (stdlib `log/slog`): `gogige.WithLogger(gogige.Slog(slog.Default()))` is the new logging setup, so `github.com/rs/zerolog` is no longer a library dependency or a transitive dependency of any example

### Changed
- `gogige.Zerolog(...)` removed (breaking): replace `gogige.Zerolog(z)` with `gogige.Slog(...)` backed by `log/slog`; `go.mod`/`go.sum` dropped zerolog + `mattn/go-colorable`/`go-isatty`/`x/sys` from root and all example modules
- `.githooks/pre-commit` now gofmt's the whole project (not just staged files) before `go test ./...`
- `.gitattributes` marks `_references/**` linguist-vendored + export-ignore so third-party spec PDFs/headers stay out of release archives and language stats
- README overhauled: shields badges, table of contents, features, project structure, contributing, and license sections; roadmap header wording updated

## [0.7.0] - 2026-08-08

### Added
- GenICam introspection reads: `NodeMap.Kind`, `NodeMap.EnumEntries`, `NodeMap.CurrentEnum`, `NodeMap.ReadInteger`, `NodeMap.ReadBoolean`, plus `Camera.BooleanFeature` — probe feature kinds, enumeration entries/current value, and integer/boolean values without exposing the whole tree
- `Session.GrabComponents`: receive one frame and return raw BSCF components (Data, dimensions, pixel format) without JPEG encoding; non-BSCF payloads come back as a single `ComponentUnknown` sample — enables probing what a camera actually streams
- New example `examples/probe-streams`: probes stream features (pixel formats, payload/chunk/component), sets `PixelFormat`, grabs one frame, and reports per-component JPEG encodability
- New example `examples/configure-camera` (replaces `configure-volume-tcp`): reads the Huaray volume-camera `ImageStoreEnable`, toggles it, then restores it via `BooleanFeature` + `ApplyControlPair`

### Changed
- All examples simplified to a single `main` with inline GigE discovery (dump-xml, features, grab, grab-components, smoke, stream-measurements, websocket-stream); removed `resolveIP`/`deviceLabel` helpers
- `.gitignore` now ignores example and `cmd/` build artifacts generically instead of listing binaries one by one

## [0.6.1] - 2026-08-08

### Changed
- Consolidate agent rules: merge `.agents/rules/*.mdc` (architecture, package layout, GVCP/GVSP/GenApi phases, Go conventions, perf guardrails) into a single root `AGENTS.md`; `ROADMAP.md` refs now point at the corresponding `AGENTS.md` sections
- Relocate the `commit-changes` skill from `.cursor/skills/` to `.agents/skills/`

## [0.6.0] - 2026-08-08

### Changed
- Rename root package `gige` → `gogige` so the package name matches its import path `github.com/aaronmurniadi/gogige` (breaking: importers must update `gige.X` → `gogige.X`)
- `doc.go` now documents the Phase 4 happy path (`gogige.OpenDevice` / `StartStream` / `Frames`) as the primary API; Grabber/JPEG path noted as BSCF-specific
- All examples and `cmd/` CLIs updated to the new package name

## [0.5.0] - 2026-08-08

### Changed
- Rename BSCF `ImageKind` API to SFNC-style `Component` (`ComponentColor` / `ComponentDepth` / `ComponentMono`)
- Options/methods: `WithComponent`, `GrabComponent`, `SetComponent`, `WithLiveComponent`, `ParseComponent`, `SampleFromBSCFComponent`
- CLI flags `-image` → `-component`; example `grab-images` → `grab-components`
- `ParseComponent` also accepts `range` (depth) and `intensity` (mono)

### Added
- `Live.SetComponent` / `Live.Component` for mid-stream component switch
- `examples/websocket-stream -component color|depth|mono`

## [0.4.0] - 2026-08-08

### Added
- BSCF `ImageKind` (color / depth / mono): parse all image blocks; `SampleFromBSCFKind` / `SampleAllFromBSCF` / `IsBSCF`
- Grab selection: `WithImageKind`, `GrabImageKind`, `SetImageKind`, `WithLiveImageKind`; `Grabber.GrabAll` for one frame → all modes
- Mono16 JPEG preview (high byte) for depth maps
- `examples/grab-images`: discover BSCF modes and write `<kind>.jpg`
- `gogige-stream -image color|depth|mono`

### Fixed
- Grab no longer treats a missing BSCF image kind as a raw non-BSCF frame
- `GrabAll` skips empty/zero-size BSCF placeholder blocks instead of failing the whole grab

## [0.3.0] - 2026-08-08

### Added
- `examples/smoke`: end-to-end CCP/heartbeat/XML/stream/pause/Live exercise on a live camera
- `examples/features`: GenICam `Has` / `ApplyControlPair` / `Execute` CLI
- `cmd/gogige-stream`: capture N frames (JPEG + BSCF measurements)

## [0.2.1] - 2026-08-08

### Added
- GenCP Technology-Agnostic ABRM constants (`0x0000–0x0250`) alongside GigE Vision ABRM/SBRM in `gvcp/register_map.go`
- `PENDING_ACK` parses GenCP `temporary_timeout` (ms) and extends the UDP read deadline
- `SyncImplementationEndianness` / `DeviceByteOrder` (probe `0x020C` on TakeControl; GenApi WriteMem uses device order)

## [0.2.0] - 2026-08-08

### Added
- GVSP packet resend: gap detection, hole-fill reassembly past trailer, `gvcp.EncodePacketResend` / `RequestResend`, and `Stream.SetResender` (Session wires channel 0)
- Path MTU discovery (`gvsp.PathMTU` / `PacketSizeForMTU`) and `GevSCPSPacketSize` negotiate with device clamp (RMW on `0x0D04`)
- `SO_RCVBUF` default 16 MiB on the stream socket with warn when the kernel grants less than 8 MiB
- `Camera.Logger()` for stream setup / clamp / rcvbuf warnings

### Changed
- Trailer no longer finalizes a frame until missing payload packets are filled (or resend is disabled and the contiguous prefix is complete)

## [0.1.0] - 2026-08-08

### Changed
- `examples/dump-xml` default filename is `<vendor>-<serial>-genicam.xml` (path-safe tokens from discovery ABRM) instead of an IP-based `*-genapi.xml` name

### Added
- GigE Vision local references: `GigE_Vision_for_Realtime_MV_11052010.pdf`, `GigE_Features_Reference.pdf`
- Root `.gitignore` for example camera dumps (`*.xml` / `*.jpg` / `*.log` under `examples/`)
- Cursor skill `.agents/skills/commit-changes` for topic-split commits, changelog, roadmap, and version bumps
