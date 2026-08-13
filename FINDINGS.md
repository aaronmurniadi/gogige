# FINDINGS — spec-alignment audit vs `_references/`

Checked the current implementation against the pinned reference standards in
`_references/` (GenCP 1.3.1, GenApi 2.1.1, GenTL 1.6, GenDC 1.1, PFNC 2.4) and
the authoritative headers `GenDC/GenDC.h`, `GenTL/GenTL.h`, `SFNC/PFNC.h`.

Items marked **BUG** are confirmed deviations (many reproduced numerically).
Items marked **NOTE** are lower-confidence / documentation or spec-target gaps.

Priority: **HIGH** = misdecodes real data / can panic. **MED** = latent wrong
constant, no current effect. **LOW** = style/status only.

---

## HIGH — confirmed deviations

### H1. ~~`internal/genDC/genDC.go` — container `ComponentOffsets[]` read at wrong offset~~ **FIXED 2026-08-13**
- Status: `ContainerHeaderBaseSize` set to **56** (`internal/genDC/genDC.go:166`); unit
  tests updated to the correct byte-56 layout and pass (`go test ./internal/genDC ./gvsp`).
- Live camera note: DS5131 streams BSCF (not `PAYLOAD_TYPE_GENDC`) and only the non-config
  BSCF path was exercised on-camera (frame grab OK, no camera config changed). The GenDC
  fix was validated against the spec header + corrected unit tests, not a live GenDC frame.
- File: `internal/genDC/genDC.go:166` (`ContainerHeaderBaseSize = 64`)
- Spec: `GenDC/GenDC.h` `GenDCContainerHeaderBase` is **56 bytes** (packed);
  `ComponentOffsets[ComponentCount]` starts at **byte 56**, not 64.
- `parseContainerHeader` (genDC.go:415-421) reads the offset array from
  `ContainerHeaderBaseSize + i*8` = 64 + i*8 → **8 bytes past the real array**,
  so every component offset is read from the wrong window (typically 0).
- Effect: `ParseGenDCContainer` → component `parseComponent` uses `ComponentOffsets[i]`
  which is 0 for real GenDC payloads → components skipped / empty. **GenDC payloads
  (incl. `PAYLOAD_TYPE_GENDC`, BSCF GenDC frames) fail to extract image data.**
- Fix: `ContainerHeaderBaseSize = 56`. (All per-field reads at bytes 4–52 are already correct.)
- Verified numerically: stored 0x1000 at byte 56, impl reads 0x0 at byte 64.

### H2. ~~`internal/genDC/genDC.go` — part header min-size + `DataOffset` read guard~~ **FIXED 2026-08-13**
- Status: `PartHeaderBaseSize` fixed to **40** (`internal/genDC/genDC.go:178`),
  `PartHeader2DBaseSize`/`PartHeader2DSize` fixed to **56**. New regression test
  `TestPartHeaderTooShort` (genDC_test.go) reproduces the old OOB panic
  (`index out of range` on a 36-byte part header) and now passes
  (`go test ./internal/genDC ./gvsp`). Same live-camera caveat as H1: GenDC path
  only validated against the spec header + synthetic containers, not a live GenDC frame.
- File: `internal/genDC/genDC.go:178` (`PartHeaderBaseSize = 32`)
- Spec: `GenDCPartHeaderBase` is **40 bytes** (packed);
  `GenDCPartHeader2DBase` is **56 bytes**.
- `parsePart` (genDC.go:487-502) guards `len(buf) < 32` but then reads
  `binary.LittleEndian.Uint64(buf[32:])` (DataOffset, bytes 32–39). A part header
  between 32–39 bytes passes the guard and reads out of slice range → **panic**.
  Also the size constants `PartHeader2DBaseSize = 44` are off for the 2D base.
- Fix: `PartHeaderBaseSize = 40`; PartHeader2DBaseSize = `PartHeader2DSize` = 56.

### H3. ~~`internal/color/packed.go` — `Unpack14P` bit-reconstructs pixels 2/3/4 wrong~~ **FIXED 2026-08-13**
- Status: `Unpack14P` rewritten to decode 4 samples per 7-byte LSB-packed group
  per the PFNC formulas below; the stray 5th pixel and out-of-range `b7` read were
  removed. New regression test `TestUnpack14P` (color_test.go) reproduces the old
  wrong output (`0105 0234 0c35 3beb`) and now matches the reference
  (`0105 1234 2bcd 3def`); `go test ./internal/color ./internal/genDC ./gvsp` passes.
  Same live-camera caveat as H1/H2: 14p path validated against a reference decoder
  and synthetic packed data, not a live camera frame.
- File: `internal/color/packed.go:61-86`
- Spec: PFNC default **lsb Packed** (`GenICam_PFNC_2_4.pdf` §6.3.1, Figure 6-9).
  `Mono14p` / `Bayer*14p`: 4 samples packed into 7 bytes.
- Verified against a reference LSB-packed decoder with known values
  (`px = 0105 1234 2bcd 3def`):
  - reference: `0105 1234 2bcd 3def`
  - `Unpack14P`: `0105 0234 0c35 3beb`  ← pixels 2..4 wrong.
- Correct 14-bit (LSB-packed) unpack per 7-byte group `b0..b6`:
  - p0 = `b0 | ((b1&0x3F)<<8)`
  - p1 = `((b1&0xC0)>>6) | (b2<<2) | ((b3&0x0F)<<10)`
  - p2 = `((b3&0xF0)>>4) | (b4<<4) | ((b5&0x03)<<12)`
  - p3 = `((b5&0xFC)>>2) | (b6<<6)`
- Note: `Unpack10P` and `Unpack12P` were verified **correct**.

---

## MED — wrong constants (unused today, fix safely)

### M1. `gvcp/register_map.go` — wrong GigE ABRM subnet/gateway offsets
- File: `gvcp/register_map.go:50-51`
- `gevCurrentSubnet = 0x0034` and `gevCurrentGateway = 0x0044` do **not** match the
  GigE Vision ABRM layout (Current Subnet = `0x0028`, Current Gateway = `0x002C`).
- These two constants are not used by `parseDiscoveryAck` (which uses the correct
  `0x0008/0x000C/0x0024/0x0048/0x0068/0x00D8/0x00E8`), so no live bug — but the
  recorded addresses are wrong and any future user would misbehave.

---

## LOW — spec-target gaps / documentation notes

### L1. GVCP command IDs are classic GigE Vision, not GenCP 1.3.1 Table 7
- `gvcp/packet.go` uses classic GigE Vision GVCP command IDs
  (READREG 0x0080, READMEM 0x0084/0x0085, WRITEMEM 0x0086/0x0087, PENDING 0x0089, DISCOVERY 0x0002/0x0003, PACKETRESEND 0x0040).
- GenCP 1.3.1 §4.3.3 defines a technology-agnostic standardized command_id space
  (`README_MEM_CMD` = 0x0800, `WRITEMEM_CMD` = 0x0802, `PENDING_ACK` = 0x0805, …).
- For interop with real GigE Vision cameras the classic IDs are correct (they also
  follow the LSB = ack-flag convention). Clarify in docs that the **GigE Vision**
  command space is used, and the full GenCP 1.3.1 spec PDF (`_references/GenCP`)
  does not themselves define GVSP packet framing (missing GigE Vision 2.0 PDF).

### L2. GenApi: `Category` / `StructReg` parsed but skipped; `**` exponent missing
- `genapi/types.go:303` parses `Category`/`StructReg` as skippable kinds (matches
  ROADMAP `[ ]`). `genapi/evaluator.go` lacks `**` exponent (ROADMAP `[~]`).
  Already tracked in `ROADMAP.md`; no new work, listing for completeness.

### L3. GVSP data-leader field offsets need verification against GigE Vision 2.x
- `gvsp/receiver.go:155-157` reads leader image fields at `data[12/16/20]`
  (pixel format / size_x / size_y). The commonly-cited GVSP Data Leader layout
  places pixel format at `data[5]`, size_x at `data[9]`, size_y at `data[13]`
  (after a 4-byte payload-type + 1 reserved byte).
- **Not asserted as a bug**: the exact GVSP leader layout lives in the GigE Vision
  2.0 spec, which is **not present** in `_references/` (only a whitepaper and the
  SFNC feature ref). Verify against the AIA GigE Vision 2.0 GVSP chapter before
  changing, and if wrong the Huaray image path would read wrong W/H/format.

### L4. `gvsp/multi_part.go` header layout unverifiable from bundled refs
- Assumes 32-bit `NumParts` + 32-bit reserved = 8-byte header, then 32-byte part
  descriptors at offset 8. The GenTL 1.6 PDF describes multi-part handling at the
  buffer-API level, not the wire byte layout (GigE Vision 2.x owns it). Confirm
  against the authoritative GigE Vision spec; low confidence either way.

### L5. BSCF field offsets are vendor-specific (no `_references/` coverage)
- `gvsp/payload.go` BSCF block/pack offsets (536/576 headers, 64-strided blocks,
  `packDetSize=1472`, orientation at 12.., stable at 1276) are Huaray-proprietary.
  The reference docs do not define BSCF, so "aligned vs references" is N/A; keep
  tests as the source of truth for these field positions.

---

## Verified-correct (spot checks)

- **gvcp ABRM** (`register_map.go`): all GenCP 1.3.1 Table-19 offsets 0x0000–0x0250 exact.
- GVCP header layout, command/ack LSB convention, READMEM_ACK 4-byte reserved prefix.
- PENDING_ACK `temporary_timeout` at SCD offset 2 (uint16, ms) — exact (`pending.go`).
- PACKETRESEND_CMD non-extended (12B) and extended (20B) layouts.
- GVSP payload-type IDs = GenTL.h `PAYLOAD_TYPE_*` exactly.
- SFNC *naming* used for streaming features (`AcquisitionStart/Stop`).
- All PFNC pixel-format **IDs** (color.go/packed.go/genDC.go) match `PFNC.h`.
- `Unpack10P`, `Unpack12P`, `Unpack14P` count/stride logic (H3 now fixed).
- GenDC signature/version, component and part header base field offsets
  (H1/H2 now fixed), flow-table parsing (`GenDCFlowTableHeader` = 16 bytes).
  `GenDCPartHeaderBase` = 40, `GenDCPartHeader2DBase` = 56 — correct.
