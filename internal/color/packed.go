package color

import (
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
		dst[dstIdx] = uint16(src[i]) | uint16(src[i+1]&0x3F)<<8
		dstIdx++
		if dstIdx < len(dst) {
			dst[dstIdx] = uint16(src[i+1]&0xC0>>6) | uint16(src[i+2])<<2 | uint16(src[i+3]&0x03)<<10
			dstIdx++
		}
		if dstIdx < len(dst) {
			dst[dstIdx] = uint16(src[i+3]&0xFC>>2) | uint16(src[i+4]&0x0F)<<8
			dstIdx++
		}
		if dstIdx < len(dst) {
			dst[dstIdx] = uint16(src[i+4]&0xF0>>4) | uint16(src[i+5])<<4 | uint16(src[i+6]&0x03)<<12
			dstIdx++
		}
		if dstIdx < len(dst) {
			dst[dstIdx] = uint16(src[i+6]&0xFC>>2) | uint16(src[i+7])<<6
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
