package gvcp

import (
	"encoding/binary"
	"testing"
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
