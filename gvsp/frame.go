package gvsp

// Frame is one reassembled GVSP image buffer plus leader metadata.
type Frame struct {
	ID          uint64
	Width       uint32
	Height      uint32
	PixelFormat uint32
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

type frameBuild struct {
	id          uint64
	width       uint32
	height      uint32
	pixelFormat uint32
	parts       map[uint32][]byte
	buf         []byte // pooled contiguous payload for packets [1, nextPkt)
	nextPkt     uint32
	trailerPkt  uint32 // trailer packet_id once seen (0 = unset / unknown)
	extended    bool
	resendNext  uint32 // first packet_id not yet covered by a resend request
	broken      bool
	pool        *BufferPool
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
	if len(fb.parts) > 0 {
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
		Data:        data,
		pool:        pool,
	}
}

func assembleFromParts(fb *frameBuild) *Frame {
	var maxID uint32
	for id := range fb.parts {
		if id > maxID {
			maxID = id
		}
	}
	start := uint32(1)
	if _, ok := fb.parts[0]; ok {
		start = 0
	}
	need := 0
	for id := start; id <= maxID; id++ {
		p, ok := fb.parts[id]
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
		data = append(data, fb.parts[id]...)
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
		Data:        data,
		pool:        pool,
	}
}
