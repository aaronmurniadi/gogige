# ROADMAP

Progress tracker against `.cursor/rules` and the EMVA/AIA specs under `_references/` (gitignored).

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
| `gvsp/frame.go`                | [x]    | Frame + reassembly helper                                                |
| `gvsp/receiver.go`             | [~]    | UDP receive path; ordered payloads use pool (OOO still allocates)        |
| `gvsp/payload.go`              | [~]    | BSCF vendor payload + `Sample`; Image/Multi-Part/GenDC still pending     |
| `gvsp/buffer_pool.go`          | [x]    | Pre-allocated frame buffers + `Frame.Release()`                          |
| `gvsp/resend.go`               | [x]    | Missing-packet tracking + `RESEND_CMD` via `Stream.SetResender`          |
| `genapi/camera_description.go` | [x]    | FirstURL fetch + zip/deflate XML                                         |
| `genapi/evaluator.go`          | [x]    | SwissKnife formula evaluator                                             |
| `genapi/nodemap.go`            | [x]    | Parse + feature get/set (monolithic)                                     |
| `genapi/node.go`               | [ ]    | Extract core Node interface from nodemap                                 |
| `genapi/types.go`              | [ ]    | Typed node kinds (IntReg, MaskedInt, Category, …)                        |
| `genapi/port.go`               | [ ]    | Explicit Port binding layer → `gvcp.Port`                                |
| `gentl/cti.go`                 | [ ]    | Optional GenTL `.cti` loader (`dlopen` / CGO)                            |
| `gentl/types.go`               | [ ]    | Mirror `GenTL.h` handles / enums                                         |
| `cmd/gogige-discover/`         | [x]    | CLI discovery utility                                                    |
| `cmd/gogige-stream/`           | [x]    | CLI N-frame JPEG + BSCF measurements                                     |
| `camera.go`                    | [x]    | High-level `Camera` + feature setters                                    |
| `discovery.go`                 | [x]    | Root `Discover` → `gvcp.Discover`                                        |
| `stream.go`                    | [~]    | `Session` / Grabber; not yet `StartStream` + `Frames()` channel API      |
| `options.go`                   | [x]    | `WithLogger` / `WithTimeout`                                             |

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
| Package name `gogige` (today: `gige`) | [ ]    | Align import path package name                              |
| `OpenDevice`                          | [ ]    | Today: `Open` → `Device`                                    |
| `SetInteger` / `SetEnum` on camera    | [~]    | Exists as `SetIntFeature` / `SetStringFeature` / `Features` |
| `StartStream` + `<-chan *Frame`       | [ ]    | Today: `StartGrabber` + `Grab`                              |
| `frame.Release()` buffer return       | [x]    | `gvsp.BufferPool` + `Frame.Release`                         |

---

## Protocol rigor (by phase)

### Phase 1 — GVCP / GenCP 1.3.1

Refs: `_references/GenCP/…`, architecture + `gvcp.mdc`. GigE Vision packet framing still needs the AIA PDF.

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

Refs: `_references/GenDC/*`, `_references/SFNC/PFNC.h`, `gvsp.mdc`. Payload-type IDs also appear in `GenTL.h` (`PAYLOAD_TYPE_*`).

| Item                                         | Status | Spec cue                                                |
| -------------------------------------------- | ------ | ------------------------------------------------------- |
| Leader / payload / trailer reassembly        | [x]    | Standard + extended (EI) headers                        |
| 64-bit `block_id` / packet ID tracking       | [~]    | GEV 2.0 extended ID path present                        |
| Zero-alloc hot path + ring buffers           | [~]    | `buffer_pool.go` + in-order path; OOO/oversize still alloc |
| MTU / `GevSCPSPacketSize` + `SO_RCVBUF` warn | [x]    | Path MTU → negotiate SCPS (device clamp); 16MiB rcvbuf warn |
| Packet resend (`RESEND_CMD`)                 | [x]    | Gap detect + `gvcp.RequestResend`; hold frame past trailer until filled |
| `PAYLOAD_TYPE_IMAGE`                         | [~]    | Image leader fields; not typed as GenTL enum yet        |
| `PAYLOAD_TYPE_CHUNK_DATA` / `CHUNK_ONLY`     | [ ]    | GenTL v1.2 / v1.4                                       |
| `PAYLOAD_TYPE_MULTI_PART`                    | [ ]    | GenTL v1.5                                              |
| `PAYLOAD_TYPE_GENDC`                         | [ ]    | GenTL v1.6 + GenDC 1.1                                  |
| Vendor BSCF payload                          | [x]    | Huaray/Dahua; stays in `payload.go` alongside standards |

#### GenDC 1.1 checklist (`GenDC.h`)

| Item                                                    | Status | Spec cue                     |
| ------------------------------------------------------- | ------ | ---------------------------- |
| Detect signature `GNDC` (`0x43444E47`)                  | [ ]    | `GDC_SIGNATURE`              |
| Parse `GenDCContainerHeader` (v1.1)                     | [ ]    | `GDC_CONTAINER_HEADER`       |
| Component headers (`Intensity`, `Range`, `Metadata`, …) | [ ]    | `GDC_*` component types      |
| Part headers: 2D / JPEG / JPEG2000 / H264 / Chunk / XML | [ ]    | `GDC_2D_*`, `GDC_METADATA_*` |
| Flow table header                                       | [ ]    | `GDC_FLOW_TABLE_HEADER`      |
| Map GenDC 2D intensity → `Frame` / PFNC pixel format    | [ ]    | Component → part → buffer    |

#### PFNC decode matrix (`PFNC.h` → `internal/color`)

| Pixel format                  | ID             | EncodeJPEG | Notes                |
| ----------------------------- | -------------- | ---------- | -------------------- |
| `Mono8`                       | `0x01080001`   | [x]        |                      |
| `RGB8`                        | `0x02180014`   | [x]        |                      |
| `BGR8`                        | `0x02180015`   | [x]        | Default / heuristic  |
| `YUV422_8` (YUYV)             | `0x02100032`   | [x]        |                      |
| `YUV422_8_UYVY`               | `0x0210001F`   | [ ]        | Distinct packing     |
| `BayerRG8` / other Bayer      | `0x01080009` … | [ ]        | Debayer before JPEG  |
| Packed Mono/Bayer (10p/12p/…) | various        | [ ]        | As needed by devices |

### Phase 3 — GenApi 2.1.1 (+ SFNC 2.7 naming)

Refs: `_references/GenApi/GenICam_Standard_v2_1_1.pdf`, `_references/SFNC/GenICam_SFNC_v2_7.pdf`, `genapi.mdc`.

| Item                                                             | Status     | Spec cue                                                                     |
| ---------------------------------------------------------------- | ---------- | ---------------------------------------------------------------------------- | ---- | --- | ------------------------ |
| Local:/HTTP XML fetch + unzip                                    | [x]        | FirstURL / device memory                                                     |
| Core node kinds + set/get                                        | [~]        | Integer, Boolean, Float, String, Enum, Command, \*Reg, SwissKnife, Converter |
| `Category` / `StructReg` as first-class types                    | [ ]        | Parsed skip today                                                            |
| Pointers: `pAddress`, `pMin`/`pMax`/`pInc`, `pValue`             | [~]        | `pAddress`/`pValue` used; min/max/inc TBD                                    |
| `pIsImplemented` / `pIsAvailable` / `pIsLocked` / `pInvalidator` | [ ]        | Cache invalidation                                                           |
| ManifestTable (`0x01D0`) path                                    | [ ]        | Preferred over FirstURL when present                                         |
| SwissKnife ops (`+ - \* / % \*\* &                               | ^ << >> && |                                                                              | ?:`) | [~] | Subset in `evaluator.go` |
| SwissKnife funcs (`SQRT`, `FLOOR`, `CEIL`, `ABS`)                | [~]        | Verify against Standard § formula grammar                                    |
| Dedicated `port.go` binding + endianness                         | [ ]        | Port node → `gvcp.Port` Read/Write                                           |
| SFNC-required features for streaming                             | [~]        | `AcquisitionStart/Stop` used; formal Gev\* coverage TBD                      |

### Phase 4 — High-level + GenTL 1.6

Refs: `_references/GenTL/GenTL.h`, `GenICam_GenTL_1_6.pdf`, GenTL SFNC 1.2.

| Item                                      | Status | Spec cue                |
| ----------------------------------------- | ------ | ----------------------- |
| Channel stream API + buffer release       | [ ]    | Phase 4 consumer API    |
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

- **2026-08-08** — Enforced package layout: `control/gvcp` → `gvcp/`, `control/genicam` → `genapi/`, `vision/gvsp` → `gvsp/`; lifted camera/device/session/grab/live into root `gige` package (`camera.go`, `stream.go`, `options.go`, …).
- **2026-08-08** — Cleared layout debt: `vision/bscf` → `gvsp/payload.go`; `vision/color` → `internal/color`; removed `vision/`.
- **2026-08-08** — Enriched roadmap from `_references/` (GenCP/GenApi/GenTL/GenDC/SFNC/PFNC); noted missing GigE Vision PDF.
- **2026-08-08** — Phase 1/2 practical slice: DISCOVERY_ACK ABRM parse, `gvcp.StartHeartbeat` (`HeartbeatTimeout/2`), `gvsp/buffer_pool.go` + `Frame.Release`.
- **2026-08-08** — Phase 2 packet resend: `gvsp` gap tracking + hole-fill reassembly; `gvcp.EncodePacketResend` / `RequestResend`; Session wires resender.
- **2026-08-08** — Phase 2 MTU/SCPS: `gvsp.PathMTU` + `PacketSizeForMTU`; `SO_RCVBUF` 16MiB with warn below 8MiB; acquisition RMW on `0x0D04`.
- **2026-08-08** — Phase 1 complete: GenCP ABRM + GigE ABRM in `register_map.go`; PENDING_ACK `temporary_timeout`; `SyncImplementationEndianness` on TakeControl + GenApi device byte-order for WriteMem.
- **2026-08-08** — Examples: `examples/smoke`, `examples/features`; CLI `cmd/gogige-stream`.
