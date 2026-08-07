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

// ImageKind is Huaray/Dahua BSCF Image::ImageType (Frame.h iota).
type ImageKind uint32

const (
	ImageUnknown ImageKind = 0
	ImageMono    ImageKind = 1 // gray / mono sensor
	ImageDepth   ImageKind = 5
	ImageColor   ImageKind = 6
)

// ParseImageKind maps names used by CLIs / options ("color", "depth", "mono").
func ParseImageKind(s string) (ImageKind, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "", "color", "colour":
		return ImageColor, nil
	case "depth":
		return ImageDepth, nil
	case "mono", "gray", "grey":
		return ImageMono, nil
	default:
		return 0, fmt.Errorf("gige: unknown image kind %q (want color|depth|mono)", s)
	}
}

func (k ImageKind) String() string {
	switch k {
	case ImageMono:
		return "mono"
	case ImageDepth:
		return "depth"
	case ImageColor:
		return "color"
	case ImageUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("image(%d)", uint32(k))
	}
}

// ImageBlock is one BSCF image payload (mono / depth / colour / …).
type ImageBlock struct {
	Kind        ImageKind
	Data        []byte
	Width       int
	Height      int
	PixelFormat uint32
}

// Sample is one grabbed frame: JPEG (filled by caller) + volume fields from BSCF.
type Sample struct {
	JPEG        []byte
	RawColor    []byte // selected image bytes (name kept for API compat)
	Width       int
	Height      int
	PixelFormat uint32
	ImageKind   ImageKind
	PackCount   int
	Length      float64
	WidthMm     float64
	HeightMm    float64
	Stable      bool
}

// BSCFFrame is a parsed BSCF multi-result chunk.
type BSCFFrame struct {
	Version   int32
	Width     int32
	Height    int32
	ID        int32
	Images    []ImageBlock
	Color     []byte // convenience: ImageColor bytes (may be nil)
	ColorW    int
	ColorH    int
	ColorFmt  uint32
	PackCount int
	Packs     []PackDet
}

// Image returns the first block of kind, or false if absent.
func (f *BSCFFrame) Image(kind ImageKind) (ImageBlock, bool) {
	if f == nil {
		return ImageBlock{}, false
	}
	for _, im := range f.Images {
		if im.Kind == kind {
			return im, true
		}
	}
	return ImageBlock{}, false
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
		imgType := ImageKind(binary.LittleEndian.Uint32(buf[off+20:]))
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
			im := ImageBlock{
				Kind:        imgType,
				Data:        cp,
				Width:       bw,
				Height:      bh,
				PixelFormat: imgFmt,
			}
			f.Images = append(f.Images, im)
			if imgType == ImageColor {
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

// SampleFromBSCF extracts Sample volume fields + colour image (default).
func SampleFromBSCF(buf []byte) (Sample, error) {
	return SampleFromBSCFKind(buf, ImageColor)
}

// SampleFromBSCFKind extracts Sample for the given BSCF image kind.
// ImageColor remains the default for preview; Depth/Mono pick those blocks when present.
func SampleFromBSCFKind(buf []byte, kind ImageKind) (Sample, error) {
	if kind == ImageUnknown {
		kind = ImageColor
	}
	f, err := ParseBSCF(buf)
	if err != nil {
		return Sample{PackCount: -1, ImageKind: kind}, err
	}
	return sampleFromFrame(f, kind)
}

// SampleAllFromBSCF returns one Sample per usable image block in the BSCF
// (skips zero-size / empty placeholders). Pack metrics are copied onto each.
func SampleAllFromBSCF(buf []byte) ([]Sample, error) {
	f, err := ParseBSCF(buf)
	if err != nil {
		return nil, err
	}
	out := make([]Sample, 0, len(f.Images))
	for _, im := range f.Images {
		if im.Width <= 0 || im.Height <= 0 || len(im.Data) == 0 {
			continue
		}
		out = append(out, sampleFromImage(f, im))
	}
	if len(out) == 0 {
		return nil, errors.New("gige: bscf has no usable image blocks")
	}
	return out, nil
}

func sampleFromFrame(f *BSCFFrame, kind ImageKind) (Sample, error) {
	im, ok := f.Image(kind)
	if !ok {
		return Sample{PackCount: -1, ImageKind: kind}, fmt.Errorf("gige: bscf %s image empty", kind)
	}
	s := sampleFromImage(f, im)
	if len(s.RawColor) == 0 {
		return s, errors.New("gige: bscf image empty")
	}
	return s, nil
}

func sampleFromImage(f *BSCFFrame, im ImageBlock) Sample {
	s := Sample{
		RawColor:    im.Data,
		Width:       im.Width,
		Height:      im.Height,
		PixelFormat: im.PixelFormat,
		ImageKind:   im.Kind,
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
	return BuildTestBSCFImages([]ImageBlock{{
		Kind: ImageColor, Data: color, Width: w, Height: h, PixelFormat: fmtPix,
	}}, packs)
}

// BuildTestBSCFImages builds a V1 BSCF with the given image blocks + optional packs.
func BuildTestBSCFImages(images []ImageBlock, packs []PackDet) []byte {
	hdr := make([]byte, bscfHeaderV1)
	binary.LittleEndian.PutUint32(hdr[0:], bscfMagic)
	binary.LittleEndian.PutUint32(hdr[4:], 1) // version
	nBlocks := len(images)
	if len(packs) > 0 {
		nBlocks++
	}
	binary.LittleEndian.PutUint32(hdr[8:], uint32(nBlocks))
	w, h := 0, 0
	if len(images) > 0 {
		w, h = images[0].Width, images[0].Height
	}
	binary.LittleEndian.PutUint32(hdr[12:], uint32(w))
	binary.LittleEndian.PutUint32(hdr[16:], uint32(h))
	binary.LittleEndian.PutUint32(hdr[20:], 1) // id

	payload := make([]byte, 0)
	writeBlock := func(i int, dataType, offset, size uint32, bw, bh int, imgType ImageKind, imgFmt uint32, packCount int) {
		off := 24 + i*bscfBlockStride
		binary.LittleEndian.PutUint32(hdr[off:], dataType)
		binary.LittleEndian.PutUint32(hdr[off+4:], offset)
		binary.LittleEndian.PutUint32(hdr[off+8:], size)
		binary.LittleEndian.PutUint32(hdr[off+12:], uint32(bw))
		binary.LittleEndian.PutUint32(hdr[off+16:], uint32(bh))
		binary.LittleEndian.PutUint32(hdr[off+20:], uint32(imgType))
		binary.LittleEndian.PutUint32(hdr[off+32:], imgFmt)
		binary.LittleEndian.PutUint32(hdr[off+36:], uint32(packCount))
		binary.LittleEndian.PutUint32(hdr[off+40:], 0)
	}
	off := uint32(bscfHeaderV1)
	for i, im := range images {
		writeBlock(i, blockTypeImage, off, uint32(len(im.Data)), im.Width, im.Height, im.Kind, im.PixelFormat, 0)
		payload = append(payload, im.Data...)
		off += uint32(len(im.Data))
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
		writeBlock(len(images), blockTypePackLocResult, packOff, uint32(len(packs)*packDetSize), 0, 0, 0, 0, len(packs))
	}
	return append(hdr, payload...)
}
