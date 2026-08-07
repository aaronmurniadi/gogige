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

	gvcpDataSizeMax = 512
	gvcpHeaderSize  = 8
)

func encodeGVCPHeader(cmd uint16, dataSize uint16, id uint16) []byte {
	b := make([]byte, gvcpHeaderSize)
	b[0] = gvcpPacketTypeCMD
	b[1] = gvcpFlagAckRequired
	binary.BigEndian.PutUint16(b[2:], cmd)
	binary.BigEndian.PutUint16(b[4:], dataSize)
	binary.BigEndian.PutUint16(b[6:], id)
	return b
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
