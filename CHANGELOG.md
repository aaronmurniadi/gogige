# Changelog

## [1.5.0] - 2026-08-21

### Added

- `calib.CamCalib` pinhole projection: `ProjectPoint3D` / `DeprojectPixel` convert between camera-frame millimetres and pixels, rescaling from calibration resolution by width ratio only — mirroring the vendor SDK's `stereoConvetPoint3dToDepth` / `stereoConvetDepthToPoint3d`.

### Tests

- `TestProjectPoint3DHandComputed`, `TestProjectRescalesToOutputResolution`, `TestDeprojectIsInverseOfProject`, `TestInvalidInputsYieldNaN`, `TestDS5131SampleProjection` (pins a live-measured BSCF pack centre against the exported color intrinsics).

## [1.4.0] - 2026-08-13

### Added

- `Camera` one-shot grab helpers so a consumer already on the `OpenDevice` path can grab samples without juggling `Device`/`Grabber`: `Camera.GrabSample` (single component + JPEG), `Camera.GrabAllSamples` (every BSCF component), `Camera.GrabComponents` (raw, no JPEG), and `Camera.GrabJPEG`. Each opens a transient GVSP stream and closes it before returning; the `Camera` stays open.
- `Camera.Features()` returns a `Features` view backed by the `Camera`, so a `*Camera` and a `Device` share the same feature vocabulary.
- Camera feature getters: `Camera.Integer`/`Boolean`/`Float`/`String`/`Enum`, plus `Features.Bool`/`Int`/`Float`/`String`/`Enum`. Reads are backed by new `NodeMap.ReadFloat` and `NodeMap.ReadString` register-read paths (previously only write paths existed).
- `grab.FromCamera(ctx, cam, comp)` one-shot JPEG grab from an already-connected `*Camera`.

### Changed

- `Camera` feature control unified onto the short Phase-4 names: `SetInteger`, `SetEnum`, `SetBoolean`, `SetFloat`, `SetString` are now primary; the previous `Set*Feature` forms remain as aliases so existing code keeps compiling. `SetInteger`/`SetEnum`/`Camera.Features` share one consistent vocabulary with `Device.Features`.
- `gvsp.PackDet` measurement fields renamed for consistency with `Sample`: `Length` → `LengthMm`, `Width` → `WidthMm`, `Height` → `HeightMm` (breaking). `gvsp.Sample` already used `*Mm`.
- Sink flow control decoupled from frame delivery (breaking): `FrameSink` now only requires `SendJPEG`; it no longer declares `Freeze`/`Resume`. Flow control moved to a new optional `Throttler` interface (`Throttle`/`Unthrottle`). `Live` type-asserts the sink to `Throttler`, so stateless sinks like `JPEGFunc` no longer carry no-op stubs. `hub` in `examples/websocket-stream` implements `Throttler` directly, and the client "freeze"/"resume" wire messages now call `Throttle`/`Unthrottle`.

### Tests

- Existing suite still passes; no new tests (interface contract refactor exercised by `go build ./...` and `go test ./...`).

## [1.3.1] - 2026-08-13

### Fixed

- `internal/color` `Unpack14P` (`Mono14p` / `Bayer*14p`) reconstructed pixels 2/3/4
  wrong per the PFNC lsb-packed layout, and had a stray 5th pixel that read past
  the 7-byte group. Rewritten to decode the 4 samples per group. Verified against a
  reference decoder (`0105 1234 2bcd 3def`).
- `internal/genDC` `PartHeaderBaseSize` was `32` but the packed `GenDCPartHeaderBase`
  is `40` bytes, so `parsePart`'s `len < 32` guard passed for 32–39 byte headers and
  the `DataOffset` read at `buf[32:]` went out of slice range → panic. `PartHeaderBaseSize`
  fixed to `40`; `PartHeader2DBaseSize` fixed to `56` (matching `PartHeader2DSize`).

### Tests

- `TestUnpack14P` (`internal/color`) reproduces the old wrong output
  (`0105 0234 0c35 3beb`) and now matches the reference.
- `TestPartHeaderTooShort` (`internal/genDC`) reproduces the old OOB panic on a
  36-byte part header and now passes.

## [1.3.0] - 2026-08-13

### Fixed

- `internal/genDC` container `ComponentOffsets[]` read at the wrong offset: `ContainerHeaderBaseSize` was `64` but the packed `GenDCContainerHeaderBase` is `56` bytes (per `_references/GenDC/GenDC.h`), so every component offset was read 8 bytes past the real array (typically `0`), causing GenDC payloads (incl. `PAYLOAD_TYPE_GENDC`) to skip/empty their components and fail to extract image data on real payloads.

### Tests

- `TestParseGenDCContainer` (`internal/genDC`) and `TestParsePayloadByTypeDispatch`/`buildGenDCContainer` (`gvsp`) updated to the corrected byte-56 component-offset layout; both packages pass.

## [1.2.0] - 2026-08-10

### Added

- `gvcp.StatusError` with `Code`/`Cmd`, returned when a device ACKs a GVCP request with a non-zero status (e.g. `INVALID_ACCESS`); message format matches the previous error text.
- `gogige.Sample.Packs []PackDet` exposing every volume pack on the frame; `PackDet` now carries the pack center (`CenterX`/`CenterY`/`CenterZ`) and a 3x3 `Orientation` (axis0, axis1, surface normal) alongside `Length`/`Width`/`Height`/`Volume`/`Stable`.

### Changed

- `gogige.Sample` field renames (breaking): `Width` → `PixelWidth`, `Height` → `PixelHeight`, `Length` → `LengthMm`, disambiguating pixel vs. physical measurement dimensions now that per-pack measurements exist.
- BSCF pack count is derived from the densely-packed `packDetSize` payload records rather than the descriptor slot — DS5131 writes 1 there regardless of record count; the descriptor value remains a fallback for empty payloads.

### Fixed

- `gvcp.ReadManifestTable` no longer fails on cameras without GenCP ManifestTable support: a bootstrap read rejected with `INVALID_ACCESS`/`WRITE_PROTECT` is treated as "no table" (`nil, nil`), so callers fall back to `FirstURL`.

### Tests

- `TestStatusErrorString`, `TestReadManifestTableInaccessibleIsNoTable`, `TestBSCFPackCountFromPayload`, and center/orientation round-trip assertions in `TestBSCFRoundTrip`.

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
