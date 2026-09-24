package gvcp

import (
	"bytes"
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

	mu          sync.Mutex
	reqID       uint16
	closed      bool
	deviceOrder binary.ByteOrder // non-bootstrap device regs (default BE)
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
	return &GVCP{
		conn:        conn,
		addr:        raddr,
		timeout:     timeout,
		reqID:       1,
		deviceOrder: binary.BigEndian,
	}, nil
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
			wait := pendingAckTimeout(buf[gvcpHeaderSize:gvcpHeaderSize+size], g.timeout)
			_ = g.conn.SetDeadline(time.Now().Add(wait))
			continue
		}
		if pktType == gvcpPacketTypeError || (pktType&0x80) != 0 {
			code := buf[1]
			return nil, &StatusError{Code: code, Cmd: cmd}
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

// StatusError is a GVCP ACK status returned by a device (e.g. INVALID_ACCESS).
type StatusError struct {
	Code byte
	Cmd  uint16
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("gige: gvcp error %s (0x%02x) cmd=0x%04x", gvcpErrorName(e.Code), e.Code, e.Cmd)
}

// isAddressInaccessible reports whether err is a device ACK rejecting access to
// an address (INVALID_ACCESS / WRITE_PROTECT), signaling an unmapped register.
func isAddressInaccessible(err error) bool {
	var se *StatusError
	if !errors.As(err, &se) {
		return false
	}
	return se.Code == 0x03 || se.Code == 0x04
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
			// Probe GenCP ImplementationEndianness once control is held.
			_ = g.SyncImplementationEndianness()
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

// ManifestEntry describes a single entry in the device ManifestTable. Two
// on-wire layouts exist:
//
//   - GenCP-conformant (GenCP 1.3.1 Table 33/34): an 8-byte entry count
//     followed by 64-byte entries { FileVersion u32, Schema/filetype/fileformat
//     bitfield u32, Register Address u64, File Size u64, SHA1 [20], Reserved }.
//   - Vendor "MTAB" table (Huaray BSCF quirk, not GenCP): a "MTAB" magic + u32
//     count followed by 12-byte entries { Address u32, Length u32, Type u32 }.
//
// Fields that the layout that produced the entry does not carry are zero;
// GenCP reports whether the entry came from a conformant table.
type ManifestEntry struct {
	FileVersion uint32   // forward: sub-minor|<<8 minor|... major (GenCP)
	Schema      uint32   // schema/filetype/fileformat bitfield (GenCP)
	Address     uint64   // register address of the file (both layouts; u32 in MTAB)
	Length      uint32   // vendor MTAB: file length in bytes
	Type        uint32   // vendor MTAB: file type; ManifestEntryTypeXML is the only defined value
	FileSize    uint64   // GenCP: file size in bytes
	SHA1        [20]byte // GenCP: SHA1 hash, zero when not available
	GenCP       bool     // parsed from a GenCP-conformant table
}

const ManifestEntryTypeXML = 0x00000001

// ReadManifestTable reads the device ManifestTable pointed to by the
// AbrmManifestTableAddress bootstrap register (0x01D0).
// Returns nil, nil if the register is zero, the table is not present, or the
// address points at neither a conformant GenCP table nor a vendor "MTAB"
// table. Devices without GenCP ManifestTable support (many GigE Vision
// cameras) reject the bootstrap read with INVALID_ACCESS; that is treated as
// "no table" so callers fall back to the classic FirstURL.
func ReadManifestTable(p Port) ([]ManifestEntry, error) {
	addrBytes, err := p.ReadMem(AbrmManifestTableAddress, 8)
	if err != nil {
		if isAddressInaccessible(err) {
			return nil, nil
		}
		return nil, err
	}
	tableAddr := binary.BigEndian.Uint64(addrBytes)
	if tableAddr == 0 {
		return nil, nil
	}
	head, err := p.ReadMem(uint32(tableAddr), 12)
	if err != nil {
		return nil, err
	}
	if string(head[0:4]) == "MTAB" {
		return readVendorMTAB(p, uint32(tableAddr), head)
	}
	return readGenCPManifestTable(p, uint32(tableAddr), head)
}

// readVendorMTAB parses the Huaray-specific "MTAB" manifest layout (magic +
// u32 count + 12-byte {Address, Length, Type} entries). This is NOT the GenCP
// 1.3.1 Manifest Table (Table 33/34), but the tested Huaray BSCF devices
// expose it, so it is kept as a documented vendor quirk, detected by its magic.
func readVendorMTAB(p Port, tableAddr uint32, head []byte) ([]ManifestEntry, error) {
	count := binary.BigEndian.Uint32(head[8:12])
	if count == 0 || count > 1024 {
		return nil, nil
	}
	entries := make([]ManifestEntry, 0, count)
	entrySize := 12
	raw, err := p.ReadMem(tableAddr+12, int(count)*entrySize)
	if err != nil {
		return nil, err
	}
	for i := 0; i < int(count); i++ {
		off := i * entrySize
		if off+entrySize > len(raw) {
			break
		}
		entries = append(entries, ManifestEntry{
			Address: uint64(binary.BigEndian.Uint32(raw[off : off+4])),
			Length:  binary.BigEndian.Uint32(raw[off+4 : off+8]),
			Type:    binary.BigEndian.Uint32(raw[off+8 : off+12]),
		})
	}
	return entries, nil
}

// readGenCPManifestTable parses a conformant GenCP 1.3.1 Manifest Table
// (Table 33/34): an 8-byte big-endian entry count followed by 64-byte entries.
func readGenCPManifestTable(p Port, tableAddr uint32, head []byte) ([]ManifestEntry, error) {
	count := binary.BigEndian.Uint64(head[0:8])
	if count == 0 || count > 1024 {
		return nil, nil
	}
	const entrySize = 64
	raw, err := p.ReadMem(tableAddr+8, int(count)*entrySize)
	if err != nil {
		return nil, err
	}
	entries := make([]ManifestEntry, 0, count)
	for i := uint64(0); i < count; i++ {
		off := int(i * entrySize)
		if off+entrySize > len(raw) {
			break
		}
		var sha1 [20]byte
		copy(sha1[:], raw[off+24:off+44])
		entries = append(entries, ManifestEntry{
			FileVersion: binary.BigEndian.Uint32(raw[off+0 : off+4]),
			Schema:      binary.BigEndian.Uint32(raw[off+4 : off+8]),
			Address:     binary.BigEndian.Uint64(raw[off+8 : off+16]),
			FileSize:    binary.BigEndian.Uint64(raw[off+16 : off+24]),
			SHA1:        sha1,
			GenCP:       true,
		})
	}
	return entries, nil
}

// ManifestTableURL returns the first GenICam XML URL from the device
// ManifestTable, preferred over FirstURL when present.
func ManifestTableURL(p Port) (string, error) {
	entries, err := ReadManifestTable(p)
	if err != nil {
		return "", err
	}
	for _, e := range entries {
		length := uint64(e.Length)
		if e.GenCP {
			if e.FileSize == 0 {
				continue
			}
			length = e.FileSize
		} else if e.Type != ManifestEntryTypeXML || e.Length == 0 {
			continue
		}
		if e.Address > uint64(^uint32(0)) || length > 1<<20 {
			continue
		}
		data, err := p.ReadMem(uint32(e.Address), int(length))
		if err != nil {
			continue
		}
		trimmed := bytes.TrimRight(data, "\x00")
		s := string(trimmed)
		if strings.HasPrefix(s, "<?xml") || strings.HasPrefix(s, "local:") ||
			strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://") {
			return s, nil
		}
	}
	return "", nil
}

// ReadManifestTable reads the manifest table from this GVCP connection.
func (g *GVCP) ReadManifestTable() ([]ManifestEntry, error) {
	return ReadManifestTable(g)
}

// ManifestTableURL returns the manifest table URL from this GVCP connection.
func (g *GVCP) ManifestTableURL() (string, error) {
	return ManifestTableURL(g)
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

// MaximumDeviceResponseTime reads the GenCP MaximumDeviceResponseTime
// bootstrap register (0x01CC) and returns it as a time.Duration in
// milliseconds. The spec caps this at 300 ms. Returns an error if the
// register is unreadable or the value is zero.
func (g *GVCP) MaximumDeviceResponseTime() (time.Duration, error) {
	v, err := g.ReadReg(AbrmMaximumDeviceResponseTime)
	if err != nil {
		return 0, err
	}
	if v == 0 {
		return 0, errors.New("gige: MaximumDeviceResponseTime is zero")
	}
	return time.Duration(v) * time.Millisecond, nil
}

// SetTimeout adjusts the GVCP transaction timeout. This allows the caller to
// tighten or loosen the deadline after construction (e.g. based on
// MaximumDeviceResponseTime). The change takes effect for all subsequent
// ReadReg/WriteReg/ReadMem/WriteMem calls.
func (g *GVCP) SetTimeout(d time.Duration) {
	if d <= 0 {
		return
	}
	g.mu.Lock()
	g.timeout = d
	g.mu.Unlock()
}

// RequestResend sends PACKETRESEND_CMD for inclusive packet_id range [first, last].
// Fire-and-forget: cameras do not ACK; resent GVSP packets arrive on the stream socket.
func (g *GVCP) RequestResend(streamChannel uint16, blockID uint64, first, last uint32, extended bool) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.closed || g.conn == nil {
		return errors.New("gige: gvcp closed")
	}
	if last < first {
		return nil
	}
	req := EncodePacketResend(g.nextIDLocked(), streamChannel, blockID, first, last, extended)
	if _, err := g.conn.Write(req); err != nil {
		return fmt.Errorf("gige: packet resend: %w", err)
	}
	return nil
}
