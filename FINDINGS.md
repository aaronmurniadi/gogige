# FINDINGS — gogige Implementation vs _references Standards Audit

Standards-compliance audit of the Go implementation against the reference documents under `_references/`.
Audit date: 2026-09-12.

**Coverage.** The GigE Vision spec itself (GVCP command set / GVSP streaming / packet formats) is **NOT** present in `_references/`. GVCP was therefore audited against the GenCP 1.3.1 spec (whose register map and packet _framing_ apply), and GVSP reassembly was audited against GenDC v1.1 + GenTL 1.6 only. Anything GVSP-specific is flagged `UNVERIFIABLE`.

| Module                      | Reference used                                                         | Verdict                                                                                                                |
| --------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `gvcp/`                     | `_references/GenCP/GenICam_GenCP_1.3.1.pdf.ocr/markdown.md`            | Framing + ABRM correct; protocol is GVCP-on-GenCP-framing; several gaps (see §1)                                       |
| `gvsp/` + `internal/genDC/` | `_references/GenDC/GenDC.h` + `GenICam_GenDC_v1_1.pdf.ocr/markdown.md` | GenDC header/part parsing exact; PartOffset base wrong; chunk parser non-standard; GenDC payload-type alias wrong (§2) |
| `genapi/`                   | `_references/GenApi/GenICam_Standard_v2_1_1.pdf.ocr/markdown.md`       | Solid skeleton, read-side fragment; MaskedIntReg reads, Converter, Endianess, floating nodes missing (§3)              |
| `internal/color/`           | `_references/SFNC/PFNC.h` + `GenICam_PFNC_2_4.pdf.ocr/markdown.md`     | Constants exact; **RGGB & GBRG demosaic R/B-swapped**; Mono10/12/14 unpacked near-black (§4)                           |
| `gentl/`                    | `_references/GenTL/GenTL.h` + `GenICam_GenTL_1_6.pdf.ocr/markdown.md`  | Every constant numerically exact; only gap is missing `_CUSTOM_ID = 1000` sentinels (§5)                               |

---

## 1. GVCP (`gvcp/`) vs GenCP 1.3.1

Implementation files: `client.go`, `discovery.go`, `heartbeat.go`, `packet.go`, `pending.go`, `register_map.go`, `endianness.go`, `acquisition.go`.

### MATCHES

- Packet/CCD framing is GenCP-shaped: 8-byte header (flags, command_id, length, request_id), request/ack pairs, 16-bit `reqID`. Layout matches GenCP §4.3.1 CCD format (`_references/GenCP/GenICam_GenCP_1.3.1.pdf.ocr/markdown.md`, CCD Table 5).
- `pendingAckTimeout` (`gvcp/pending.go:14`) parses the GenCP PENDING_ACK SCD — `temporary_timeout` in ms at SCD offset 2, fallback when SCD length < 4 — matches GenCP Table 12.
- `Abrm*` bootstrap-register constants (`gvcp/register_map.go:20-47`) match GenCP Table 19 exactly: `GenCPVersion=0x0000`, `Manufacturer=0x0004`, `Model=0x0044`, `Family=0x0084`, `DeviceVersion=0x00C4`, `ManufacturerInfo=0x0104`, `Serial=0x0144`, `UserDefined=0x0184`, `DeviceCapability=0x01C4`, `MDRT=0x01CC`, `ManifestTableAddress=0x01D0`, `SBRM=0x01D8`, `DeviceConfiguration=0x01E0`, `HeartbeatTimeout=0x01E8`, `MessageChannelID=0x01EC`, `Timestamp=0x01F0`, `TimestampLatch=0x01F8`, `TimestampIncrement=0x01FC`, `AccessPrivilege=0x0204`, `ImplementationEndianness=0x020C`, `DeviceSoftware=0x0210`…
- `AbrmCapEndiannessRegister = 1<<10` (`gvcp/endianness.go:16`) matches GenCP DeviceCapability bit 10 (endianness register supported).
- ImplementationEndianness semantics: only `0` (big-endian) and `0xFFFFFFFF` (little-endian) accepted, anything else → default unchanged (`gvcp/endianness.go:48-75`). Matches GenCP; bootstrap/CCD fields always big-endian regardless.
- ReadMem chunking to 512 bytes with 4-byte alignment (`gvcp/client.go:198-238`) matches the GVCP max data size (AUX data 512) — correct for GigE, and the GenCP write-mask 4-byte rule.
- `register_map.go` comment correctly documents the ABRM overlap: GigE first-URL prefix at 0x0200–0x03FF overlaps GenCP `AccessPrivilege`/`ImplementationEndianness` — hence the fallback behavior in `SyncImplementationEndianness`.

### MISMATCHES / BUGS

- **Protocol is GigE Vision GVCP, not GenCP.** Command IDs are the GVCP set — `0x0080/0x0081` READREG, `0x0082/0x0083` WRITEREG, `0x0084/0x0085` READMEM, `0x0086/0x0087` WRITEMEM, `0x0089` PENDING_ACK, `0x0004` BYE, `0x0040` PACKETRESEND, `0x0002/0x0003` DISCOVERY (`gvcp/packet.go:29-45`). GenCP defines _different_ command IDs (`0x0800` READMEM, `0x0802` WRITEMEM, `0x0805` PENDING_ACK, `0x0806/0x0808` READMEM/WRITEMEM_STACKED, `0x0C00` EVENT). AGENTS.md claims "gvcp → GenCP (control channel over GigE UDP) 1.3.1" — the layer is GVCP whose GenCP framing improved; only GenCP's ABRM/register semantics are re-used. Not necessarily a defect, but the documentation should say "GVCP 2.x with GenCP-derived register map".
- **ACK status is the GVCP single byte at offset 1** (packet type at `buf[0]`, status/error code at `buf[1]`, `gvcp/client.go:96-112`; codes 0x01..0x07, e.g. `StatusError=0x03`, `StatusNoPermission=0x07`). GenCP ACKs use a **16-bit status word at offset 0** with code/namespace/severity bits and values `0x8001..0x800F`. Statuscode values do not coincide (GVCP `0x03`=ERROR, GenCP `0x8003`=WRITE_PROTECT has different meaning). `isAddressInaccessible` (`gvcp/client.go:262-269`) treats `StatusError`/`StatusWriteProtect` as "address inaccessible" — a GenCP-style semantic grafted onto GVCP codes. Consistent enough for the wire format in use, but any attempt to talk pure GenCP with this client would misparse ACKs.
- **Heartbeat register is GigE, not GenCP.** Heartbeat sends on `gvbsHeartbeatTO` (0x0938, `gvcp/heartbeat.go:41`); GenCP §3.2 uses ABRM `Heartbeat Interval` 0x01E8 (default 3000 ms) plus the heartbeat-enable bit in `DeviceConfiguration` (0x01E0). The GenCP `AbrmHeartbeatTimeout` constant is defined (`register_map.go:44`) but never used.
- **MDRT never consulted.** `AbrmMaximumDeviceResponseTime` (0x01CC) constant exists but timeouts are hardcoded ~2 s defaults; GenCP requires response timeout = MDRT + transfer-time basis. `pendingAckTimeout` is the only spec-sourced timeout.
- **No retransmit on timeout.** GenCP §3.1.4.3 requires retrying a timed-out command 3 times before erroring; `transact` returns immediately on timeout (`gvcp/client.go:113-123`).
- **Command `reqID` starts at 1** and skips 0/0xFFFF (`gvcp/pending.go:24-27`); GenCP recommends starting at 0. (0/0xFFFF are "invalid/unknown" in some GVCP impls — harmless but non-default.)
- **`cString` uses `strings.TrimSpace`** (`gvcp/register_map.go:62-66`); GenCP says unused register bytes are zero-filled, so trailing whitespace is not a defined value — trimming could strip a legitimate trailing space from a Name/Serial string. Prefer `bytes.TrimRight(b, "\x00")`.
- TakeControl writes the GigE CCP register `gvbsCCP` (0x0A00, `gvcp/client.go:151-180`) with 8 × 400 ms retries on ACCESS_DENIED, then `SyncImplementationEndianness`. GenCP §3.2 instead says set `AccessPrivilege` (0x0204) at startup. CCP is the correct GigE register; note that `AbrmAccessPrivilege` is defined but unused, so `IsMMonitoringOrControl`-style reads aren't done.

### GAPS / UNIMPLEMENTED (GenCP 1.3 features)

- READMEM_STACKED / WRITEMEM_STACKED (GenCP 1.3, `0x0806/0x0808`) — not implemented.
- EVENT (`0x0C00`) / message-channel — not implemented; the GenCP `MessageChannelID` register constant is defined but unused.
- Heartbeat enable/disable bit in `DeviceConfiguration` (0x01E0) never read — heartbeat always-on.
- Manifest Table is parsed via GigE-style First URL / a **vendor "MTAB" layout** (see below); GenCP Table 33/34 manifest entries (8-byte count; 64-byte entries of version+schema/type+address+size+SHA1) are not used.

### UNVERIFIABLE / DEVIATION FROM GENCP

- **`ReadManifestTable` expects a non-GenCP layout** (`gvcp/client.go:283-308`): 12-byte header with `"MTAB"` magic at `[0:4]`, count at `[8:12]`, then 12-byte entries `{address, length, type}`. GenCP's Manifest Table (§ Table 33/34) has **no magic, no `MTAB`**, an 8-byte entry count, and 64-byte entries. The `"MTAB"` + 12-byte-entry scheme is vendor-proprietary (Huaray?). Also `ManifestEntryTypeXML = 0x00000001` — in GenCP the schema/file-type field is a _bitfield_ (bitwise: 0=Device XML, 1=Buffer XML — `0x00000001` would be "Buffer XML"), not a scalar "the XML type". This path only matches Huaray devices.
- Discovery (0x0002/0x0003) and PACKETRESEND (0x0040, extended with 20-byte data + 0x10 flag) are HTTP-free GVCP-only commands — not defined by GenCP. `parseDiscoveryAck` reads GigE ABRM offsets (`gvcp/discovery.go:45-96`, `discoveryAckMinSize=0x00F8`). Cannot be checked against references (no GigE Vision spec). Flagged `UNVERIFIABLE` but functionally correct for GigE.
- PACKETRESEND data/statuses, `gvspPacketIDMask=0x00ffffff` extended framing — GVSP-side, unverifiable.

---

## 2. GVSP streaming + `internal/genDC/` vs GenDC v1.1 (+ GenTL 1.6 constants)

Audited everything in `gvsp/` and `internal/genDC/`. References: `_references/GenDC/GenDC.h`, `GenICam_GenDC_v1_1.pdf.ocr/markdown.md`, `_references/GenTL/GenTL.h`.

### MATCHES

- **Container signature** `GDC_SIGNATURE = 0x43444E47` ("GNDC") per GenDC.h:47 / OCR §2.2.2; impl `Signature=0x43444E47`, checked by `IsGenDC` (`internal/genDC/genDC.go:10,348-353`). (No 8-byte "GENDC"-string signature exists.)
- **Container header field offsets exact** (OCR §2.2.2 Table 2-1): Signature@0, Version(3B)@4, Reserved@7, HeaderType@8, Flags@10, HeaderSize@12, Id@16, VariableFields@24, DataSize@32, DataOffset@40, DescriptorSize@48, ComponentCount@52, ComponentOffset[]@56; impl `parseContainerHeader` (`genDC.go:385-427`, base 56).
- **Component header field offsets exact** (OCR §2.2.4 Table 2-2): HeaderType@0, Flags@2, HeaderSize@4, GroupId@10, SourceId@12, RegionId@14, RegionOffsetX@16, RegionOffsetY@20, Timestamp@24, TypeId@32, Format@40, PartCount@46, PartOffset[]@48; impl `parseComponent` (`genDC.go:434-448`, base 48).
- **Part header common offsets exact** (OCR Table 2-3): HeaderType@0, Flags@2, HeaderSize@4, Format@8, FlowId@14, FlowOffset@16, DataSize@24, DataOffset@32; impl `parsePart` (`genDC.go:495-505`, base 40). 2D fields SizeX@40/SizeY@44 (Table 2-4, `PartHeader2DSize=56`) read at `genDC.go:508-511`.
- **Endianness: GenDC is always little-endian** (R-004/R-005); Flags bit0=TimestampPTP, bit1=ComponentInvalid, bits2-15 reserved (OCR §2.2.2). Impl hardcodes LE and never treats Flags as endianness — conformant.
- **Timestamp = 64-bit signed ns**, read as int64 LE (`genDC.go:444`) — conformant.
- **Component TypeId & Part header values** — `GDC_INTENSITY=0x01`, `GDC_RANGE=0x04`, `GDC_METADATA=0x8001` (GenDC.h:64-74) = `genDC.go:31,34,41`; `GDC_PART_HEADER` 0x4000-0x4FFF mask, `GDC_2D=0x4200`, `GDC_2D_JPEG=0x4201`, `GDC_METADATA_GENICAM_CHUNK=0x4000` (GenDC.h:57-59,88-97) = `genDC.go:50-65`.
- **Flow mapping table header layout** (OCR §3.2.1 Table 3-1): HeaderType@0=0x7000, Flags@2, HeaderSize@4, VersionMajor@8, VersionMinor@9, FlowCount@12, FlowSize[]@16; impl `ParseFlowTable` (`genDC.go:283-309`, base 16). `GDC_FLOW_TABLE_HEADER=0x7000`.
- **`PayloadTypeGenDC = 0x0000000B`** (`gvsp/payloadtype.go:19`) — matches `PAYLOAD_TYPE_GENDC` = 0x0B in GenTL.h:466-479. (The 0x0023 guess is wrong.)

### MISMATCHES / BUGS

- **[FIXED]** **`PartOffset[]` applied component-relative, but the spec says container-relative.** OCR §2.2.4 Table 2-2: PartOffset[] is _relative to the start of the Container's Header_ (same basis as ComponentOffset[], DataOffset). `parseComponent` calls `parsePart(buf[pOff:])` (`genDC.go:469-473`) but `buf` is already the component-sliced buffer (`genDC.go:372`), so conformant data is read from `componentOffset+partOffset` instead of `partOffset` — garbage parts. **The test fixture masks it**: `genDC_test.go:32` itself writes a component-relative offset (56), i.e. the fixture is non-conformant.
- **[FIXED]** **`ParseChunkPayload` is not the GenDC/GenICam chunk format.** `gvsp/chunk_data.go:74-114` expects a big-endian 16-byte header (u64 PayloadSize + u32 ChunkCount + u32 Reserved) + 16-byte entries {ChunkID, Offset, Size, Version u16, Reserved u16}. GenDC metadata chunks are little-endian **trailing tags** (4-byte ChunkID + 4-byte length appended after each chunk, 0xFF.. sentinel, with empty-alignment chunks) per OCR §2.2.8.1 (GenDC metadata part layout for GenICam chunk) — inverted endianness AND wrong layout.
- **[FIXED]** **Payload-type helper constants contradict GenTL and this package.** `GenDCPayloadType()=0x80000008`, `MultiPartPayloadType()=0x80000007`, `ChunkPayloadType()=0x80000009` (`gvsp/genDC_payload.go:97`, `multi_part.go:98`, `chunk_data.go:129`, all commented "per GenTL"). GenTL.h has GENDC=0x0B, MULTI_PART=0x0A, CHUNK_DATA=0x04, CHUNK_ONLY=0x08. The vendor-alias low-byte heuristic is self-defeating: `payloadTypeAliasGenDC=0x80000008` has low byte 8 = CHUNK_ONLY, **not** GENDC(11); a genuine vendor-encoded GenDC would be `0x8000000B`, which `IsPayloadTypeGenDC` (`payloadtype.go:32`) rejects → `receiver.go:150-153` zeroes the payload type.
- **Multi-part part-type constants don't match GenTL.** `MultiPartPartTypeImage=0, Chunk=1, ExtendedChunk=2` (`multi_part.go:33-35`, "GenTL v1.5") — GenTL.h:508-521 `PART_DATATYPE_*` are 0=UNKNOWN, 1=2D_IMAGE, 2=2D_PLANE_BIPLANAR…; GenTL has no chunk/extended-chunk part datatype. `[FIXED]` — comment corrected (`multi_part.go:31-33`): these are GVSP part-type codes (0=image,1=chunk,2=extended chunk), not GenTL `PART_DATATYPE`.
- **`ComponentDepth=5` labeled "GenDC Range".** `gvsp/payload.go:33` — GenDC.h:67 `GDC_RANGE=0x04`(4); the repo's own `internal/genDC/genDC.go:34` `ComponentRange=0x04`. Wire value 5 is a Huaray BSCF vendor value, wrongly attributed to GenDC. `[FIXED]` — comment corrected (`payload.go:31-35`): 5 is the Huaray BSCF wire value (`Frame.h` iota), GenDC Range is `0x04`.
- **Flow-table search mechanism is non-spec.** `FlowTableFromContainer` (`genDC.go:311-322`) scans inside container data for 4-byte-aligned LE 0x7000. GenDC §3.2.1 says the flow table is delivered _in a Transport Layer-specific way_ (device XML / bootstrap registers), **not** inside the container's descriptor+data (R-006) — false-positive-prone.

### GAPS / UNIMPLEMENTED

- Component/container `Invalid` flags parsed (`genDC.go:397,436,497`) but never honored — `ParseGenDcPayload` (`genDC_payload.go:37-69`) uses components regardless; "ComponentInvalid" must make the component unusable.
- `VariableFields` parsed (`genDC.go:400`) but ignored; no preliminary/final descriptor handling (CR-012/CR-013) — a preliminary descriptor's DataSize (only an upper limit) would be treated as final. GenTL 5.7.3 requires exactly one final descriptor per delivered GENDC buffer.
- FlowId/FlowOffset parsed (`genDC.go:501-502`) but unused; part data resolved only against one linear buffer (`genDC_payload.go:48-52`), so parts on Flows≠0 (R-008/9/10) can't be located.
- Version parsed (`genDC.go:391-394`) but never validated — non-1.x majors accepted; flow-table version not checked (must be 1.0).
- DataSize/DataOffset/DescriptorSize/HeaderSize never used for bounds validation → truncated/malformed descriptors degrade silently (`continue` at `genDC.go:374,475`).
- `Id` never cross-checked against transport frame/block ID; planar/multi-part components: only first 2D part surfaced (`genDC_payload.go:59-67`).

### UNVERIFIABLE (no GVSP/GigE Vision spec in `_references`)

- Leader payload-type @ `data[0:]` as 32-bit (`receiver.go:149`) vs GVSP 16-bit field + variant offset (`payloadtype.go:6-7` is self-authored).
- Content-type `uint8(infos>>24)&0x7f`, packetID `infos&0x00ffffff`, extended-mode handling (`receiver.go:107-126`).
- Content IDs 0x01/0x02/0x03 (leader/trailer/payload), frame matching, gap detect + RESEND ranges (`receiver.go:142-160`, `resend.go`), OOO ring/reset (`frame.go`), SCPS/MTU (`socket.go`).
- GVSP multi-part 32-byte part-header wire format (`multi_part.go:44-68`) — and GenTL `PART_DATATYPE` doesn't match these anyway.
- `ChunkIDTimestamp=0x00000001…` (`chunk_data.go:34-66`) are GVSP/GEV chunk IDs, not GenICam/GenDC chunk semantics; the layout they'd have to handle is the trailing-tag GenICam one (see MISMATCHES).
- `gvsp/doc.go`'s GigE Vision 2.0/2.1 compliance claims untestable without the reference.

---

## 3. GenApi (`genapi/`) vs GenICam Standard 2.1.1

Files: `node.go`, `types.go`, `nodemap.go`, `evaluator.go`, `port.go`, `camera_description.go`. Reference: `_references/GenApi/GenICam_Standard_v2_1_1.pdf.ocr/markdown.md` (2159 lines).

### MATCHES

- Address computation = sum of `Address`+`pAddress` (§2.8.3) — `types.go:75-83`, `port.go:27-34`.
- `Bit` ≡ `LSB==MSB` (§2.8.5) — `types.go:100-113`.
- MaskedIntReg write mask w/ lo/hi swap (§2.8.5 BE/LE bit numbering) — `types.go:201-211`, `port.go:86-91`.
- `pIsImplemented/pIsAvailable/pIsLocked`: zero ⇒ not accessible (§2.5/§2.8.1) — `nodemap.go:416-453`, `evalBoolish` `nodemap.go:180-208`.
- Formula ops `( ) + - * / % & | ^ ~ < > = <> <= >= && || << >> ?:` + unary minus (§2.8.13) — `evaluator.go:72-453`.
- Escaped/CDATA formulas handled by `encoding/xml`.
- `pVariable` Name→feature mapping (§2.8.13).
- Boolean/Enum resolve through `pValue` (§2.8.7/§2.8.10) — READ path `nodemap.go:61-77`, `port.go:274-281`.
- `local:name;addr;length` scheme + deflate/zip (PK/stdlib) — `camera_description.go:43-54,70-97,129-135`.
- Big-endian default (§3.1.1/3.1.2) — `port.go:199-209`.

### MISMATCHES / BUGS

- **MaskedIntReg READ ignores the mask** — full register value returned. `port.go:272-273` routes IntReg/MaskedIntReg to `readIntReg`, never applying `maskFromBits`; §2.8.5 requires extracting bits [LSB,MSB] + sign extension. Write masks only for `length==4` (`port.go:81-92`, mask dropped for length 2 at `port.go:93-114`). `[FIXED]`
- **`<Sign>` never parsed/applied** — no sign extension anywhere (§2.8.5 "The Sign element is used to manage sign bit extension"; default Unsigned §3.2). `[FIXED]`
- **Per-register `<Endianess>` ignored.** Only device-level `DeviceByteOrder()` used (`port.go:61,123,140,199-209`). Worse: length-4 int reads go through `ReadReg` hardwired `binary.BigEndian.Uint32` (`gvcp/client.go:173`) — for a LE device this returns network order, violating §3.1.1 (must flip each 4-byte word before GenApi). `[FIXED]` per-register `<Endianess>` now honored.
- **Boolean hardcodes true=1/false=0**, rejects floating boolean (`nodemap.go:54-71`); `OnValue/OffValue` not parsed (`types.go:74-140`) — §2.8.7 default is 1/0 but values are configurable. `[FIXED]` — `<OnValue>/<OffValue>` parsed (`types.go`, `node.go`) and honored by `SetBoolean`/`ReadBoolean` (`nodemap.go:58-100`); true iff register == OnValue.
- **Converter direction broken both ways.** `FormulaFrom` not parsed (`types.go:96-99` only Formula/FormulaTo); parseNodeXML substitutes `FormulaTo` for the read formula (`types.go:253-256`); write path ignores FormulaTo, writes raw value to the `pValue` register (`port.go:289-306`). §2.8.13: FormulaFrom=register→user, FormulaTo=user→register. `[FIXED]`
- **`**` power operator missing** (listed in §2.8.13; `doc.go:41` claims support) — `parseMul` only `* / %` (`evaluator.go:334-369`). `[FIXED §7 #9]` — `**` via `parsePow`, right-assoc, binds tighter than unary (`evaluator.go`).
- **IntSwissKnife function set deviates.** §2.8.13: Int functions `SGN, NEG` only; float adds `ABS, EXP, LN, LG, SQRT, TRUNC, FLOOR, CEIL, ROUND, ATAN, ASIN, ACOS, SIN, COS, TAN, E, PI`. Impl: `ABS/FLOOR/CEIL/SQRT` on both (`evaluator.go:11-26`), **no SGN**, `FLOOR/CEIL` are identity no-ops (`evaluator.go:18-19`). `[FIXED §7 #9]` — full §2.8.13 int+float table, `SGN/NEG`, variadic `ROUND(x,prec)`, consts `E/PI`.
- **`==` rejected** — the spec's own example is `TRIGGER==1` (§2.5, p.14); `parseEquality` accepts only `<>`, `!=`, `=` (`evaluator.go:202-210`). (Spec internally mixes `=` and `==`.) `[FIXED §7 #9]` — both accepted, `==` checked first.
- **`<>` not-equal broken at runtime** — real Huaray formulas (`TriggerDelayAvailExpr` = `VAR_TRIGGERTYPE <> 4`, `TCPPortAvailExpr`, `LaserBrightnessLockedExpr`, etc.) failed `bad token at "> 0 …"`: `parseRel` greedily consumed the `<` as less-than and `parsePrimary` hit the leftover `>`. `[FIXED §7 #10]` — `<>` treated as a single spacing-tolerant token at the equality tier (peek-guarded in `parseRel`), exact camera formulas locked by `TestEvalFormulaNotEqual`.
- **Converter reserved `FROM`/`TO` never bound** — camera `IntConverter`s (`GevIEEE1588ValConv`, `GevGVCPPendingAckValConv`, `GevGVCPHeartbeatDisableValConv`, `GevGVCPExtendedStatusCodesValConv`) use the GenApi §2.8.13 variables: `FormulaFrom=(TO & 0x00080000) >> 19`, `FormulaTo=(VAR_CFG & 0xFFF7FFFF) | (FROM << 19)`. Read failed `unknown var "TO"`; the write heuristic also clobbered bitfields (bound the pValue var to the user value instead of preserving other bits). `[FIXED §7 #10]` — read binds `TO`=register value (`port.go`); write binds `FROM`=user value with pVariables read as current values (true RMW), legacy single-var converters still honored via `formulaUses`; locked by `TestConverterReservedFromTo`.
- **`resolveAddr` refuses address 0** (`port.go:39-42`) — legitimate GVCP/ABRM addresses; also truncates 64→32-bit (`port.go:43`), dropping `0xFFFF…` ConfRom/SmartFeature space (§2.8.14/15). `[FIXED]` — address 0 (e.g. `GevVersionReg`) is addressable; sums > `math.MaxUint32` now return an error instead of silently truncating. True 64-bit ConfRom/SmartFeature space (§2.8.14/15) remains a GAP below.
- **Floating nodes never stored** — `<Integer><Value>` fails on Set/Read because node has no address (`nodemap.go:161-167`); same for floating String/Float/Converter (`nodemap.go:286-318`). `[FIXED]` — constant `Float`/`String` nodes (a `<Value>` with no address) read their constant in `readFloatReg`/`readStringReg` (`port.go`); verified on the offline Huaray replay (`ReadFloat`/`ReadString` at 433 fewer errors than recorded).
- **Runtime access mode never derived/enforced.** §2.5 defines NI/NA/WO/RO/RW. Impl: `GetAccess()` returns raw XML string (`node.go:113-115`); writes don't check access/locked/available (`nodemap.go:80-86,133-167`); RO write or WO read unguarded; `ImposedAccessMode` (§2.8.1) unparsed. `[FIXED]` — `effectiveAccess` (NI→NA→modes∩`ImposedAccessMode`→locked RW→RO / locked WO→NA) enforced at Set*/Read*/CurrentEnum; target register access checked on writes; `ImposedAccessMode` parsed (`types.go`); locked by `genapi/gaps_test.go:TestLockedDowngrade`/`TestImposedAccessMode`.
- **Float-domain SwissKnife evaluated as int or errored.** All math went through the `int64` domain (`evaluator.go`); float `SwissKnife`/`pVariable`→`Float` targets (`GainRawMaxExpr*`, `ExposureTime*Expr`, `AcquisitionFrameRate*Expr`, `CC*Expr`) hit `cannot evaluate integer (kind=Float)` on the offline Huaray replay (9 errors). `[FIXED]` — float-domain evaluator `evalFormulaFloat` (`floatParser` + §2.8.13 float functions `ATAN2`/`MIN`/`MAX`/`POW`/`LOG` and `E`/`PI`), float suffix resolver (`formulaContextFloat`, `readFormulaVariableFloat`), `ReadFloat` SwissKnife case, integer-read truncation (negative → error) in `resolveIntegerReference`; locked by `evaluator_test.go:TestEvalFormulaFloat`/`TestEvalFormulaFloatSuffix` + `gaps_test.go:TestSwissKnifeFloatDomain`, replay errors 333→324 (all remaining are guard-correct).

### GAPS / UNIMPLEMENTED

### GAPS / UNIMPLEMENTED

- **Node types** (`§2.8`): Category (parsed→skipped `types.go:303-306`, no Root/pFeature tree), Register, Port (§2.8.16), Node, StructReg→MaskedInt expansion (§2.8.6), ConfRom/TextDesc/IntKey (§2.8.14), DcamLock/SmartFeature (§2.8.15). No selector/array addressing (`pSelected`, `pIndex`, `pIndex Offset`) or embedded SwissKnife address contributions (§2.8.4).
- **Node elements**: `pLength`, `pPort`, `pValueCopy` (§2.8.5), `ValueIndexed/pValueIndexed/ValueDefault/pValueDefault` (§2.8.5/9), `ValidValueSet`, `Unit`, `Representation`, `DisplayNotation`, `DisplayPrecision`, `PollingTime`, `Visibility`, `NameSpace`, `MergePriority`/merge (§2.8.1), `pError`, `pAlias/pCastAlias`, `IsDeprecated`, `EventID` invalidation, `pBlockPolling`.
- **Enumeration** (§2.8.10): `NumericValue`, `IsSelfClearing`, `Symbolic`, EnumEntry `pIsImplemented/pIsAvailable`, polling semantics.
- **SwissKnife** (§2.8.13): `Constant`, `Expression` sub-exprs, `.Entry.Name` qualified suffix; variable suffixes `.Min/.Max/.Inc/.Value/.Entry` resolved in float domain too (`evaluator.go` `floatParser` + `formulaContextFloat`), float-domain formulas (e.g. `GainRawMaxExpr` = `VAR_GAINMAX`, a Float feature) now evaluate via `ReadFloat` (`ReadFloat` SwissKnife case) and truncate cleanly for the integer path (`resolveIntegerReference`); verified on the offline Huaray replay (9 float-SwissKnife errors cleared).
- **Command** (§2.8.8): `CommandValue/pCommandValue` unparsed (`nodemap.go:141-147` writes 1 or `<Value>` text); **no `IsDone()`**, no self-clearing polling.
- **Port** (§2.8.16): chunk ports, `CHUNK_BASE_ADDRESS_REGISTER`=INT64_MAX / `CHUNK_LENGTH_REGISTER`, `CacheChunkData`, `SwapEndianess`, max-chunk emulation missing.
- **Caching/invalidation** (§2.6): `PInvalidator` parsed/exposed (`node.go:93`, `nodemap.go:457-463`) but nothing invalidates; acceptable only on "caching optional" (§2.8.3).
- **Schema versioning** (§2.7.1): `SchemaMajorVersion` never read → future-schema rejection impossible; no `DeviceEndianessMechanism` handoff (§3.1.3); name-uniqueness unenforced (`nodemap.go:38` overwrites duplicates).

### UNVERIFIABLE

- `=` vs `==` equality (§2.8.13 table lists `=`; §2.5 example uses `TRIGGER==1`) — OCR can't disambiguate; impl now accepts both (§7 #9).
- `local:` URI + zip-of-XML: not in GenICam 2.1.1 (lives in GenCP/GEV specs) — `Local:name;hexaddr;hexlen` + `.zip`/PK handling can't be validated here.
- `FormulaFrom` spelling only OCR-garbled mention (`markdown.md:1275`).
- `CacheAs`, `IsCachable`, `IsElseReg`, `Task`, `ImposedMin/Max` don't exist as XML attributes in this spec.

---

## 4. `internal/color/` vs PFNC 2.4

References: `_references/SFNC/PFNC.h`, `GenICamPixelFormatValues.md`, `GenICam_PFNC_2_4.pdf.ocr/markdown.md`.

### MISMATCHES / BUGS (RED FLAGS — visual output wrong)

- **RGGB debayer is R/B-swapped** (8-bit `bayer.go:136-141,199-223`, 16-bit `packed.go:214-219,274-293`). PFNC §3.1.7 RGGB tile = `[R G; G B]` → R at (even,even), B at (odd,odd). Code samples B from (even,even) and takes R from the current pixel with no R interpolation. Simulated RGGB (R=100,G=10,G=20,B=50): p(0,0)→(100,15,**100**) [want (100,15,50)]; p(1,1)→(50,15,25) [want (100,15,50)].
- **GBRG debayer is R/B-swapped** (8-bit `bayer.go:150-155,306-354`, 16-bit `packed.go:228-233,366-406`). PFNC GBRG = `[G B; R G]` → R at (even,odd), B at (odd,even). Code samples R from (odd,odd) and B from (even,odd): p(0,0)→(5,75,50) [want (100,10,50)].
- **BGGR and GRBG debayers are correct** (site parities verified — `bayer.go:253-276,384-432`, `packed.go:320-339,433-473`).
- **Unpacked Mono10/12/14 preview near-black.** `DecodeHighDepth` → `mono16Preview` (`packed.go:552-556`) takes the high byte of an LSB-aligned 16-bit word _without_ `shiftTo16`; packed mono shifts (`packed.go:577`), Bayer paths shift (`packed.go:601,610`). Data lives in the low 2-6 bits.
- **Misleading test**: `TestDecodeHighDepth_MonoUnpacked` uses out-of-spec 0x1122 for Mono10 (max 0x3FF) and its error text says `want 0x22` while asserting 0x11 (`color_test.go:103`). No test asserts color _fidelity_ — all sample tests only check encode-success/non-black.

### MATCHES

- **All pixel-format constants verified numerically** against PFNC.h: Mono8=0x01080001, Mono10=0x01100003, Mono10p=0x010A0046, Mono12=0x01100005, Mono12p=0x010C0047, Mono14=0x01100025, Mono14p=0x010E0104, Mono16=0x01100007, Mono32=0x01200111, BayerBG4p=0x01040110, Bayer*14/14p/16=0x01100109..0C/0x010E0105..108/0x0110002F (`color.go:15-59`); RGB8=0x02180014, BGR8=0x02180015 (`color.go`); YUV422_8=0x02100032, YUV422_8_UYVY=0x0210001F (`color.go:130-152`) match LMN422 byte order.
- **Unpack10P/12P/14P bit layouts verified letter-for-letter** against §6.3.1 (lsb-packed, no padding) — all 12 sample-assembly expressions correct (`packed.go:13-83`); `TestUnpack14P` encodes via official layout.
- **`bayer16PFNC`/`BayerPFNCMap` pattern IDs correct** — only the RGGB/GBRG demosaic math is broken.
- Length guards correct per format: `(w*h*bits+7)/8` packed, `w*h*2` unpacked, `w*h*3` RGB (`color.go:86-171`, `packed.go:553-610`).
- RGB/BGR/Mono routing and channel order correct.

### GAPS

- `Debayer10P/12P/14P` (`packed.go:86-119`) are DEAD CODE with no `shiftTo16` — latent truncation if wired up.
- Legacy GEV-2.0 codes `Mono10Packed`=0x010C0004, `Mono12Packed`=0x010C0006 (GenICamPixelFormatValues.md) — no handling.
- Defined in `genDC.go` but no decode path: RGB10/12/16, BGR10/12/14/16, RGB10p/12p, RGB8a32/BGR8a32, YCbCr422_8/10/12(+_CbYCrY), Mono32, Bayer*4p.
- Green channel at green sites averages adjacent R/B (`getGreen*`) rather than using site's own value — quality artifact even on correct BGGR/GRBG paths.

### UNVERIFIABLE

- YUV→RGB uses studio-range BT.601 (`color.go:188-196`); consistent with PFNC §3.2 legacy YUV notes, but PFNC prescribes no exact transform.
- Unpackers assume LSB-first byte order; PFNC defines no endianness (§1.4) — correct for GEV targets only.

---

## 5. `gentl/` vs GenTL 1.6

`gentl/types.go` (405 lines), `gentl/doc.go`. `gentl.` has no consumers beyond doc comments — pure constant library.

### MATCHES

**Every exported constant verified numerically against GenTL.h — zero mismatches.**

| Family                                                                      | Verdict                                    |
| --------------------------------------------------------------------------- | ------------------------------------------ |
| GC_ERROR (25, err 0..-1023, CustomID -10000)                                | match (`types.go:8-34` vs GenTL.h:156-184) |
| GenTL 1/6/0 version                                                         | match                                      |
| TLType names (11) / Module names (6)                                        | match                                      |
| INFO_DATATYPE 0-14                                                          | match                                      |
| TL_CHAR_ENCODING, TL_INFO_CMD 0-10                                          | match                                      |
| INTERFACE_INFO, DEVICE_ACCESS_FLAGS/STATUS, DEVICE_INFO                     | match                                      |
| ACQ_STOP/START, ACQ_QUEUE 0-4                                               | match                                      |
| STREAM_INFO 0-15 (FLOW_TABLE=14, GENDC_PREFETCH=15)                         | match                                      |
| BUFFER_INFO 0-31 (DATA_SIZE=27, IS_COMPOSITE=31), BUFFER_PART_INFO 0-13     | match                                      |
| PAYLOADTYPE_INFO 0-11 (MULTI_PART=10, GENDC=11)                             | match                                      |
| PIXELFORMAT_NAMESPACE 0-4, PIXELENDIANNESS 0-2, PARTDATATYPE 0-11           | match                                      |
| PORT_INFO, URL_SCHEME, URL_INFO, EVENT_TYPE 0-5, EVENT_INFO/EVENT_DATA_INFO | match                                      |
| FLOW_INFO, SEGMENT_INFO 0-4 (v1.6)                                          | match                                      |
| InvalidHandle=0, Infinite=0xFFFFFFFFFFFFFFFF, handle=uintptr                | match                                      |

### MISMATCHES

None.

### GAPS

- **All 23 `*_CUSTOM_ID = 1000` sentinels missing** — every GenTL.h enum except PIXELENDIANNESS ends with one (spec md:4779,4821,4846,4867,4890,4923,4947,4962,4974,5096,5129,5152,5211,5273,5325,5338,5371,5412,5446,5465,5492,5517,5551). `ErrCustomID=-10000` is present; the positive family boundary isn't. `[FIXED §7 #10]`
- **`INFO_DATATYPE_CUSTOM_ID=1000` missing** (GenTL.h:259). `[FIXED §7 #10]`
- **No typed alias groups** — C typedefs are `int32_t`; impl exposes raw untyped `int` constants, so nothing encodes the base type or enum grouping (values are correct).
- Out of declared scope ("constants only"): `EVENT_NEW_BUFFER_DATA`, `PORT_REGISTER_STACK_ENTRY`, `SINGLE_CHUNK_DATA`, `DS_BUFFER_INFO_STACKED`, `DS_BUFFER_PART_INFO_STACKED`, and GC_API prototypes — noted, not defects.

### UNVERIFIABLE / NOT APPLICABLE

- `GC_INFO_*`, `GC_STREAM_INFO_*`, `GC_MODULE_*`, `GC_PARAM_*`, `GC_EVENT_*`, stream callback types **do not exist in GenTL 1.6** (spec md:960 explicitly rejects callbacks in favor of OS events). These are 1.7+/2.0 concepts — correct to omit.

---

## 6. Cross-cutting root package observations

Loose, integration-level notes (not deep per-file audits):

- `device.go:32 Open()` → `connectCamera()` → `gvcp.TakeControl`; `TakeControl/LeaveControl` also used by `examples/smoke-test/main.go` and `examples/dump-xml/main.go`.
- Component defaulting in `Open`: defaults to `ComponentColor`; `options.go` carries `WithComponent`/`GrabComponent`.
- `camera.go:251 BooleanFeature`, `device.go:80 Features` — GenApi Feature interface in root package.
- Neither GVCP access-privilege (`AbrmAccessPrivilege`) nor GenCP heartbeat-enable is consulted at the root layer.

## 7. Quick fix list (highest ROI, by severity)

1. **[DONE]** **R/B swap in RGGB + GBRG debayers** (`internal/color/bayer.go`, `packed.go`) — visual correctness regression; add a color-fidelity test. (Parities + green-site short-circuits + OOB-neighbor-excluded interpolation; `TestBayerFidelityUniform`.)
2. **[DONE]** **Monochrome unpacked 10/12/14 near-black** (`internal/color/packed.go:552-556`) — add `shiftTo16`.
3. **[DONE]** **`PartOffset` base (component- vs container-relative)** in `internal/genDC/genDC.go:469-473` + fix non-conformant fixture `genDC_test.go:32` — decide against GenICam reference implementation. (Container-absolute; both fixtures use container-relative offset now.)
4. **[DONE]** **Bogus payload-type constants** `0x80000008/07/09` (`gvsp/payloadtype.go:23`, `genDC_payload.go:97`, `multi_part.go:98`, `chunk_data.go:129`) — real values per GenTL.h (GENDC=0x0B, MULTI_PART=0x0A) or document vendor semantics. (Helpers now return real ids; vendor aliases = |0x80000000 of the real ids.)
5. **[DONE]** **Chunk parser endianness/layout** (`gvsp/chunk_data.go:74-114`) — GenDC trailing-tag little-endian format, not big-endian offset table. (Backward-walk parser, 0xFFFFFFFF padding tags, copied chunk data, truncated-test added.)
6. **[DONE]** **Manifest `MTAB` vendor layout** (`gvcp/client.go:283-308`) — divergent from GenCP Table 33/34; refactor to spec or document as Huaray quirk. (Dual-layout: GenCP tables parsed per Table 33/34; `MTAB` magic kept as documented Huaray quirk; `ManifestEntry` extended with FileVersion/Schema/Address u64/FileSize/SHA1/GenCP.)
7. **[DONE]** **MaskedIntReg read ignores mask + `<Sign>` + per-register `<Endianess>`** (`genapi/port.go`, `types.go`) — GenApi spec §2.8.5. (Mask applied on read `(raw & mask) >> shift`; `<Sign>Signed</Sign>` sign-extends via `intBitWidth`; `<Endianess>` honored on read+write incl. LE 4-byte flip; mask RMW extended to length<=4; tests: masked read signed/unsigned, LE round-trip.)
8. **[DONE]** **Converter FormulaFrom/FormulaTo wiring** (`genapi/nodemap.go`, `port.go`) — both directions broken. (`FormulaFrom` parsed; read uses FormulaFrom, write applies FormulaTo via `convertUserToRegister`, legacy `Formula` fallback; test `TestConverterFormulaDirections`.)
9. **[DONE]** **`==` in SwissKnife formulae + missing float math + `**`** (`genapi/evaluator.go`) — GenICam §2.8.13. (`==` accepted as `=`; `**` power right-assoc via `parsePow`; full §2.8.13 function set SGN/NEG/ABS/FLOOR/CEIL/TRUNC/ROUND(x,prec)/SQRT/EXP/LN/LG/SIN/COS/TAN/ATAN/ASIN/ACOS + E/PI constants, float-domain math truncated to the int domain; variadic args for ROUND; tests added.)
10. **[DONE]** **`_CUSTOM_ID=1000` sentinels** in `gentl/types.go` (23 enums + INFO_DATATYPE_CUSTOM_ID). (Added per GenTL.h: `InfoDataTypeCustomID`, `TLInfoCustomID`, `InterfaceInfoCustomID`, `DeviceAccessCustomID`, `DeviceAccessStatusCustomID`, `DeviceInfoCustomID`, `AcqStopFlagsCustomID`, `AcqStartFlagsCustomID`, `AcqQueueCustomID`, `StreamInfoCustomID`, `BufferInfoCustomID`, `BufferPartInfoCustomID`, `PayloadTypeCustomID`, `PixelFormatNamespaceCustomID`, `PartDataTypeCustomID`, `PortInfoCustomID`, `URLSchemeCustomID`, `URLInfoCustomID`, `EventCustomID`, `EventInfoCustomID`, `EventDataCustomID`, `FlowInfoCustomID`, `SegmentInfoCustomID`.)
