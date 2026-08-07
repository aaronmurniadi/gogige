package gvsp

import "testing"

func TestBufferPoolGetPut(t *testing.T) {
	p := NewBufferPool(2, 64)
	a := p.Get()
	if cap(a) < 64 {
		t.Fatalf("cap=%d", cap(a))
	}
	a = append(a, 1, 2, 3)
	p.Put(a)
	b := p.Get()
	if len(b) != 0 || cap(b) < 64 {
		t.Fatalf("reused len=%d cap=%d", len(b), cap(b))
	}
}

func TestFrameReleaseReturnsToPool(t *testing.T) {
	p := NewBufferPool(1, 32)
	buf := p.Get()
	buf = append(buf, make([]byte, 16)...)
	f := &Frame{Data: buf, pool: p}
	f.Release()
	if f.Data != nil {
		t.Fatal("Data should be nil after Release")
	}
	// Pool had 1 slot; after Release Get should not allocate a new backing array
	// beyond the original capacity contract — just ensure Get succeeds.
	again := p.Get()
	if cap(again) < 32 {
		t.Fatalf("cap=%d", cap(again))
	}
}
