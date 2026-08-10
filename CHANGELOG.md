# Changelog

## [1.1.0] - 2026-08-10

### Added

- Package-level `gvcp.ReadManifestTable(port)` and `gvcp.ManifestTableURL(port)` for reading the GenCP ManifestTable without requiring a live `*GVCP` connection; `*GVCP` convenience wrappers preserved. Tests cover zero-address, full table, and URL extraction.
- `genapi.FetchXML` now prefers `URLReader.ManifestTableURL()` over `FirstURL()` when the device exposes a ManifestTable.
- GenApi availability and lock pointers: `pIsImplemented`, `pIsAvailable`, `pIsLocked`, and `pInvalidator` are parsed and exposed via `NodeMap.IsImplemented()`, `NodeMap.IsAvailable()`, `NodeMap.IsLocked()`, `NodeMap.GetInvalidator()`, and `Node.GetInvalidator()`. Missing pointers default to implemented/available/not-locked.
- SwissKnife math functions: `ABS`, `FLOOR`, `CEIL`, `SQRT` supported in formula evaluation (`genapi/evaluator.go`). `SQRT` returns 0 for negative inputs; all others pass-through for integer values.

## [1.0.0] - 2026-08-10

### Added

- `gogige/grab` subpackage with `GrabJPEG` one-shot capture.
- `gogige/live` subpackage with `NewLive`, `WithSink`, `WithOnSample`, `WithLiveComponent`, and the continuous preview loop.

### Changed

- Root package reduced from 14 to 9 files: `grab.go` → `grab/`, `live.go` → `live/`, `framestream.go` merged into `stream.go`, `logger.go` + `logger_iface.go` merged into `log.go`, `version.go` merged into `options.go`.
- Removed root-level convenience wrappers `GrabJPEG`, `NewLive`, `WithSink`, `WithOnSample`, `WithLiveComponent`; import subpackages directly (`gogige/grab`, `gogige/live`).
- Phase 4 high-level API (`OpenDevice`, `StartStream`, `Frames()`, `Discover`, `Camera`, `Session`, `Stream`) remains in the root package.

### Removed

- `gogige.GrabJPEG` (use `gogige/grab.GrabJPEG`).
- `gogige.NewLive` / `gogige.WithSink` / `gogige.WithOnSample` / `gogige.WithLiveComponent` (use `gogige/live.*`).

## [0.13.2] - 2026-08-09

### Changed

- Renamed the full-stack example `examples/smoke` → `examples/smoke-test` and expanded it into five sequential phases exercising the whole current stack on one camera: (1) raw GVCP control channel — CCP access privilege, byte order, register-map reads, heartbeat hold/pulse, FirstURL, GenApi XML build; (2) GenApi feature surface — `Has`/`Kind`, integer reads, `GetMin`/`GetMax`/`GetInc` constraints, `EnumEntries`/`CurrentEnum`, and non-destructive `SetInteger`/`SetEnum` round-trips; (3) streaming Session — `StartGrabber`/`Grab`, `GrabAll` component enumeration, `SetComponent`, `Pause`/`Resume`, JPEG write-out; (4) Phase-4 channel API `StartStream`/`Frames()`/`frame.Release()` (opt-in `-channel`); (5) `Live` preview with `OnSample`/`LatestSample` (opt-in `-live`).
- The example module path is now `github.com/aaronmurniadi/gogige/examples/smoke-test`; README/ROADMAP references updated.

## [0.13.1] - 2026-08-09

### Added

- `gentl` opaque handle types mirroring `GenTL.h` (`TL_HANDLE`, `IF_HANDLE`, `DEV_HANDLE`, `DS_HANDLE`, `PORT_HANDLE`, `BUFFER_HANDLE`, `EVENTSRC_HANDLE`, `EVENT_HANDLE`) plus `IsValidHandle` — no CGO, zero value is the invalid handle.

## [0.13.0] - 2026-08-09

### Added

- GVSP payload type constants (GenTL/GigE Vision leader field IDs) in `gvsp/payloadtype.go`: `PayloadTypeImage`, `PayloadTypeChunkData`, `PayloadTypeChunkOnly`, `PayloadTypeMultiPart`, `PayloadTypeGenDC`, plus vendor aliases and `PayloadTypeName` / `IsPayloadType*` helpers.

- `gvsp.Frame.PayloadType` populated from the GVSP leader payload-type-specific header; vendor/custom leaders (e.g. BSCF) keep it `0`.

- GenDC flow table (`GDC_FLOW_TABLE_HEADER`) parsing: `internal/genDC` `IsFlowTable`, `ParseFlowTable`, `FlowTableFromContainer`; `gvsp.GenDCPayload.FlowTable`.

- Correct 2D part dimension extraction (SizeX/SizeY) and absolute data-offset handling in GenDC parts; `ParsedGenDcComponent.Width/Height` now derived from the 2D part header.

- PFNC decoder matrix completion in `internal/color`: Bayer8 debayering wired into `EncodeJPEG`, plus `DecodeHighDepth` for unpacked/packed Mono (10/12/14 bit) and packed/unpacked Bayer (10/12/14/16 bit) with LSB→MSB alignment.

### Changed

- `ParsePayloadByType` dispatches on the GVSP payload type constants instead of ad-hoc 0x800000xx values.

### Fixed

- GenDC part `DataOffset` now holds the absolute container data offset (from the part header) rather than the part-header position, so `ParseGenDcPayload` extracts the correct image bytes.

- OOM on long-lived streams: `OOOPacketRing` pre-allocated every slot at `DefaultFrameSize` (8 MiB) — 256 slots × 8 MiB ≈ 2 GiB per in-flight `frameBuild` — ballooning RSS to >10 GiB on the websocket/live examples until the kernel killed the process. Slots are now capped at 16 KiB (one GVSP transport packet), and `gvsp.Stream` bounds concurrent in-flight frames (`maxInFlightFrames=64`), evicting the oldest incomplete build when full.

### Tests
- `TestPayloadTypeNames`, `TestParsePayloadByTypeDispatch`, `TestGVSPLeaderPayloadType`, `FlowParseGenDCContainer`/`TestParseFlowTable` (internal/genDC), and color Bayer/packed decode tests.

## [0.12.0] - 2026-08-09

### Changed

- GVSP out-of-order packet reassembly now uses a pre-allocated `OOOPacketRing` in place of `map[uint32][]byte`, removing per-packet heap allocations on the OOO gap-refill path (zero-alloc hot path per AGENTS.md). Ring spills to a lazily-created overflow map only past `MaxOOOPackets` (256).

- `gvsp/resend.go` adds `MissingPayloadRangesRing`, the ring-backed equivalent of `MissingPayloadRanges`, used for gap/trailer resend computation.

### Fixed

- `OOOPacketRing` middle-delete no longer drops the ring-head packet (previous compact-on-delete advanced head unconditionally and lost the oldest entry).

- `appendPayload` ring-full fallback now stores the packet via the ring overflow map instead of allocating a copy and discarding it (packets were silently dropped).

### Added

- Exported `OOOPacketRing`, `NewOOOPacketRing`, `RingBufferSlot`, and `MaxOOOPackets` in `gvsp/frame.go`.

- Tests: `TestOOOPacketRingPutGetDelete`, `TestOOOPacketRingOverflowSpill`, and `TestGVSPOutOfOrder` (OOO arrival + resend hole-fill).

## [0.11.1] - 2026-08-09

### Removed

- CGO from gentl package: removed gentl/cti.go (dlopen/dlclose cannot be pure Go), rewrote gentl/types.go with pure Go constants

### Added

- doc.go with protocol references to gvcp, gvsp, genapi, and root packages

### Changed

- gentl: now provides only pure Go constants (no CGO, no producer loading)
- doc.go: added all protocol specification references (GenCP v1.3.1, GenTL v1.6, GenDC v1.1, GenApi v2.1.1, GenICam v2.1.1, SFNC v2.7)

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
