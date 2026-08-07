package gvsp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
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

	// Frame.h Image::ImageType (iota): Unknown=0 … Depth=5, Color=6.
	imageTypeColor = 6

	packDetSize = 1472
)

// Sample is one grabbed frame: JPEG (filled by caller) + volume fields from BSCF.
type Sample struct {
	JPEG        []byte
	RawColor    []byte // decoded RGB/BGR planar-ish bytes before JPEG
	Width       int
	Height      int
	PixelFormat uint32
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
	Color     []byte
	ColorW    int
	ColorH    int
	ColorFmt  uint32
	PackCount int
	Packs     []PackDet
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
		imgType := binary.LittleEndian.Uint32(buf[off+20:])
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
			// Prefer Color; never keep Depth/IR as the preview buffer.
			if imgType != imageTypeColor {
				continue
			}
			cp := make([]byte, len(payload))
			copy(cp, payload)
			f.Color = cp
			f.ColorW = bw
			f.ColorH = bh
			f.ColorFmt = imgFmt
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

// SampleFromBSCF extracts Sample volume fields + raw color from a BSCF buffer.
func SampleFromBSCF(buf []byte) (Sample, error) {
	s := Sample{PackCount: -1}
	f, err := ParseBSCF(buf)
	if err != nil {
		return s, err
	}
	s.RawColor = f.Color
	s.Width = f.ColorW
	s.Height = f.ColorH
	s.PixelFormat = f.ColorFmt
	s.PackCount = f.PackCount
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
	if len(f.Color) == 0 {
		return s, errors.New("gige: bscf color image empty")
	}
	return s, nil
}

// BuildTestBSCF builds a minimal V1 BSCF for tests.
func BuildTestBSCF(color []byte, w, h int, fmtPix uint32, packs []PackDet) []byte {
	hdr := make([]byte, bscfHeaderV1)
	binary.LittleEndian.PutUint32(hdr[0:], bscfMagic)
	binary.LittleEndian.PutUint32(hdr[4:], 1) // version
	nBlocks := 1
	if len(packs) > 0 {
		nBlocks = 2
	}
	binary.LittleEndian.PutUint32(hdr[8:], uint32(nBlocks))
	binary.LittleEndian.PutUint32(hdr[12:], uint32(w))
	binary.LittleEndian.PutUint32(hdr[16:], uint32(h))
	binary.LittleEndian.PutUint32(hdr[20:], 1) // id

	payload := make([]byte, 0, len(color)+len(packs)*packDetSize)
	// color block @ header size
	colorOff := uint32(bscfHeaderV1)
	payload = append(payload, color...)
	writeBlock := func(i int, dataType, offset, size uint32, bw, bh int, imgType, imgFmt uint32, packCount int) {
		off := 24 + i*bscfBlockStride
		binary.LittleEndian.PutUint32(hdr[off:], dataType)
		binary.LittleEndian.PutUint32(hdr[off+4:], offset)
		binary.LittleEndian.PutUint32(hdr[off+8:], size)
		binary.LittleEndian.PutUint32(hdr[off+12:], uint32(bw))
		binary.LittleEndian.PutUint32(hdr[off+16:], uint32(bh))
		binary.LittleEndian.PutUint32(hdr[off+20:], imgType)
		binary.LittleEndian.PutUint32(hdr[off+32:], imgFmt)
		binary.LittleEndian.PutUint32(hdr[off+36:], uint32(packCount))
		binary.LittleEndian.PutUint32(hdr[off+40:], 0)
	}
	writeBlock(0, blockTypeImage, colorOff, uint32(len(color)), w, h, imageTypeColor, fmtPix, 0)

	if len(packs) > 0 {
		packOff := colorOff + uint32(len(color))
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
		writeBlock(1, blockTypePackLocResult, packOff, uint32(len(packs)*packDetSize), 0, 0, 0, 0, len(packs))
	}
	return append(hdr, payload...)
}
