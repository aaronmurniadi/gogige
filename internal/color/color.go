package color

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
)

// GigE Vision pixel format IDs (PFNC).
const (
	PixelFormatYUV422_8      = 0x02100032 // YUYV
	PixelFormatYUV422_8_UYVY = 0x0210001F // UYVY
	PixelFormatYUV422        = PixelFormatYUV422_8
	PixelFormatMono8         = 0x01080001
	PixelFormatMono16        = 0x01100007
	PixelFormatBGR8          = 0x02180015
	PixelFormatRGB8          = 0x02180014
)

// EncodeJPEG converts camera payload bytes to JPEG.
// Supports YUV422_8 (YUYV/UYVY), BGR8, RGB8, Mono8, Mono16 (high byte preview).
func EncodeJPEG(raw []byte, w, h int, pixelFormat uint32, quality int) ([]byte, error) {
	rgba, err := toRGBA(raw, w, h, pixelFormat)
	if err != nil {
		return nil, err
	}
	if quality <= 0 || quality > 100 {
		quality = 100
	}
	var out bytes.Buffer
	if err := jpeg.Encode(&out, rgba, &jpeg.Options{Quality: quality}); err != nil {
		return nil, fmt.Errorf("gige: jpeg encode: %w", err)
	}
	return out.Bytes(), nil
}

func toRGBA(raw []byte, w, h int, pixelFormat uint32) (*image.RGBA, error) {
	if w <= 0 || h <= 0 {
		return nil, fmt.Errorf("gige: invalid size %dx%d", w, h)
	}
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	switch pixelFormat {
	case 0, PixelFormatBGR8:
		need := w * h * 3
		if len(raw) < need {
			return nil, fmt.Errorf("gige: bgr short (%d < %d)", len(raw), need)
		}
		for i, o := 0, 0; i < need; i, o = i+3, o+4 {
			rgba.Pix[o+0] = raw[i+2]
			rgba.Pix[o+1] = raw[i+1]
			rgba.Pix[o+2] = raw[i+0]
			rgba.Pix[o+3] = 255
		}
	case PixelFormatRGB8:
		need := w * h * 3
		if len(raw) < need {
			return nil, fmt.Errorf("gige: rgb short (%d < %d)", len(raw), need)
		}
		for i, o := 0, 0; i < need; i, o = i+3, o+4 {
			rgba.Pix[o+0] = raw[i+0]
			rgba.Pix[o+1] = raw[i+1]
			rgba.Pix[o+2] = raw[i+2]
			rgba.Pix[o+3] = 255
		}
	case PixelFormatMono8:
		need := w * h
		if len(raw) < need {
			return nil, fmt.Errorf("gige: mono short (%d < %d)", len(raw), need)
		}
		for i, o := 0, 0; i < need; i, o = i+1, o+4 {
			v := raw[i]
			rgba.Pix[o+0] = v
			rgba.Pix[o+1] = v
			rgba.Pix[o+2] = v
			rgba.Pix[o+3] = 255
		}
	case PixelFormatMono16:
		need := w * h * 2
		if len(raw) < need {
			return nil, fmt.Errorf("gige: mono16 short (%d < %d)", len(raw), need)
		}
		for i, o := 0, 0; i < need; i, o = i+2, o+4 {
			v := raw[i+1] // LE Mono16 → high byte preview
			rgba.Pix[o+0] = v
			rgba.Pix[o+1] = v
			rgba.Pix[o+2] = v
			rgba.Pix[o+3] = 255
		}
	case PixelFormatYUV422_8:
		need := w * h * 2
		if len(raw) < need {
			return nil, fmt.Errorf("gige: yuv422 short (%d < %d)", len(raw), need)
		}
		for x := 0; x < w*h; x += 2 {
			i := x * 2
			y0, u, y1, v := int(raw[i]), int(raw[i+1]), int(raw[i+2]), int(raw[i+3])
			r0, g0, b0 := yuvToRGB(y0, u, v)
			r1, g1, b1 := yuvToRGB(y1, u, v)
			o := x * 4
			rgba.Pix[o+0], rgba.Pix[o+1], rgba.Pix[o+2], rgba.Pix[o+3] = r0, g0, b0, 255
			rgba.Pix[o+4], rgba.Pix[o+5], rgba.Pix[o+6], rgba.Pix[o+7] = r1, g1, b1, 255
		}
	case PixelFormatYUV422_8_UYVY:
		need := w * h * 2
		if len(raw) < need {
			return nil, fmt.Errorf("gige: uyvy short (%d < %d)", len(raw), need)
		}
		for x := 0; x < w*h; x += 2 {
			i := x * 2
			u, y0, v, y1 := int(raw[i]), int(raw[i+1]), int(raw[i+2]), int(raw[i+3])
			r0, g0, b0 := yuvToRGB(y0, u, v)
			r1, g1, b1 := yuvToRGB(y1, u, v)
			o := x * 4
			rgba.Pix[o+0], rgba.Pix[o+1], rgba.Pix[o+2], rgba.Pix[o+3] = r0, g0, b0, 255
			rgba.Pix[o+4], rgba.Pix[o+5], rgba.Pix[o+6], rgba.Pix[o+7] = r1, g1, b1, 255
		}
	default:
		// Heuristic: if buffer looks like BGR packed for WxH, treat as BGR
		need := w * h * 3
		if len(raw) >= need {
			return toRGBA(raw, w, h, PixelFormatBGR8)
		}
		need2 := w * h * 2
		if len(raw) >= need2 {
			return toRGBA(raw, w, h, PixelFormatYUV422_8)
		}
		return nil, fmt.Errorf("gige: unsupported pixel format 0x%08x", pixelFormat)
	}
	return rgba, nil
}

func yuvToRGB(y, u, v int) (uint8, uint8, uint8) {
	c := y - 16
	d := u - 128
	e := v - 128
	r := (298*c + 409*e + 128) >> 8
	g := (298*c - 100*d - 208*e + 128) >> 8
	b := (298*c + 516*d + 128) >> 8
	return clamp8(r), clamp8(g), clamp8(b)
}

func clamp8(v int) uint8 {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return uint8(v)
}
