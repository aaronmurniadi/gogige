package gvsp

import "testing"

func TestOOOPacketRingPutGetDelete(t *testing.T) {
	r := NewOOOPacketRing()
	if r.Count() != 0 {
		t.Fatalf("count=%d want 0", r.Count())
	}
	// Put out of order
	if !r.Put(5, []byte("e")) {
		t.Fatal("put 5 failed")
	}
	if !r.Put(2, []byte("b")) {
		t.Fatal("put 2 failed")
	}
	if !r.Put(8, []byte("h")) {
		t.Fatal("put 8 failed")
	}
	if r.Count() != 3 {
		t.Fatalf("count=%d want 3", r.Count())
	}
	// Get by id
	if p, ok := r.Get(2); !ok || string(p) != "b" {
		t.Fatalf("get 2 = %q,%v", p, ok)
	}
	if p, ok := r.Get(8); !ok || string(p) != "h" {
		t.Fatalf("get 8 = %q,%v", p, ok)
	}
	if _, ok := r.Get(9); ok {
		t.Fatal("get 9 should miss")
	}
	// Delete middle entry - must not drop head
	if p, ok := r.Delete(2); !ok || string(p) != "b" {
		t.Fatalf("delete 2 = %q,%v", p, ok)
	}
	if r.Count() != 2 {
		t.Fatalf("count after delete=%d want 2", r.Count())
	}
	if p, ok := r.Get(5); !ok || string(p) != "e" {
		t.Fatalf("get 5 after delete = %q,%v", p, ok)
	}
	if p, ok := r.Get(8); !ok || string(p) != "h" {
		t.Fatalf("get 8 after delete = %q,%v", p, ok)
	}
	if r.maxID() != 8 {
		t.Fatalf("maxID=%d want 8", r.maxID())
	}
}

func TestOOOPacketRingOverflowSpill(t *testing.T) {
	r := NewOOOPacketRing()
	// Fill ring to capacity
	for i := 0; i < MaxOOOPackets; i++ {
		if !r.Put(uint32(i+1), []byte{byte(i)}) {
			t.Fatalf("put %d failed", i)
		}
	}
	if r.Count() != MaxOOOPackets {
		t.Fatalf("count=%d want %d", r.Count(), MaxOOOPackets)
	}
	// Overflow spills to map
	if r.Put(uint32(MaxOOOPackets+1), []byte("x")) {
		t.Fatal("overflow put should return false")
	}
	if r.Count() != MaxOOOPackets+1 {
		t.Fatalf("count=%d want %d", r.Count(), MaxOOOPackets+1)
	}
	if p, ok := r.Get(uint32(MaxOOOPackets + 1)); !ok || string(p) != "x" {
		t.Fatalf("overflow get = %q,%v", p, ok)
	}
	// Ring entries still intact
	if p, ok := r.Get(1); !ok || string(p) != string([]byte{0}) {
		t.Fatalf("get 1 after spill = %q,%v", p, ok)
	}
	// Delete range through ring slots (holes) keeps remaining intact
	for i := 1; i < MaxOOOPackets; i += 2 {
		if _, ok := r.Delete(uint32(i)); !ok {
			t.Fatalf("delete %d failed", i)
		}
	}
	if r.Count() != MaxOOOPackets+1-(MaxOOOPackets/2) {
		t.Fatalf("count after hole delete=%d", r.Count())
	}
	if p, ok := r.Get(2); !ok || string(p) != string([]byte{1}) {
		t.Fatalf("get 2 after hole-deletes = %q,%v", p, ok)
	}
	if p, ok := r.Get(uint32(MaxOOOPackets + 1)); !ok || string(p) != "x" {
		t.Fatalf("get overflow after hole-deletes = %q,%v", p, ok)
	}
	// Ring slots refillable after delete
	if !r.Put(1000, []byte("fresh")) {
		t.Fatal("reuse after delete failed")
	}
	if p, ok := r.Get(1000); !ok || string(p) != "fresh" {
		t.Fatalf("get reused slot = %q,%v", p, ok)
	}
}
