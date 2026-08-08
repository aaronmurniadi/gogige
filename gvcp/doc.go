// Package gvcp provides GigE Vision Control Protocol (GVCP) implementation.
//
// GigE Vision Version 2.0/2.1
// GenCP Version 1.3.1
// Reference: GigE Vision for Realtime MV, GenICam_GenCP_1.3.1.pdf
//
// # Overview
//
// GVCP (GigE Vision Control Protocol) provides the control channel for camera
// configuration and command execution. It is part of the GigE Vision standard
// and uses the GenCP (GenICam Control Protocol) for register access.
//
// # Package Content
//
// gvcp/ provides:
//   - Client: Control channel client for register read/write operations
//   - Discovery: UDP broadcast DISCOVERY_CMD and DISCOVERY_ACK parser
//   - Heartbeat: Background thread for maintaining access privilege
//   - Packet: GVCP packet serialization/deserialization
//   - Register map: Device Bootstrap Register Map (ABRM/SBRM)
//
// # Protocol Details
//
// GVCP uses UDP port 3956 for control commands. The protocol includes:
//   - READREG_CMD / READREG_ACK: Register read operations
//   - WRITEREG_CMD / WRITEREG_ACK: Register write operations
//   - READMEM_CMD / READMEM_ACK: Memory read operations
//   - WRITEMEM_CMD / WRITEMEM_ACK: Memory write operations
//   - PENDING_ACK: Dynamic timeout extension mechanism
//
// GenCP (GenICam Control Protocol) defines the register map:
//   - Bootstrap Register Map (ABRM): 0x0000-0x0250
//   - Standard Register Map (SBRM): Device-specific
//   - Access privilege control: Available (0), Open (Exclusive) (1)
//
// References
//
//   - GigE Vision for Realtime MV (2010)
//   - GenICam GenCP Standard v1.3.1
//   - https://www.emva.org/standards-technical-resources/gige-vision-documentation/
package gvcp
