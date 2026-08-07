# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
- Cursor skill `.cursor/skills/commit-changes` for topic-split commits, changelog, roadmap, and version bumps
