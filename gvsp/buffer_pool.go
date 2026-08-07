package gvsp

// Default pool sizing for GigE frames (enough for ~1080p BGR8 + headroom).
const (
	DefaultPoolFrames = 8
	DefaultFrameSize  = 4 << 20 // 4 MiB
)

// BufferPool is a fixed set of pre-sized frame buffers for zero-alloc reuse.
type BufferPool struct {
	ch chan []byte
	sz int
}

// NewBufferPool pre-allocates count buffers of size bytes each.
func NewBufferPool(count, size int) *BufferPool {
	if count <= 0 {
		count = DefaultPoolFrames
	}
	if size <= 0 {
		size = DefaultFrameSize
	}
	p := &BufferPool{ch: make(chan []byte, count), sz: size}
	for i := 0; i < count; i++ {
		p.ch <- make([]byte, 0, size)
	}
	return p
}

// Get returns a buffer with length 0 and capacity >= pool size.
// Under exhaustion it allocates a fresh buffer (hot path stays alloc-free
// when the pool is sized for concurrent in-flight frames).
func (p *BufferPool) Get() []byte {
	if p == nil {
		return make([]byte, 0, DefaultFrameSize)
	}
	select {
	case b := <-p.ch:
		return b[:0]
	default:
		return make([]byte, 0, p.sz)
	}
}

// Put returns a buffer to the pool. Undersized or nil buffers are dropped.
func (p *BufferPool) Put(b []byte) {
	if p == nil || b == nil || cap(b) < p.sz {
		return
	}
	select {
	case p.ch <- b[:0]:
	default:
	}
}

// Size returns the capacity of each pooled buffer.
func (p *BufferPool) Size() int {
	if p == nil {
		return 0
	}
	return p.sz
}
