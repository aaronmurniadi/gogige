package gvcp

import (
	"encoding/binary"
	"fmt"
)

// DeviceByteOrder returns the endianness used for non-bootstrap device
// register byte arrays (GenCP ImplementationEndianness). Bootstrap / GVCP
// protocol fields are always big-endian.
func (g *GVCP) DeviceByteOrder() binary.ByteOrder {
	if g == nil || g.deviceOrder == nil {
		return binary.BigEndian
	}
	return g.deviceOrder
}

// SetDeviceByteOrder overrides ImplementationEndianness (tests / known devices).
func (g *GVCP) SetDeviceByteOrder(order binary.ByteOrder) {
	if g == nil || order == nil {
		return
	}
	g.mu.Lock()
	g.deviceOrder = order
	g.mu.Unlock()
}

// SyncImplementationEndianness reads GenCP ABRM 0x020C and caches the result.
// Only the two defined values are accepted (0 = BE, 0xFFFFFFFF = LE). Any other
// value — including GigE Vision First URL bytes that overlap 0x020C — leaves
// the current order unchanged.
func (g *GVCP) SyncImplementationEndianness() error {
	if g == nil {
		return fmt.Errorf("gige: gvcp nil")
	}
	v, err := g.ReadReg(AbrmImplementationEndianness)
	if err != nil {
		return err
	}
	applyImplementationEndianness(g, v)
	return nil
}

func applyImplementationEndianness(g *GVCP, v uint32) {
	switch v {
	case EndiannessBig:
		g.SetDeviceByteOrder(binary.BigEndian)
	case EndiannessLittle:
		g.SetDeviceByteOrder(binary.LittleEndian)
	}
}

// EncodeDeviceUint32 writes v using DeviceByteOrder (register-map sync helper).
func (g *GVCP) EncodeDeviceUint32(v uint32) []byte {
	var b [4]byte
	g.DeviceByteOrder().PutUint32(b[:], v)
	return b[:]
}

// DecodeDeviceUint32 reads a 4-byte device register using DeviceByteOrder.
func (g *GVCP) DecodeDeviceUint32(b []byte) uint32 {
	if len(b) < 4 {
		return 0
	}
	return g.DeviceByteOrder().Uint32(b)
}
