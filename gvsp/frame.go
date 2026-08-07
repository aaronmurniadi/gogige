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
	buf         []byte // pooled contiguous payload when packets arrive in order
	nextPkt     uint32
	ordered     bool
	broken      bool
	done        bool
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

	var data []byte
	var pool *BufferPool

	if fb.ordered && fb.buf != nil {
		data = fb.buf
		pool = fb.pool
		fb.buf = nil
	} else {
		if len(fb.parts) == 0 {
			if fb.buf != nil && fb.pool != nil {
				fb.pool.Put(fb.buf)
				fb.buf = nil
			}
			return nil
		}
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
				if fb.buf != nil && fb.pool != nil {
					fb.pool.Put(fb.buf)
					fb.buf = nil
				}
				return nil
			}
			need += len(p)
		}
		pool = fb.pool
		if fb.buf != nil && cap(fb.buf) >= need {
			data = fb.buf[:0]
		} else {
			if fb.buf != nil && pool != nil {
				pool.Put(fb.buf)
				fb.buf = nil
			}
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
		}
		for id := start; id <= maxID; id++ {
			data = append(data, fb.parts[id]...)
		}
		fb.buf = nil
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
