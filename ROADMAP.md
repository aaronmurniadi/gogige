# ROADMAP

Progress tracker against `AGENTS.md` and GiGE Vision specs under `_references/`.

Legend: `[x]` done · `[ ]` not started · `[~]` partial

---

## Spec index (`_references/`)

Target versions from architecture rules, mapped to local artifacts:

| Standard         | Target    | Local reference                                                                   | Used by                                                    |
| ---------------- | --------- | --------------------------------------------------------------------------------- | ---------------------------------------------------------- |
| GenCP            | 1.3.1     | `_references/GenCP/GenICam_GenCP_1.3.1.pdf`                                       | `gvcp/` (control channel is GenCP over GigE UDP)           |
| GenICam / GenApi | 2.1.1     | `_references/GenApi/GenICam_Standard_v2_1_1.pdf`                                  | `genapi/`                                                  |
| GenTL            | 1.6       | `_references/GenTL/GenICam_GenTL_1_6.pdf`, `GenTL/GenTL.h`                        | `gentl/` (future)                                          |
| GenTL SFNC       | 1.2       | `_references/GenTL SFNC/GenICam_GenTL_SFNC_1_2.pdf`                               | GenTL feature naming                                       |
| GenDC            | 1.1       | `_references/GenDC/GenICam_GenDC_v1_1.pdf`, `GenDC/GenDC.h`                       | `gvsp/payload` GenDC path                                  |
| SFNC             | 2.7       | `_references/SFNC/GenICam_SFNC_v2_7.pdf`                                          | Feature names (`AcquisitionStart`, `GevSCPSPacketSize`, …) |
| PFNC             | 2.4       | `_references/SFNC/PFNC.h`, `GenICam_PFNC_2_4.pdf`, `GenICamPixelFormatValues.pdf` | `internal/color`, pixel format IDs                         |
| CLProtocol       | 1.2       | `_references/GenApi/GenICam_CLProtocol_Standard_v1.2.pdf`                         | Out of scope (Camera Link)                                 |
| GigE Vision      | 2.0 / 2.1 | `_references/GigE_Vision_for_Realtime_MV_11052010.pdf`, `GigE_Features_Reference.pdf` | `gvcp/`, `gvsp/`                                           |

Authoritative machine-readable headers for implementers: `GenDC/GenDC.h`, `GenTL/GenTL.h`, `SFNC/PFNC.h`.

---

## Package layout

### Target tree

| Path                           | Status | Notes                                                                    |
| ------------------------------ | ------ | ------------------------------------------------------------------------ |
| `gvcp/client.go`               | [x]    | Control channel, register R/W                                            |
| `gvcp/discovery.go`            | [x]    | Per-iface bind + directed/limited broadcast; DISCOVERY_ACK ABRM parse    |
| `gvcp/heartbeat.go`            | [x]    | `PulseHeartbeat` + background tick at `HeartbeatTimeout/2`               |
| `gvcp/packet.go`               | [x]    | Header encode + cmd constants                                            |
| `gvcp/register_map.go`         | [x]    | GenCP ABRM 0x0000–0x0250 + GigE Vision ABRM/SBRM                     |
| `gvsp/frame.go`                | [x]    | Frame + reassembly; `OOOPacketRing` zero-alloc OOO store (rare overflow map) |
| `gvsp/receiver.go`             | [x]    | UDP receive path; OOO via ring, contiguous via pool                          |
| `gvsp/payload.go`              | [x]    | BSCF + `Component` / `GrabAll`; payload-type dispatch (Image/Multi-Part/GenDC/Chunk) |
| `gvsp/buffer_pool.go`          | [x]    | Pre-allocated frame buffers + `Frame.Release()`                          |
| `gvsp/resend.go`               | [x]    | Missing-packet tracking + `RESEND_CMD` via `Stream.SetResender`          |
| `genapi/camera_description.go` | [x]    | FirstURL fetch + zip/deflate XML                                         |
| `genapi/evaluator.go`          | [x]    | SwissKnife formula evaluator                                             |
| `genapi/nodemap.go`            | [x]    | Parse + feature get/set (orchestration layer)                            |
| `genapi/node.go`               | [x]    | Core Node interface + gcNode struct + attributes                         |
| `genapi/types.go`              | [x]    | Node parsing: nodeFields, parseNodeXML, parseNodeMapXML stream           |
| `genapi/port.go`               | [x]    | Port binding layer: portAdapter → gvcp.Port I/O + byte order             |
| `gentl/cti.go`                 | [ ]    | Optional GenTL `.cti` loader (`dlopen` / CGO)                            |
| `gentl/types.go`               | [x]    | Mirror `GenTL.h` handles / enums                                         |
| `cmd/gogige-discover/`         | [x]    | CLI discovery utility                                                    |
| `cmd/gogige-stream/`           | [x]    | CLI N-frame JPEG + BSCF measurements                                     |
| `camera.go`                    | [x]    | High-level `Camera` + feature setters                                    |
| `discovery.go`                 | [x]    | Root `Discover` → `gvcp.Discover`                                        |
| `stream.go`                    | [x]    | `Session` / `Grab` / `GrabAll`; `StartStream` + `Frames()` in `framestream.go` |
| `options.go`                   | [x]    | `WithLogger` / `WithTimeout` / `WithComponent` / `GrabComponent` |

### Layout debt (outside target tree)

| Path                                                           | Status | Notes                                                     |
| -------------------------------------------------------------- | ------ | --------------------------------------------------------- |
| Root extras (`device.go`, `grab.go`, `live.go`, `alias.go`, …) | [~]    | Working high-level API; converge on Phase 4 surface below |

Cleared: `vision/bscf` → `gvsp/payload.go`; `vision/color` → `internal/color` (`vision/` removed).

---

## Architecture / Phase 4 high-level API

```go
cam, err := gogige.OpenDevice(ctx, "192.168.1.100")
stream, err := cam.StartStream(ctx)
for frame := range stream.Frames() { frame.Release() }
```

| Item                                  | Status | Notes                                                       |
| ------------------------------------- | ------ | ----------------------------------------------------------- |
| Package name `gogige` (was `gige`)    | [x]    | Renamed; name now matches import path                       |
| `OpenDevice`                          | [x]    | Returns `*Camera`; `Open`→`Device` kept for legacy consumers |
| `SetInteger` / `SetEnum` on camera    | [x]    | `camera.go` (wrap `SetIntFeature` / `SetStringFeature`)     |
| `StartStream` + `<-chan *Frame`       | [x]    | `framestream.go`; pooled frames, `Stop`/`Pause`/`Resume`    |
| `frame.Release()` buffer return       | [x]    | `gvsp.BufferPool` + `Frame.Release`                         |

---

## Protocol rigor (by phase)

### Phase 1 — GVCP / GenCP 1.3.1

Refs: `_references/GenCP/…`, architecture + GVCP section of `AGENTS.md`. GigE Vision packet framing still needs the AIA PDF.

| Item                                                       | Status | Spec cue                                           |
| ---------------------------------------------------------- | ------ | -------------------------------------------------- |
| READ/WRITE REG/MEM + PENDING_ACK                           | [x]    | GenCP memory/register cmds over UDP/3956           |
| CCP take/leave control                                     | [x]    | Access privilege / CCP                             |
| DISCOVERY_CMD broadcast                                    | [x]    | Per-iface bind + directed + limited broadcast      |
| Full DISCOVERY_ACK TLV parse                               | [x]    | ABRM dump in ACK: MAC, serial, manufacturer, model, user name |
| Background heartbeat goroutine                             | [x]    | `HeartbeatTimeout/2`; pulse CCP                        |
| `ImplementationEndianness` (`0x020C`) aware reg sync       | [x]    | Probe on TakeControl; ignore non-0/0xFFFFFFFF (GigE FirstURL overlap) |
| PENDING_ACK extends read deadline from `temporary_timeout` | [x]    | GenCP Table 12: reserved(2)+timeout_ms(2); fallback to client timeout |
| Bootstrap map completeness (`0x0000–0x0250`)               | [x]    | GenCP ABRM + GigE Vision ABRM/SBRM constants in `register_map.go` |

### Phase 2 — GVSP (+ GenDC 1.1, PFNC 2.4)

Refs: `_references/GenDC/*`, `_references/SFNC/PFNC.h`, GVSP section of `AGENTS.md`. Payload-type IDs also appear in `GenTL.h` (`PAYLOAD_TYPE_*`).

| Item                                         | Status | Spec cue                                                |
| -------------------------------------------- | ------ | ------------------------------------------------------- |
| Leader / payload / trailer reassembly        | [x]    | Standard + extended (EI) headers                        |
| 64-bit `block_id` / packet ID tracking       | [x]    | GEV 2.0 extended ID path present                        |
| Zero-alloc hot path + ring buffers           | [x]    | `buffer_pool.go` + `OOOPacketRing`; OOO refill zero-alloc, overflow spills only past 256 pkts |
| MTU / `GevSCPSPacketSize` + `SO_RCVBUF` warn | [x]    | Path MTU → negotiate SCPS (device clamp); 16MiB rcvbuf warn |
| Packet resend (`RESEND_CMD`)                 | [x]    | Gap detect + `gvcp.RequestResend`; hold frame past trailer until filled |
| `PAYLOAD_TYPE_IMAGE`                         | [x]    | Leader payload-type field → `Frame.PayloadType`; GenTL enum constants in `payloadtype.go` |
| `PAYLOAD_TYPE_CHUNK_DATA` / `CHUNK_ONLY`     | [x]    | GenTL v1.2 / v1.4; chunk-only dispatch                    |
| `PAYLOAD_TYPE_MULTI_PART`                    | [x]    | GenTL v1.5; first image part dispatch                    |
| `PAYLOAD_TYPE_GENDC`                         | [x]    | GenTL v1.6 + GenDC 1.1; container → image mapping         |
| Vendor BSCF payload                          | [x]    | Huaray/Dahua; `Component` select (color/depth/mono) in `payload.go` |

#### GenDC 1.1 checklist (`GenDC.h`)

| Item                                                    | Status | Spec cue                     |
| ------------------------------------------------------- | ------ | ---------------------------- |
| Detect signature `GNDC` (`0x43444E47`)                  | [x]    | `GDC_SIGNATURE`              |
| Parse `GenDCContainerHeader` (v1.1)                     | [x]    | `GDC_CONTAINER_HEADER`       |
| Component headers (`Intensity`, `Range`, `Metadata`, …) | [x]    | `GDC_*` component types      |
| Part headers: 2D / JPEG / JPEG2000 / H264 / Chunk / XML | [x]    | `GDC_2D_*`, `GDC_METADATA_*` |
| Flow table header                                       | [x]    | `GDC_FLOW_TABLE_HEADER` (`FlowTableFromContainer`) |
| Map GenDC 2D intensity → `Frame` / PFNC pixel format    | [x]    | Component → part → buffer (absolute DataOffset, SizeX/SizeY) |

#### PFNC decode matrix (`PFNC.h` → `internal/color`)

| Pixel format                  | ID             | EncodeJPEG | Notes                |
| ----------------------------- | -------------- | ---------- | -------------------- |
| `Mono8`                       | `0x01080001`   | [x]        |                      |
| `Mono16`                      | `0x01100007`   | [x]        | High-byte JPEG preview |
| `RGB8`                        | `0x02180014`   | [x]        |                      |
| `BGR8`                        | `0x02180015`   | [x]        | Default / heuristic  |
| `YUV422_8` (YUYV)             | `0x02100032`   | [x]        |                      |
| `YUV422_8_UYVY`               | `0x0210001F`   | [x]        | Distinct packing     |
| `BayerRG8` / other Bayer      | `0x01080009` … | [x]        | Debayer before JPEG (`DebayerToRGBA`) |
| Packed Mono/Bayer (10p/12p/…) | various        | [x]        | `DecodeHighDepth` (LSB→MSB align) |

### Phase 3 — GenApi 2.1.1 (+ SFNC 2.7 naming)

Refs: `_references/GenApi/GenICam_Standard_v2_1_1.pdf`, `_references/SFNC/GenICam_SFNC_v2_7.pdf`, GenApi section of `AGENTS.md`.

| Item                                                             | Status     | Spec cue                                                                     |
| ---------------------------------------------------------------- | ---------- | ---------------------------------------------------------------------------- | ---- | --- | ------------------------ |
| Local:/HTTP XML fetch + unzip                                    | [x]        | FirstURL / device memory                                                     |
| Core node kinds + set/get                                        | [x]        | Integer, Boolean, Float, String, Enum, Command, \*Reg, SwissKnife, Converter |
| `Category` / `StructReg` as first-class types                    | [ ]        | Parsed skip today                                                            |
| Pointers: `pAddress`, `pMin`/`pMax`/`pInc`, `pValue`             | [x]        | `pAddress`/`pValue` + `pMin`/`pMax`/`pInc` implemented; min/max/inc static values |
| `pIsImplemented` / `pIsAvailable` / `pIsLocked` / `pInvalidator` | [ ]        | Cache invalidation                                                           |
| ManifestTable (`0x01D0`) path                                    | [ ]        | Preferred over FirstURL when present                                         |
| SwissKnife ops (`+ - \* / % \*\* &                               | ^ << >> && |                                                                              | ?:`) | [~] | Subset in `evaluator.go` |
| SwissKnife funcs (`SQRT`, `FLOOR`, `CEIL`, `ABS`)                | [~]        | Verify against Standard § formula grammar                                    |
| Dedicated `port.go` binding + endianness                         | [x]        | Port node → `gvcp.Port` Read/Write; complete with byte order awareness       |
| SFNC-required features for streaming                             | [~]        | `AcquisitionStart/Stop` used; formal Gev\* coverage TBD                      |

### Phase 4 — High-level + GenTL 1.6

Refs: `_references/GenTL/GenTL.h`, `GenICam_GenTL_1_6.pdf`, GenTL SFNC 1.2.

| Item                                      | Status | Spec cue                |
| ----------------------------------------- | ------ | ----------------------- |
| Channel stream API + buffer release       | [x]    | `OpenDevice` / `StartStream` / `Frames()` in `framestream.go` |
| Optional `gentl/` CGO producer / consumer | [ ]    | See module ladder below |
| `cmd/` CLIs                               | [~]    | discover done; stream done                  |

#### GenTL module ladder (`GenTL.h`)

Produce or consume via `.cti` — pure-Go path can stay primary; GenTL is optional interop.

| Step                                              | Status | API                                 |
| ------------------------------------------------- | ------ | ----------------------------------- |
| Load `.cti` + `GCInitLib`                         | [ ]    | `gentl/cti.go`                      |
| `TLOpen` → `TLOpenInterface`                      | [ ]    | `TL_HANDLE` / `IF_HANDLE`           |
| `IFOpenDevice` (`DEVICE_ACCESS_*`)                | [ ]    | Exclusive / Control / ReadOnly      |
| `DevOpenDataStream`                               | [ ]    | `DS_HANDLE`                         |
| Announce / queue buffers                          | [ ]    | `DSAnnounceBuffer`, `DSQueueBuffer` |
| `DSStartAcquisition` + `EVENT_NEW_BUFFER`         | [ ]    | `EventGetData`                      |
| Map `PAYLOAD_TYPE_GENDC` / `MULTI_PART` / `IMAGE` | [ ]    | Align with `gvsp/payload`           |
| GenTL SFNC feature names on modules               | [ ]    | `_references/GenTL SFNC/…`          |

---

## Migration log

- **2026-08-09** — Fixed streaming OOM: OOO ring slots were preallocated at 8 MiB each (256 × 8 MiB ≈ 2 GiB per `frameBuild`), so the websocket/live examples ballooned to >10 GiB RSS and got SIGKILL'd. Slots now capped at 16 KiB (a single GVSP transport packet); `gvsp.Stream` additionally bounds concurrent in-flight frames (`maxInFlightFrames=64`) and evicts the oldest incomplete build when full.

- **2026-08-09** — Phase 2 payload typing complete: GVSP payload-type constants + `Frame.PayloadType` from leader (`gvsp/payloadtype.go`); `ParsePayloadByType` dispatches GenDC/Multi-Part/Chunk/Image. GenDC flow table parsing (`internal/genDC`) + 2D part SizeX/SizeY + absolute DataOffset fix. PFNC decode matrix finished in `internal/color` (`DecodeHighDepth` for Bayer/packed Mono+Bayer).

- **2026-08-09** — GVSP OOO zero-alloc: replaced `map[uint32][]byte` in `frameBuild` with pre-allocated `OOOPacketRing` (`gvsp/frame.go`), ring spills to a lazily-created overflow map only past `MaxOOOPackets` (256). `receiver.go` appendPayload uses ring `Put`/`Get`/`Delete`; `resend.go` adds `MissingPayloadRangesRing`. Fixed middle-delete ring compaction dropping the head packet; added `TestOOOPacketRing*` + `TestGVSPOutOfOrder`. Duplicate `frame_assemble.go` removed.

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
- **2026-08-08** — Examples: `examples/smoke`, `examples/features`; CLI `cmd/gogige-stream`.
- **2026-08-08** — BSCF/SFNC `Component` (color/depth/mono): parse all component blocks; `GrabComponent` / `WithComponent` / `SetComponent`; Mono16 JPEG preview.
- **2026-08-08** — Root package `gige` → `gogige`; name now matches import path. Phase 4 happy path documented in `doc.go` (`OpenDevice` / `StartStream` / `Frames`).
- **2026-08-08** — GenApi introspection reads (`Kind`, `EnumEntries`, `CurrentEnum`, `ReadInteger`, `ReadBoolean`) + `Session.GrabComponents`; new `examples/probe-streams` and `examples/configure-camera`.
