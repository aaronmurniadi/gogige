package genDC

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// buildContainer builds a minimal GenDC 1.1 container with one Intensity
// component containing a single 2D Mono8 part, plus an optional flow table.
func buildContainer(w, h int, pixels []byte, withFlow bool) []byte {
	var buf bytes.Buffer

	hdr := make([]byte, 64) // 56 base + 8 component offset
	binary.LittleEndian.PutUint32(hdr[0:], Signature)
	hdr[4] = 1 // major
	hdr[5] = 1 // minor
	hdr[6] = 0 // sub minor
	binary.LittleEndian.PutUint16(hdr[8:], HeaderContainer)
	binary.LittleEndian.PutUint32(hdr[12:], 64) // header size incl. offset
	binary.LittleEndian.PutUint32(hdr[52:], 1)  // component count
	binary.LittleEndian.PutUint64(hdr[56:], 64) // component offset
	buf.Write(hdr)

	// Component header + one part offset at abs 72.
	comp := make([]byte, 56)
	binary.LittleEndian.PutUint16(comp[0:], HeaderComponent)
	binary.LittleEndian.PutUint32(comp[4:], 56) // header size incl. offset
	binary.LittleEndian.PutUint64(comp[32:], ComponentIntensity)
	binary.LittleEndian.PutUint32(comp[40:], PixelFormatMono8)
	binary.LittleEndian.PutUint16(comp[46:], 1)  // part count
	binary.LittleEndian.PutUint64(comp[48:], 56) // part header offset (rel. to component)
	buf.Write(comp)

	partAbs := 64 + 56
	dataAbs := partAbs + 56

	part := make([]byte, 56)
	binary.LittleEndian.PutUint16(part[0:], Part2D)
	binary.LittleEndian.PutUint32(part[4:], 56)
	binary.LittleEndian.PutUint32(part[8:], PixelFormatMono8)
	binary.LittleEndian.PutUint64(part[24:], uint64(len(pixels))) // data size
	binary.LittleEndian.PutUint64(part[32:], uint64(dataAbs))     // absolute data offset
	binary.LittleEndian.PutUint32(part[40:], uint32(w))           // size x
	binary.LittleEndian.PutUint32(part[44:], uint32(h))           // size y
	buf.Write(part)
	buf.Write(pixels)

	if withFlow {
		ft := make([]byte, 16+2*8)
		binary.LittleEndian.PutUint16(ft[0:], HeaderFlowTable)
		binary.LittleEndian.PutUint32(ft[4:], 16+2*8)
		ft[8] = 1
		ft[9] = 1
		binary.LittleEndian.PutUint32(ft[12:], 2) // flow count
		binary.LittleEndian.PutUint64(ft[16:], 1024)
		binary.LittleEndian.PutUint64(ft[24:], 2048)
		buf.Write(ft)
	}
	return buf.Bytes()
}

func TestParseGenDCContainer(t *testing.T) {
	pixels := make([]byte, 6*4)
	for i := range pixels {
		pixels[i] = byte(i)
	}
	data := buildContainer(6, 4, pixels, false)

	if !IsGenDC(data) {
		t.Fatal("IsGenDC false for valid container")
	}

	f, err := ParseGenDCContainer(data)
	if err != nil {
		t.Fatal(err)
	}
	if f.Container == nil || f.Container.ComponentCount != 1 {
		t.Fatalf("component count %+v", f.Container)
	}
	if len(f.Components) != 1 {
		t.Fatalf("components=%d", len(f.Components))
	}
	c := f.Components[0]
	if c.TypeName != "intensity" {
		t.Fatalf("type=%s", c.TypeName)
	}
	if len(c.Parts) != 1 {
		t.Fatalf("parts=%d", len(c.Parts))
	}
	p := c.Parts[0]
	if p.SizeX != 6 || p.SizeY != 4 {
		t.Fatalf("size %dx%d", p.SizeX, p.SizeY)
	}
	if int(p.DataSize) != len(pixels) {
		t.Fatalf("dataSize=%d", p.DataSize)
	}
	off := int(p.DataOffset)
	if off+len(pixels) > len(data) || !bytes.Equal(data[off:off+len(pixels)], pixels) {
		t.Fatal("part data does not point at pixels")
	}
}

func TestParseFlowTable(t *testing.T) {
	data := buildContainer(2, 2, make([]byte, 4), true)

	ft, err := FlowTableFromContainer(data)
	if err != nil {
		t.Fatal(err)
	}
	if ft == nil || ft.Header == nil {
		t.Fatal("nil flow table")
	}
	if ft.Header.HeaderType != HeaderFlowTable {
		t.Fatalf("header type 0x%04x", ft.Header.HeaderType)
	}
	if ft.Header.FlowCount != 2 || len(ft.FlowSizes) != 2 {
		t.Fatalf("flow count %d sizes %v", ft.Header.FlowCount, ft.FlowSizes)
	}
	if ft.FlowSizes[0] != 1024 || ft.FlowSizes[1] != 2048 {
		t.Fatalf("flow sizes %v", ft.FlowSizes)
	}
	if !IsFlowTable(data[len(data)-32:]) {
		t.Fatal("IsFlowTable false")
	}

	_, err = ParseFlowTable(data) // data starts with a container header, not a flow table
	if err == nil {
		t.Fatal("expected error parsing container start as flow table")
	}
}

// TestPartHeaderTooShort guards against the H2 OOB read: a part header shorter
// than the real 40-byte GenDCPartHeaderBase (but >= the old 32-byte guard) must
// be rejected, never trigger a slice out-of-range panic when reading DataOffset.
func TestPartHeaderTooShort(t *testing.T) {
	data := buildContainer(2, 2, make([]byte, 4), false)
	// Find the part header (absolute offset 120 = 64 + 56) and truncate the
	// container so only 36 bytes of the 40-byte part header remain.
	end := 64 + 56 + 36
	short := data[:end]

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("ParseGenDCContainer panicked on truncated part header: %v", r)
		}
	}()

	f, err := ParseGenDCContainer(short)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// The truncated part must be skipped (no panic), leaving zero parts.
	if len(f.Components) != 1 || len(f.Components[0].Parts) != 0 {
		t.Fatalf("expected component with no parseable parts, got %d", len(f.Components[0].Parts))
	}
}
