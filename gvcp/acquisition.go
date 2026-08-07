package gvcp

import (
	"encoding/binary"
	"errors"
	"net"
)

type Commander interface {
	Has(name string) bool
	Execute(name string) error
}

func StartAcquisition(g *GVCP, cmd Commander, destIP net.IP, destPort, packetSize int) error {
	if g == nil {
		return errors.New("gige: no gvcp")
	}
	ip4 := destIP.To4()
	if ip4 == nil {
		return errors.New("gige: need IPv4 stream destination")
	}
	ipVal := binary.BigEndian.Uint32(ip4)
	if err := g.WriteReg(Stream0IP, ipVal); err != nil {
		return err
	}
	if err := g.WriteReg(Stream0Port, uint32(destPort)); err != nil {
		return err
	}
	if _, err := NegotiatePacketSize(g, packetSize); err != nil {
		return err
	}
	if cmd != nil && cmd.Has("AcquisitionStart") {
		return cmd.Execute("AcquisitionStart")
	}
	return nil
}

// NegotiatePacketSize programs GevSCPSPacketSize (low 16 bits of 0x0D04).
// want is typically the path MTU; falls back via binary search when the device
// rejects larger sizes (common when the camera NIC is 1500 despite host jumbo).
func NegotiatePacketSize(p Port, want int) (int, error) {
	if p == nil {
		return 0, errors.New("gige: no gvcp")
	}
	if want < 576 {
		want = 1500
	}
	if want > 16384 {
		want = 16384
	}
	cur, err := p.ReadReg(Stream0PacketSize)
	if err != nil {
		return 0, err
	}
	write := func(sz int) error {
		val := (cur &^ 0xffff) | (uint32(sz) & 0xffff)
		return p.WriteReg(Stream0PacketSize, val)
	}
	if err := write(want); err == nil {
		return want, nil
	}
	lo, hi, best := 576, want-1, 0
	if err := write(1500); err == nil {
		best = 1500
	}
	for lo <= hi {
		mid := (lo + hi) / 2
		if err := write(mid); err == nil {
			best = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if best == 0 {
		return 0, errors.New("gige: GevSCPSPacketSize negotiate failed")
	}
	if err := write(best); err != nil {
		return 0, err
	}
	return best, nil
}

func StopAcquisition(cmd Commander) error {
	if cmd == nil {
		return nil
	}
	if cmd.Has("AcquisitionStop") {
		return cmd.Execute("AcquisitionStop")
	}
	return nil
}
