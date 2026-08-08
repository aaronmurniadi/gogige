package genDC

import (
	"encoding/binary"
	"fmt"
)

// GenDC constants from GenDC.h
const (
	Signature       = 0x43444E47 // "GNDC" LE
	VersionMajor    = 0x01
	VersionMinor    = 0x01
	VersionSubMinor = 0x00
)

// Header types
const (
	HeaderTypeMask  = 0xF000
	HeaderUndefined = 0x0000
	HeaderContainer = 0x1000
	HeaderComponent = 0x2000
	HeaderPartMin   = 0x4000
	HeaderPartMax   = 0x4FFF
	HeaderPart      = 0x4000
	HeaderFlowTable = 0x7000
)

// Component types (from SFNC)
const (
	ComponentUndefined     = 0x00
	ComponentIntensity     = 0x01
	ComponentInfrared      = 0x02
	ComponentUltraviolet   = 0x03
	ComponentRange         = 0x04 // Depth
	ComponentReflectance   = 0x05
	ComponentConfidence    = 0x06
	ComponentScatter       = 0x07
	ComponentDisparity     = 0x08
	ComponentMultispectral = 0x09
	ComponentExtended      = 0x8000
	ComponentMetadata      = ComponentExtended + 0x0001
	ComponentCustomMin     = ComponentExtended + 0x7F00
	ComponentCustomMax     = ComponentCustomMin + 0x00FE
	ComponentCustom        = ComponentCustomMin + 0x0000
	ComponentReserved      = ComponentExtended + 0x7FFF
)

// Part types
const (
	PartUndefined            = 0x0000
	PartGenericMetadata      = 0x4000
	PartGeneric1D            = 0x4100
	PartGeneric2D            = 0x4200
	PartGenericCustom        = 0x4F00
	PartMetadataGenICamChunk = PartGenericMetadata + 0x00
	PartMetadataGenICamXML   = PartGenericMetadata + 0x01
	PartMetadataCustomBase   = PartGenericMetadata + 0xF0
	Part1D                   = PartGeneric1D + 0x00
	Part1DCustomBase         = PartGeneric1D + 0xF0
	Part2D                   = PartGeneric2D + 0x00
	Part2DJPEG               = PartGeneric2D + 0x01
	Part2DJPEG2000           = PartGeneric2D + 0x02
	Part2DH264               = PartGeneric2D + 0x03
	Part2DCustomBase         = PartGeneric2D + 0xF0
	PartCustomBase           = 0x4F00
)

// PFNC pixel format IDs
const (
	PixelFormatMono1p        = 0x01010037
	PixelFormatMono2p        = 0x01020038
	PixelFormatMono4p        = 0x01040039
	PixelFormatMono8         = 0x01080001
	PixelFormatMono8s        = 0x01080002
	PixelFormatMono10        = 0x01100003
	PixelFormatMono10p       = 0x010A0046
	PixelFormatMono12        = 0x01100005
	PixelFormatMono12p       = 0x010C0047
	PixelFormatMono14        = 0x01100025
	PixelFormatMono14p       = 0x010E0104
	PixelFormatMono16        = 0x01100007
	PixelFormatMono32        = 0x01200111
	PixelFormatBayerBG4p     = 0x01040110
	PixelFormatBayerBG8      = 0x0108000B
	PixelFormatBayerBG10     = 0x0110000F
	PixelFormatBayerBG10p    = 0x010A0052
	PixelFormatBayerBG12     = 0x01100013
	PixelFormatBayerBG12p    = 0x010C0053
	PixelFormatBayerBG14     = 0x0110010C
	PixelFormatBayerBG14p    = 0x010E0108
	PixelFormatBayerBG16     = 0x01100031
	PixelFormatBayerGB4p     = 0x0104010F
	PixelFormatBayerGB8      = 0x0108000A
	PixelFormatBayerGB10     = 0x0110000E
	PixelFormatBayerGB10p    = 0x010A0054
	PixelFormatBayerGB12     = 0x01100012
	PixelFormatBayerGB12p    = 0x010C0055
	PixelFormatBayerGB14     = 0x0110010B
	PixelFormatBayerGB14p    = 0x010E0107
	PixelFormatBayerGB16     = 0x01100030
	PixelFormatBayerGR4p     = 0x0104010D
	PixelFormatBayerGR8      = 0x01080008
	PixelFormatBayerGR10     = 0x0110000C
	PixelFormatBayerGR10p    = 0x010A0056
	PixelFormatBayerGR12     = 0x01100010
	PixelFormatBayerGR12p    = 0x010C0057
	PixelFormatBayerGR14     = 0x01100109
	PixelFormatBayerGR14p    = 0x010E0105
	PixelFormatBayerGR16     = 0x0110002E
	PixelFormatBayerRG4p     = 0x0104010E
	PixelFormatBayerRG8      = 0x01080009
	PixelFormatBayerRG10     = 0x0110000D
	PixelFormatBayerRG10p    = 0x010A0058
	PixelFormatBayerRG12     = 0x01100011
	PixelFormatBayerRG12p    = 0x010C0059
	PixelFormatBayerRG14     = 0x0110010A
	PixelFormatBayerRG14p    = 0x010E0106
	PixelFormatBayerRG16     = 0x0110002F
	PixelFormatRGBa8         = 0x02200016
	PixelFormatRGBa10        = 0x0240005F
	PixelFormatRGBa10p       = 0x02280060
	PixelFormatRGBa12        = 0x02400061
	PixelFormatRGBa12p       = 0x02300062
	PixelFormatRGBa14        = 0x02400063
	PixelFormatRGBa16        = 0x02400064
	PixelFormatRGB8          = 0x02180014
	PixelFormatRGB8_Planar   = 0x02180021
	PixelFormatRGB8a32       = 0x0220012D
	PixelFormatRGB10         = 0x02300018
	PixelFormatRGB10_Planar  = 0x02300022
	PixelFormatRGB10p        = 0x021E005C
	PixelFormatRGB10p32      = 0x0220001D
	PixelFormatRGB12         = 0x0230001A
	PixelFormatRGB12_Planar  = 0x02300023
	PixelFormatRGB12p        = 0x0224005D
	PixelFormatRGB14         = 0x0230005E
	PixelFormatRGB16         = 0x02300033
	PixelFormatRGB16_Planar  = 0x02300024
	PixelFormatRGB565p       = 0x02100035
	PixelFormat_BGRa8        = 0x02200017
	PixelFormat_BGRa10       = 0x0240004C
	PixelFormat_BGRa10p      = 0x0228004D
	PixelFormat_BGRa12       = 0x0240004E
	PixelFormat_BGRa12p      = 0x0230004F
	PixelFormat_BGRa14       = 0x02400050
	PixelFormat_BGRa16       = 0x02400051
	PixelFormat_BGR8         = 0x02180015
	PixelFormat_BGR8a32      = 0x0220012E
	PixelFormat_BGR10        = 0x02300019
	PixelFormat_BGR10p       = 0x021E0048
	PixelFormat_BGR12        = 0x0230001B
	PixelFormat_BGR12p       = 0x02240049
	PixelFormat_BGR14        = 0x0230004A
	PixelFormat_BGR16        = 0x0230004B
	PixelFormat_BGR565p      = 0x02100036
	PixelFormatYUV422_8      = 0x02100032 // YUYV
	PixelFormatYUV422_8_UYVY = 0x0210001F
	PixelFormatYUV422_10     = 0x02200065
	PixelFormatYUV422_10p    = 0x02140087
	PixelFormatYUV422_12     = 0x02200066
	PixelFormatYUV422_12p    = 0x02180088
)

// GenDCContainerHeader from GenDC.h (packed struct)
const (
	ContainerHeaderBaseSize      = 64
	ContainerComponentOffsetSize = 8
)

// GenDCComponentHeader from GenDC.h (packed struct)
const (
	ComponentHeaderBaseSize = 48
	ComponentPartOffsetSize = 8
)

// GenDCPartHeaderBase from GenDC.h (packed struct)
const (
	PartHeaderBaseSize = 32
)

// GenDCPartHeader2DBase extends PartHeaderBase with SizeX, SizeY, Padding
const (
	PartHeader2DBaseSize = 44
)

// GenDCContainerHeader represents the container header
type ContainerHeader struct {
	Signature        uint32
	VersionMajor     byte
	VersionMinor     byte
	VersionSubMinor  byte
	Reserved         byte
	HeaderType       uint16
	Flags            uint16
	HeaderSize       uint32
	Id               uint64
	VariableFields   uint16
	Reserved16       uint16
	Reserved32       uint32
	DataSize         uint64
	DataOffset       int64
	DescriptorSize   uint32
	ComponentCount   uint32
	ComponentOffsets []int64
}

// GenDCComponentHeader represents a component header
type ComponentHeader struct {
	HeaderType    uint16
	Flags         uint16
	HeaderSize    uint32
	Reserved      uint16
	GroupId       uint16
	SourceId      uint16
	RegionId      uint16
	RegionOffsetX uint32
	RegionOffsetY uint32
	Timestamp     int64
	TypeId        uint64
	Format        uint32
	Reserved2     uint16
	PartCount     uint16
	PartOffsets   []int64
}

// GenDCPartHeader represents a part header
type PartHeader struct {
	HeaderType uint16
	Flags      uint16
	HeaderSize uint32
	Format     uint32
	Reserved   uint16
	FlowId     uint16
	FlowOffset int64
	DataSize   uint64
	DataOffset int64
}

// GenDCPartHeader2D extends PartHeader with 2D dimensions
type PartHeader2D struct {
	PartHeader
	SizeX        uint32
	SizeY        uint32
	PaddingX     uint16
	PaddingY     uint16
	InfoReserved uint32
}

// GenDCComponent represents a parsed component
type Component struct {
	Header     *ComponentHeader
	Parts      []Part
	TypeName   string
	FormatName string
}

// Part represents a parsed part
type Part struct {
	Header     *PartHeader
	DataOffset int64
	DataSize   uint64
}

// GenDCFrame represents a parsed GenDC container
type GenDCFrame struct {
	Container  *ContainerHeader
	Components []Component
}

// IsGenDC reports whether buf starts with a GenDC signature
func IsGenDC(buf []byte) bool {
	if len(buf) < 4 {
		return false
	}
	return binary.LittleEndian.Uint32(buf) == Signature
}

// ParseGenDCContainer parses a GenDC container from buffer
func ParseGenDCContainer(buf []byte) (*GenDCFrame, error) {
	if len(buf) < ContainerHeaderBaseSize {
		return nil, fmt.Errorf("genDC: buffer too short for container header")
	}

	ch, _, err := parseContainerHeader(buf)
	if err != nil {
		return nil, err
	}

	components := make([]Component, 0, ch.ComponentCount)
	for i := uint32(0); i < ch.ComponentCount; i++ {
		off := ch.ComponentOffsets[i]
		if int(off) < 0 || int(off) >= len(buf) {
			continue
		}
		comp, err := parseComponent(buf[int(off):], int64(len(buf)-int(off)))
		if err != nil {
			continue
		}
		components = append(components, comp)
	}

	return &GenDCFrame{
		Container:  ch,
		Components: components,
	}, nil
}

func parseContainerHeader(buf []byte) (*ContainerHeader, []int64, error) {
	if len(buf) < ContainerHeaderBaseSize {
		return nil, nil, fmt.Errorf("genDC: container header too short")
	}

	ch := &ContainerHeader{
		Signature:       binary.LittleEndian.Uint32(buf[0:]),
		VersionMajor:    buf[4],
		VersionMinor:    buf[5],
		VersionSubMinor: buf[6],
		Reserved:        buf[7],
		HeaderType:      binary.LittleEndian.Uint16(buf[8:]),
		Flags:           binary.LittleEndian.Uint16(buf[10:]),
		HeaderSize:      binary.LittleEndian.Uint32(buf[12:]),
		Id:              binary.LittleEndian.Uint64(buf[16:]),
		VariableFields:  binary.LittleEndian.Uint16(buf[24:]),
		Reserved16:      binary.LittleEndian.Uint16(buf[26:]),
		Reserved32:      binary.LittleEndian.Uint32(buf[28:]),
		DataSize:        binary.LittleEndian.Uint64(buf[32:]),
		DataOffset:      int64(binary.LittleEndian.Uint64(buf[40:])),
		DescriptorSize:  binary.LittleEndian.Uint32(buf[48:]),
		ComponentCount:  binary.LittleEndian.Uint32(buf[52:]),
	}

	numOffsets := ch.ComponentCount
	if numOffsets == 0 {
		numOffsets = 1
	}
	offsetsSize := numOffsets * 8
	if ContainerHeaderBaseSize+int(offsetsSize) > len(buf) {
		return nil, nil, fmt.Errorf("genDC: container header truncated")
	}

	offsets := make([]int64, numOffsets)
	for i := uint32(0); i < numOffsets; i++ {
		off := ContainerHeaderBaseSize + int(i*8)
		val := binary.LittleEndian.Uint64(buf[off:])
		offsets[i] = int64(val)
		ch.ComponentOffsets = append(ch.ComponentOffsets, offsets[i])
	}

	return ch, offsets, nil
}

func parseComponent(buf []byte, maxLen int64) (Component, error) {
	if len(buf) < ComponentHeaderBaseSize {
		return Component{}, fmt.Errorf("genDC: component header too short")
	}

	header := &ComponentHeader{
		HeaderType:    binary.LittleEndian.Uint16(buf[0:]),
		Flags:         binary.LittleEndian.Uint16(buf[2:]),
		HeaderSize:    binary.LittleEndian.Uint32(buf[4:]),
		Reserved:      binary.LittleEndian.Uint16(buf[8:]),
		GroupId:       binary.LittleEndian.Uint16(buf[10:]),
		SourceId:      binary.LittleEndian.Uint16(buf[12:]),
		RegionId:      binary.LittleEndian.Uint16(buf[14:]),
		RegionOffsetX: binary.LittleEndian.Uint32(buf[16:]),
		RegionOffsetY: binary.LittleEndian.Uint32(buf[20:]),
		Timestamp:     int64(binary.LittleEndian.Uint64(buf[24:])),
		TypeId:        binary.LittleEndian.Uint64(buf[32:]),
		Format:        binary.LittleEndian.Uint32(buf[40:]),
		Reserved2:     binary.LittleEndian.Uint16(buf[44:]),
		PartCount:     binary.LittleEndian.Uint16(buf[46:]),
	}

	numParts := header.PartCount
	if numParts == 0 {
		numParts = 1
	}
	offsetsSize := int(numParts * 8)
	if ComponentHeaderBaseSize+offsetsSize > len(buf) {
		return Component{}, fmt.Errorf("genDC: component header truncated")
	}

	header.PartOffsets = make([]int64, numParts)
	for i := uint16(0); i < numParts; i++ {
		off := ComponentHeaderBaseSize + int(i*8)
		val := binary.LittleEndian.Uint64(buf[off:])
		header.PartOffsets[i] = int64(val)
	}

	parts := make([]Part, 0, numParts)
	for i := uint16(0); i < numParts; i++ {
		pOff := int(header.PartOffsets[i])
		if pOff < 0 || pOff >= len(buf) {
			continue
		}
		p, err := parsePart(buf[pOff:])
		if err != nil {
			continue
		}
		p.DataOffset = header.PartOffsets[i]
		p.DataSize = p.Header.DataSize
		parts = append(parts, p)
	}

	return Component{
		Header:     header,
		Parts:      parts,
		TypeName:   ComponentTypeName(header.TypeId),
		FormatName: PixelFormatName(header.Format),
	}, nil
}

func parsePart(buf []byte) (Part, error) {
	if len(buf) < PartHeaderBaseSize {
		return Part{}, fmt.Errorf("genDC: part header too short")
	}

	header := &PartHeader{
		HeaderType: binary.LittleEndian.Uint16(buf[0:]),
		Flags:      binary.LittleEndian.Uint16(buf[2:]),
		HeaderSize: binary.LittleEndian.Uint32(buf[4:]),
		Format:     binary.LittleEndian.Uint32(buf[8:]),
		Reserved:   binary.LittleEndian.Uint16(buf[12:]),
		FlowId:     binary.LittleEndian.Uint16(buf[14:]),
		FlowOffset: int64(binary.LittleEndian.Uint64(buf[16:])),
		DataSize:   binary.LittleEndian.Uint64(buf[24:]),
		DataOffset: int64(binary.LittleEndian.Uint64(buf[32:])),
	}

	return Part{Header: header}, nil
}

// ComponentTypeName returns human-readable component type name
func ComponentTypeName(t uint64) string {
	switch t {
	case ComponentUndefined:
		return "undefined"
	case ComponentIntensity:
		return "intensity"
	case ComponentInfrared:
		return "infrared"
	case ComponentUltraviolet:
		return "ultraviolet"
	case ComponentRange:
		return "range"
	case ComponentReflectance:
		return "reflectance"
	case ComponentConfidence:
		return "confidence"
	case ComponentScatter:
		return "scatter"
	case ComponentDisparity:
		return "disparity"
	case ComponentMultispectral:
		return "multispectral"
	case ComponentMetadata:
		return "metadata"
	case ComponentCustom:
		return "custom"
	case ComponentReserved:
		return "reserved"
	default:
		return fmt.Sprintf("unknown(%d)", t)
	}
}

// PixelFormatName returns human-readable PFNC format name
func PixelFormatName(f uint32) string {
	switch f {
	case PixelFormatMono8:
		return "Mono8"
	case PixelFormatMono16:
		return "Mono16"
	case PixelFormatRGB8:
		return "RGB8"
	case PixelFormat_BGR8:
		return "BGR8"
	case PixelFormatYUV422_8:
		return "YUV422_8_YUYV"
	case PixelFormatYUV422_8_UYVY:
		return "YUV422_8_UYVY"
	default:
		return fmt.Sprintf("0x%08x", f)
	}
}
