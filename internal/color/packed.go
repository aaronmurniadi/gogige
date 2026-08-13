package color

import (
	"encoding/binary"
	"fmt"
	"image"
)

// Packed format support for Mono/Bayer formats
// Supports: Mono10p, Mono12p, Mono14p, Mono16p, Bayer10p, Bayer12p, Bayer14p

// Unpack10P converts 10-bit packed data to 16-bit
func Unpack10P(src []byte, dst []uint16) {
	if len(src) < 5 || len(dst) < 1 {
		return
	}
	dstIdx := 0
	for i := 0; i+4 < len(src); i += 5 {
		dst[dstIdx] = uint16(src[i]) | uint16(src[i+1]&0x03)<<8
		dstIdx++
		if dstIdx < len(dst) {
			dst[dstIdx] = uint16(src[i+1]&0xFC>>2) | uint16(src[i+2]&0x0F)<<6
			dstIdx++
		}
		if dstIdx < len(dst) {
			dst[dstIdx] = uint16(src[i+2]&0xF0>>4) | uint16(src[i+3]&0x3F)<<4
			dstIdx++
		}
		if dstIdx < len(dst) {
			dst[dstIdx] = uint16(src[i+3]&0xC0>>6) | uint16(src[i+4])<<2
			dstIdx++
		}
	}
}

// Unpack12P converts 12-bit packed data to 16-bit
func Unpack12P(src []byte, dst []uint16) {
	if len(src) < 6 || len(dst) < 1 {
		return
	}
	dstIdx := 0
	for i := 0; i+5 < len(src); i += 6 {
		dst[dstIdx] = uint16(src[i]) | uint16(src[i+1]&0x0F)<<8
		dstIdx++
		if dstIdx < len(dst) {
			dst[dstIdx] = uint16(src[i+1]&0xF0>>4) | uint16(src[i+2])<<4
			dstIdx++
		}
		if dstIdx < len(dst) {
			dst[dstIdx] = uint16(src[i+3]) | uint16(src[i+4]&0x0F)<<8
			dstIdx++
		}
		if dstIdx < len(dst) {
			dst[dstIdx] = uint16(src[i+4]&0xF0>>4) | uint16(src[i+5])<<4
			dstIdx++
		}
	}
}

// Unpack14P converts 14-bit packed data to 16-bit
func Unpack14P(src []byte, dst []uint16) {
	if len(src) < 7 || len(dst) < 1 {
		return
	}
	dstIdx := 0
	for i := 0; i+6 < len(src); i += 7 {
		// 4 pixels per 7-byte group, LSB-packed (14 bits each).
		dst[dstIdx] = uint16(src[i]) | uint16(src[i+1]&0x3F)<<8
		dstIdx++
		if dstIdx < len(dst) {
			dst[dstIdx] = uint16(src[i+1]&0xC0>>6) | uint16(src[i+2])<<2 | uint16(src[i+3]&0x0F)<<10
			dstIdx++
		}
		if dstIdx < len(dst) {
			dst[dstIdx] = uint16(src[i+3]&0xF0>>4) | uint16(src[i+4])<<4 | uint16(src[i+5]&0x03)<<12
			dstIdx++
		}
		if dstIdx < len(dst) {
			dst[dstIdx] = uint16(src[i+5]&0xFC>>2) | uint16(src[i+6])<<6
			dstIdx++
		}
	}
}

// Debayer10P debayers 10-bit packed Bayer to RGBA
func Debayer10P(raw []byte, w, h int, pattern byte) (*image.RGBA, error) {
	need := (w*h*10 + 7) / 8
	if len(raw) < need {
		return nil, fmt.Errorf("color: 10p data too short (%d < %d)", len(raw), need)
	}

	dst := make([]uint16, w*h)
	Unpack10P(raw, dst)
	return Debayer16(dst, w, h, pattern)
}

// Debayer12P debayers 12-bit packed Bayer to RGBA
func Debayer12P(raw []byte, w, h int, pattern byte) (*image.RGBA, error) {
	need := (w*h*12 + 7) / 8
	if len(raw) < need {
		return nil, fmt.Errorf("color: 12p data too short (%d < %d)", len(raw), need)
	}

	dst := make([]uint16, w*h)
	Unpack12P(raw, dst)
	return Debayer16(dst, w, h, pattern)
}

// Debayer14P debayers 14-bit packed Bayer to RGBA
func Debayer14P(raw []byte, w, h int, pattern byte) (*image.RGBA, error) {
	need := (w*h*14 + 7) / 8
	if len(raw) < need {
		return nil, fmt.Errorf("color: 14p data too short (%d < %d)", len(raw), need)
	}

	dst := make([]uint16, w*h)
	Unpack14P(raw, dst)
	return Debayer16(dst, w, h, pattern)
}

// Debayer16 debayers 16-bit unpacked data to RGBA
func Debayer16(raw []uint16, w, h int, pattern byte) (*image.RGBA, error) {
	need := w * h
	if len(raw) < need {
		return nil, fmt.Errorf("color: 16-bit data too short (%d < %d)", len(raw), need)
	}

	switch pattern {
	case BayerPatternRGGB:
		return Debayer16RGGB(raw, w, h), nil
	case BayerPatternBGGR:
		return Debayer16BGGR(raw, w, h), nil
	case BayerPatternGBRG:
		return Debayer16GBRG(raw, w, h), nil
	case BayerPatternGRBG:
		return Debayer16GRBG(raw, w, h), nil
	default:
		return nil, fmt.Errorf("color: unknown Bayer pattern %d", pattern)
	}
}

func Debayer16RGGB(raw []uint16, w, h int) *image.RGBA {
	if w < 2 || h < 2 {
		return image.NewRGBA(image.Rect(0, 0, w, h))
	}
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			o := y*w + x
			r, g, b := debayer16RGGBPixel(raw, w, h, x, y, o)
			rgba.Pix[o*4+0] = uint8(r)
			rgba.Pix[o*4+1] = uint8(g)
			rgba.Pix[o*4+2] = uint8(b)
			rgba.Pix[o*4+3] = 255
		}
	}
	return rgba
}

func Debayer16BGGR(raw []uint16, w, h int) *image.RGBA {
	if w < 2 || h < 2 {
		return image.NewRGBA(image.Rect(0, 0, w, h))
	}
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			o := y*w + x
			r, g, b := debayer16BGGRPixel(raw, w, h, x, y, o)
			rgba.Pix[o*4+0] = uint8(r)
			rgba.Pix[o*4+1] = uint8(g)
			rgba.Pix[o*4+2] = uint8(b)
			rgba.Pix[o*4+3] = 255
		}
	}
	return rgba
}

func Debayer16GBRG(raw []uint16, w, h int) *image.RGBA {
	if w < 2 || h < 2 {
		return image.NewRGBA(image.Rect(0, 0, w, h))
	}
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			o := y*w + x
			r, g, b := debayer16GBRGPixel(raw, w, h, x, y, o)
			rgba.Pix[o*4+0] = uint8(r)
			rgba.Pix[o*4+1] = uint8(g)
			rgba.Pix[o*4+2] = uint8(b)
			rgba.Pix[o*4+3] = 255
		}
	}
	return rgba
}

func Debayer16GRBG(raw []uint16, w, h int) *image.RGBA {
	if w < 2 || h < 2 {
		return image.NewRGBA(image.Rect(0, 0, w, h))
	}
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			o := y*w + x
			r, g, b := debayer16GRBGPixel(raw, w, h, x, y, o)
			rgba.Pix[o*4+0] = uint8(r)
			rgba.Pix[o*4+1] = uint8(g)
			rgba.Pix[o*4+2] = uint8(b)
			rgba.Pix[o*4+3] = 255
		}
	}
	return rgba
}

func debayer16RGGBPixel(raw []uint16, w, h, x, y, o int) (int, int, int) {
	r := getPixel16(raw, w, h, x, y)
	g := getGreen16RGGB(raw, w, h, x, y)
	b := getBlue16RGGB(raw, w, h, x, y)
	return r, g, b
}

func debayer16BGGRPixel(raw []uint16, w, h, x, y, o int) (int, int, int) {
	b := getPixel16(raw, w, h, x, y)
	g := getGreen16BGGR(raw, w, h, x, y)
	r := getRed16BGGR(raw, w, h, x, y)
	return r, g, b
}

func debayer16GBRGPixel(raw []uint16, w, h, x, y, o int) (int, int, int) {
	g := getGreen16GBRG(raw, w, h, x, y)
	r := getRed16GBRG(raw, w, h, x, y)
	b := getBlue16GBRG(raw, w, h, x, y)
	return r, g, b
}

func debayer16GRBGPixel(raw []uint16, w, h, x, y, o int) (int, int, int) {
	g := getGreen16GRBG(raw, w, h, x, y)
	b := getBlue16GRBG(raw, w, h, x, y)
	r := getRed16GRBG(raw, w, h, x, y)
	return r, g, b
}

func getPixel16(raw []uint16, w, h, x, y int) int {
	if x < 0 || x >= w || y < 0 || y >= h {
		return 0
	}
	return int(raw[y*w+x])
}

func getGreen16RGGB(raw []uint16, w, h, x, y int) int {
	g := 0
	count := 0
	if x > 0 {
		g += int(raw[y*w+(x-1)])
		count++
	}
	if x < w-1 {
		g += int(raw[y*w+(x+1)])
		count++
	}
	if y > 0 {
		g += int(raw[(y-1)*w+x])
		count++
	}
	if y < h-1 {
		g += int(raw[(y+1)*w+x])
		count++
	}
	if count > 0 {
		return g / count
	}
	return 0
}

func getBlue16RGGB(raw []uint16, w, h, x, y int) int {
	b := 0
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx%2 == 0 && ny%2 == 0 {
				b += getPixel16(raw, w, h, nx, ny)
				count++
			}
		}
	}
	if count > 0 {
		return b / count
	}
	return 0
}

func getGreen16BGGR(raw []uint16, w, h, x, y int) int {
	g := 0
	count := 0
	if x > 0 {
		g += int(raw[y*w+(x-1)])
		count++
	}
	if x < w-1 {
		g += int(raw[y*w+(x+1)])
		count++
	}
	if y > 0 {
		g += int(raw[(y-1)*w+x])
		count++
	}
	if y < h-1 {
		g += int(raw[(y+1)*w+x])
		count++
	}
	if count > 0 {
		return g / count
	}
	return 0
}

func getRed16BGGR(raw []uint16, w, h, x, y int) int {
	r := 0
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx%2 == 1 && ny%2 == 1 {
				r += getPixel16(raw, w, h, nx, ny)
				count++
			}
		}
	}
	if count > 0 {
		return r / count
	}
	return 0
}

func getGreen16GBRG(raw []uint16, w, h, x, y int) int {
	g := 0
	count := 0
	if x > 0 {
		g += int(raw[y*w+(x-1)])
		count++
	}
	if x < w-1 {
		g += int(raw[y*w+(x+1)])
		count++
	}
	if y > 0 {
		g += int(raw[(y-1)*w+x])
		count++
	}
	if y < h-1 {
		g += int(raw[(y+1)*w+x])
		count++
	}
	if count > 0 {
		return g / count
	}
	return 0
}

func getRed16GBRG(raw []uint16, w, h, x, y int) int {
	r := 0
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx%2 == 1 && ny%2 == 1 {
				r += getPixel16(raw, w, h, nx, ny)
				count++
			}
		}
	}
	if count > 0 {
		return r / count
	}
	return 0
}

func getBlue16GBRG(raw []uint16, w, h, x, y int) int {
	b := 0
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx%2 == 0 && ny%2 == 1 {
				b += getPixel16(raw, w, h, nx, ny)
				count++
			}
		}
	}
	if count > 0 {
		return b / count
	}
	return 0
}

func getGreen16GRBG(raw []uint16, w, h, x, y int) int {
	g := 0
	count := 0
	if x > 0 {
		g += int(raw[y*w+(x-1)])
		count++
	}
	if x < w-1 {
		g += int(raw[y*w+(x+1)])
		count++
	}
	if y > 0 {
		g += int(raw[(y-1)*w+x])
		count++
	}
	if y < h-1 {
		g += int(raw[(y+1)*w+x])
		count++
	}
	if count > 0 {
		return g / count
	}
	return 0
}

func getBlue16GRBG(raw []uint16, w, h, x, y int) int {
	b := 0
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx%2 == 0 && ny%2 == 1 {
				b += getPixel16(raw, w, h, nx, ny)
				count++
			}
		}
	}
	if count > 0 {
		return b / count
	}
	return 0
}

func getRed16GRBG(raw []uint16, w, h, x, y int) int {
	r := 0
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx%2 == 1 && ny%2 == 0 {
				r += getPixel16(raw, w, h, nx, ny)
				count++
			}
		}
	}
	if count > 0 {
		return r / count
	}
	return 0
}

// shiftTo16 left-aligns LSB-packed sample values into the top bits of a 16-bit
// word so the display path (which previews via the high byte) shows them scaled.
func shiftTo16(dst []uint16, bits int) {
	if bits >= 16 {
		return
	}
	sh := 16 - bits
	for i := range dst {
		dst[i] <<= sh
	}
}

// bayer16PFNC maps high bit-depth PFNC Bayer IDs to their Bayer pattern.
var bayer16PFNC = map[uint32]byte{
	PixelFormatBayerRG10: BayerPatternRGGB,
	PixelFormatBayerRG12: BayerPatternRGGB,
	PixelFormatBayerRG14: BayerPatternRGGB,
	PixelFormatBayerRG16: BayerPatternRGGB,
	PixelFormatBayerGR10: BayerPatternGRBG,
	PixelFormatBayerGR12: BayerPatternGRBG,
	PixelFormatBayerGR14: BayerPatternGRBG,
	PixelFormatBayerGR16: BayerPatternGRBG,
	PixelFormatBayerGB10: BayerPatternGBRG,
	PixelFormatBayerGB12: BayerPatternGBRG,
	PixelFormatBayerGB14: BayerPatternGBRG,
	PixelFormatBayerGB16: BayerPatternGBRG,
	PixelFormatBayerBG10: BayerPatternBGGR,
	PixelFormatBayerBG12: BayerPatternBGGR,
	PixelFormatBayerBG14: BayerPatternBGGR,
	PixelFormatBayerBG16: BayerPatternBGGR,

	PixelFormatBayerRG10p: BayerPatternRGGB,
	PixelFormatBayerRG12p: BayerPatternRGGB,
	PixelFormatBayerRG14p: BayerPatternRGGB,
	PixelFormatBayerGR10p: BayerPatternGRBG,
	PixelFormatBayerGR12p: BayerPatternGRBG,
	PixelFormatBayerGR14p: BayerPatternGRBG,
	PixelFormatBayerGB10p: BayerPatternGBRG,
	PixelFormatBayerGB12p: BayerPatternGBRG,
	PixelFormatBayerGB14p: BayerPatternGBRG,
	PixelFormatBayerBG10p: BayerPatternBGGR,
	PixelFormatBayerBG12p: BayerPatternBGGR,
	PixelFormatBayerBG14p: BayerPatternBGGR,
}

// bayerBits returns the nominal bit depth for a Bayer PFNC ID (packed or not).
func bayerBits(pf uint32) (bits int, packed bool, ok bool) {
	switch pf {
	case PixelFormatBayerRG10, PixelFormatBayerGR10, PixelFormatBayerGB10, PixelFormatBayerBG10:
		return 10, false, true
	case PixelFormatBayerRG12, PixelFormatBayerGR12, PixelFormatBayerGB12, PixelFormatBayerBG12:
		return 12, false, true
	case PixelFormatBayerRG14, PixelFormatBayerGR14, PixelFormatBayerGB14, PixelFormatBayerBG14:
		return 14, false, true
	case PixelFormatBayerRG16, PixelFormatBayerGR16, PixelFormatBayerGB16, PixelFormatBayerBG16:
		return 16, false, true
	case PixelFormatBayerRG10p, PixelFormatBayerGR10p, PixelFormatBayerGB10p, PixelFormatBayerBG10p:
		return 10, true, true
	case PixelFormatBayerRG12p, PixelFormatBayerGR12p, PixelFormatBayerGB12p, PixelFormatBayerBG12p:
		return 12, true, true
	case PixelFormatBayerRG14p, PixelFormatBayerGR14p, PixelFormatBayerGB14p, PixelFormatBayerBG14p:
		return 14, true, true
	default:
		return 0, false, false
	}
}

// DecodeHighDepth converts high bit-depth (>=10 bit) mono/Bayer data — both
// unpacked 16-bit and packed (10p/12p/14p) forms — into RGBA. Returns
// ok=false when the format is not a high-depth format handled here.
func DecodeHighDepth(raw []byte, w, h int, pf uint32) (*image.RGBA, bool) {
	if w <= 0 || h <= 0 {
		return nil, false
	}

	// Packed / unpacked mono.
	switch pf {
	case PixelFormatMono10, PixelFormatMono12, PixelFormatMono14, PixelFormatMono16:
		if len(raw) < w*h*2 {
			return nil, false
		}
		return mono16Preview(raw, w, h), true
	case PixelFormatMono10p, PixelFormatMono12p, PixelFormatMono14p:
		bits := 10
		switch pf {
		case PixelFormatMono12p:
			bits = 12
		case PixelFormatMono14p:
			bits = 14
		}
		if len(raw) < (w*h*bits+7)/8 {
			return nil, false
		}
		dst := make([]uint16, w*h)
		switch pf {
		case PixelFormatMono10p:
			Unpack10P(raw, dst)
		case PixelFormatMono12p:
			Unpack12P(raw, dst)
		case PixelFormatMono14p:
			Unpack14P(raw, dst)
		}
		shiftTo16(dst, bits)
		return mono16Preview16(dst, w, h), true
	}

	// Bayer (packed and unpacked).
	pattern, ok := bayer16PFNC[pf]
	if !ok {
		return nil, false
	}
	bits, packed, _ := bayerBits(pf)
	var dst []uint16
	if packed {
		if len(raw) < (w*h*bits+7)/8 {
			return nil, false
		}
		dst = make([]uint16, w*h)
		switch {
		case bits == 10:
			Unpack10P(raw, dst)
		case bits == 12:
			Unpack12P(raw, dst)
		case bits == 14:
			Unpack14P(raw, dst)
		}
		shiftTo16(dst, bits)
	} else {
		if len(raw) < w*h*2 {
			return nil, false
		}
		dst = make([]uint16, w*h)
		for i := 0; i < w*h; i++ {
			dst[i] = binary.LittleEndian.Uint16(raw[i*2:])
		}
		shiftTo16(dst, bits)
	}
	img, err := Debayer16(dst, w, h, pattern)
	if err != nil {
		return nil, false
	}
	return img, true
}

// mono16Preview16 renders LSB-aligned 16-bit mono data using its high byte.
func mono16Preview16(raw []uint16, w, h int) *image.RGBA {
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	for i, o := 0, 0; i < w*h; i, o = i+1, o+4 {
		v := uint8(raw[i] >> 8)
		rgba.Pix[o+0] = v
		rgba.Pix[o+1] = v
		rgba.Pix[o+2] = v
		rgba.Pix[o+3] = 255
	}
	return rgba
}
