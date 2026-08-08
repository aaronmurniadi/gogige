package gvsp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strings"
)

const (
	bscfMagic       = 0x46435342 // 'BSCF' LE
	bscfHeaderV0    = 536
	bscfHeaderV1    = 576
	bscfBlockStride = 64
	bscfMaxBlocks   = 8

	blockTypeImage         = 0
	blockTypePackLocResult = 1
	blockTypeExtend        = 2

	packDetSize = 1472
)

// Component is a GenICam/SFNC-style imaging component selector.
// Wire values follow Huaray/Dahua BSCF Image::ImageType (Frame.h iota);
// Color/Mono ≈ Intensity, Depth ≈ Range in GenDC.
type Component uint32

const (
	ComponentUnknown Component = 0
	ComponentMono    Component = 1 // gray / mono sensor
	ComponentDepth   Component = 5 // GenDC Range
	ComponentColor   Component = 6
)

// ParseComponent maps names used by CLIs / options ("color", "depth", "mono").
func ParseComponent(s string) (Component, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "", "color", "colour":
		return ComponentColor, nil
	case "depth", "range":
		return ComponentDepth, nil
	case "mono", "gray", "grey", "intensity":
		return ComponentMono, nil
	default:
		return 0, fmt.Errorf("gige: unknown component %q (want color|depth|mono)", s)
	}
}

func (c Component) String() string {
	switch c {
	case ComponentMono:
		return "mono"
	case ComponentDepth:
		return "depth"
	case ComponentColor:
		return "color"
	case ComponentUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("component(%d)", uint32(c))
	}
}

// ComponentBlock is one BSCF component payload (mono / depth / colour / …).
type ComponentBlock struct {
	Component   Component
	Data        []byte
	Width       int
	Height      int
	PixelFormat uint32
}

// Sample is one grabbed frame: JPEG (filled by caller) + volume fields from BSCF.
type Sample struct {
	JPEG        []byte
	RawColor    []byte // selected component bytes (name kept for API compat)
	Width       int
	Height      int
	PixelFormat uint32
	Component   Component
	PackCount   int
	Length      float64
	WidthMm     float64
	HeightMm    float64
	Stable      bool
}

// BSCFFrame is a parsed BSCF multi-result chunk.
type BSCFFrame struct {
	Version    int32
	Width      int32
	Height     int32
	ID         int32
	Components []ComponentBlock
	Color      []byte // convenience: ComponentColor bytes (may be nil)
	ColorW     int
	ColorH     int
	ColorFmt   uint32
	PackCount  int
	Packs      []PackDet
}

// Block returns the first block for component c, or false if absent.
func (f *BSCFFrame) Block(c Component) (ComponentBlock, bool) {
	if f == nil {
		return ComponentBlock{}, false
	}
	for _, b := range f.Components {
		if b.Component == c {
			return b, true
		}
	}
	return ComponentBlock{}, false
}

// PackDet is PackDetVolumeInfo / ScVolumePackInfo fields we care about.
type PackDet struct {
	Length float32
	Width  float32
	Height float32
	Volume float32
	Stable bool
}

// ParseBSCF parses a BSCF V0/V1 buffer.
func ParseBSCF(buf []byte) (*BSCFFrame, error) {
	if len(buf) < bscfHeaderV0 {
		return nil, fmt.Errorf("gige: bscf too short (%d)", len(buf))
	}
	magic := binary.LittleEndian.Uint32(buf[0:])
	if magic != bscfMagic {
		return nil, fmt.Errorf("gige: bscf magic 0x%08x", magic)
	}
	version := int32(binary.LittleEndian.Uint32(buf[4:]))
	need := bscfHeaderV0
	if version == 1 {
		need = bscfHeaderV1
	}
	if len(buf) < need {
		return nil, fmt.Errorf("gige: bscf v%d size %d < %d", version, len(buf), need)
	}
	dataNum := int32(binary.LittleEndian.Uint32(buf[8:]))
	f := &BSCFFrame{
		Version: version,
		Width:   int32(binary.LittleEndian.Uint32(buf[12:])),
		Height:  int32(binary.LittleEndian.Uint32(buf[16:])),
		ID:      int32(binary.LittleEndian.Uint32(buf[20:])),
	}
	_ = dataNum
	for i := 0; i < bscfMaxBlocks; i++ {
		off := 24 + i*bscfBlockStride
		if off+bscfBlockStride > len(buf) {
			break
		}
		dataType := binary.LittleEndian.Uint32(buf[off:])
		if dataType > blockTypeExtend {
			continue
		}
		blockOff := binary.LittleEndian.Uint32(buf[off+4:])
		blockSize := binary.LittleEndian.Uint32(buf[off+8:])
		bw := int(binary.LittleEndian.Uint32(buf[off+12:]))
		bh := int(binary.LittleEndian.Uint32(buf[off+16:]))
		comp := Component(binary.LittleEndian.Uint32(buf[off+20:]))
		imgFmt := binary.LittleEndian.Uint32(buf[off+32:])
		// PackLocResult: Huaray stores pack count at +36; +40 is unused/0 on DS5131.
		packCount := int(binary.LittleEndian.Uint32(buf[off+36:]))
		if packCount == 0 {
			packCount = int(binary.LittleEndian.Uint32(buf[off+40:]))
		}

		if blockOff == 0 && blockSize == 0 {
			continue
		}
		end := int(blockOff + blockSize)
		if int(blockOff) > len(buf) || end > len(buf) {
			continue
		}
		payload := buf[blockOff:end]

		switch dataType {
		case blockTypeImage:
			if bw <= 0 || bh <= 0 || len(payload) == 0 {
				continue
			}
			cp := make([]byte, len(payload))
			copy(cp, payload)
			blk := ComponentBlock{
				Component:   comp,
				Data:        cp,
				Width:       bw,
				Height:      bh,
				PixelFormat: imgFmt,
			}
			f.Components = append(f.Components, blk)
			if comp == ComponentColor {
				f.Color = cp
				f.ColorW = bw
				f.ColorH = bh
				f.ColorFmt = imgFmt
			}
		case blockTypePackLocResult:
			f.PackCount = packCount
			n := packCount
			if n < 0 {
				n = 0
			}
			maxPacks := len(payload) / packDetSize
			if n > maxPacks {
				n = maxPacks
			}
			for p := 0; p < n; p++ {
				base := p * packDetSize
				pd := parsePackDet(payload[base : base+packDetSize])
				f.Packs = append(f.Packs, pd)
			}
			if f.PackCount < 0 {
				f.PackCount = len(f.Packs)
			}
		}
	}
	return f, nil
}

func parsePackDet(b []byte) PackDet {
	// length@48, width@52, height@56, volume@60, isStable@1276
	pd := PackDet{
		Length: math.Float32frombits(binary.LittleEndian.Uint32(b[48:])),
		Width:  math.Float32frombits(binary.LittleEndian.Uint32(b[52:])),
		Height: math.Float32frombits(binary.LittleEndian.Uint32(b[56:])),
		Volume: math.Float32frombits(binary.LittleEndian.Uint32(b[60:])),
	}
	if len(b) >= 1280 {
		pd.Stable = int32(binary.LittleEndian.Uint32(b[1276:])) != 0
	}
	return pd
}

// SampleFromBSCF extracts Sample volume fields + colour component (default).
func SampleFromBSCF(buf []byte) (Sample, error) {
	return SampleFromBSCFComponent(buf, ComponentColor)
}

// SampleFromBSCFComponent extracts Sample for the given BSCF/SFNC component.
// ComponentColor is the default preview; Depth/Mono select those blocks when present.
func SampleFromBSCFComponent(buf []byte, c Component) (Sample, error) {
	if c == ComponentUnknown {
		c = ComponentColor
	}
	f, err := ParseBSCF(buf)
	if err != nil {
		return Sample{PackCount: -1, Component: c}, err
	}
	return sampleFromFrame(f, c)
}

// SampleAllFromBSCF returns one Sample per usable component block in the BSCF
// (skips zero-size / empty placeholders). Pack metrics are copied onto each.
func SampleAllFromBSCF(buf []byte) ([]Sample, error) {
	f, err := ParseBSCF(buf)
	if err != nil {
		return nil, err
	}
	out := make([]Sample, 0, len(f.Components))
	for _, blk := range f.Components {
		if blk.Width <= 0 || blk.Height <= 0 || len(blk.Data) == 0 {
			continue
		}
		out = append(out, sampleFromBlock(f, blk))
	}
	if len(out) == 0 {
		return nil, errors.New("gige: bscf has no usable component blocks")
	}
	return out, nil
}

func sampleFromFrame(f *BSCFFrame, c Component) (Sample, error) {
	blk, ok := f.Block(c)
	if !ok {
		return Sample{PackCount: -1, Component: c}, fmt.Errorf("gige: bscf %s component empty", c)
	}
	s := sampleFromBlock(f, blk)
	if len(s.RawColor) == 0 {
		return s, errors.New("gige: bscf component empty")
	}
	return s, nil
}

func sampleFromBlock(f *BSCFFrame, blk ComponentBlock) Sample {
	s := Sample{
		RawColor:    blk.Data,
		Width:       blk.Width,
		Height:      blk.Height,
		PixelFormat: blk.PixelFormat,
		Component:   blk.Component,
		PackCount:   f.PackCount,
	}
	if len(f.Packs) > 0 {
		p := f.Packs[0]
		s.Length = float64(p.Length)
		s.WidthMm = float64(p.Width)
		s.HeightMm = float64(p.Height)
		s.Stable = p.Stable
	}
	if s.PackCount < 0 && len(f.Packs) > 0 {
		s.PackCount = len(f.Packs)
	}
	return s
}

// IsBSCF reports whether buf starts with a BSCF magic.
func IsBSCF(buf []byte) bool {
	return len(buf) >= 4 && binary.LittleEndian.Uint32(buf) == bscfMagic
}

// BuildTestBSCF builds a minimal V1 BSCF for tests (colour + optional packs).
func BuildTestBSCF(color []byte, w, h int, fmtPix uint32, packs []PackDet) []byte {
	return BuildTestBSCFComponents([]ComponentBlock{{
		Component: ComponentColor, Data: color, Width: w, Height: h, PixelFormat: fmtPix,
	}}, packs)
}

// BuildTestBSCFComponents builds a V1 BSCF with the given component blocks + optional packs.
func BuildTestBSCFComponents(blocks []ComponentBlock, packs []PackDet) []byte {
	hdr := make([]byte, bscfHeaderV1)
	binary.LittleEndian.PutUint32(hdr[0:], bscfMagic)
	binary.LittleEndian.PutUint32(hdr[4:], 1) // version
	nBlocks := len(blocks)
	if len(packs) > 0 {
		nBlocks++
	}
	binary.LittleEndian.PutUint32(hdr[8:], uint32(nBlocks))
	w, h := 0, 0
	if len(blocks) > 0 {
		w, h = blocks[0].Width, blocks[0].Height
	}
	binary.LittleEndian.PutUint32(hdr[12:], uint32(w))
	binary.LittleEndian.PutUint32(hdr[16:], uint32(h))
	binary.LittleEndian.PutUint32(hdr[20:], 1) // id

	payload := make([]byte, 0)
	writeBlock := func(i int, dataType, offset, size uint32, bw, bh int, comp Component, imgFmt uint32, packCount int) {
		off := 24 + i*bscfBlockStride
		binary.LittleEndian.PutUint32(hdr[off:], dataType)
		binary.LittleEndian.PutUint32(hdr[off+4:], offset)
		binary.LittleEndian.PutUint32(hdr[off+8:], size)
		binary.LittleEndian.PutUint32(hdr[off+12:], uint32(bw))
		binary.LittleEndian.PutUint32(hdr[off+16:], uint32(bh))
		binary.LittleEndian.PutUint32(hdr[off+20:], uint32(comp))
		binary.LittleEndian.PutUint32(hdr[off+32:], imgFmt)
		binary.LittleEndian.PutUint32(hdr[off+36:], uint32(packCount))
		binary.LittleEndian.PutUint32(hdr[off+40:], 0)
	}
	off := uint32(bscfHeaderV1)
	for i, blk := range blocks {
		writeBlock(i, blockTypeImage, off, uint32(len(blk.Data)), blk.Width, blk.Height, blk.Component, blk.PixelFormat, 0)
		payload = append(payload, blk.Data...)
		off += uint32(len(blk.Data))
	}
	if len(packs) > 0 {
		packOff := off
		for _, p := range packs {
			pd := make([]byte, packDetSize)
			binary.LittleEndian.PutUint32(pd[48:], math.Float32bits(p.Length))
			binary.LittleEndian.PutUint32(pd[52:], math.Float32bits(p.Width))
			binary.LittleEndian.PutUint32(pd[56:], math.Float32bits(p.Height))
			binary.LittleEndian.PutUint32(pd[60:], math.Float32bits(p.Volume))
			st := uint32(0)
			if p.Stable {
				st = 1
			}
			binary.LittleEndian.PutUint32(pd[1276:], st)
			payload = append(payload, pd...)
		}
		writeBlock(len(blocks), blockTypePackLocResult, packOff, uint32(len(packs)*packDetSize), 0, 0, 0, 0, len(packs))
	}
	return append(hdr, payload...)
}

// ParsePayloadByType parses a payload based on its type
// Returns: data, format, width, height, isPacked, error
func ParsePayloadByType(data []byte, payloadType uint32) ([]byte, uint32, int, int, bool, error) {
	switch payloadType {
	case 0x80000008: // PAYLOAD_TYPE_GENDC
		return ParseGenDCPayload(data)
	case 0x80000007: // PAYLOAD_TYPE_MULTI_PART
		return ParseMultiPartPayloadAsImage(data)
	case 0x80000009: // PAYLOAD_TYPE_CHUNK_DATA
		return ParseChunkPayloadAsImage(data)
	default:
		// Default to treating as raw image data
		return data, 0, 0, 0, false, nil
	}
}

// ParseGenDCPayload parses GenDC payload and returns image data
func ParseGenDCPayload(data []byte) ([]byte, uint32, int, int, bool, error) {
	if !IsGenDCPayload(data) {
		return nil, 0, 0, 0, false, fmt.Errorf("gige: not a GenDC payload")
	}
	payload, err := ParseGenDcPayload(data)
	if err != nil {
		return nil, 0, 0, 0, false, err
	}
	if len(payload.Components) == 0 {
		return nil, 0, 0, 0, false, fmt.Errorf("gige: no components in GenDC payload")
	}
	c := payload.Components[0]
	return c.Data, c.Format, c.Width, c.Height, false, nil
}

// ParseMultiPartPayloadAsImage parses multi-part payload and returns first image part
func ParseMultiPartPayloadAsImage(data []byte) ([]byte, uint32, int, int, bool, error) {
	payload, err := ParseMultiPartPayload(data)
	if err != nil {
		return nil, 0, 0, 0, false, err
	}
	if len(payload.Parts) == 0 {
		return nil, 0, 0, 0, false, fmt.Errorf("gige: no parts in multi-part payload")
	}
	// Find first image part
	for _, part := range payload.Parts {
		if part.PartType == MultiPartPartTypeImage && len(part.Data) > 0 {
			return part.Data, part.PixelFormat, int(part.Width), int(part.Height), false, nil
		}
	}
	return nil, 0, 0, 0, false, fmt.Errorf("gige: no image part in multi-part payload")
}

// ParseChunkPayloadAsImage parses chunk data payload and returns image data
func ParseChunkPayloadAsImage(data []byte) ([]byte, uint32, int, int, bool, error) {
	// Chunk data doesn't contain image data directly, only chunk metadata
	// The actual image is separate; return nil for now
	return nil, 0, 0, 0, false, fmt.Errorf("gige: chunk data payload contains no image data")
}
