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

// ChunkHeader is the payload type specific header for chunk data. There is no
// wire header in the GenDC trailing-tag format; PayloadSize and ChunkCount are
// computed while parsing.
type ChunkHeader struct {
	PayloadSize uint64
	ChunkCount  uint32
}

// Chunk represents one chunk entry
type Chunk struct {
	ChunkID uint32
	Size    uint32
	Data    []byte
}

// Chunk ID constants (GenTL v1.2/v1.4). ChunkIDUnknown doubles as the empty
// padding/sentinel marker of the GenDC trailing-tag chunk format (GenDC 1.1
// §2.2.8.1): a tag carrying this ID wraps alignment padding, not a real chunk.
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

// ParseChunkPayload parses a GenICam-chunk data blob in the GenDC trailing-tag
// format (GenDC 1.1 §2.2.8.1). A chunk blob is a sequence of tagged blocks:
//
//	[chunkData_1][tag_1] [chunkData_2][tag_2] ... [chunkData_n][tag_n]
//
// Each trailing tag is 8 bytes, little-endian: { ChunkID u32, Length u32 },
// where Length counts the bytes of the preceding chunk data only (excluding the
// tag itself) and is a multiple of 4. Tags with ChunkID 0xFFFFFFFF (ChunkIDUnknown)
// wrap empty alignment padding and carry no chunk data.
//
// Because each tag trails its own data, parsing walks backwards from the end of
// the blob; gaps are detected when a Length would reach before the previous tag.
func ParseChunkPayload(data []byte) (*ChunkPayload, error) {
	chunks := make([]Chunk, 0, 4)
	pos := len(data)
	for pos >= 8 {
		tag := pos - 8
		id := binary.LittleEndian.Uint32(data[tag:])
		size := binary.LittleEndian.Uint32(data[tag+4:])
		if int(size) > tag {
			return nil, fmt.Errorf("gige: chunk %#x length %d overruns preceding data", id, size)
		}
		start := tag - int(size)
		if id == ChunkIDUnknown {
			pos = start // empty alignment/padding tag, no chunk
			continue
		}
		c := make([]byte, size)
		copy(c, data[start:tag])
		chunks = append(chunks, Chunk{ChunkID: id, Size: size, Data: c})
		pos = start
	}
	if pos != 0 {
		return nil, fmt.Errorf("gige: chunk payload has %d trailing bytes without a tag", pos)
	}

	// The backward walk produces the chunks in reverse stream order.
	for i, j := 0, len(chunks)-1; i < j; i, j = i+1, j-1 {
		chunks[i], chunks[j] = chunks[j], chunks[i]
	}

	return &ChunkPayload{
		Header: ChunkHeader{
			PayloadSize: uint64(len(data)),
			ChunkCount:  uint32(len(chunks)),
		},
		ChunkData: data,
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
	return PayloadTypeChunkData // 0x00000004 per GenTL 1.4
}

// IsChunkData reports whether the data looks like a GenICam-chunk blob in the
// GenDC trailing-tag format: a chain of little-endian tags that parses cleanly
// back to the start of the buffer.
func IsChunkData(data []byte) bool {
	pos := len(data)
	for pos >= 8 {
		tag := pos - 8
		size := int(binary.LittleEndian.Uint32(data[tag+4:]))
		if size > tag {
			return false
		}
		pos = tag - size
	}
	return pos == 0
}
