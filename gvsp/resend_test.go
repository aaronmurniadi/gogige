package gvsp

import "testing"

func TestMissingPayloadRanges(t *testing.T) {
	have := map[uint32][]byte{3: {}, 5: {}}
	got := MissingPayloadRanges(1, have, 1, 5)
	if len(got) != 2 {
		t.Fatalf("ranges=%v", got)
	}
	if got[0] != (PacketRange{1, 2}) || got[1] != (PacketRange{4, 4}) {
		t.Fatalf("ranges=%v", got)
	}
	// Contiguous prefix covers 1..2
	got = MissingPayloadRanges(3, have, 1, 5)
	if len(got) != 1 || got[0] != (PacketRange{4, 4}) {
		t.Fatalf("after prefix ranges=%v", got)
	}
}

func TestResendFillsHoleBeforeTrailer(t *testing.T) {
	s := &Stream{frames: map[uint64]*frameBuild{}, pool: NewBufferPool(2, 256)}
	var reqs []PacketRange
	s.resend = func(blockID uint64, first, last uint32, extended bool) {
		if blockID != 7 || extended {
			t.Fatalf("block=%d ext=%v", blockID, extended)
		}
		reqs = append(reqs, PacketRange{first, last})
	}

	s.handlePacket(EncodeGVSPPayload(7, 1, gvspContentPayload, []byte("AA")))
	s.handlePacket(EncodeGVSPPayload(7, 3, gvspContentPayload, []byte("CC")))
	if len(reqs) != 1 || reqs[0] != (PacketRange{2, 2}) {
		t.Fatalf("gap resend=%v", reqs)
	}
	// Resent packet arrives before trailer
	s.handlePacket(EncodeGVSPPayload(7, 2, gvspContentPayload, []byte("BB")))
	s.handlePacket(EncodeGVSPPayload(7, 4, gvspContentTrailer, nil))

	if s.last == nil {
		t.Fatal("expected assembled frame after hole fill")
	}
	if string(s.last.Data) != "AABBCC" {
		t.Fatalf("data=%q", s.last.Data)
	}
	s.last.Release()
}

func TestResendOnTrailerStillMissing(t *testing.T) {
	s := &Stream{frames: map[uint64]*frameBuild{}, pool: NewBufferPool(2, 256)}
	var reqs []PacketRange
	s.resend = func(_ uint64, first, last uint32, _ bool) {
		reqs = append(reqs, PacketRange{first, last})
	}

	s.handlePacket(EncodeGVSPPayload(1, 1, gvspContentPayload, []byte("AA")))
	s.handlePacket(EncodeGVSPPayload(1, 3, gvspContentPayload, []byte("CC")))
	reqs = nil
	s.handlePacket(EncodeGVSPPayload(1, 4, gvspContentTrailer, nil))
	if len(reqs) != 1 || reqs[0] != (PacketRange{2, 2}) {
		t.Fatalf("trailer resend=%v", reqs)
	}
	if s.last != nil {
		t.Fatal("must not finish incomplete frame")
	}
	if _, ok := s.frames[1]; !ok {
		t.Fatal("incomplete frame should remain pending")
	}
}
