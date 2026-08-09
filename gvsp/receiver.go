package gvsp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	gvspContentLeader  = 0x01
	gvspContentTrailer = 0x02
	gvspContentPayload = 0x03
)

// Stream receives GVSP packets and reassembles frames.
type Stream struct {
	conn   *net.UDPConn
	port   int
	pool   *BufferPool
	resend ResendFunc
	rcvBuf int

	mu     sync.Mutex
	frames map[uint64]*frameBuild
	last   *Frame
}

// ListenStream binds a UDP socket for GVSP (port 0 = ephemeral).
func ListenStream(port int) (*Stream, error) {
	return ListenStreamPool(port, NewBufferPool(DefaultPoolFrames, DefaultFrameSize))
}

// ListenStreamPool is ListenStream with an explicit frame buffer pool.
func ListenStreamPool(port int, pool *BufferPool) (*Stream, error) {
	addr := &net.UDPAddr{IP: net.IPv4zero, Port: port}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		return nil, fmt.Errorf("gige: listen gvsp: %w", err)
	}
	la := conn.LocalAddr().(*net.UDPAddr)
	if pool == nil {
		pool = NewBufferPool(DefaultPoolFrames, DefaultFrameSize)
	}
	s := &Stream{conn: conn, port: la.Port, pool: pool, frames: map[uint64]*frameBuild{}}
	if _, err := s.SetRecvBuffer(DefaultRecvBufSize); err != nil {
		_ = conn.Close()
		return nil, err
	}
	go s.readLoop()
	return s, nil
}

// SetResender registers a GVCP PACKETRESEND_CMD sender. nil disables resend.
func (s *Stream) SetResender(fn ResendFunc) {
	if s == nil {
		return
	}
	s.mu.Lock()
	s.resend = fn
	s.mu.Unlock()
}

// Port returns the local UDP port.
func (s *Stream) Port() int {
	if s == nil {
		return 0
	}
	return s.port
}

// Close stops the stream socket.
func (s *Stream) Close() error {
	if s == nil || s.conn == nil {
		return nil
	}
	return s.conn.Close()
}

func (s *Stream) readLoop() {
	buf := make([]byte, 65535)
	for {
		n, _, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			return
		}
		s.handlePacket(buf[:n])
	}
}

func (s *Stream) handlePacket(pkt []byte) {
	if len(pkt) < 8 {
		return
	}
	status := binary.BigEndian.Uint16(pkt[0:])
	if status&0x8000 != 0 {
		return
	}
	infos := binary.BigEndian.Uint32(pkt[4:])
	ext := infos&0x80000000 != 0
	var (
		frameID  uint64
		packetID uint32
		content  uint8
		dataOff  int
	)
	if ext {
		if len(pkt) < 20 {
			return
		}
		content = uint8((infos >> 24) & 0x7f)
		frameID = binary.BigEndian.Uint64(pkt[8:])
		packetID = binary.BigEndian.Uint32(pkt[16:])
		dataOff = 20
	} else {
		frameID = uint64(binary.BigEndian.Uint16(pkt[2:]))
		content = uint8((infos >> 24) & 0x7f)
		packetID = infos & 0x00ffffff
		dataOff = 8
	}
	data := pkt[dataOff:]

	s.mu.Lock()
	defer s.mu.Unlock()
	fb := s.frames[frameID]
	if fb == nil {
		fb = newFrameBuild()
		fb.id = frameID
		fb.pool = s.pool
		fb.nextPkt = 1
		s.frames[frameID] = fb
	}
	fb.extended = ext

	switch content {
	case gvspContentLeader:
		if len(data) >= 36 {
			// GVSP leader payload-type-specific header: payload type lives in
			// the first 4 bytes, followed by the transport header fields.
			fb.payloadType = binary.BigEndian.Uint32(data[0:])
			if fb.payloadType > 0x0B && !IsPayloadTypeGenDC(fb.payloadType) &&
				!IsPayloadTypeMultiPart(fb.payloadType) && !IsPayloadTypeChunk(fb.payloadType) {
				// vendor/custom leaders (e.g. BSCF) carry other bytes here
				fb.payloadType = 0
			}
			fb.pixelFormat = binary.BigEndian.Uint32(data[12:])
			fb.width = binary.BigEndian.Uint32(data[16:])
			fb.height = binary.BigEndian.Uint32(data[20:])
		}
	case gvspContentPayload:
		s.appendPayload(fb, packetID, data)
		if fb.trailerPkt != 0 {
			s.tryFinish(frameID, fb)
		}
	case gvspContentTrailer:
		fb.trailerPkt = packetID
		s.requestMissing(fb)
		s.tryFinish(frameID, fb)
	}
}

func (s *Stream) tryFinish(frameID uint64, fb *frameBuild) {
	frame := assembleFrame(fb)
	if frame == nil {
		return
	}
	if s.last != nil {
		s.last.Release()
	}
	s.last = frame
	delete(s.frames, frameID)
}

func (s *Stream) requestMissing(fb *frameBuild) {
	if s.resend == nil || fb == nil {
		return
	}
	var to uint32
	switch {
	case fb.trailerPkt >= 2:
		to = fb.trailerPkt - 1
	default:
		return
	}
	for _, r := range fb.missingPayloadRanges(1, to) {
		s.resend(fb.id, r.First, r.Last, fb.extended)
		if r.Last+1 > fb.resendNext {
			fb.resendNext = r.Last + 1
		}
	}
}

func (s *Stream) requestGap(fb *frameBuild, first, last uint32) {
	if s.resend == nil || last < first {
		return
	}
	if first < fb.resendNext {
		first = fb.resendNext
	}
	if last < first {
		return
	}
	s.resend(fb.id, first, last, fb.extended)
	fb.resendNext = last + 1
}

// appendPayload copies payload into the pooled contiguous buffer when possible;
// out-of-order packets are held in parts ring buffer and trigger PACKETRESEND for the gap.
func (s *Stream) appendPayload(fb *frameBuild, packetID uint32, data []byte) {
	if fb.broken || packetID == 0 {
		return
	}
	if packetID < fb.nextPkt {
		return // duplicate already flushed
	}
	if packetID > fb.nextPkt {
		// OOO packet - store in ring (spills to rare overflow map when full)
		fb.parts.Put(packetID, data)
		s.requestGap(fb, fb.nextPkt, packetID-1)
		return
	}

	s.appendContiguous(fb, data)
	fb.nextPkt++
	for {
		p, ok := fb.parts.Get(fb.nextPkt)
		if !ok {
			break
		}
		_, _ = fb.parts.Delete(fb.nextPkt)
		s.appendContiguous(fb, p)
		fb.nextPkt++
	}
}

func (s *Stream) appendContiguous(fb *frameBuild, data []byte) {
	if fb.buf == nil {
		fb.buf = s.pool.Get()
	}
	need := len(fb.buf) + len(data)
	if need > cap(fb.buf) {
		capHint := need
		if c := cap(fb.buf) * 2; c > capHint {
			capHint = c
		}
		nb := make([]byte, len(fb.buf), capHint)
		copy(nb, fb.buf)
		if fb.pool != nil {
			fb.pool.Put(fb.buf)
			fb.pool = nil
		}
		fb.buf = nb
	}
	fb.buf = append(fb.buf, data...)
}

// Recv waits up to timeout for a complete frame.
func (s *Stream) Recv(timeout time.Duration) (*Frame, error) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		s.mu.Lock()
		f := s.last
		if f != nil {
			s.last = nil
			s.mu.Unlock()
			return f, nil
		}
		s.mu.Unlock()
		time.Sleep(5 * time.Millisecond)
	}
	return nil, errors.New("gige: gvsp recv timeout")
}

// EncodeGVSPPayload builds a minimal non-extended payload packet for tests.
func EncodeGVSPPayload(frameID uint16, packetID uint32, content uint8, data []byte) []byte {
	pkt := make([]byte, 8+len(data))
	binary.BigEndian.PutUint16(pkt[0:], 0) // OK
	binary.BigEndian.PutUint16(pkt[2:], frameID)
	infos := (uint32(content) << 24) | (packetID & 0x00ffffff)
	binary.BigEndian.PutUint32(pkt[4:], infos)
	copy(pkt[8:], data)
	return pkt
}
