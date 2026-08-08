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
	if s.Component != ComponentColor {
		t.Fatalf("component=%v", s.Component)
	}
	jpeg, err := color.EncodeJPEG(s.RawColor, s.Width, s.Height, s.PixelFormat, 60)
	if err != nil {
		t.Fatal(err)
	}
	if len(jpeg) < 2 || jpeg[0] != 0xff || jpeg[1] != 0xd8 {
		t.Fatalf("bad jpeg")
	}
}

func TestBSCFSelectDepth(t *testing.T) {
	w, h := 2, 2
	depth := make([]byte, w*h*2)
	for i := range depth {
		depth[i] = byte(i * 17)
	}
	colorPix := []byte{
		0, 0, 255, 0, 255, 0,
		255, 0, 0, 128, 128, 128,
	}
	buf := BuildTestBSCFComponents([]ComponentBlock{
		{Component: ComponentDepth, Data: depth, Width: w, Height: h, PixelFormat: color.PixelFormatMono16},
		{Component: ComponentColor, Data: colorPix, Width: w, Height: h, PixelFormat: color.PixelFormatBGR8},
	}, nil)

	s, err := SampleFromBSCF(buf)
	if err != nil {
		t.Fatal(err)
	}
	if s.PixelFormat != color.PixelFormatBGR8 || len(s.RawColor) != len(colorPix) {
		t.Fatalf("default: fmt=0x%x len=%d want BGR color", s.PixelFormat, len(s.RawColor))
	}

	d, err := SampleFromBSCFComponent(buf, ComponentDepth)
	if err != nil {
		t.Fatal(err)
	}
	if d.Component != ComponentDepth || d.PixelFormat != color.PixelFormatMono16 || len(d.RawColor) != len(depth) {
		t.Fatalf("depth: component=%v fmt=0x%x len=%d", d.Component, d.PixelFormat, len(d.RawColor))
	}
	jpeg, err := color.EncodeJPEG(d.RawColor, d.Width, d.Height, d.PixelFormat, 60)
	if err != nil || len(jpeg) < 2 {
		t.Fatalf("depth jpeg: %v len=%d", err, len(jpeg))
	}

	_, err = SampleFromBSCFComponent(buf, ComponentMono)
	if err == nil {
		t.Fatal("expected missing mono error")
	}

	all, err := SampleAllFromBSCF(buf)
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 2 {
		t.Fatalf("all=%d", len(all))
	}
	if all[0].Component != ComponentDepth || all[1].Component != ComponentColor {
		t.Fatalf("components %v %v", all[0].Component, all[1].Component)
	}
	if !IsBSCF(buf) {
		t.Fatal("IsBSCF")
	}
}

func TestParseComponent(t *testing.T) {
	c, err := ParseComponent("Depth")
	if err != nil || c != ComponentDepth {
		t.Fatalf("got %v %v", c, err)
	}
	c, err = ParseComponent("range")
	if err != nil || c != ComponentDepth {
		t.Fatalf("range: %v %v", c, err)
	}
	if ComponentColor.String() != "color" {
		t.Fatal(ComponentColor.String())
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
	writeBlock := func(i int, dataType, offset, size uint32, bw, bh int, comp, imgFmt uint32) {
		off := 24 + i*bscfBlockStride
		binary.LittleEndian.PutUint32(hdr[off:], dataType)
		binary.LittleEndian.PutUint32(hdr[off+4:], offset)
		binary.LittleEndian.PutUint32(hdr[off+8:], size)
		binary.LittleEndian.PutUint32(hdr[off+12:], uint32(bw))
		binary.LittleEndian.PutUint32(hdr[off+16:], uint32(bh))
		binary.LittleEndian.PutUint32(hdr[off+20:], comp)
		binary.LittleEndian.PutUint32(hdr[off+32:], imgFmt)
	}
	depthOff := uint32(bscfHeaderV1)
	colorOff := depthOff + uint32(len(depth))
	writeBlock(0, blockTypeImage, depthOff, uint32(len(depth)), w, h, uint32(ComponentDepth), 0x01100007)
	writeBlock(1, blockTypeImage, colorOff, uint32(len(colorPix)), w, h, uint32(ComponentColor), color.PixelFormatBGR8)
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

func TestIsGenDCPayload(t *testing.T) {
	buf := []byte{0x47, 0x4E, 0x44, 0x43}
	if !IsGenDCPayload(buf) {
		t.Fatal("expected GenDC signature detection")
	}
	buf2 := []byte{0x42, 0x53, 0x43, 0x46}
	if IsGenDCPayload(buf2) {
		t.Fatal("expected non-GenDC detection")
	}
}

func TestMultiPartPayload(t *testing.T) {
	hdr := make([]byte, 8)
	binary.BigEndian.PutUint32(hdr[0:], 1)
	binary.BigEndian.PutUint32(hdr[4:], 0)

	part := make([]byte, 32)
	binary.BigEndian.PutUint32(part[0:], 0x00000000)
	binary.BigEndian.PutUint64(part[4:], 40)
	binary.BigEndian.PutUint64(part[12:], 12)
	binary.BigEndian.PutUint32(part[20:], 4)
	binary.BigEndian.PutUint32(part[24:], 3)
	binary.BigEndian.PutUint32(part[28:], 0x01080001)

	imgData := make([]byte, 12)
	for i := range imgData {
		imgData[i] = byte(i)
	}

	buf := append(append(hdr, part...), imgData...)

	payload, err := ParseMultiPartPayload(buf)
	if err != nil {
		t.Fatalf("ParseMultiPartPayload failed: %v", err)
	}
	if payload.Header.NumParts != 1 {
		t.Fatalf("numParts=%d", payload.Header.NumParts)
	}
	if len(payload.Parts) != 1 {
		t.Fatalf("parts len=%d", len(payload.Parts))
	}
	if payload.Parts[0].PartType != 0x00000000 {
		t.Fatalf("partType=0x%x", payload.Parts[0].PartType)
	}
	if payload.Parts[0].Width != 3 {
		t.Fatalf("width=%d", payload.Parts[0].Width)
	}
	if payload.Parts[0].Height != 4 {
		t.Fatalf("height=%d", payload.Parts[0].Height)
	}
	if len(payload.Parts[0].Data) != 12 {
		t.Fatalf("data len=%d", len(payload.Parts[0].Data))
	}

	partImg, ok := payload.GetPartByType(0x00000000)
	if !ok {
		t.Fatal("GetPartByType not found")
	}
	if partImg.Width != 3 {
		t.Fatalf("width=%d", partImg.Width)
	}
}

func TestChunkPayload(t *testing.T) {
	hdr := make([]byte, 16)
	binary.BigEndian.PutUint64(hdr[0:], 100)
	binary.BigEndian.PutUint32(hdr[8:], 1)
	binary.BigEndian.PutUint32(hdr[12:], 0)

	chunk := make([]byte, 16)
	binary.BigEndian.PutUint32(chunk[0:], 0x00000001)
	binary.BigEndian.PutUint32(chunk[4:], 16)
	binary.BigEndian.PutUint32(chunk[8:], 8)
	binary.BigEndian.PutUint16(chunk[12:], 1)
	binary.BigEndian.PutUint16(chunk[14:], 0)

	chunkData := make([]byte, 8)
	binary.BigEndian.PutUint64(chunkData, 1234567890)

	buf := append(append(hdr, chunk...), chunkData...)

	payload, err := ParseChunkPayload(buf)
	if err != nil {
		t.Fatalf("ParseChunkPayload failed: %v", err)
	}
	if payload.Header.ChunkCount != 1 {
		t.Fatalf("chunkCount=%d", payload.Header.ChunkCount)
	}
	if len(payload.Chunks) != 1 {
		t.Fatalf("chunks len=%d", len(payload.Chunks))
	}
	if payload.Chunks[0].ChunkID != 0x00000001 {
		t.Fatalf("chunkID=0x%x", payload.Chunks[0].ChunkID)
	}
	if payload.Chunks[0].Size != 8 {
		t.Fatalf("chunkSize=%d", payload.Chunks[0].Size)
	}
	if len(payload.Chunks[0].Data) != 8 {
		t.Fatalf("chunkData len=%d", len(payload.Chunks[0].Data))
	}

	chunkTs, ok := payload.GetChunkByID(0x00000001)
	if !ok {
		t.Fatal("GetChunkByID not found")
	}
	if chunkTs.ChunkID != 0x00000001 {
		t.Fatalf("chunkID=0x%x", chunkTs.ChunkID)
	}
}
