package gvsp

import (
	"encoding/binary"
	"testing"

	"github.com/aaronmurniadi/gogige/internal/color"
)

func TestBSCFRoundTrip(t *testing.T) {
	w, h := 4, 2
	colorBuf := make([]byte, w*h*3)
	for i := range colorBuf {
		colorBuf[i] = byte(i)
	}
	buf := BuildTestBSCF(colorBuf, w, h, color.PixelFormatBGR8, []PackDet{
		{Length: 253.5, Width: 106.6, Height: 111, Volume: 1, Stable: true},
	})
	f, err := ParseBSCF(buf)
	if err != nil {
		t.Fatal(err)
	}
	if f.Version != 1 {
		t.Fatalf("version=%d", f.Version)
	}
	if len(f.Color) != len(colorBuf) {
		t.Fatalf("color len %d", len(f.Color))
	}
	if f.PackCount != 1 || len(f.Packs) != 1 {
		t.Fatalf("packs=%d count=%d", len(f.Packs), f.PackCount)
	}
	if !f.Packs[0].Stable {
		t.Fatal("expected stable")
	}
	if f.Packs[0].Length < 253 || f.Packs[0].Length > 254 {
		t.Fatalf("length=%v", f.Packs[0].Length)
	}
	s, err := SampleFromBSCF(buf)
	if err != nil {
		t.Fatal(err)
	}
	jpeg, err := color.EncodeJPEG(s.RawColor, s.Width, s.Height, s.PixelFormat, 60)
	if err != nil {
		t.Fatal(err)
	}
	if len(jpeg) < 2 || jpeg[0] != 0xff || jpeg[1] != 0xd8 {
		t.Fatalf("bad jpeg")
	}
}

func TestBSCFPrefersColorOverDepth(t *testing.T) {
	w, h := 2, 2
	depth := make([]byte, w*h*2)
	colorPix := []byte{
		0, 0, 255, 0, 255, 0,
		255, 0, 0, 128, 128, 128,
	}
	hdr := make([]byte, bscfHeaderV1)
	binary.LittleEndian.PutUint32(hdr[0:], bscfMagic)
	binary.LittleEndian.PutUint32(hdr[4:], 1)
	binary.LittleEndian.PutUint32(hdr[8:], 2)
	writeBlock := func(i int, dataType, offset, size uint32, bw, bh int, imgType, imgFmt uint32) {
		off := 24 + i*bscfBlockStride
		binary.LittleEndian.PutUint32(hdr[off:], dataType)
		binary.LittleEndian.PutUint32(hdr[off+4:], offset)
		binary.LittleEndian.PutUint32(hdr[off+8:], size)
		binary.LittleEndian.PutUint32(hdr[off+12:], uint32(bw))
		binary.LittleEndian.PutUint32(hdr[off+16:], uint32(bh))
		binary.LittleEndian.PutUint32(hdr[off+20:], imgType)
		binary.LittleEndian.PutUint32(hdr[off+32:], imgFmt)
	}
	depthOff := uint32(bscfHeaderV1)
	colorOff := depthOff + uint32(len(depth))
	writeBlock(0, blockTypeImage, depthOff, uint32(len(depth)), w, h, 5, 0x01100007) // Depth
	writeBlock(1, blockTypeImage, colorOff, uint32(len(colorPix)), w, h, imageTypeColor, color.PixelFormatBGR8)
	buf := append(append(hdr, depth...), colorPix...)
	s, err := SampleFromBSCF(buf)
	if err != nil {
		t.Fatal(err)
	}
	if s.PixelFormat != color.PixelFormatBGR8 || len(s.RawColor) != len(colorPix) {
		t.Fatalf("got fmt=0x%x len=%d want BGR color", s.PixelFormat, len(s.RawColor))
	}
	jpeg, err := color.EncodeJPEG(s.RawColor, s.Width, s.Height, s.PixelFormat, 60)
	if err != nil || len(jpeg) < 2 {
		t.Fatalf("jpeg: %v len=%d", err, len(jpeg))
	}
}
