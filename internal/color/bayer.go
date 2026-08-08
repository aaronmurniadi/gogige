package color

import (
	"fmt"
	"image"
)

// Bayer pattern constants
const (
	BayerPatternRGGB = 0 // RG at (0,0)
	BayerPatternBGGR = 1 // BG at (0,0)
	BayerPatternGBRG = 2 // GB at (0,0)
	BayerPatternGRBG = 3 // GR at (0,0)
)

// BayerPFNCMap maps PFNC pixel format IDs to Bayer patterns
var BayerPFNCMap = map[uint32]byte{
	0x01080009: BayerPatternRGGB, // BayerRG8
	0x01080008: BayerPatternGRBG, // BayerGR8
	0x0108000A: BayerPatternGBRG, // BayerGB8
	0x0108000B: BayerPatternBGGR, // BayerBG8
}

// IsBayerFormat returns true if the pixel format is a Bayer format
func IsBayerFormat(pixelFormat uint32) bool {
	_, ok := BayerPFNCMap[pixelFormat]
	return ok
}

// GetBayerPattern returns the Bayer pattern for a given pixel format
func GetBayerPattern(pixelFormat uint32) (byte, bool) {
	pattern, ok := BayerPFNCMap[pixelFormat]
	return pattern, ok
}

// DebayerRGGB8 debayers RGGB pattern to RGB
func DebayerRGGB8(raw []byte, w, h int) *image.RGBA {
	if w < 2 || h < 2 {
		return image.NewRGBA(image.Rect(0, 0, w, h))
	}

	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	need := w * h
	if len(raw) < need {
		return rgba
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			o := y*w + x
			r, g, b := debayerRGGBPixel(raw, w, h, x, y, o)
			rgba.Pix[o*4+0] = r
			rgba.Pix[o*4+1] = g
			rgba.Pix[o*4+2] = b
			rgba.Pix[o*4+3] = 255
		}
	}
	return rgba
}

// DebayerBGGR8 debayers BGGR pattern to RGB
func DebayerBGGR8(raw []byte, w, h int) *image.RGBA {
	if w < 2 || h < 2 {
		return image.NewRGBA(image.Rect(0, 0, w, h))
	}

	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	need := w * h
	if len(raw) < need {
		return rgba
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			o := y*w + x
			r, g, b := debayerBGGRPixel(raw, w, h, x, y, o)
			rgba.Pix[o*4+0] = r
			rgba.Pix[o*4+1] = g
			rgba.Pix[o*4+2] = b
			rgba.Pix[o*4+3] = 255
		}
	}
	return rgba
}

// DebayerGBRG8 debayers GBRG pattern to RGB
func DebayerGBRG8(raw []byte, w, h int) *image.RGBA {
	if w < 2 || h < 2 {
		return image.NewRGBA(image.Rect(0, 0, w, h))
	}

	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	need := w * h
	if len(raw) < need {
		return rgba
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			o := y*w + x
			r, g, b := debayerGBRGPixel(raw, w, h, x, y, o)
			rgba.Pix[o*4+0] = r
			rgba.Pix[o*4+1] = g
			rgba.Pix[o*4+2] = b
			rgba.Pix[o*4+3] = 255
		}
	}
	return rgba
}

// DebayerGRBG8 debayers GRBG pattern to RGB
func DebayerGRBG8(raw []byte, w, h int) *image.RGBA {
	if w < 2 || h < 2 {
		return image.NewRGBA(image.Rect(0, 0, w, h))
	}

	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	need := w * h
	if len(raw) < need {
		return rgba
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			o := y*w + x
			r, g, b := debayerGRBGPixel(raw, w, h, x, y, o)
			rgba.Pix[o*4+0] = r
			rgba.Pix[o*4+1] = g
			rgba.Pix[o*4+2] = b
			rgba.Pix[o*4+3] = 255
		}
	}
	return rgba
}

func debayerRGGBPixel(raw []byte, w, h, x, y, o int) (uint8, uint8, uint8) {
	r := getPixel2D(raw, w, h, x, y)
	g := getGreenRGGB(raw, w, h, x, y)
	b := getBlueRGGB(raw, w, h, x, y)
	return clamp8(r), clamp8(g), clamp8(b)
}

func debayerBGGRPixel(raw []byte, w, h, x, y, o int) (uint8, uint8, uint8) {
	b := getPixel2D(raw, w, h, x, y)
	g := getGreenBGGR(raw, w, h, x, y)
	r := getRedBGGR(raw, w, h, x, y)
	return clamp8(r), clamp8(g), clamp8(b)
}

func debayerGBRGPixel(raw []byte, w, h, x, y, o int) (uint8, uint8, uint8) {
	g := getGreenGBRG(raw, w, h, x, y)
	r := getRedGBRG(raw, w, h, x, y)
	b := getBlueGBRG(raw, w, h, x, y)
	return clamp8(r), clamp8(g), clamp8(b)
}

func debayerGRBGPixel(raw []byte, w, h, x, y, o int) (uint8, uint8, uint8) {
	g := getGreenGRBG(raw, w, h, x, y)
	b := getBlueGRBG(raw, w, h, x, y)
	r := getRedGRBG(raw, w, h, x, y)
	return clamp8(r), clamp8(g), clamp8(b)
}

func getPixel2D(raw []byte, w, h, x, y int) int {
	if x < 0 || x >= w || y < 0 || y >= h {
		return 0
	}
	return int(raw[y*w+x])
}

func getGreenRGGB(raw []byte, w, h, x, y int) int {
	// RGGB: R G R G ...
	//       B G B G ...
	// G neighbors: (x-1,y), (x+1,y), (x,y-1), (x,y+1)
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

func getBlueRGGB(raw []byte, w, h, x, y int) int {
	// Blue is at (even x, even y) in RGGB
	if x%2 == 0 && y%2 == 0 {
		return getPixel2D(raw, w, h, x, y)
	}
	// Neighbors
	b := 0
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx%2 == 0 && ny%2 == 0 {
				b += getPixel2D(raw, w, h, nx, ny)
				count++
			}
		}
	}
	if count > 0 {
		return b / count
	}
	return 0
}

func getGreenBGGR(raw []byte, w, h, x, y int) int {
	// BGGR: B G B G ...
	//       G R G R ...
	// Green at (odd x, even y) and (even x, odd y)
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

func getRedBGGR(raw []byte, w, h, x, y int) int {
	// Red is at (odd x, odd y) in BGGR
	if x%2 == 1 && y%2 == 1 {
		return getPixel2D(raw, w, h, x, y)
	}
	r := 0
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx%2 == 1 && ny%2 == 1 {
				r += getPixel2D(raw, w, h, nx, ny)
				count++
			}
		}
	}
	if count > 0 {
		return r / count
	}
	return 0
}

func getGreenGBRG(raw []byte, w, h, x, y int) int {
	// GBRG: G B G B ...
	//       R G R G ...
	// Green at (odd x, even y) and (even x, odd y)
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

func getRedGBRG(raw []byte, w, h, x, y int) int {
	// Red is at (odd x, odd y) in GBRG
	if x%2 == 1 && y%2 == 1 {
		return getPixel2D(raw, w, h, x, y)
	}
	r := 0
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx%2 == 1 && ny%2 == 1 {
				r += getPixel2D(raw, w, h, nx, ny)
				count++
			}
		}
	}
	if count > 0 {
		return r / count
	}
	return 0
}

func getBlueGBRG(raw []byte, w, h, x, y int) int {
	// Blue is at (even x, odd y) in GBRG
	if x%2 == 0 && y%2 == 1 {
		return getPixel2D(raw, w, h, x, y)
	}
	b := 0
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx%2 == 0 && ny%2 == 1 {
				b += getPixel2D(raw, w, h, nx, ny)
				count++
			}
		}
	}
	if count > 0 {
		return b / count
	}
	return 0
}

func getGreenGRBG(raw []byte, w, h, x, y int) int {
	// GRBG: G R G R ...
	//       B G B G ...
	// Green at (odd x, even y) and (even x, odd y)
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

func getBlueGRBG(raw []byte, w, h, x, y int) int {
	// Blue is at (even x, odd y) in GRBG
	if x%2 == 0 && y%2 == 1 {
		return getPixel2D(raw, w, h, x, y)
	}
	b := 0
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx%2 == 0 && ny%2 == 1 {
				b += getPixel2D(raw, w, h, nx, ny)
				count++
			}
		}
	}
	if count > 0 {
		return b / count
	}
	return 0
}

func getRedGRBG(raw []byte, w, h, x, y int) int {
	// Red is at (odd x, even y) in GRBG
	if x%2 == 1 && y%2 == 0 {
		return getPixel2D(raw, w, h, x, y)
	}
	r := 0
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx%2 == 1 && ny%2 == 0 {
				r += getPixel2D(raw, w, h, nx, ny)
				count++
			}
		}
	}
	if count > 0 {
		return r / count
	}
	return 0
}

// DebayerToRGBA debayers a Bayer pattern to RGBA
func DebayerToRGBA(raw []byte, w, h int, pattern byte) (*image.RGBA, error) {
	if w < 2 || h < 2 {
		return image.NewRGBA(image.Rect(0, 0, w, h)), nil
	}
	switch pattern {
	case BayerPatternRGGB:
		return DebayerRGGB8(raw, w, h), nil
	case BayerPatternBGGR:
		return DebayerBGGR8(raw, w, h), nil
	case BayerPatternGBRG:
		return DebayerGBRG8(raw, w, h), nil
	case BayerPatternGRBG:
		return DebayerGRBG8(raw, w, h), nil
	default:
		return nil, fmt.Errorf("color: unknown Bayer pattern %d", pattern)
	}
}
