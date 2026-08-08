// Package gentl provides GenTL (GenICam Transport Layer) definitions.
//
// # GenTL Version 1.6
//
// # Overview
//
// This package provides Go definitions for the GenICam GenTL standard, which defines
// a common interface for accessing image acquisition devices (cameras). The GenTL
// standard is part of the GenICam family of standards and works alongside:
//   - GigE Vision (GEV) for Gigabit Ethernet cameras
//   - GenCP for camera control protocol
//   - GenApi for feature access and XML description
//   - GenDC for multi-part/GenDC payloads
//
// # Package Content
//
// gentl/ defines:
//   - Error codes (GC_ERROR)
//   - Transport layer types (GEV, CL, U3V, etc.)
//   - Module names (TLSystem, TLInterface, etc.)
//   - All GenTL v1.6 enums and constants
//
// This package does NOT include GenTL producer (.cti) loading.
// GenTL producers are provided by camera manufacturers as shared libraries
// and require CGO/dlopen for dynamic loading, which is outside the scope
// of this pure-Go project.
//
// # Usage
//
// Use the constants directly:
//
//	// Get the transport layer type for GigE Vision
//	tlType := gentl.TLTypeGEV // "GEV"
//
//	// Check payload types
//	isGenDC := payloadType == gentl.PayloadTypeGenDC
//
// References
//
//   - GenICam GenTL Standard v1.6
//   - https://www.emva.org/standards-technical-resources/generic-interface-documentation/
package gentl
