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
