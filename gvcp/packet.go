package gvcp

import "encoding/binary"

const (
	gvcpPort = 3956

	gvcpPacketTypeCMD   = 0x42
	gvcpPacketTypeACK   = 0x00
	gvcpPacketTypeError = 0x80

	gvcpFlagAckRequired = 0x01

	gvcpCmdReadReg      = 0x0080
	gvcpCmdReadRegAck   = 0x0081
	gvcpCmdWriteReg     = 0x0082
	gvcpCmdWriteRegAck  = 0x0083
	gvcpCmdReadMem      = 0x0084
	gvcpCmdReadMemAck   = 0x0085
	gvcpCmdWriteMem     = 0x0086
	gvcpCmdWriteMemAck  = 0x0087
	gvcpCmdPendingAck   = 0x0089
	gvcpCmdBye          = 0x0004
	gvcpCmdDiscovery    = 0x0002
	gvcpCmdDiscoveryAck = 0x0003
	gvcpCmdPacketResend = 0x0040

	gvcpFlagExtendedIDs = 0x10

	gvcpDataSizeMax = 512
	gvcpHeaderSize  = 8

	gvspPacketIDMask = 0x00ffffff
)

func encodeGVCPHeader(cmd uint16, dataSize uint16, id uint16) []byte {
	return encodeGVCPHeaderFlags(cmd, dataSize, id, gvcpFlagAckRequired)
}

func encodeGVCPHeaderFlags(cmd uint16, dataSize uint16, id uint16, flags byte) []byte {
	b := make([]byte, gvcpHeaderSize)
	b[0] = gvcpPacketTypeCMD
	b[1] = flags
	binary.BigEndian.PutUint16(b[2:], cmd)
	binary.BigEndian.PutUint16(b[4:], dataSize)
	binary.BigEndian.PutUint16(b[6:], id)
	return b
}

// EncodePacketResend builds a GVCP PACKETRESEND_CMD (0x0040).
// No ACK is required. extended selects GigE Vision 2.0 64-bit block_id layout.
func EncodePacketResend(reqID, streamChannel uint16, blockID uint64, first, last uint32, extended bool) []byte {
	if extended {
		const dataSize = 20
		req := encodeGVCPHeaderFlags(gvcpCmdPacketResend, dataSize, reqID, gvcpFlagExtendedIDs)
		var data [20]byte
		binary.BigEndian.PutUint32(data[0:], uint32(streamChannel)<<16)
		binary.BigEndian.PutUint32(data[4:], first)
		binary.BigEndian.PutUint32(data[8:], last)
		binary.BigEndian.PutUint64(data[12:], blockID)
		return append(req, data[:]...)
	}
	const dataSize = 12
	req := encodeGVCPHeaderFlags(gvcpCmdPacketResend, dataSize, reqID, 0)
	var data [12]byte
	binary.BigEndian.PutUint32(data[0:], uint32(streamChannel)<<16|uint32(blockID&0xffff))
	binary.BigEndian.PutUint32(data[4:], first&gvspPacketIDMask)
	binary.BigEndian.PutUint32(data[8:], last&gvspPacketIDMask)
	return append(req, data[:]...)
}

func gvcpErrorName(code byte) string {
	switch code {
	case 0x01:
		return "NOT_IMPLEMENTED"
	case 0x02:
		return "INVALID_PARAMETER"
	case 0x03:
		return "INVALID_ACCESS"
	case 0x04:
		return "WRITE_PROTECT"
	case 0x05:
		return "BAD_ALIGNMENT"
	case 0x06:
		return "ACCESS_DENIED"
	case 0x07:
		return "BUSY"
	default:
		return "UNKNOWN"
	}
}

// encode helpers used by tests
func encodeReadReg(addr uint32, id uint16) []byte {
	req := encodeGVCPHeader(gvcpCmdReadReg, 4, id)
	var ab [4]byte
	binary.BigEndian.PutUint32(ab[:], addr)
	return append(req, ab[:]...)
}

func encodeWriteReg(addr, value uint32, id uint16) []byte {
	req := encodeGVCPHeader(gvcpCmdWriteReg, 8, id)
	var ab [8]byte
	binary.BigEndian.PutUint32(ab[0:], addr)
	binary.BigEndian.PutUint32(ab[4:], value)
	return append(req, ab[:]...)
}

func encodeReadMem(addr uint32, size uint32, id uint16) []byte {
	aligned := ((size + 3) / 4) * 4
	req := encodeGVCPHeader(gvcpCmdReadMem, 8, id)
	var ab [8]byte
	binary.BigEndian.PutUint32(ab[0:], addr)
	binary.BigEndian.PutUint32(ab[4:], aligned)
	return append(req, ab[:]...)
}
