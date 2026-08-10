package gvcp

import (
	"encoding/binary"
	"fmt"
	"strings"
	"testing"
	"time"
)

type memPort struct {
	regs map[uint32]uint32
	mem  map[uint32]byte
}

func (m *memPort) ReadReg(addr uint32) (uint32, error) {
	return m.regs[addr], nil
}
func (m *memPort) WriteReg(addr, value uint32) error {
	if m.regs == nil {
		m.regs = map[uint32]uint32{}
	}
	m.regs[addr] = value
	return nil
}
func (m *memPort) ReadMem(addr uint32, n int) ([]byte, error) {
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = m.mem[addr+uint32(i)]
	}
	return out, nil
}
func (m *memPort) WriteMem(addr uint32, data []byte) error {
	if m.mem == nil {
		m.mem = map[uint32]byte{}
	}
	for i, b := range data {
		m.mem[addr+uint32(i)] = b
	}
	return nil
}

func (m *memPort) FirstURL() (string, error) {
	b, err := m.ReadMem(gvbsXMLURL0, gvbsXMLURLSize)
	if err != nil {
		return "", err
	}
	s := string(b)
	if i := strings.IndexByte(s, 0); i >= 0 {
		s = s[:i]
	}
	return strings.TrimSpace(s), nil
}

func (m *memPort) ManifestTableURL() (string, error) {
	return "", nil
}

func TestEncodeReadReg(t *testing.T) {
	pkt := encodeReadReg(0xa00, 7)
	if len(pkt) != 12 {
		t.Fatalf("len=%d", len(pkt))
	}
	if pkt[0] != gvcpPacketTypeCMD || pkt[1] != gvcpFlagAckRequired {
		t.Fatalf("header bytes %x %x", pkt[0], pkt[1])
	}
	if binary.BigEndian.Uint16(pkt[2:]) != gvcpCmdReadReg {
		t.Fatalf("cmd=%x", binary.BigEndian.Uint16(pkt[2:]))
	}
	if binary.BigEndian.Uint16(pkt[4:]) != 4 {
		t.Fatalf("size=%d", binary.BigEndian.Uint16(pkt[4:]))
	}
	if binary.BigEndian.Uint16(pkt[6:]) != 7 {
		t.Fatalf("id=%d", binary.BigEndian.Uint16(pkt[6:]))
	}
	if binary.BigEndian.Uint32(pkt[8:]) != 0xa00 {
		t.Fatalf("addr=%x", binary.BigEndian.Uint32(pkt[8:]))
	}
}

func TestEncodeWriteReg(t *testing.T) {
	pkt := encodeWriteReg(0xd00, 0x1234, 3)
	if len(pkt) != 16 {
		t.Fatalf("len=%d", len(pkt))
	}
	if binary.BigEndian.Uint32(pkt[8:]) != 0xd00 {
		t.Fatalf("addr")
	}
	if binary.BigEndian.Uint32(pkt[12:]) != 0x1234 {
		t.Fatalf("value")
	}
}

func TestEncodeReadMemAlign(t *testing.T) {
	pkt := encodeReadMem(0x1000, 5, 1)
	size := binary.BigEndian.Uint32(pkt[12:])
	if size != 8 {
		t.Fatalf("aligned size=%d want 8", size)
	}
}

func TestEncodePacketResend(t *testing.T) {
	pkt := EncodePacketResend(9, 0, 0x1234, 2, 5, false)
	if len(pkt) != 20 {
		t.Fatalf("len=%d", len(pkt))
	}
	if pkt[0] != gvcpPacketTypeCMD || pkt[1] != 0 {
		t.Fatalf("hdr %x %x", pkt[0], pkt[1])
	}
	if binary.BigEndian.Uint16(pkt[2:]) != gvcpCmdPacketResend {
		t.Fatalf("cmd=%x", binary.BigEndian.Uint16(pkt[2:]))
	}
	if binary.BigEndian.Uint32(pkt[8:]) != 0x1234 {
		t.Fatalf("channel|block=%x", binary.BigEndian.Uint32(pkt[8:]))
	}
	if binary.BigEndian.Uint32(pkt[12:]) != 2 || binary.BigEndian.Uint32(pkt[16:]) != 5 {
		t.Fatalf("first/last")
	}

	ext := EncodePacketResend(1, 0, 0x1122334455667788, 1, 3, true)
	if len(ext) != 28 {
		t.Fatalf("ext len=%d", len(ext))
	}
	if ext[1] != gvcpFlagExtendedIDs {
		t.Fatalf("flags=%x", ext[1])
	}
	if binary.BigEndian.Uint64(ext[20:]) != 0x1122334455667788 {
		t.Fatalf("block64")
	}
}

func TestPendingAckTimeout(t *testing.T) {
	fallback := 2 * time.Second
	if got := pendingAckTimeout(nil, fallback); got != fallback {
		t.Fatalf("nil: %v", got)
	}
	if got := pendingAckTimeout([]byte{0, 0}, fallback); got != fallback {
		t.Fatalf("short: %v", got)
	}
	scd := []byte{0, 0, 0, 0} // timeout 0 → fallback
	if got := pendingAckTimeout(scd, fallback); got != fallback {
		t.Fatalf("zero: %v", got)
	}
	binary.BigEndian.PutUint16(scd[2:], 1500)
	if got := pendingAckTimeout(scd, fallback); got != 1500*time.Millisecond {
		t.Fatalf("1500ms: %v", got)
	}
}

func TestImplementationEndiannessHelpers(t *testing.T) {
	g := &GVCP{deviceOrder: binary.BigEndian}
	if g.DeviceByteOrder() != binary.BigEndian {
		t.Fatal("default BE")
	}
	enc := g.EncodeDeviceUint32(0x01020304)
	if binary.BigEndian.Uint32(enc) != 0x01020304 {
		t.Fatalf("BE encode %x", enc)
	}
	g.SetDeviceByteOrder(binary.LittleEndian)
	enc = g.EncodeDeviceUint32(0x01020304)
	if binary.LittleEndian.Uint32(enc) != 0x01020304 {
		t.Fatalf("LE encode %x", enc)
	}
	if g.DecodeDeviceUint32(enc) != 0x01020304 {
		t.Fatal("LE decode")
	}
}

func TestApplyImplementationEndiannessValue(t *testing.T) {
	// Only the two GenCP-defined values flip the cached order.
	cases := []struct {
		v    uint32
		want binary.ByteOrder
	}{
		{EndiannessBig, binary.BigEndian},
		{EndiannessLittle, binary.LittleEndian},
		{0x4c6f6361, binary.BigEndian}, // "Loca" from FirstURL — ignored
	}
	for _, tc := range cases {
		g := &GVCP{deviceOrder: binary.BigEndian}
		applyImplementationEndianness(g, tc.v)
		if g.DeviceByteOrder() != tc.want {
			t.Fatalf("v=%#x: got %v want %v", tc.v, g.DeviceByteOrder(), tc.want)
		}
	}
}

func TestReadManifestTable(t *testing.T) {
	tableAddr := uint32(0x1000)
	m := &memPort{mem: map[uint32]byte{}}
	// Write manifest table address to 0x01D0 (big-endian)
	addrBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(addrBytes, uint64(tableAddr))
	for i, b := range addrBytes {
		m.mem[AbrmManifestTableAddress+uint32(i)] = b
	}
	// Write MTAB header at 0x1000
	header := []byte("MTAB")
	for i, b := range header {
		m.mem[tableAddr+uint32(i)] = b
	}
	countBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(countBytes, 2)
	for i, b := range countBytes {
		m.mem[tableAddr+8+uint32(i)] = b
	}
	// Entry 1: XML at 0x2000, len 5, type 1
	e1 := make([]byte, 12)
	binary.BigEndian.PutUint32(e1[0:], 0x2000)
	binary.BigEndian.PutUint32(e1[4:], 5)
	binary.BigEndian.PutUint32(e1[8:], ManifestEntryTypeXML)
	for i, b := range e1 {
		m.mem[tableAddr+12+uint32(i)] = b
	}
	// Entry 2: other at 0x3000, len 4, type 2
	e2 := make([]byte, 12)
	binary.BigEndian.PutUint32(e2[0:], 0x3000)
	binary.BigEndian.PutUint32(e2[4:], 4)
	binary.BigEndian.PutUint32(e2[8:], 2)
	for i, b := range e2 {
		m.mem[tableAddr+24+uint32(i)] = b
	}
	// Write XML data at 0x2000
	m.WriteMem(0x2000, []byte("<?xml"))

	entries, err := ReadManifestTable(m)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 2 {
		t.Fatalf("entries=%d", len(entries))
	}
	if entries[0].Address != 0x2000 || entries[0].Length != 5 || entries[0].Type != ManifestEntryTypeXML {
		t.Fatalf("entry0=%+v", entries[0])
	}
	if entries[1].Address != 0x3000 || entries[1].Length != 4 || entries[1].Type != 2 {
		t.Fatalf("entry1=%+v", entries[1])
	}
}

func TestReadManifestTableZeroAddr(t *testing.T) {
	m := &memPort{mem: map[uint32]byte{}}
	// Zero address in 0x01D0
	for i := 0; i < 8; i++ {
		m.mem[AbrmManifestTableAddress+uint32(i)] = 0
	}
	entries, err := ReadManifestTable(m)
	if err != nil {
		t.Fatal(err)
	}
	if entries != nil {
		t.Fatalf("want nil, got %v", entries)
	}
}

// rejectManifestPort rejects the ManifestTable bootstrap read with the same
// INVALID_ACCESS (0x03) status a GigE Vision camera without GenCP
// ManifestTable support returns.
type rejectManifestPort struct {
	*memPort
}

func (rejectManifestPort) ReadMem(addr uint32, n int) ([]byte, error) {
	return nil, &StatusError{Code: 0x03, Cmd: gvcpCmdReadMemAck}
}

func TestReadManifestTableInaccessibleIsNoTable(t *testing.T) {
	entries, err := ReadManifestTable(&rejectManifestPort{&memPort{}})
	if err != nil {
		t.Fatal(err)
	}
	if entries != nil {
		t.Fatalf("want nil, got %v", entries)
	}
}

func TestStatusErrorString(t *testing.T) {
	err := &StatusError{Code: 0x03, Cmd: gvcpCmdReadMemAck}
	if want := "gige: gvcp error INVALID_ACCESS (0x03) cmd=0x0085"; err.Error() != want {
		t.Fatalf("got %q want %q", err.Error(), want)
	}
	if !isAddressInaccessible(err) {
		t.Fatal("0x03 should be address-inaccessible")
	}
	if isAddressInaccessible(fmt.Errorf("network down")) {
		t.Fatal("non-status error must not match")
	}
}
func TestManifestTableURL(t *testing.T) {
	tableAddr := uint32(0x1000)
	m := &memPort{mem: map[uint32]byte{}}
	addrBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(addrBytes, uint64(tableAddr))
	for i, b := range addrBytes {
		m.mem[AbrmManifestTableAddress+uint32(i)] = b
	}
	m.WriteMem(tableAddr, []byte("MTAB"))
	countBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(countBytes, 1)
	for i, b := range countBytes {
		m.mem[tableAddr+8+uint32(i)] = b
	}
	e1 := make([]byte, 12)
	binary.BigEndian.PutUint32(e1[0:], 0x2000)
	binary.BigEndian.PutUint32(e1[4:], 20)
	binary.BigEndian.PutUint32(e1[8:], ManifestEntryTypeXML)
	for i, b := range e1 {
		m.mem[tableAddr+12+uint32(i)] = b
	}
	m.WriteMem(0x2000, []byte("local:x.xml;100;20"))

	url, err := ManifestTableURL(m)
	if err != nil {
		t.Fatal(err)
	}
	if url != "local:x.xml;100;20" {
		t.Fatalf("url=%q", url)
	}
}
