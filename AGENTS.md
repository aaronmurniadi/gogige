# gogige Project Rules

You are a senior systems engineer expert in GenICam protocols (GigE Vision/GVCP/GVSP, GenApi, GenTL, and GenDC). Apply these rules to write, review, or refactor any code in the gogige repository.

## Architectural hierarchy

```
               +------------------------------------------------+
               |                  gogige API                    |
               |       (High-Level Camera & Stream Interface)   |
               +-----------------------+------------------------+
                                       |
          +----------------------------+----------------------------+
          |                                                         |
+---------v----------+                                   +----------v---------+
|     GenApi XML     |                                   |    GVSP Stream     |
|   Parser & Tree    |                                   | Engine (UDP Data)  |
+---------+----------+                                   +----------+---------+
          |                                                         |
+---------v----------+                                   +----------v---------+
|     GVCP Client    |                                   |  Frame Buffer /    |
| (Control & Reg Map)|                                   | GenDC Reconstruct  |
+---------+----------+                                   +--------------------+
          |
+---------v----------+
| Physical Network / |
|   UDP Sockets      |
+--------------------+
```

## Protocol rigor (strict EMVA/A3)

- GigE Vision Version 2.0/2.1
- GenCP Version 1.3.1
- GenICam Standard Version 2.1.1 (GenApi Node Map)
- GenTL Version 1.6
- GenDC Version 1.1

## Package layout

Enforce the following project package layout:

```
gogige/
├── .agents/
│   └── skills/
├── cmd/
│   ├── gogige-discover/         # CLI discovery utility
│   └── gogige-stream/           # CLI stream capture utility
├── gvcp/                        # GigE Vision Control Protocol (GenCP)
│   ├── client.go                # Control channel client, register read/write
│   ├── discovery.go             # UDP broadcast DISCOVERY_CMD and parser
│   ├── heartbeat.go             # Background heartbeat maintenance thread
│   ├── packet.go                # GVCP packet serialization / deserialization
│   └── register_map.go          # Device Bootstrap Register Map (ABRM/SBRM)
├── gvsp/                        # GigE Vision Streaming Protocol
│   ├── buffer_pool.go           # Lock-free/ring frame buffer management
│   ├── frame.go                 # Frame construction & metadata handling
│   ├── payload.go               # Payload parsing (Image, Multi-Part, GenDC)
│   ├── receiver.go              # Hot-path UDP frame receiver (zero-alloc)
│   └── resend.go                # Packet resend logic (RESEND_CMD generation)
├── genapi/                      # GenICam GenApi XML Module
│   ├── camera_description.go    # Manifest parsing & zip/deflate decompression
│   ├── evaluator.go             # SwissKnife / IntSwissKnife expression parser
│   ├── node.go                  # Core Node interface and attributes
│   ├── nodemap.go               # GenApi NodeMap tree structure
│   ├── types.go                 # Node types (IntReg, MaskedInt, Category, etc.)
│   └── port.go                  # Port binding GenApi nodes -> GVCP client
├── gentl/                       # Optional GenTL Producer/Consumer CGO bindings
│   ├── cti.go                   # .cti shared library interface loader
│   └── types.go                 # GenTL C-interface bindings
├── camera.go                    # High-Level Camera abstraction API
├── discovery.go                 # High-Level Network Discovery API
├── stream.go                    # High-Level Frame Stream Channel API
├── options.go                   # Functional Option patterns for gogige Config
├── go.mod
└── go.sum
```

Target layout is enforced for protocol packages (`gvcp/`, `gvsp/`, `genapi/`) and root high-level files (`camera.go`, `discovery.go`, `stream.go`, `options.go`). Track missing files and Phase 4 API gaps in `ROADMAP.md`. Prefer target names for all new code; do not reintroduce `control/` or `vision/gvsp`.

## GVCP Protocol & Discovery (Phase 1)

### Discovery protocol

- Broadcast `DISCOVERY_CMD` packet over UDP to port 3956 (or `255.255.255.255`)
- Parse `DISCOVERY_ACK` to extract IP, MAC, Serial Number, User Defined Name, Manufacturer, and Model Name

### Control channel setup

- Open a dedicated UDP control socket
- Command IDs: `READREG_CMD` (0x0800), `READREG_ACK` (0x0801), `WRITEREG_CMD` (0x0802), `WRITEREG_ACK` (0x0803), `READMEM_CMD`, `READMEM_ACK`, `WRITEMEM_CMD`, `WRITEMEM_ACK`, `PENDING_ACK` (0x0805)
- Handle `PENDING_ACK` by dynamically extending client network read timeouts to match the `temporary_timeout` sent by the camera
- Implement request ID synchronization (sequentially incremented `request_id`)
- Implement Access Privilege control: Available (0), Open (Exclusive) (1)

### Heartbeat maintenance

- Spin a background goroutine ticking at `HeartbeatTimeout / 2` (default timeout = 3000ms)
- Issue periodic `READREG_CMD` on AccessPrivilege (`0x0204`) or HeartbeatTimeout (`0x01E8`) to keep control channel privilege alive

### Endianness

- GVCP network headers are BigEndian
- Bootstrap register map (`0x0000–0x0250`) fields use BigEndian protocol ordering
- Device implementation registers depend on `ImplementationEndianness` (`0x020C`) (`0` = BigEndian, `0xFFFFFFFF` = LittleEndian). Convert byte arrays accordingly during register map synchronization.

## GVSP Streaming Protocol (Phase 2)

### Stream channel setup

- Negotiate network MTU and set `GevSCPSPacketSize` (e.g., Jumbo frames 9000 bytes or standard 1500 bytes)
- Configure UDP socket buffer sizes using sysctl OS hints (`SO_RCVBUF` set to 8MB–64MB) to minimize kernel-level packet drops
- Check MTU capability dynamically before establishing stream channels. Warn users if socket buffers are constrained below recommended GigE streaming thresholds.

### Payload handling

- Parse packet headers: Leader, Data, Trailer, Extended Chunk Data, Multi-Part, and GenDC
- Track `block_id` (64-bit in GEV 2.0) and `packet_id` (24-bit/32-bit)
- Support `PAYLOAD_TYPE_IMAGE`, `PAYLOAD_TYPE_CHUNK_DATA`, `PAYLOAD_TYPE_MULTI_PART`, and `PAYLOAD_TYPE_GENDC`

### Buffer pool & resend logic

- Pre-allocate frame buffers based on device payload size
- Track missing `packet_id` per `block_id`. Send `RESEND_CMD` over GVCP control channel if missing packets are detected before frame completion timeout.

### Memory allocation guardrail

NEVER allocate slices or structs within `gvsp.Receiver.Listen()` loops. Allocate all packet memory upfront into ring buffers.

```go
// ❌ BAD — alloc inside Listen loop
buf := make([]byte, n)

// ✅ GOOD — ring / pool allocated upfront; Listen only writes into slots
```

## GenApi XML Handling (Phase 3)

### Manifest & description retrieval

- Query Bootstrap Register `ManifestTableAddress` (`0x01D0`)
- Read Manifest Entries to locate device XML
- Decompress ZIP (DEFLATE/STORE) XML or raw XML payloads from device memory space using `READMEM_CMD`

### Node map construction

Parse node topologies: Category, Integer, IntReg, MaskedIntReg, Float, FloatReg, Enumeration, EnumEntry, Boolean, Command, StringReg, SwissKnife, IntSwissKnife, Converter, Port.

Handle `pAddress`, `pMin`, `pMax`, `pInc`, `pValue`, `pIsImplemented`, `pIsAvailable`, `pIsLocked`, `pInvalidator`.

### SwissKnife evaluator

Implement dynamic AST tokenization and formula evaluation supporting operators (`+`, `-`, `*`, `/`, `%`, `**`, `&`, `|`, `^`, `<<`, `>>`, `&&`, `||`, comparison, ternary `? :`) and mathematical functions (`SQRT`, `FLOOR`, `CEIL`, `ABS`).

### Port binding

Bind Port nodes directly to `gvcp.Client` byte read/write calls (`Read(addr, len)` and `Write(addr, data)`). Respect device register endianness (BigEndian / LittleEndian).

## Idiomatic Go conventions

- Explicit `context.Context` cancellation for all blocking and stream routines
- Channel-based stream consumption (`<-chan *Frame`)
- Standard library primitives over unvetted 3rd-party dependencies
- Comprehensive Go error wrapping (`fmt.Errorf("...: %w", err)`)

## High-level API surface (Phase 4)

```go
cam, err := gogige.OpenDevice(ctx, "192.168.1.100")
defer cam.Close()

// Node operations
err = cam.SetInteger("Width", 1920)
err = cam.SetEnum("PixelFormat", "Mono8")

// Streaming
stream, err := cam.StartStream(ctx)
defer stream.Stop()

for frame := range stream.Frames() {
    // Process frame
    frame.Release() // Return buffer to pre-allocated pool
}
```

## Performance & safety (global)

### Zero allocation in hot loops

The hot frame reception path (`gvsp/`) must achieve zero heap allocations per packet under normal streaming conditions using pre-allocated packet/frame buffers.

### Endianness awareness

- GVCP network headers are BigEndian
- Bootstrap register map (`0x0000–0x0250`) fields use BigEndian protocol ordering
- Device implementation registers depend on `ImplementationEndianness` (`0x020C`) (`0` = BigEndian, `0xFFFFFFFF` = LittleEndian). Convert byte arrays accordingly during register map synchronization.

### Network sockets & MTU

Check MTU capability dynamically before establishing stream channels. Warn users if socket buffers are constrained below recommended GigE streaming thresholds.
