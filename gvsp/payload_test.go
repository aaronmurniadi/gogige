package gvsp

import (
	"bytes"
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
		{
			CenterX: -130.75, CenterY: -184.09, CenterZ: 2133.96,
			LengthMm: 253.5, WidthMm: 106.6, HeightMm: 111, Volume: 1, Stable: true,
			Orientation: [3][3]float32{
				{0.5, 0, 0.866}, {0, 1, 0}, {-0.866, 0, 0.5},
			},
		},
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
	if f.Packs[0].LengthMm < 253 || f.Packs[0].LengthMm > 254 {
		t.Fatalf("length=%v", f.Packs[0].LengthMm)
	}
	if f.Packs[0].CenterX < -131 || f.Packs[0].CenterX > -130 {
		t.Fatalf("centerX=%v", f.Packs[0].CenterX)
	}
	if f.Packs[0].Orientation[0][2] < 0.865 || f.Packs[0].Orientation[0][2] > 0.867 {
		t.Fatalf("orientation=%v", f.Packs[0].Orientation)
	}
	s, err := SampleFromBSCF(buf)
	if err != nil {
		t.Fatal(err)
	}
	if s.Component != ComponentColor {
		t.Fatalf("component=%v", s.Component)
	}
	jpeg, err := color.EncodeJPEG(s.RawColor, s.PixelWidth, s.PixelHeight, s.PixelFormat, 60)
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
	jpeg, err := color.EncodeJPEG(d.RawColor, d.PixelWidth, d.PixelHeight, d.PixelFormat, 60)
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

func TestBSCFPackCountFromPayload(t *testing.T) {
	// DS5131 always writes 1 in the descriptor pack-count slot while the
	// payload carries many densely-packed packDetSize records. The parser must
	// derive the count from the payload size, not the descriptor.
	w, h := 2, 2
	colorPix := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	packs := []PackDet{
		{LengthMm: 100, WidthMm: 20, HeightMm: 30, Volume: 60000, Stable: true},
		{LengthMm: 40, WidthMm: 10, HeightMm: 15, Volume: 6000, Stable: false},
		{LengthMm: 80, WidthMm: 5, HeightMm: 12, Volume: 4800, Stable: true},
	}

	// Reuse the builder but overwrite the descriptor slot back to 1.
	buf := BuildTestBSCFComponents([]ComponentBlock{
		{Component: ComponentColor, Data: colorPix, Width: w, Height: h, PixelFormat: color.PixelFormatBGR8},
	}, packs)
	if len(buf) < bscfHeaderV1 {
		t.Fatal("bscf too short")
	}
	binary.LittleEndian.PutUint32(buf[24+2*bscfBlockStride+36:], 1) // block[2] = packloc slot
	binary.LittleEndian.PutUint32(buf[24+2*bscfBlockStride+40:], 0)

	f, err := ParseBSCF(buf)
	if err != nil {
		t.Fatal(err)
	}
	if f.PackCount != 3 {
		t.Fatalf("PackCount=%d want 3", f.PackCount)
	}
	if len(f.Packs) != 3 {
		t.Fatalf("len(Packs)=%d want 3", len(f.Packs))
	}
	if f.Packs[2].LengthMm != 80 {
		t.Fatalf("packs[2].LengthMm=%v want 80", f.Packs[2].LengthMm)
	}

	s, err := SampleFromBSCF(buf)
	if err != nil {
		t.Fatal(err)
	}
	if s.PackCount != 3 {
		t.Fatalf("sample PackCount=%d want 3", s.PackCount)
	}
	if len(s.Packs) != 3 {
		t.Fatalf("sample len(Packs)=%d want 3", len(s.Packs))
	}
	if s.Packs[2].LengthMm != 80 || s.Packs[2].CenterY != 0 {
		t.Fatalf("sample packs[2]=%+v", s.Packs[2])
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
	jpeg, err := color.EncodeJPEG(s.RawColor, s.PixelWidth, s.PixelHeight, s.PixelFormat, 60)
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

// buildGenDCContainer mirrors the container built by internal/genDC tests:
// one Intensity component, one 2D Mono8 part.
func buildGenDCContainer(w, h int, pixels []byte) []byte {
	var buf bytes.Buffer
	hdr := make([]byte, 64)
	binary.LittleEndian.PutUint32(hdr[0:], 0x43444E47)
	hdr[4], hdr[5] = 1, 1
	binary.LittleEndian.PutUint16(hdr[8:], 0x1000)
	binary.LittleEndian.PutUint32(hdr[12:], 64)
	binary.LittleEndian.PutUint32(hdr[52:], 1)
	binary.LittleEndian.PutUint64(hdr[56:], 64)
	buf.Write(hdr)

	comp := make([]byte, 56)
	binary.LittleEndian.PutUint16(comp[0:], 0x2000)
	binary.LittleEndian.PutUint32(comp[4:], 56)
	binary.LittleEndian.PutUint64(comp[32:], 1) // intensity
	binary.LittleEndian.PutUint32(comp[40:], color.PixelFormatMono8)
	binary.LittleEndian.PutUint16(comp[46:], 1)
	binary.LittleEndian.PutUint64(comp[48:], 56)
	buf.Write(comp)

	dataAbs := 64 + 56 + 56
	part := make([]byte, 56)
	binary.LittleEndian.PutUint16(part[0:], 0x4200)
	binary.LittleEndian.PutUint32(part[4:], 56)
	binary.LittleEndian.PutUint32(part[8:], color.PixelFormatMono8)
	binary.LittleEndian.PutUint64(part[24:], uint64(len(pixels)))
	binary.LittleEndian.PutUint64(part[32:], uint64(dataAbs))
	binary.LittleEndian.PutUint32(part[40:], uint32(w))
	binary.LittleEndian.PutUint32(part[44:], uint32(h))
	buf.Write(part)
	buf.Write(pixels)
	return buf.Bytes()
}

func TestPayloadTypeNames(t *testing.T) {
	cases := map[uint32]string{
		PayloadTypeImage:          "IMAGE",
		PayloadTypeChunkData:      "CHUNK_DATA",
		PayloadTypeChunkOnly:      "CHUNK_ONLY",
		PayloadTypeMultiPart:      "MULTI_PART",
		PayloadTypeGenDC:          "GENDC",
		payloadTypeAliasGenDC:     "GENDC",
		payloadTypeAliasMultiPart: "MULTI_PART",
		payloadTypeAliasChunk:     "CHUNK_DATA",
	}
	for id, name := range cases {
		if PayloadTypeName(id) != name {
			t.Fatalf("name for 0x%08x: got %q want %q", id, PayloadTypeName(id), name)
		}
	}
	_ = cases
	if PayloadTypeName(0xdeadbeef) == "" {
		t.Fatal("expected fallback name")
	}
	if !IsPayloadTypeGenDC(PayloadTypeGenDC) || !IsPayloadTypeGenDC(payloadTypeAliasGenDC) {
		t.Fatal("IsPayloadTypeGenDC")
	}
	if !IsPayloadTypeMultiPart(PayloadTypeMultiPart) || !IsPayloadTypeMultiPart(payloadTypeAliasMultiPart) {
		t.Fatal("IsPayloadTypeMultiPart")
	}
	if !IsPayloadTypeChunk(PayloadTypeChunkData) || !IsPayloadTypeChunk(PayloadTypeChunkOnly) || !IsPayloadTypeChunk(payloadTypeAliasChunk) {
		t.Fatal("IsPayloadTypeChunk")
	}
}

func TestParsePayloadByTypeDispatch(t *testing.T) {
	// IMAGE passes through as raw data.
	out, _, _, _, _, err := ParsePayloadByType([]byte{1, 2, 3}, PayloadTypeImage)
	if err != nil || len(out) != 3 || out[0] != 1 {
		t.Fatalf("image passthrough: %v %x", err, out)
	}

	pixels := make([]byte, 4*4)
	for i := range pixels {
		pixels[i] = byte(i)
	}
	container := buildGenDCContainer(4, 4, pixels)
	for _, pt := range []uint32{PayloadTypeGenDC, payloadTypeAliasGenDC} {
		out, fmtID, w, h, _, err := ParsePayloadByType(container, pt)
		if err != nil {
			t.Fatalf("gends 0x%x: %v", pt, err)
		}
		if !bytes.Equal(out, pixels) {
			t.Fatalf("gends data mismatch")
		}
		if fmtID != color.PixelFormatMono8 || w != 4 || h != 4 {
			t.Fatalf("gends meta %x %dx%d", fmtID, w, h)
		}
	}

	// MULTI_PART routes to the multi-part parser.
	hdr := make([]byte, 8)
	binary.BigEndian.PutUint32(hdr[0:], 1)
	part := make([]byte, 32)
	binary.BigEndian.PutUint32(part[0:], 0)
	binary.BigEndian.PutUint64(part[4:], 40)
	binary.BigEndian.PutUint64(part[12:], 4)
	binary.BigEndian.PutUint32(part[20:], 1)
	binary.BigEndian.PutUint32(part[24:], 1)
	binary.BigEndian.PutUint32(part[28:], color.PixelFormatMono8)
	mp := append(append(hdr, part...), []byte{9, 8, 7, 6}...)
	out, fmtID, w, h, _, err := ParsePayloadByType(mp, PayloadTypeMultiPart)
	if err != nil {
		t.Fatalf("mp: %v", err)
	}
	if len(out) != 4 || fmtID != color.PixelFormatMono8 || w != 1 || h != 1 {
		t.Fatalf("mp result %x %dx%d", fmtID, w, h)
	}
}
