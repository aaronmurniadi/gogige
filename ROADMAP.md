# ROADMAP

Progress tracker against `AGENTS.md` and GiGE Vision specs under `_references/`. Single source of truth for outstanding work.

Legend: `[x]` done · `[ ]` not started · `[~]` partial

## Outstanding work (single source of truth)

Hand-complete rows still open against the reference specs; this is the canonical pending list.

| Area              | Item                                                                      | Status |
| ----------------- | ------------------------------------------------------------------------- | ------ |
| GenApi 2.1.1      | `Category` / `StructReg` as first-class node types (parsed/skipped today) | [x]    |
| GenApi 2.1.1      | SwissKnife `**` exponent (rest of § formula grammar done)                 | [x]    |
| GenApi + SFNC 2.7 | Formal `Gev*` / `Device*` streaming-feature coverage                      | [x]    |
| CLI               | `cmd/` CLIs beyond discover/stream                                        | [~]    |

---

## Spec index (`_references/`)

Target versions from architecture rules, mapped to local artifacts:

| Standard         | Target    | Local reference                                                                       | Used by                                                    |
| ---------------- | --------- | ------------------------------------------------------------------------------------- | ---------------------------------------------------------- |
| GenCP            | 1.3.1     | `_references/GenCP/GenICam_GenCP_1.3.1.pdf`                                           | `gvcp/` (control channel is GenCP over GigE UDP)           |
| GenICam / GenApi | 2.1.1     | `_references/GenApi/GenICam_Standard_v2_1_1.pdf`                                      | `genapi/`                                                  |
| GenDC            | 1.1       | `_references/GenDC/GenICam_GenDC_v1_1.pdf`, `GenDC/GenDC.h`                           | `gvsp/payload` GenDC path                                  |
| SFNC             | 2.7       | `_references/SFNC/GenICam_SFNC_v2_7.pdf`                                              | Feature names (`AcquisitionStart`, `GevSCPSPacketSize`, …) |
| PFNC             | 2.4       | `_references/SFNC/PFNC.h`, `GenICam_PFNC_2_4.pdf`, `GenICamPixelFormatValues.pdf`     | `internal/color`, pixel format IDs                         |
| CLProtocol       | 1.2       | `_references/GenApi/GenICam_CLProtocol_Standard_v1.2.pdf`                             | Out of scope (Camera Link)                                 |
| GigE Vision      | 2.0 / 2.1 | `_references/GigE_Vision_for_Realtime_MV_11052010.pdf`, `GigE_Features_Reference.pdf` | `gvcp/`, `gvsp/`                                           |

Authoritative machine-readable headers for implementers: `GenDC/GenDC.h`, `SFNC/PFNC.h`.

GenTL (`.cti` interop) is intentionally out of scope: gogige speaks GVCP/GVSP/GenApi directly, keeping the pure-Go path portable to any GigE Vision + GenICam camera without vendor producer libraries.

---

## Package layout

### Target tree

| Path                           | Status | Notes                                                                                                                                                   |
| ------------------------------ | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `gvcp/client.go`               | [x]    | Control channel, register R/W                                                                                                                           |
| `gvcp/discovery.go`            | [x]    | Per-iface bind + directed/limited broadcast; DISCOVERY_ACK ABRM parse                                                                                   |
| `gvcp/heartbeat.go`            | [x]    | `PulseHeartbeat` + background tick at `HeartbeatTimeout/2`                                                                                              |
| `gvcp/packet.go`               | [x]    | Header encode + cmd constants                                                                                                                           |
| `gvcp/register_map.go`         | [x]    | GenCP ABRM 0x0000–0x0250 + GigE Vision ABRM/SBRM                                                                                                        |
| `gvsp/frame.go`                | [x]    | Frame + reassembly; `OOOPacketRing` zero-alloc OOO store (rare overflow map)                                                                            |
| `gvsp/receiver.go`             | [x]    | UDP receive path; OOO via ring, contiguous via pool                                                                                                     |
| `gvsp/payload.go`              | [x]    | BSCF + `Component` / `GrabAll`; payload-type dispatch (Image/Multi-Part/GenDC/Chunk)                                                                    |
| `gvsp/buffer_pool.go`          | [x]    | Pre-allocated frame buffers + `Frame.Release()`                                                                                                         |
| `gvsp/resend.go`               | [x]    | Missing-packet tracking + `RESEND_CMD` via `Stream.SetResender`                                                                                         |
| `genapi/camera_description.go` | [x]    | FirstURL fetch + zip/deflate XML                                                                                                                        |
| `genapi/evaluator.go`          | [x]    | SwissKnife formula evaluator                                                                                                                            |
| `genapi/nodemap.go`            | [x]    | Parse + feature get/set (orchestration layer)                                                                                                           |
| `genapi/node.go`               | [x]    | Core Node interface + gcNode struct + attributes                                                                                                        |
| `genapi/types.go`              | [x]    | Node parsing: nodeFields, parseNodeXML, parseNodeMapXML stream                                                                                          |
| `genapi/port.go`               | [x]    | Port binding layer: portAdapter → gvcp.Port I/O + byte order                                                                                            |
| `cmd/gogige-discover/`         | [x]    | CLI discovery utility                                                                                                                                   |
| `cmd/gogige-stream/`           | [x]    | CLI N-frame JPEG + BSCF measurements                                                                                                                    |
| `camera.go`                    | [x]    | High-level `Camera`: get/set features (short names + getters), one-shot grabs (`GrabSample`/`GrabAllSamples`/`GrabComponents`/`GrabJPEG`), `Features()` |
| `discovery.go`                 | [x]    | Root `Discover` → `gvcp.Discover`                                                                                                                       |
| `stream.go`                    | [x]    | `Session` / `Grab` / `GrabAll`; `StartStream` + `Frames()`                                                                                              |
| `options.go`                   | [x]    | `WithLogger` / `WithTimeout` / `WithComponent` / `GrabComponent`                                                                                        |
| `grab/grab.go`                 | [x]    | One-shot `GrabJPEG` convenience + `FromCamera` for an open `Camera`                                                                                     |
| `live/live.go`                 | [x]    | Continuous preview loop (`NewLive` / `WithSink` / `Start` / `Stop`)                                                                                     |
| `log.go`                       | [x]    | `Logger` interface + `NopLogger` + `Slog` adapter                                                                                                       |
| `interfaces.go`                | [x]    | Core interfaces: `Device`, `Features` (get+set), `Grabber`, `FrameSink`, `Throttler`, `JPEGFunc`                                                        |
| `device.go`                    | [x]    | `device` struct, `Open`, `OpenDevice`, `connectCamera`                                                                                                  |
| `alias.go`                     | [x]    | Type aliases + re-exports for ergonomic `gogige.Sample` etc.                                                                                            |
| `doc.go`                       | [x]    | Package documentation                                                                                                                                   |

### Layout debt (outside target tree)

| Path | Status | Notes                                                                                                                                                        |
| ---- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| None | [x]    | `grab.go` → `grab/`, `live.go` → `live/`, `framestream.go` merged into `stream.go`, `logger*.go` merged into `log.go`, `version.go` merged into `options.go` |

Cleared: `vision/bscf` → `gvsp/payload.go`; `vision/color` → `internal/color` (`vision/` removed).

---

## Architecture / Phase 4 high-level API

```go
cam, err := gogige.OpenDevice(ctx, "192.168.1.100")
stream, err := cam.StartStream(ctx)
for frame := range stream.Frames() { frame.Release() }
```

| Item                               | Status | Notes                                                                                                                                                                                   |
| ---------------------------------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Package name `gogige` (was `gige`) | [x]    | Renamed; name now matches import path                                                                                                                                                   |
| `OpenDevice`                       | [x]    | Returns `*Camera`; `Open`→`Device` kept for legacy consumers                                                                                                                            |
| `SetInteger` / `SetEnum` on camera | [x]    | `camera.go` primary setter/getter set (`SetInteger`/`SetEnum`/`SetBoolean`/`SetFloat`/`SetString` + `Integer`/`Enum`/`Boolean`/`Float`/`String`); old `Set*Feature` retained as aliases |
| `StartStream` + `<-chan *Frame`    | [x]    | `framestream.go`; pooled frames, `Stop`/`Pause`/`Resume`                                                                                                                                |
| `frame.Release()` buffer return    | [x]    | `gvsp.BufferPool` + `Frame.Release`                                                                                                                                                     |

---

## Protocol rigor (by phase)

### Phase 1 — GVCP / GenCP 1.3.1

Refs: `_references/GenCP/…`, architecture + GVCP section of `AGENTS.md`. GigE Vision packet framing still needs the AIA PDF.

| Item                                                       | Status | Spec cue                                                              |
| ---------------------------------------------------------- | ------ | --------------------------------------------------------------------- |
| READ/WRITE REG/MEM + PENDING_ACK                           | [x]    | GenCP memory/register cmds over UDP/3956                              |
| CCP take/leave control                                     | [x]    | Access privilege / CCP                                                |
| DISCOVERY_CMD broadcast                                    | [x]    | Per-iface bind + directed + limited broadcast                         |
| Full DISCOVERY_ACK TLV parse                               | [x]    | ABRM dump in ACK: MAC, serial, manufacturer, model, user name         |
| Background heartbeat goroutine                             | [x]    | `HeartbeatTimeout/2`; pulse CCP                                       |
| `ImplementationEndianness` (`0x020C`) aware reg sync       | [x]    | Probe on TakeControl; ignore non-0/0xFFFFFFFF (GigE FirstURL overlap) |
| PENDING_ACK extends read deadline from `temporary_timeout` | [x]    | GenCP Table 12: reserved(2)+timeout_ms(2); fallback to client timeout |
| Bootstrap map completeness (`0x0000–0x0250`)               | [x]    | GenCP ABRM + GigE Vision ABRM/SBRM constants in `register_map.go`     |

### Phase 2 — GVSP (+ GenDC 1.1, PFNC 2.4)

Refs: `_references/GenDC/*`, `_references/SFNC/PFNC.h`, GVSP section of `AGENTS.md`. Payload-type IDs are the GVSP leader-field values shared with the GigE Vision spec.

| Item                                         | Status | Spec cue                                                                                      |
| -------------------------------------------- | ------ | --------------------------------------------------------------------------------------------- |
| Leader / payload / trailer reassembly        | [x]    | Standard + extended (EI) headers                                                              |
| 64-bit `block_id` / packet ID tracking       | [x]    | GEV 2.0 extended ID path present                                                              |
| Zero-alloc hot path + ring buffers           | [x]    | `buffer_pool.go` + `OOOPacketRing`; OOO refill zero-alloc, overflow spills only past 256 pkts |
| MTU / `GevSCPSPacketSize` + `SO_RCVBUF` warn | [x]    | Path MTU → negotiate SCPS (device clamp); 16MiB rcvbuf warn                                   |
| Packet resend (`RESEND_CMD`)                 | [x]    | Gap detect + `gvcp.RequestResend`; hold frame past trailer until filled                       |
| `PAYLOAD_TYPE_IMAGE`                         | [x]    | Leader payload-type field → `Frame.PayloadType`; constants in `payloadtype.go`                |
| `PAYLOAD_TYPE_CHUNK_DATA` / `CHUNK_ONLY`     | [x]    | Chunk-data / chunk-only dispatch                                                              |
| `PAYLOAD_TYPE_MULTI_PART`                    | [x]    | First image part dispatch                                                                     |
| `PAYLOAD_TYPE_GENDC`                         | [x]    | GenDC 1.1; container → image mapping                                                          |
| Vendor BSCF payload                          | [x]    | Huaray/Dahua; `Component` select (color/depth/mono) in `payload.go`                           |

#### GenDC 1.1 checklist (`GenDC.h`)

| Item                                                    | Status | Spec cue                                                     |
| ------------------------------------------------------- | ------ | ------------------------------------------------------------ |
| Detect signature `GNDC` (`0x43444E47`)                  | [x]    | `GDC_SIGNATURE`                                              |
| Parse `GenDCContainerHeader` (v1.1)                     | [x]    | `GDC_CONTAINER_HEADER`                                       |
| Component headers (`Intensity`, `Range`, `Metadata`, …) | [x]    | `GDC_*` component types                                      |
| Part headers: 2D / JPEG / JPEG2000 / H264 / Chunk / XML | [x]    | `GDC_2D_*`, `GDC_METADATA_*`                                 |
| Flow table header                                       | [x]    | `GDC_FLOW_TABLE_HEADER` (`FlowTableFromContainer`)           |
| Map GenDC 2D intensity → `Frame` / PFNC pixel format    | [x]    | Component → part → buffer (absolute DataOffset, SizeX/SizeY) |

#### PFNC decode matrix (`PFNC.h` → `internal/color`)

| Pixel format                  | ID             | EncodeJPEG | Notes                                 |
| ----------------------------- | -------------- | ---------- | ------------------------------------- |
| `Mono8`                       | `0x01080001`   | [x]        |                                       |
| `Mono16`                      | `0x01100007`   | [x]        | High-byte JPEG preview                |
| `RGB8`                        | `0x02180014`   | [x]        |                                       |
| `BGR8`                        | `0x02180015`   | [x]        | Default / heuristic                   |
| `YUV422_8` (YUYV)             | `0x02100032`   | [x]        |                                       |
| `YUV422_8_UYVY`               | `0x0210001F`   | [x]        | Distinct packing                      |
| `BayerRG8` / other Bayer      | `0x01080009` … | [x]        | Debayer before JPEG (`DebayerToRGBA`) |
| Packed Mono/Bayer (10p/12p/…) | various        | [x]        | `DecodeHighDepth` (LSB→MSB align)     |

### Phase 3 — GenApi 2.1.1 (+ SFNC 2.7 naming)

Refs: `_references/GenApi/GenICam_Standard_v2_1_1.pdf`, `_references/SFNC/GenICam_SFNC_v2_7.pdf`, GenApi section of `AGENTS.md`.

| Item                                                             | Status | Spec cue                                                                                                                                                                                                                                                             |
| ---------------------------------------------------------------- | ------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Local:/HTTP XML fetch + unzip                                    | [x]    | FirstURL / device memory                                                                                                                                                                                                                                             |
| Core node kinds + set/get                                        | [x]    | Integer, Boolean, Float, String, Enum, Command, \*Reg, SwissKnife, Converter                                                                                                                                                                                         |
| `Category` / `StructReg` as first-class types                    | [x]    | `Category` node kind + `Features` list + `NodeMap.Category`/`CategoryTree`/`RootCategories` (§2.8.2); `StructReg` → MaskedIntReg expansion (§2.8.6)                                                                                                                  |
| Pointers: `pAddress`, `pMin`/`pMax`/`pInc`, `pValue`             | [x]    | `pAddress`/`pValue` + `pMin`/`pMax`/`pInc` implemented; min/max/inc static values                                                                                                                                                                                    |
| `pIsImplemented` / `pIsAvailable` / `pIsLocked` / `pInvalidator` | [x]    | `IsImplemented`/`IsAvailable`/`IsLocked` + `GetInvalidator`; runtime access derived per §2.5 (NI→NA→RW/RO/WO narrowed by `ImposedAccessMode`, locked RW→RO / WO→NA) and enforced at Set*/Read*/CurrentEnum                                                           |
| Constant `Float`/`String`/`Integer` nodes                        | [x]    | `<Value>` with no address reads its constant (floating nodes; `readFloatReg`/`readStringReg`); `resolveAddr` accepts address 0 (ABRM) and errors on >32-bit sums                                                                                                     |
| ManifestTable (`0x01D0`) path                                    | [x]    | `ReadManifestTable` + `ManifestTableURL` preferred over FirstURL                                                                                                                                                                                                     |
| SwissKnife ops                                                   | [x]    | `+ - * / % & \| ^ << >> ~ ( ) = == != < > <= >= && \|\| ?: **` in `evaluator.go` (full § grammar, `<>` alias, `E`/`PI`, 24 math funcs)                                                                                                                               |
| SwissKnife variable suffixes (`.Min/.Max/.Inc/.Value/.Entry`)    | [x]    | `formulaContext` + `suffixResolver` + float-domain `formulaContextFloat`/`FloatSuffixResolver` in `port.go` + `floatParser`/`evalFormulaFloat` in `evaluator.go` (float `pVariable`s and float `SwissKnife` nodes now evaluated in float domain; int reads truncate) |
| SwissKnife funcs (`SQRT`, `FLOOR`, `CEIL`, `ABS`)                | [x]    | `ABS`, `FLOOR`, `CEIL`, `SQRT` in `evaluator.go`                                                                                                                                                                                                                     |
| Dedicated `port.go` binding + endianness                         | [x]    | Port node → `gvcp.Port` Read/Write; complete with byte order awareness                                                                                                                                                                                               |
| SFNC-required features for streaming                             | [x]    | `AcquisitionStart/Stop`, `AcquisitionMode`, `AcquisitionFrameRate` wired; `Gev*`/`Device*` constants + getters in `sfnc.go`/`camera.go`; `DeviceLinkHeartbeatTimeout` GenApi path; MDRT-aware GVCP timeout                                                           |

### Phase 4 — High-level API

| Item                                | Status | Spec cue                                                      |
| ----------------------------------- | ------ | ------------------------------------------------------------- |
| Channel stream API + buffer release | [x]    | `OpenDevice` / `StartStream` / `Frames()` in `framestream.go` |
| `cmd/` CLIs                         | [~]    | discover done; stream done                                    |

GenTL `.cti` producer/consumer interop is dropped: the pure-Go GVCP/GVSP/GenApi path is the portable standard-compliant route to any GigE Vision camera.

---

## Migration log

- **2026-09-24** — Formal `Gev*`/`Device*` streaming-feature coverage (v1.10.0): `sfnc.go` exports well-known SFNC feature-name constants; `Camera` gains typed convenience getters (`DeviceVendorName`, `DeviceModelName`, `DeviceSerialNumber`, `DeviceUserID`, `GevSCPSPacketSize`, `GevSCPD`); `gvcp.GVCP.MaximumDeviceResponseTime` reads the GenCP MDRT register and `connectCamera` uses it to widen the GVCP deadline for slow devices; `StartHeartbeat` accepts an optional timeout override and `Session` prefers the SFNC `DeviceLinkHeartbeatTimeout` GenApi feature when present, falling back to the GigE Vision SBRM register.
- **2026-09-24** — Dropped GenTL interop: removed `gentl/` (constants only; `cti.go` already gone) and `_references/GenTL/`, `_references/GenTL SFNC/`; cleared the GenTL rows, spec-index entries, and the Phase 4 `.cti` module ladder from this roadmap. Stance: gogige stays pure-Go GVCP/GVSP/GenApi — portable to any GigE Vision + GenICam camera without vendor `.cti` producers; README notes the portability rationale.
- **2026-09-24** — `gvsp.Sample.Overlay` (`Box2D`): `calib.ProjectPack`/`OverlayBoxes` project a pack's 3D oriented box onto the color image plane; `gogige.WithOverlay(bool, calib.CamCalib)` GrabOption toggles it (v1.8.0). Also `fix: genapi` command execution ignores the `pIsLocked` gate.
- **2026-09-13** — GenApi `Category` + `StructReg` first-class nodes (v1.7.0): parsed `Category` with ordered `Features`, `NodeMap.Category`/`Categories`/`RootCategories`/`CategoryTree`; `StructReg` expands to `MaskedIntReg` entries at parse time. Offline Huaray XML replay `1777 → 1821` nodes, 0 regressions.
- **2026-09-12** — GenApi conformance batch (v1.6.0–1.6.2): full SwissKnife function table + `**`/`==`/`<>`; `Converter` FROM/TO reserved variables; `Boolean` OnValue/OffValue; `IntReg` sign + endianness; `gvcp.ReadManifestTable` conformant GenCP Manifest Table + vendor MTAB; access-mode enforcement per §2.5; `resolveAddr` ABRM-0/32-bit fixes; constant node reads. Plus `gvsp` GenDC trailing-tag/part-offset fixes, payload-type IDs corrected to GenTL/GigE ids, `internal/color` Bayer R/B swaps + Mono10/12/14 decode. Audit recorded in `FINDINGS.md`.
- **2026-09-12** — Float-domain SwissKnife support: float alternative to the § formula evaluator (`evalFormulaFloat`, `floatParser`, `ATAN2`/`MIN`/`MAX`/`POW`/`LOG` funcs, `E`/`PI`), float variable context + suffix resolver (`formulaContextFloat`, `readFormulaVariableFloat`), `ReadFloat` SwissKnike case, integer-read truncation of float results with negative-guard in `resolveIntegerReference`. Offline replay of the Huaray `DS5131MG30CE` XML: the 9 float-SwissKnife errors cleared (333→324; remaining all guard-correct NI/NA/WO), 0 regressions. Locked by `TestEvalFormulaFloat`, `TestEvalFormulaFloatSuffix`, `TestSwissKnifeFloatDomain`.
- **2026-09-12** — GenApi core gaps closed (FINDINGS §3 items 1-4): `resolveAddr` accepts ABRM address 0 and rejects >32-bit sums; constant `Float`/`String`/`Integer` nodes read via `<Value>`; runtime access mode derived per §2.5 (NI/NA, `ImposedAccessMode`, locked RW→RO and WO→NA) and enforced at Set*/Read*/CurrentEnum incl. target-register writes; SwissKnife variable suffixes `.Min/.Max/.Inc/.Value/.Entry` resolved. Offline replay of the Huaray `DS5131MG30CE` XML shows 0 regressions (remaining errors are guard-correct NI/NA/WO or pre-existing float-domain SwissKnife gaps). Locked by `genapi/gaps_test.go`.
- **2026-08-09** — Fixed streaming OOM: OOO ring slots were preallocated at 8 MiB each (256 × 8 MiB ≈ 2 GiB per `frameBuild`), so the websocket/live examples ballooned to >10 GiB RSS and got SIGKILL'd. Slots now capped at 16 KiB (a single GVSP transport packet); `gvsp.Stream` additionally bounds concurrent in-flight frames (`maxInFlightFrames=64`) and evicts the oldest incomplete build when full.
- **2026-08-09** — Phase 2 payload typing complete: GVSP payload-type constants + `Frame.PayloadType` from leader (`gvsp/payloadtype.go`); `ParsePayloadByType` dispatches GenDC/Multi-Part/Chunk/Image. GenDC flow table parsing (`internal/genDC`) + 2D part SizeX/SizeY + absolute DataOffset fix. PFNC decode matrix finished in `internal/color` (`DecodeHighDepth` for Bayer/packed Mono+Bayer).
- **2026-08-09** — GVSP OOO zero-alloc: replaced `map[uint32][]byte` in `frameBuild` with pre-allocated `OOOPacketRing` (`gvsp/frame.go`), ring spills to a lazily-created overflow map only past `MaxOOOPackets` (256). `receiver.go` appendPayload uses ring `Put`/`Get`/`Delete`; `resend.go` adds `MissingPayloadRangesRing`. Fixed middle-delete ring compaction dropping the head packet; added `TestOOOPacketRing*` + `TestGVSPOutOfOrder`. Duplicate `frame_assemble.go` removed.
- **2026-08-13** — API ergonomics (1.4.0): `Camera` is now the unified handle — one-shot grabs (`GrabSample`/`GrabAllSamples`/`GrabComponents`/`GrabJPEG`), `Features()`, and consistent short get/set feature names (`SetInteger`/`SetEnum`/… plus `Integer`/`Enum`/`Float`/`String`/`Boolean` getters); added `NodeMap.ReadFloat`/`ReadString`; `Features` gained getters; `grab.FromCamera`; `PackDet` renamed `Length`/`Width`/`Height` → `LengthMm`/`WidthMm`/`HeightMm`.
- **2026-08-08** — Phase 3: Constraint pointers (pMin, pMax, pInc) complete. Added Node.GetConstraints(), NodeMap.GetMin/Max/Inc() methods. Parser now extracts Min/Max/Inc static values + pMin/pMax/pInc feature references. Enables parameter bounds validation. Test: TestConstraintPointers.
- **2026-08-08** — GenApi refactoring complete: `node.go` (Node interface + gcNode), `types.go` (node parsing: nodeFields, parseNodeXML, parseNodeMapXML), `port.go` (portAdapter binding → gvcp.Port); `nodemap.go` now clean orchestration layer; zero-alloc architecture with explicit separation of concerns per AGENTS.md.
- **2026-08-08** — Phase 4 API: `OpenDevice` → `*Camera`, `Camera.SetInteger` / `SetEnum`, `Camera.StartStream` → `Stream.Frames()` channel of pooled `*gvsp.Frame` with `Stop`/`Pause`/`Resume` (`framestream.go`); `Stream` alias → `GVSPStream`. New `examples/frames`.
- **2026-08-08** — Enforced package layout: `control/gvcp` → `gvcp/`, `control/genicam` → `genapi/`, `vision/gvsp` → `gvsp/`; lifted camera/device/session/grab/live into root `gige` package (`camera.go`, `stream.go`, `options.go`, …).
- **2026-08-08** — Cleared layout debt: `vision/bscf` → `gvsp/payload.go`; `vision/color` → `internal/color`; removed `vision/`.
- **2026-08-08** — Enriched roadmap from `_references/` (GenCP/GenApi/GenTL/GenDC/SFNC/PFNC); noted missing GigE Vision PDF.
- **2026-08-08** — Phase 1/2 practical slice: DISCOVERY_ACK ABRM parse, `gvcp.StartHeartbeat` (`HeartbeatTimeout/2`), `gvsp/buffer_pool.go` + `Frame.Release`.
- **2026-08-08** — Phase 2 packet resend: `gvsp` gap tracking + hole-fill reassembly; `gvcp.EncodePacketResend` / `RequestResend`; Session wires resender.
- **2026-08-08** — Phase 2 MTU/SCPS: `gvsp.PathMTU` + `PacketSizeForMTU`; `SO_RCVBUF` 16MiB with warn below 8MiB; acquisition RMW on `0x0D04`.
- **2026-08-08** — Phase 1 complete: GenCP ABRM + GigE ABRM in `register_map.go`; PENDING_ACK `temporary_timeout`; `SyncImplementationEndianness` on TakeControl + GenApi device byte-order for WriteMem.
- **2026-08-08** — Examples: `examples/smoke` (now `examples/smoke-test`), `examples/features`; CLI `cmd/gogige-stream`.
- **2026-08-08** — BSCF/SFNC `Component` (color/depth/mono): parse all component blocks; `GrabComponent` / `WithComponent` / `SetComponent`; Mono16 JPEG preview.
- **2026-08-08** — Root package `gige` → `gogige`; name now matches import path. Phase 4 happy path documented in `doc.go` (`OpenDevice` / `StartStream` / `Frames`).
- **2026-08-08** — GenApi introspection reads (`Kind`, `EnumEntries`, `CurrentEnum`, `ReadInteger`, `ReadBoolean`) + `Session.GrabComponents`; new `examples/probe-streams` and `examples/configure-camera`.
