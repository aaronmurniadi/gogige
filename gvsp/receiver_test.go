package gvsp

import (
	"encoding/binary"
	"testing"
	"time"
)

func TestHandlePacketHighBlockIDNotExtended(t *testing.T) {
	// Real camera packets use standard (non-EI) headers. Once the 16-bit
	// block ID reaches ≥0x8000, pkt[2]&0x80 is set — that must NOT flip EI.
	const blockID uint16 = 0x9c80
	s := &Stream{frames: map[uint64]*frameBuild{}, pool: NewBufferPool(2, 256)}

	payload := []byte{'B', 'S', 'C', 'F', 1, 0, 0, 0}
	s.handlePacket(EncodeGVSPPayload(blockID, 1, gvspContentPayload, payload))
	s.handlePacket(EncodeGVSPPayload(blockID, 0, gvspContentTrailer, nil))

	if s.last == nil {
		t.Fatal("expected assembled frame; high block ID was treated as extended ID")
	}
	if s.last.ID != uint64(blockID) {
		t.Fatalf("frame ID: got %d want %d", s.last.ID, blockID)
	}
	if string(s.last.Data[:4]) != "BSCF" {
		t.Fatalf("payload magic: got %x want BSCF", s.last.Data[:4])
	}
	s.last.Release()
}

func TestHandlePacketExtendedIDBitInInfos(t *testing.T) {
	// True EI: bit 31 of packet_infos set; 20-byte header before data.
	s := &Stream{frames: map[uint64]*frameBuild{}, pool: NewBufferPool(2, 256)}
	pkt := make([]byte, 20+4)
	binary.BigEndian.PutUint16(pkt[0:], 0) // status OK
	// bytes 2-3: flags (unused here)
	infos := uint32(gvspContentPayload)<<24 | 0x80000000 // EI + payload
	binary.BigEndian.PutUint32(pkt[4:], infos)
	binary.BigEndian.PutUint64(pkt[8:], 0x123456789abcdef0)
	binary.BigEndian.PutUint32(pkt[16:], 1) // packet id
	copy(pkt[20:], []byte("DATA"))

	s.handlePacket(pkt)
	fb := s.frames[0x123456789abcdef0]
	if fb == nil {
		t.Fatal("expected frame build under extended frame ID")
	}
	if got := string(fb.buf); got != "DATA" {
		t.Fatalf("payload: got %q", got)
	}
}

func TestGVSPDropsIncomplete(t *testing.T) {
	fb := &frameBuild{id: 1, width: 2, height: 2, parts: map[uint32][]byte{
		1: {1, 2, 3},
		3: {7, 8, 9}, // missing packet 2
	}}
	if assembleFrame(fb) != nil {
		t.Fatal("expected nil for incomplete frame")
	}
}

func TestGVSPGrowsUndersizedPool(t *testing.T) {
	s := &Stream{frames: map[uint64]*frameBuild{}, pool: NewBufferPool(2, 8)}
	s.handlePacket(EncodeGVSPPayload(1, 1, gvspContentPayload, []byte("abcdefghijkl"))) // 12 > 8
	s.handlePacket(EncodeGVSPPayload(1, 2, gvspContentTrailer, nil))
	f, err := s.Recv(50 * time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if string(f.Data) != "abcdefghijkl" {
		t.Fatalf("got %q", f.Data)
	}
	f.Release()
}

func TestGVSPAssemble(t *testing.T) {
	s := &Stream{frames: map[uint64]*frameBuild{}, pool: NewBufferPool(2, 256)}
	leader := make([]byte, 36)
	encodeU32 := func(b []byte, off int, v uint32) {
		b[off] = byte(v >> 24)
		b[off+1] = byte(v >> 16)
		b[off+2] = byte(v >> 8)
		b[off+3] = byte(v)
	}
	encodeU32(leader, 12, 0x02180015) // BGR8
	encodeU32(leader, 16, 2)
	encodeU32(leader, 20, 2)
	s.handlePacket(EncodeGVSPPayload(1, 0, gvspContentLeader, leader))
	payload := []byte{0, 0, 255, 0, 255, 0, 255, 0, 0, 128, 128, 128}
	s.handlePacket(EncodeGVSPPayload(1, 1, gvspContentPayload, payload))
	s.handlePacket(EncodeGVSPPayload(1, 2, gvspContentTrailer, nil))
	f, err := s.Recv(50 * time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if f.Width != 2 || f.Height != 2 {
		t.Fatalf("dims %dx%d", f.Width, f.Height)
	}
	if len(f.Data) != len(payload) {
		t.Fatalf("data len %d", len(f.Data))
	}
	f.Release()
}
