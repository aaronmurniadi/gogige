package gvsp

import (
	"encoding/binary"
	"fmt"
)

// MultiPartPayload represents a parsed multi-part payload
type MultiPartPayload struct {
	Header MultiPartHeader
	Parts  []MultiPartPart
}

// MultiPartHeader is the payload type specific header for multi-part payloads
type MultiPartHeader struct {
	NumParts uint32
	Reserved uint32
}

// MultiPartPart represents one part within a multi-part payload
type MultiPartPart struct {
	PartType    uint32
	Offset      uint64
	Size        uint64
	Height      uint32
	Width       uint32
	PixelFormat uint32
	Data        []byte
}

// Multi-part part type constants (GenTL v1.5)
const (
	MultiPartPartTypeImage         = 0x00000000
	MultiPartPartTypeChunk         = 0x00000001
	MultiPartPartTypeExtendedChunk = 0x00000002
)

// ParseMultiPartPayload parses a multi-part payload
func ParseMultiPartPayload(data []byte) (*MultiPartPayload, error) {
	if len(data) < 8 {
		return nil, fmt.Errorf("gige: multi-part payload too short")
	}

	header := MultiPartHeader{
		NumParts: binary.BigEndian.Uint32(data[0:]),
		Reserved: binary.BigEndian.Uint32(data[4:]),
	}

	if header.NumParts == 0 {
		return nil, fmt.Errorf("gige: multi-part payload has no parts")
	}

	// Each part header is 32 bytes
	headerSize := 8 + header.NumParts*32
	if int(headerSize) > len(data) {
		return nil, fmt.Errorf("gige: multi-part header truncated")
	}

	parts := make([]MultiPartPart, 0, header.NumParts)
	for i := uint32(0); i < header.NumParts; i++ {
		off := 8 + int(i*32)
		part := MultiPartPart{
			PartType:    binary.BigEndian.Uint32(data[off:]),
			Offset:      binary.BigEndian.Uint64(data[off+4:]),
			Size:        binary.BigEndian.Uint64(data[off+12:]),
			Height:      binary.BigEndian.Uint32(data[off+20:]),
			Width:       binary.BigEndian.Uint32(data[off+24:]),
			PixelFormat: binary.BigEndian.Uint32(data[off+28:]),
		}

		// Extract data if offset and size are valid
		if part.Offset > 0 && part.Size > 0 && part.Offset+part.Size <= uint64(len(data)) {
			part.Data = make([]byte, part.Size)
			copy(part.Data, data[part.Offset:part.Offset+part.Size])
		}

		parts = append(parts, part)
	}

	return &MultiPartPayload{
		Header: header,
		Parts:  parts,
	}, nil
}

// GetPartByType returns the first part with the specified type, or false if not found
func (m *MultiPartPayload) GetPartByType(partType uint32) (*MultiPartPart, bool) {
	for i := range m.Parts {
		if m.Parts[i].PartType == partType {
			return &m.Parts[i], true
		}
	}
	return nil, false
}

// MultiPartPayloadType returns the multi-part payload type constant
func MultiPartPayloadType() uint32 {
	return 0x80000007 // PAYLOAD_TYPE_MULTI_PART per GenTL
}
