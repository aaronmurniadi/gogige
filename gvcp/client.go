package gvcp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// Port is a GVCP register/memory access port (used by GenApi).
type Port interface {
	ReadReg(addr uint32) (uint32, error)
	WriteReg(addr, value uint32) error
	ReadMem(addr uint32, n int) ([]byte, error)
	WriteMem(addr uint32, data []byte) error
}

// GVCP is a GigE Vision Control Protocol client over UDP.
type GVCP struct {
	conn    *net.UDPConn
	addr    *net.UDPAddr
	timeout time.Duration

	mu     sync.Mutex
	reqID  uint16
	closed bool
}

// DialGVCP connects to a camera's GVCP endpoint (UDP 3956).
func DialGVCP(ip string, timeout time.Duration) (*GVCP, error) {
	if ip == "" {
		return nil, errors.New("gige: empty camera IP")
	}
	if timeout <= 0 {
		timeout = 2 * time.Second
	}
	raddr, err := net.ResolveUDPAddr("udp4", net.JoinHostPort(ip, fmt.Sprintf("%d", gvcpPort)))
	if err != nil {
		return nil, fmt.Errorf("gige: resolve %s: %w", ip, err)
	}
	conn, err := net.DialUDP("udp4", nil, raddr)
	if err != nil {
		return nil, fmt.Errorf("gige: dial %s: %w", ip, err)
	}
	_ = conn.SetDeadline(time.Now().Add(timeout))
	return &GVCP{conn: conn, addr: raddr, timeout: timeout, reqID: 1}, nil
}

// Close releases the UDP socket and attempts a BYE.
func (g *GVCP) Close() error {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.closed || g.conn == nil {
		return nil
	}
	g.closed = true
	pkt := encodeGVCPHeader(gvcpCmdBye, 0, g.nextIDLocked())
	_, _ = g.conn.Write(pkt)
	err := g.conn.Close()
	g.conn = nil
	return err
}

func (g *GVCP) nextIDLocked() uint16 {
	id := g.reqID
	if id == 0 || id == 0xffff {
		id = 1
	}
	g.reqID = id + 1
	return id
}

func (g *GVCP) nextID() uint16 {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.nextIDLocked()
}

func (g *GVCP) transact(req []byte, expectCmd uint16) ([]byte, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.closed || g.conn == nil {
		return nil, errors.New("gige: gvcp closed")
	}
	_ = g.conn.SetDeadline(time.Now().Add(g.timeout))
	if _, err := g.conn.Write(req); err != nil {
		return nil, fmt.Errorf("gige: gvcp write: %w", err)
	}

	buf := make([]byte, 1500)
	for {
		n, err := g.conn.Read(buf)
		if err != nil {
			return nil, fmt.Errorf("gige: gvcp read: %w", err)
		}
		if n < gvcpHeaderSize {
			continue
		}
		pktType := buf[0]
		cmd := binary.BigEndian.Uint16(buf[2:])
		size := int(binary.BigEndian.Uint16(buf[4:]))
		id := binary.BigEndian.Uint16(buf[6:])
		reqID := binary.BigEndian.Uint16(req[6:])
		if id != reqID {
			continue
		}
		if cmd == gvcpCmdPendingAck {
			_ = g.conn.SetDeadline(time.Now().Add(g.timeout))
			continue
		}
		if pktType == gvcpPacketTypeError || (pktType&0x80) != 0 {
			code := buf[1]
			return nil, fmt.Errorf("gige: gvcp error %s (0x%02x) cmd=0x%04x", gvcpErrorName(code), code, cmd)
		}
		if cmd != expectCmd {
			return nil, fmt.Errorf("gige: gvcp unexpected ack 0x%04x want 0x%04x", cmd, expectCmd)
		}
		if n < gvcpHeaderSize+size {
			return nil, fmt.Errorf("gige: gvcp short ack (%d < %d)", n, gvcpHeaderSize+size)
		}
		out := make([]byte, size)
		copy(out, buf[gvcpHeaderSize:gvcpHeaderSize+size])
		return out, nil
	}
}

// ReadReg reads one 32-bit bootstrap/device register.
func (g *GVCP) ReadReg(addr uint32) (uint32, error) {
	id := g.nextID()
	req := encodeGVCPHeader(gvcpCmdReadReg, 4, id)
	var ab [4]byte
	binary.BigEndian.PutUint32(ab[:], addr)
	req = append(req, ab[:]...)
	data, err := g.transact(req, gvcpCmdReadRegAck)
	if err != nil {
		return 0, err
	}
	if len(data) < 4 {
		return 0, errors.New("gige: readreg short ack")
	}
	return binary.BigEndian.Uint32(data), nil
}

// WriteReg writes one 32-bit register.
func (g *GVCP) WriteReg(addr, value uint32) error {
	id := g.nextID()
	req := encodeGVCPHeader(gvcpCmdWriteReg, 8, id)
	var ab [8]byte
	binary.BigEndian.PutUint32(ab[0:], addr)
	binary.BigEndian.PutUint32(ab[4:], value)
	req = append(req, ab[:]...)
	_, err := g.transact(req, gvcpCmdWriteRegAck)
	return err
}

// ReadMem reads device memory (chunked to GVCP max).
func (g *GVCP) ReadMem(addr uint32, n int) ([]byte, error) {
	if n <= 0 {
		return nil, nil
	}
	out := make([]byte, 0, n)
	left := n
	off := addr
	for left > 0 {
		chunk := left
		if chunk > gvcpDataSizeMax {
			chunk = gvcpDataSizeMax
		}
		aligned := ((chunk + 3) / 4) * 4
		id := g.nextID()
		req := encodeGVCPHeader(gvcpCmdReadMem, 8, id)
		var ab [8]byte
		binary.BigEndian.PutUint32(ab[0:], off)
		binary.BigEndian.PutUint32(ab[4:], uint32(aligned))
		req = append(req, ab[:]...)
		data, err := g.transact(req, gvcpCmdReadMemAck)
		if err != nil {
			return nil, err
		}
		if len(data) < 4 {
			return nil, errors.New("gige: readmem short ack")
		}
		payload := data[4:]
		if len(payload) < chunk {
			return nil, fmt.Errorf("gige: readmem short payload (%d < %d)", len(payload), chunk)
		}
		out = append(out, payload[:chunk]...)
		off += uint32(chunk)
		left -= chunk
	}
	return out, nil
}

// WriteMem writes device memory (chunked).
func (g *GVCP) WriteMem(addr uint32, data []byte) error {
	off := addr
	for len(data) > 0 {
		chunk := len(data)
		if chunk > gvcpDataSizeMax {
			chunk = gvcpDataSizeMax
		}
		aligned := ((chunk + 3) / 4) * 4
		payload := make([]byte, aligned)
		copy(payload, data[:chunk])
		id := g.nextID()
		req := encodeGVCPHeader(gvcpCmdWriteMem, uint16(4+aligned), id)
		var ab [4]byte
		binary.BigEndian.PutUint32(ab[:], off)
		req = append(req, ab[:]...)
		req = append(req, payload...)
		if _, err := g.transact(req, gvcpCmdWriteMemAck); err != nil {
			return err
		}
		off += uint32(chunk)
		data = data[chunk:]
	}
	return nil
}

// TakeControl writes CCP control privilege.
// Retries ACCESS_DENIED briefly — common when a prior process still holds CCP
// until its heartbeat expires, or another local client just released it.
func (g *GVCP) TakeControl() error {
	var err error
	for attempt := 0; attempt < 8; attempt++ {
		err = g.WriteReg(gvbsCCP, gvbsCCPControl)
		if err == nil {
			return nil
		}
		if !strings.Contains(err.Error(), "ACCESS_DENIED") {
			return err
		}
		time.Sleep(400 * time.Millisecond)
	}
	return err
}

// LeaveControl clears CCP.
func (g *GVCP) LeaveControl() error {
	return g.WriteReg(gvbsCCP, 0)
}

// FirstURL reads the GenICam XML URL string from bootstrap.
func (g *GVCP) FirstURL() (string, error) {
	b, err := g.ReadMem(gvbsXMLURL0, gvbsXMLURLSize)
	if err != nil {
		return "", err
	}
	s := string(b)
	if i := strings.IndexByte(s, 0); i >= 0 {
		s = s[:i]
	}
	return strings.TrimSpace(s), nil
}

// LocalAddr returns the local UDP address used for the GVCP socket.
func (g *GVCP) LocalAddr() *net.UDPAddr {
	if g == nil || g.conn == nil {
		return nil
	}
	a, _ := g.conn.LocalAddr().(*net.UDPAddr)
	return a
}
