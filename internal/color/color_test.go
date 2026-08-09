package color

import (
	"bytes"
	"image/jpeg"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeJPEG_BGR(t *testing.T) {
	w, h := 2, 2
	raw := []byte{
		0, 0, 255,
		0, 255, 0,
		255, 0, 0,
		255, 255, 255,
	}
	jpegBytes, err := EncodeJPEG(raw, w, h, PixelFormatBGR8, 85)
	require.NoError(t, err)
	img, err := jpeg.Decode(bytes.NewReader(jpegBytes))
	require.NoError(t, err)
	assert.Equal(t, w, img.Bounds().Dx())
	assert.Equal(t, h, img.Bounds().Dy())
}

func TestEncodeJPEG_Short(t *testing.T) {
	_, err := EncodeJPEG([]byte{1, 2, 3}, 2, 2, PixelFormatBGR8, 85)
	assert.Error(t, err)
}

func TestEncodeJPEG_YUV422(t *testing.T) {
	w, h := 4, 2
	raw := make([]byte, w*h*2)
	for i := 0; i < len(raw); i += 4 {
		raw[i] = 128
		raw[i+1] = 128
		raw[i+2] = 128
		raw[i+3] = 128
	}
	jpegBytes, err := EncodeJPEG(raw, w, h, PixelFormatYUV422, 80)
	require.NoError(t, err)
	if len(jpegBytes) < 100 {
		t.Fatalf("jpeg too small %d", len(jpegBytes))
	}
}

func TestEncodeJPEG_Bayer8(t *testing.T) {
	w, h := 4, 4
	raw := make([]byte, w*h)
	for i := range raw {
		raw[i] = byte(i * 17)
	}
	for _, pf := range []uint32{0x01080009, 0x01080008, 0x0108000A, 0x0108000B} {
		jpegBytes, err := EncodeJPEG(raw, w, h, pf, 80)
		require.NoError(t, err)
		img, err := jpeg.Decode(bytes.NewReader(jpegBytes))
		require.NoError(t, err)
		assert.Equal(t, w, img.Bounds().Dx())
		assert.Equal(t, h, img.Bounds().Dy())
	}
}

func TestDecodeHighDepth_MonoPacked(t *testing.T) {
	w, h := 2, 2 // 4 pixels = 6 packed bytes for 12-bit
	// 12-bit packed: 4 samples per 6 bytes. Value 0x123 → after <<4 shift the
	// preview high byte is 0x12.
	samples := []uint16{0x123, 0x123, 0x123, 0x123}
	packed12 := make([]byte, 6)
	packed12[0] = byte(samples[0] & 0xFF)
	packed12[1] = byte(samples[0]>>8) | byte((samples[1]&0x0F)<<4)
	packed12[2] = byte(samples[1] >> 4)
	packed12[3] = byte(samples[2] & 0xFF)
	packed12[4] = byte(samples[2]>>8) | byte((samples[3]&0x0F)<<4)
	packed12[5] = byte(samples[3] >> 4)

	jpegBytes, err := EncodeJPEG(packed12, w, h, PixelFormatMono12p, 80)
	require.NoError(t, err)
	img, err := jpeg.Decode(bytes.NewReader(jpegBytes))
	require.NoError(t, err)
	assert.Equal(t, w, img.Bounds().Dx())
	assert.Equal(t, h, img.Bounds().Dy())

	// Direct decode check on the first pixel.
	rgba, ok := DecodeHighDepth(packed12, w, h, PixelFormatMono12p)
	require.True(t, ok)
	if rgba.Pix[0] != 0x12 {
		t.Fatalf("preview byte got 0x%02x want 0x12", rgba.Pix[0])
	}
	// Grayscale.
	if rgba.Pix[0] != rgba.Pix[1] || rgba.Pix[0] != rgba.Pix[2] {
		t.Fatal("mono expected gray")
	}
}

func TestDecodeHighDepth_MonoUnpacked(t *testing.T) {
	w, h := 2, 2
	raw := []byte{0x22, 0x11, 0x22, 0x11, 0x22, 0x11, 0x22, 0x11}
	rgba, ok := DecodeHighDepth(raw, w, h, PixelFormatMono10)
	require.True(t, ok)
	if rgba.Pix[0] != 0x11 { // LE 0x1122 → high byte preview
		t.Fatalf("preview 0x%02x want 0x22", rgba.Pix[0])
	}
	_, ok = DecodeHighDepth(raw, w, h, PixelFormatRGB8)
	if ok {
		t.Fatal("RGB8 must not be handled by DecodeHighDepth")
	}
}

func TestDecodeHighDepth_BayerRG12Packed(t *testing.T) {
	w, h := 4, 4
	packed := make([]byte, (w*h*12+7)/8)
	for i := range packed {
		packed[i] = 0x22
	}
	rgba, ok := DecodeHighDepth(packed, w, h, PixelFormatBayerRG12p)
	require.True(t, ok)
	if rgba.Bounds().Dx() != w || rgba.Bounds().Dy() != h {
		t.Fatalf("bounds %v", rgba.Bounds())
	}
	nonZero := 0
	for _, v := range rgba.Pix {
		if v != 0 {
			nonZero++
		}
	}
	if nonZero == 0 {
		t.Fatal("debayer produced all-black output")
	}
}
