package gvsp

import (
	"encoding/binary"
	"fmt"
)

// ChunkPayload represents a parsed chunk data payload
type ChunkPayload struct {
	Header    ChunkHeader
	ChunkData []byte
	Chunks    []Chunk
}

// ChunkHeader is the payload type specific header for chunk data
type ChunkHeader struct {
	PayloadSize uint64
	ChunkCount  uint32
	Reserved    uint32
}

// Chunk represents one chunk entry
type Chunk struct {
	ChunkID  uint32
	Offset   uint32
	Size     uint32
	Version  uint16
	Reserved uint16
	Data     []byte
}

// Chunk ID constants (GenTL v1.2/v1.4)
const (
	ChunkIDUnknown         = 0xFFFFFFFF
	ChunkIDTimestamp       = 0x00000001
	ChunkIDFrameID         = 0x00000002
	ChunkIDOffsetX         = 0x00000003
	ChunkIDOffsetY         = 0x00000004
	ChunkIDWidth           = 0x00000005
	ChunkIDHeight          = 0x00000006
	ChunkIDPixelFormat     = 0x00000007
	ChunkIDBinningX        = 0x00000008
	ChunkIDBinningY        = 0x00000009
	ChunkIDPaddingX        = 0x0000000A
	ChunkIDPaddingY        = 0x0000000B
	ChunkIDGamma           = 0x0000000C
	ChunkIDGain            = 0x0000000D
	ChunkIDShutter         = 0x0000000E
	ChunkIDBrightness      = 0x0000000F
	ChunkIDBlackLevel      = 0x00000010
	ChunkIDTemperature     = 0x00000011
	ChunkIDExpTime         = 0x00000012
	ChunkIDAcqFrameRate    = 0x00000013
	ChunkIDLineStatus      = 0x00000014
	ChunkIDLineMode        = 0x00000015
	ChunkIDLinePulse       = 0x00000016
	ChunkIDLineSource      = 0x00000017
	ChunkIDLineTermination = 0x00000018
	ChunkIDLineFormat      = 0x00000019
	ChunkIDLineDuration    = 0x0000001A
	ChunkIDLineDelay       = 0x0000001B
	ChunkIDLineEnable      = 0x0000001C
	ChunkIDLineInversion   = 0x0000001D
	ChunkIDLinePolarity    = 0x0000001E
	ChunkIDLineValue       = 0x0000001F
)

// ParseChunkPayload parses a chunk data payload
func ParseChunkPayload(data []byte) (*ChunkPayload, error) {
	if len(data) < 16 {
		return nil, fmt.Errorf("gige: chunk payload too short")
	}

	header := ChunkHeader{
		PayloadSize: binary.BigEndian.Uint64(data[0:]),
		ChunkCount:  binary.BigEndian.Uint32(data[8:]),
		Reserved:    binary.BigEndian.Uint32(data[12:]),
	}

	if header.ChunkCount == 0 {
		return nil, fmt.Errorf("gige: chunk payload has no chunks")
	}

	// Each chunk entry is 16 bytes
	entriesSize := 16 + header.ChunkCount*16
	if int(entriesSize) > len(data) {
		return nil, fmt.Errorf("gige: chunk header truncated")
	}

	chunks := make([]Chunk, 0, header.ChunkCount)
	for i := uint32(0); i < header.ChunkCount; i++ {
		off := 16 + int(i*16)
		chunk := Chunk{
			ChunkID:  binary.BigEndian.Uint32(data[off:]),
			Offset:   binary.BigEndian.Uint32(data[off+4:]),
			Size:     binary.BigEndian.Uint32(data[off+8:]),
			Version:  binary.BigEndian.Uint16(data[off+12:]),
			Reserved: binary.BigEndian.Uint16(data[off+14:]),
		}

		// Extract data if offset and size are valid
		if chunk.Offset > 0 && chunk.Size > 0 && int(chunk.Offset+chunk.Size) <= len(data) {
			chunk.Data = make([]byte, chunk.Size)
			copy(chunk.Data, data[chunk.Offset:chunk.Offset+chunk.Size])
		}

		chunks = append(chunks, chunk)
	}

	return &ChunkPayload{
		Header:    header,
		ChunkData: data[16:],
		Chunks:    chunks,
	}, nil
}

// GetChunkByID returns the chunk with the specified ID, or false if not found
func (c *ChunkPayload) GetChunkByID(id uint32) (*Chunk, bool) {
	for i := range c.Chunks {
		if c.Chunks[i].ChunkID == id {
			return &c.Chunks[i], true
		}
	}
	return nil, false
}

// ChunkPayloadType returns the chunk data payload type constant
func ChunkPayloadType() uint32 {
	return 0x80000009 // PAYLOAD_TYPE_CHUNK_DATA per GenTL
}

// IsChunkData reports whether the payload is chunk-only data
func IsChunkData(data []byte) bool {
	// Chunk-only data has no image data, just chunk headers
	// Check for chunk header format
	if len(data) < 16 {
		return false
	}
	// Look for chunk entry markers
	payloadSize := binary.BigEndian.Uint64(data[0:])
	chunkCount := binary.BigEndian.Uint32(data[8:])
	if chunkCount > 0 && payloadSize > 16 {
		// Check if chunk entries exist
		entriesStart := 16
		if entriesStart+16*int(chunkCount) <= len(data) {
			return true
		}
	}
	return false
}
