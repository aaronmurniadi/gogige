// Package gvsp provides GigE Vision Streaming Protocol (GVSP) implementation.
//
// GigE Vision Version 2.0/2.1
// GenTL Version 1.6 (for payload types)
// GenDC Version 1.1 (for multi-part/GenDC payloads)
// Reference: GigE Vision for Realtime MV, GenICam_GenTL_1_6.pdf, GenICam_GenDC_v1_1.pdf
//
// # Overview
//
// GVSP (GigE Vision Streaming Protocol) provides the streaming channel for
// image acquisition. It operates over UDP and supports various payload types
// including images, chunk data, multi-part data, and GenDC formats.
//
// # Package Content
//
// gvsp/ provides:
//   - Receiver: Hot-path UDP frame receiver (zero-alloc per packet)
//   - Buffer pool: Lock-free/ring buffer frame management
//   - Frame: Frame construction and metadata handling
//   - Payload: Payload parsing (Image, Multi-Part, GenDC, Chunk)
//   - Resend: Packet resend logic for lost packets
//
// # Payload Types
//
// Supported payload types per GenTL v1.6:
//   - PAYLOAD_TYPE_IMAGE: Single image frame
//   - PAYLOAD_TYPE_MULTI_PART: Multi-part frames (GenTL v1.5)
//   - PAYLOAD_TYPE_GENDC: GenDC format (GenTL v1.6)
//   - PAYLOAD_TYPE_CHUNK_DATA: Chunk data (deprecated in GenTL 1.5)
//   - PAYLOAD_TYPE_JPEG, JPEG2000, H264: Compressed formats
//
// GenDC (GenICam Data Container) supports multi-part data:
//   - Color/depth/mono components
//   - Per-part format definitions (PFNC pixel formats)
//   - Chunk layout integration
//
// # Performance
//
// The Receiver is designed for zero-heap-allocation per packet:
//   - Pre-allocated ring buffers for packets
//   - Pre-allocated frame buffers based on payload size
//   - No allocations in the hot path during streaming
//
// References
//
//   - GigE Vision for Realtime MV (2010)
//   - GenICam GenTL Standard v1.6
//   - GenICam GenDC Standard v1.1
//   - GenICam Pixel Format Name Convention (PFNC)
//   - https://www.emva.org/standards-technical-documents/
package gvsp
