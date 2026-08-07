package gvcp

import (
	"errors"
	"testing"
)

func TestNegotiatePacketSizeAcceptsWant(t *testing.T) {
	g := &fakeSCPS{reg: 0x200005dc}
	got, err := NegotiatePacketSize(g, 1500)
	if err != nil || got != 1500 {
		t.Fatalf("got=%d err=%v", got, err)
	}
}

func TestNegotiatePacketSizeFallsBack(t *testing.T) {
	g := &fakeSCPS{reg: 0x5dc, maxOK: 1500}
	got, err := NegotiatePacketSize(g, 9000)
	if err != nil || got != 1500 {
		t.Fatalf("got=%d err=%v", got, err)
	}
}

type fakeSCPS struct {
	reg   uint32
	maxOK uint32 // 0 = accept any size
}

func (f *fakeSCPS) ReadReg(addr uint32) (uint32, error) {
	if addr != Stream0PacketSize {
		return 0, nil
	}
	return f.reg, nil
}

func (f *fakeSCPS) WriteReg(addr, value uint32) error {
	if addr != Stream0PacketSize {
		return nil
	}
	sz := value & 0xffff
	if f.maxOK != 0 && sz > f.maxOK {
		return errors.New("gige: gvcp error INVALID_PARAMETER (0x02) cmd=0x0083")
	}
	f.reg = value
	return nil
}

func (f *fakeSCPS) ReadMem(uint32, int) ([]byte, error) { return nil, nil }
func (f *fakeSCPS) WriteMem(uint32, []byte) error       { return nil }
