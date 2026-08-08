# Changelog

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
