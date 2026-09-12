# Changelog

## [1.7.0] - 2026-09-13

### Added

- `genapi` `Category` nodes (GenApi 2.1.1 §2.8.2) are now first-class: parsed into a `gcNode` with kind `Category` and an ordered `Features` (`<pFeature>`) list plus a `Visibility` element. New `NodeMap` traversal API exposes the grouping tree: `Category(name)`, `Categories()`, `RootCategories()`, and `CategoryTree(name)` (with a `CategoryNode{Name, Features, Categories}` type). The tree nests categories to arbitrary depth, skips dangling `pFeature` references, and guards against category cycles. `Camera.Categories()`, `Camera.RootCategories()`, and `Camera.CategoryTree()` surface it through the root facade; `CategoryNode` is re-exported from the root package.
- `genapi` `StructReg` nodes (GenApi 2.1.1 §2.8.6) are now expanded at parse time: each `<StructEntry>` becomes a `MaskedIntReg` node inheriting the `StructReg`'s shared elements (`Address`, `pAddress`, `Length`, `AccessMode`, `Endianess`, …) unless the entry defines its own, in which case the entry wins.

### Tests

- `genapi/category_test.go`: `TestCategoryTree` (spec §2.8.2 Root → nested tree, order, `RootCategories`, raw `Category` list, dangling-ref skip), `TestCategoryCycle`, `TestCategoryNoRoot`.
- `genapi/structreg_test.go`: `TestStructRegExpansion` (spec Format-7 example: inherited `Address 0x14`/`pAddress`/`Length 4`/`RO`/`BigEndian`/bit mask + functional big-endian register read), `TestStructRegEntryOverride` (entry-defined `Address`/`LSB`/`MSB` override inherited elements).
- Offline replay of the Huaray `DS5131MG30CE` XML: `nodes 1777 → 1821` (44 `Category` nodes now parsed), `readOK=1438` unchanged, error profile identical to the 1.6.2 baseline (`229` NI + `90` NA + `5` WO, `0` unexpected) — 0 regressions. `CategoryTree(Root)` resolves and `RootCategories()` returns the 23 top-level categories.

## [1.6.2] - 2026-09-12

### Added

- `genapi` float-domain SwissKnife evaluation: a float alternative `evalFormulaFloat`/`floatParser` to the §2.8.13 integer evaluator keeps arithmetic in the float domain (no `.Max` bound truncation, non-integer division), adds the float functions `ATAN2`/`MIN`/`MAX`/`POW`/`LOG` and the `E`/`PI` constants, and float-literal/scientific/hex number parsing.

### Fixed

- `genapi` evaluated every SwissKnife formula in the `int64` domain: float `SwissKnife` nodes and `pVariable` targets (a `Float`/`FloatReg` feature, e.g. `GainRawMaxExpr` = `VAR_GAINMAX` → `GainMax`) failed with `cannot evaluate integer (kind=Float)` — 9 errors on the offline Huaray replay. `ReadFloat` now evaluates `SwissKnife` via the float path, `formulaContextFloat`/`readFormulaVariableFloat` resolve float variables and `.Min/.Max/.Inc/.Value/.Entry` suffixes, and the integer path truncates the float result (with a negative-value guard) instead of erroring.

### Changed

- `genapi` SwissKnife §2.8.13 variable suffixes `.Min/.Max/.Inc/.Value/.Entry` and float-domain formulas are now fully resolved (previously int-domain only, a documented GAP in `FINDINGS.md`).

### Tests

- `genapi/evaluator_test.go`: `TestEvalFormulaFloat`, `TestEvalFormulaFloatSuffix`.
- `genapi/gaps_test.go`: `TestSwissKnifeFloatDomain` (mirrors the real Huaray XML layout: `FloatReg`→`Float`→`SwissKnife`→`<pMax>`; asserts float division stays `11.5`, integer reads truncate, and identity formulas carry the live float register value).
- Offline replay of the Huaray `DS5131MG30CE` XML: the 9 float-SwissKnife errors cleared (`333 → 324`; all remaining are guard-correct NI/NA/WO), 0 regressions.

## [1.6.1] - 2026-09-12

### Fixed

- `genapi` `resolveAddr` rejected the legitimate ABRM address `0x0` (e.g. `GevVersionReg`) and silently truncated `pAddress` sums over 32 bits; address 0 is now addressable and sums exceeding `math.MaxUint32` return an error instead of wrapping.
- `genapi` constant `Float`/`String` nodes (a declarative `<Value>` with no register address) read as their constant via `readFloatReg`/`readStringReg`; they previously failed with "no register address".
- `genapi` never derived or enforced runtime access modes per §2.5: RO writes / WO reads were unguarded and `ImposedAccessMode` was unparsed. `effectiveAccess` now reduces NI → NA → `AccessMode` ∩ `ImposedAccessMode` → locked (RW→RO, WO→NA), and `Set*`/`Read*`/`CurrentEnum` reject disallowed access — including on the resolved target register, not just the user feature.

### Changed

- `genapi` SwissKnife formulas resolve the §2.8.13 variable suffixes `.Min/.Max/.Inc/.Value/.Entry` (int-domain variables; float-domain `pVariable` targets remain a documented GAP).

### Tests

- `genapi/gaps_test.go`: `TestZeroAddressPermitted`, `TestAddressOverflowRejected`, `TestConstantFloatRead`, `TestConstantStringRead`, `TestReadOnlyWriteRejected`, `TestWriteOnlyRejected`, `TestNotImplemented`, `TestNotAvailable`, `TestLockedDowngrade`, `TestImposedAccessMode`, `TestSwissKnifeSuffixes`.

## [1.6.0] - 2026-09-12

### Added

- `genapi` formula evaluator now implements the full SwissKnife §2.8.13 grammar: exponent `**`, equality `==`/`!=` (incl. the `<>` alias), and the complete math function table `SGN NEG ABS FLOOR CEIL TRUNC ROUND SQRT EXP LN LG LOG SIN COS TAN ASIN ACOS ATAN ATAN2 MIN MAX POW` plus the `E`/`PI` constants. Parse errors report the offending token.
- `genapi` `Converter`/`IntConverter` nodes implement the GenApi §2.8.10 reserved variables: read maps the register value to `TO` (`FormulaFrom`, e.g. `(TO & 0x00080000) >> 19`); write maps the user value to `FROM` with `pVariables` evaluated as the current register domain for read-modify-write. Legacy single-variable converters keep the old `X`-binds-user-value binding.
- `genapi` `Boolean` nodes honor `<OnValue>`/`<OffValue>` when a camera overrides the §2.8.7 default of `1`/`0` (`SetBoolean`/`ReadBoolean`).
- `genapi` `IntReg`/`MaskedIntReg` apply per-node `<Sign>` sign-extension and `<Endianess>` byte order (defaulting to device order), and the read path applies the LSB/MSB mask.
- `gvcp.ReadManifestTable` parses the conformant GenCP 1.3.1 Manifest Table (u64 count + 64-byte entries carrying FileVersion/Schema/FileSize/SHA1) in addition to the vendor Huaray "MTAB" layout; `ManifestEntry.GenCP` reports which layout produced the entry. `ManifestTableURL` serves both.
- `gentl` adds the GenTL 1.6 `*_CustomID = 1000` constants: `InfoDataTypeCustomID`, `TLInfoCustomID`, `InterfaceInfoCustomID`, `DeviceAccessCustomID`, `DeviceAccessStatusCustomID`, `DeviceInfoCustomID`, `AcqStopFlagsCustomID`, `AcqStartFlagsCustomID`, `AcqQueueCustomID`, `StreamInfoCustomID`.

### Fixed

- `genapi` formula evaluator parsed `<>` as `<` then `>` ("bad token at `> 0`") on camera formulas like `GevTimestampControlResetAvailExpr`; `<>` is now a single spacing-tolerant token at the equality tier.
- `internal/color` Bayer demosaic mixed up the R/B channels on `RGGB` and `GBRG` patterns and mishandled green-site sampling; the kernels now bounds-check at image edges. Regression-tested by a uniform-tile fidelity test for all four patterns in the 8- and 16-bit paths.
- `internal/color` `DecodeHighDepth` for unpacked `Mono10`/`Mono12`/`Mono14` (LE `uint16`) now left-aligns samples to 16 bits (`<< (16-bits)`) instead of surfacing the raw low byte.
- `gvsp.ParseChunkPayload` now parses the GenDC 1.1 §2.2.8.1 trailing-tag chunk format (the actual wire format) instead of a 16-byte-header layout that matched no device; `IsChunkData` validates the tag chain back to byte zero.
- `internal/genDC` treats `PartOffset[]` as container-relative per GenDC §2.2.4 (was sliced relative to the component, misparsing any component whose part data did not immediately follow).

### Changed

- `gvsp` payload-type helpers now use real GenTL ids (`PayloadTypeChunkData = 0x00000004`, `PayloadTypeGenDC = 0x0000000B`, `PayloadTypeMultiPart = 0x0000000A`); the vendor high-bit aliases encode the same ids (`0x80000000 | id`). `ComponentDepth` comment notes the Huaray BSCF wire value vs. GenDC `Range` (`0x04`).

### Tests

- `genapi`: `TestEvalFormulaNotEqual`, `TestEvalFormulaTernaryRegression`, `TestConverterReservedFromTo`.
- `gvcp`: `TestReadManifestTableGenCP`.
- `gvsp`: `TestChunkPayload` rewritten for the trailing-tag format; `TestChunkPayloadTruncated`.
- `internal/color`: `TestBayerFidelityUniform` (four patterns, 8- and 16-bit); `TestDecodeHighDepth_MonoPacked` updated.

## [1.5.0] - 2026-08-21

### Added

- `calib.CamCalib` pinhole projection: `ProjectPoint3D` / `DeprojectPixel` convert between camera-frame millimetres and pixels, rescaling from calibration resolution by width ratio only — mirroring the vendor SDK's `stereoConvetPoint3dToDepth` / `stereoConvetDepthToPoint3d`.
- `calib.LoadVendorFile` + `VendorCalibJSON.Color`/`Left`/`RectifiedLeft` load intrinsics from the vendor calibration export ("IPC4.94 Camera Calibration.json" format); `RectifiedLeft` exposes the rectified-left projection `P`.
- `calib.ReadStereoCalib` downloads the stereo/color calibration from camera memory bank `0x20001` over GVCP, replicating the vendor `readData` protocol (bank select `0xE0000000`, size/CRC `0xE0000004`/`0xE0000008`, chunked data window `0xE0000100`, ack `0xE000000C`, IEEE CRC32 verify). Any `*gvcp.GVCP` satisfies the new `calib.RegisterPort`. Returns a parsed `StereoCalib` (full `MvSstereoCalibrateResult` layout: left/right/color intrinsics + distortion, extrinsics, rectify rotations, `P`s, disparity-to-depth `Q`, valid ROIs, RMS errors) with `Color`/`Left`/`RectifiedLeft` accessors.
- `calib.ReadCalibTypes` probes bank `0x20000` for the semicolon-separated calibration type names (e.g. `calibration_pd`).

### Tests

- `TestProjectPoint3DHandComputed`, `TestProjectRescalesToOutputResolution`, `TestDeprojectIsInverseOfProject`, `TestInvalidInputsYieldNaN`, `TestDS5131SampleProjection` (pins a live-measured BSCF pack centre against the exported color intrinsics).
- `TestLoadVendorFileAndAccessors`, `TestLoadVendorFileErrors` (vendor JSON path), `TestReadStereoCalib` (fake register port serving a 1600-byte blob built from real DS5131MG30CE export values; asserts protocol register writes and parsed fields), `TestReadStereoCalibErrors` (empty bank acks then `ErrNoCalib`; CRC mismatch; short blob), `TestReadCalibTypes`.

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
