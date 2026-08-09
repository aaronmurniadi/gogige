package gvsp

const (
	// Max OOO packets per frame (pre-allocated ring buffer)
	MaxOOOPackets = 256
)

// Frame is one reassembled GVSP image buffer plus leader metadata.
type Frame struct {
	ID          uint64
	Width       uint32
	Height      uint32
	PixelFormat uint32
	// PayloadType is the GVSP leader payload type (see PayloadType*).
	// 0 when the leader did not carry a recognized value (e.g. vendor BSCF).
	PayloadType uint32
	Data        []byte

	pool *BufferPool
}

// Release returns the frame buffer to its pool. Safe to call on nil.
// After Release, Data must not be used.
func (f *Frame) Release() {
	if f == nil || f.pool == nil || f.Data == nil {
		return
	}
	f.pool.Put(f.Data)
	f.Data = nil
	f.pool = nil
}

// frameBuild holds state while reassembling a frame from GVSP packets.
type frameBuild struct {
	id          uint64
	width       uint32
	height      uint32
	pixelFormat uint32
	payloadType uint32
	parts       *OOOPacketRing // zero-alloc OOO packet storage
	buf         []byte         // pooled contiguous payload for packets [1, nextPkt)
	nextPkt     uint32
	trailerPkt  uint32 // trailer packet_id once seen (0 = unset / unknown)
	extended    bool
	resendNext  uint32 // first packet_id not yet covered by a resend request
	broken      bool
	pool        *BufferPool
}

// newFrameBuild creates a new frameBuild with pre-allocated OOO ring buffer.
func newFrameBuild() *frameBuild {
	return &frameBuild{
		parts: NewOOOPacketRing(),
	}
}

// OOOPacketRing is a fixed-size store for out-of-order packets.
// Zero-alloc replacement for map[uint32][]byte in frameBuild for the common case;
// only overflows past MaxOOOPackets allocate (rare, spilling to a lazily-created map).
type OOOPacketRing struct {
	packets  [MaxOOOPackets]*RingBufferSlot
	ids      [MaxOOOPackets]uint32
	used     [MaxOOOPackets]bool
	count    int
	overflow map[uint32][]byte
}

// RingBufferSlot holds a pre-allocated buffer for OOO packets.
type RingBufferSlot struct {
	data []byte // buffer with length=0, capacity=DefaultFrameSize
}

// NewOOOPacketRing creates a zero-alloc ring buffer for OOO packets.
func NewOOOPacketRing() *OOOPacketRing {
	r := &OOOPacketRing{}
	for i := 0; i < MaxOOOPackets; i++ {
		r.packets[i] = &RingBufferSlot{
			data: make([]byte, 0, DefaultFrameSize),
		}
	}
	return r
}

// Put stores a packet data. Returns true if stored in the zero-alloc ring,
// false if it spilled to the overflow map.
func (r *OOOPacketRing) Put(packetID uint32, data []byte) bool {
	for i := 0; i < MaxOOOPackets; i++ {
		if !r.used[i] {
			slot := r.packets[i]
			slot.data = slot.data[:0]
			slot.data = append(slot.data, data...)
			r.ids[i] = packetID
			r.used[i] = true
			r.count++
			return true
		}
	}
	if r.overflow == nil {
		r.overflow = make(map[uint32][]byte)
	}
	cp := make([]byte, len(data))
	copy(cp, data)
	r.overflow[packetID] = cp
	return false
}

// Get retrieves packet data by ID. Returns nil if not found.
func (r *OOOPacketRing) Get(packetID uint32) ([]byte, bool) {
	for i := 0; i < MaxOOOPackets; i++ {
		if r.used[i] && r.ids[i] == packetID {
			return r.packets[i].data, true
		}
	}
	if p, ok := r.overflow[packetID]; ok {
		return p, true
	}
	return nil, false
}

// Delete removes packet by ID and returns its data.
func (r *OOOPacketRing) Delete(packetID uint32) ([]byte, bool) {
	for i := 0; i < MaxOOOPackets; i++ {
		if r.used[i] && r.ids[i] == packetID {
			data := r.packets[i].data
			r.used[i] = false
			r.count--
			return data, true
		}
	}
	if p, ok := r.overflow[packetID]; ok {
		delete(r.overflow, packetID)
		return p, true
	}
	return nil, false
}

// Has reports whether packet ID exists in ring.
func (r *OOOPacketRing) Has(packetID uint32) bool {
	_, ok := r.Get(packetID)
	return ok
}

// Count returns number of packets in ring.
func (r *OOOPacketRing) Count() int {
	return r.count + len(r.overflow)
}

// maxID returns the largest stored packet ID (0 if empty).
func (r *OOOPacketRing) maxID() uint32 {
	var max uint32
	for i := 0; i < MaxOOOPackets; i++ {
		if r.used[i] && r.ids[i] > max {
			max = r.ids[i]
		}
	}
	for id := range r.overflow {
		if id > max {
			max = id
		}
	}
	return max
}

func assembleFrame(fb *frameBuild) *Frame {
	if fb == nil || fb.broken {
		if fb != nil && fb.buf != nil && fb.pool != nil {
			fb.pool.Put(fb.buf)
			fb.buf = nil
		}
		return nil
	}

	// OOO hole still open: wait for resend (or classic parts-only assemble if no buf yet).
	if fb.parts.Count() > 0 {
		if fb.buf != nil || fb.nextPkt > 1 {
			return nil
		}
		return assembleFromParts(fb)
	}

	if fb.buf == nil {
		return nil
	}
	// Trailer told us the last payload id; refuse incomplete contiguous prefix.
	if fb.trailerPkt >= 2 && fb.nextPkt != fb.trailerPkt {
		return nil
	}

	data := fb.buf
	pool := fb.pool
	fb.buf = nil
	if len(data) == 0 {
		if pool != nil {
			pool.Put(data)
		}
		return nil
	}
	return &Frame{
		ID:          fb.id,
		Width:       fb.width,
		Height:      fb.height,
		PixelFormat: fb.pixelFormat,
		PayloadType: fb.payloadType,
		Data:        data,
		pool:        pool,
	}
}

func assembleFromParts(fb *frameBuild) *Frame {
	// Get max packet ID from ring
	maxID := fb.parts.maxID()

	start := uint32(1)
	if _, ok := fb.parts.Get(0); ok {
		start = 0
	}
	need := 0
	for id := start; id <= maxID; id++ {
		p, ok := fb.parts.Get(id)
		if !ok {
			return nil
		}
		need += len(p)
	}
	pool := fb.pool
	var data []byte
	if pool != nil {
		data = pool.Get()
		if cap(data) < need {
			pool.Put(data)
			data = make([]byte, 0, need)
			pool = nil
		}
	} else {
		data = make([]byte, 0, need)
	}
	for id := start; id <= maxID; id++ {
		p, ok := fb.parts.Get(id)
		if ok {
			data = append(data, p...)
		}
	}
	if len(data) == 0 {
		if pool != nil {
			pool.Put(data)
		}
		return nil
	}
	return &Frame{
		ID:          fb.id,
		Width:       fb.width,
		Height:      fb.height,
		PixelFormat: fb.pixelFormat,
		PayloadType: fb.payloadType,
		Data:        data,
		pool:        pool,
	}
}
