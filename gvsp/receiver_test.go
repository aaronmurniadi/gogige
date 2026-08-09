package gvsp

import (
	"encoding/binary"
	"testing"
	"time"
)

func putU32(b []byte, off int, v uint32) {
	b[off] = byte(v >> 24)
	b[off+1] = byte(v >> 16)
	b[off+2] = byte(v >> 8)
	b[off+3] = byte(v)
}

// TestGVSPLeaderPayloadType verifies the leader payload type reaches the frame,
// while vendor/custom leaders (unknown payload bytes) keep PayloadType 0.
func TestGVSPLeaderPayloadType(t *testing.T) {
	s := &Stream{frames: map[uint64]*frameBuild{}, pool: NewBufferPool(2, 256)}
	leader := make([]byte, 36)
	putU32(leader, 0, PayloadTypeGenDC) // payload type at leader[0:4]
	putU32(leader, 12, 0x01100007)
	putU32(leader, 16, 4)
	putU32(leader, 20, 4)
	s.handlePacket(EncodeGVSPPayload(1, 0, gvspContentLeader, leader))
	s.handlePacket(EncodeGVSPPayload(1, 1, gvspContentPayload, []byte("GNDC")))
	s.handlePacket(EncodeGVSPPayload(1, 2, gvspContentTrailer, nil))
	f, err := s.Recv(50 * time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if f.PayloadType != PayloadTypeGenDC {
		t.Fatalf("payload type=%d want %d", f.PayloadType, PayloadTypeGenDC)
	}
	f.Release()

	// Vendor BSCF-style leader: bytes[0:4] are not a GVSP payload type ID.
	s2 := &Stream{frames: map[uint64]*frameBuild{}, pool: NewBufferPool(2, 256)}
	leader2 := make([]byte, 36)
	copy(leader2[0:], "BSCF")
	putU32(leader2, 16, 2)
	putU32(leader2, 20, 2)
	s2.handlePacket(EncodeGVSPPayload(1, 0, gvspContentLeader, leader2))
	s2.handlePacket(EncodeGVSPPayload(1, 1, gvspContentPayload, []byte("DATA")))
	s2.handlePacket(EncodeGVSPPayload(1, 2, gvspContentTrailer, nil))
	f2, err := s2.Recv(50 * time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if f2.PayloadType != 0 {
		t.Fatalf("bscf payload type=%d want 0", f2.PayloadType)
	}
	f2.Release()
}

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
	fb := &frameBuild{id: 1, width: 2, height: 2, parts: NewOOOPacketRing()}
	fb.parts.Put(1, []byte{1, 2, 3})
	fb.parts.Put(3, []byte{7, 8, 9}) // missing packet 2
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

// TestGVSPOutOfOrder tests OOO packet handling with ring buffer
func TestGVSPOutOfOrder(t *testing.T) {
	s := &Stream{frames: map[uint64]*frameBuild{}, pool: NewBufferPool(2, 256)}

	// Resend mock to fill OOO packets
	var resendPackets map[uint64]map[uint32][]byte
	s.SetResender(func(blockID uint64, first, last uint32, extended bool) {
		if resendPackets == nil {
			resendPackets = make(map[uint64]map[uint32][]byte)
		}
		if resendPackets[blockID] == nil {
			resendPackets[blockID] = make(map[uint32][]byte)
		}
		// Simulate resend - send the missing packets
		payload := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		for id := first; id <= last; id++ {
			offset := (id - 1) * 3
			resendPackets[blockID][id] = payload[offset : offset+3]
		}
	})

	// Send packets in OOO order: packet 3 first
	payload := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	s.handlePacket(EncodeGVSPPayload(1, 3, gvspContentPayload, payload[6:9])) // packet 3 - OOO

	// Check OOO ring buffer
	fb := s.frames[1]
	if fb == nil {
		t.Fatal("frame build not created")
	}

	// Trailer arrives - triggers resend for packets 1,2
	s.handlePacket(EncodeGVSPPayload(1, 4, gvspContentTrailer, nil))

	// Resent packets arrive
	if pkt, ok := resendPackets[1][1]; ok {
		s.handlePacket(EncodeGVSPPayload(1, 1, gvspContentPayload, pkt))
	}
	if pkt, ok := resendPackets[1][2]; ok {
		s.handlePacket(EncodeGVSPPayload(1, 2, gvspContentPayload, pkt))
	}

	// Try to receive
	f, err := s.Recv(50 * time.Millisecond)
	if err != nil {
		t.Fatalf("recv failed: %v", err)
	}

	if got, want := f.Data, payload[:9]; string(got) != string(want) {
		t.Fatalf("data mismatch: got %v want %v", got, want)
	}
	f.Release()
}
