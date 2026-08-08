package gvsp

import (
	"encoding/binary"
	"fmt"
	"github.com/aaronmurniadi/gogige/internal/genDC"
)

// GenDCPayload represents a parsed GenDC payload
type GenDCPayload struct {
	Frame      *genDC.GenDCFrame
	Components []GenDcComponent
}

// GenDcComponent represents a parsed GenDC component
type GenDcComponent struct {
	TypeName   string
	Format     uint32
	FormatName string
	Width      int
	Height     int
	Data       []byte
}

// IsGenDCPayload checks if the payload is a GenDC container
func IsGenDCPayload(data []byte) bool {
	return genDC.IsGenDC(data)
}

// ParseGenDcPayload parses a GenDC payload
func ParseGenDcPayload(data []byte) (*GenDCPayload, error) {
	frame, err := genDC.ParseGenDCContainer(data)
	if err != nil {
		return nil, err
	}

	components := make([]GenDcComponent, 0, len(frame.Components))
	for _, comp := range frame.Components {
		c := GenDcComponent{
			TypeName:   comp.TypeName,
			Format:     comp.Header.Format,
			FormatName: comp.FormatName,
		}

		// Find the first part with data
		for _, part := range comp.Parts {
			if part.DataSize > 0 && part.DataOffset > 0 {
				offset := int(part.DataOffset)
				size := int(part.DataSize)
				if offset+size <= len(data) {
					c.Data = make([]byte, size)
					copy(c.Data, data[offset:offset+size])
				}
				break
			}
		}

		// Extract dimensions from part header if available
		if len(comp.Parts) > 0 {
			p := comp.Parts[0]
			if p.Header.HeaderType >= genDC.HeaderPartMin && p.Header.HeaderType <= genDC.HeaderPartMax {
				// Check if this is a 2D part header
				if len(data)-int(p.DataOffset) >= genDC.PartHeader2DBaseSize {
					headerData := data[p.DataOffset:]
					c.Width = int(binary.LittleEndian.Uint32(headerData[32:]))
					c.Height = int(binary.LittleEndian.Uint32(headerData[36:]))
				}
			}
		}

		components = append(components, c)
	}

	return &GenDCPayload{
		Frame:      frame,
		Components: components,
	}, nil
}

// ParseGenDCPayloadAsImage returns the first component as image data
func ParseGenDCPayloadAsImage(data []byte) ([]byte, uint32, int, int, error) {
	payload, err := ParseGenDcPayload(data)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	if len(payload.Components) == 0 {
		return nil, 0, 0, 0, fmt.Errorf("genDC: no components in payload")
	}

	c := payload.Components[0]
	return c.Data, c.Format, c.Width, c.Height, nil
}

// GenDCPayloadType returns the GenDC payload type constant
func GenDCPayloadType() uint32 {
	return 0x80000008 // PAYLOAD_TYPE_GENDC per GenTL
}
