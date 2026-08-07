package gvcp

import (
	"encoding/binary"
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
