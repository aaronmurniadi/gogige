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
	if packetSize <= 0 {
		packetSize = 1500
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
	if err := g.WriteReg(Stream0PacketSize, uint32(packetSize)&0xffff); err != nil {
		return err
	}
	if cmd != nil && cmd.Has("AcquisitionStart") {
		return cmd.Execute("AcquisitionStart")
	}
	return nil
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
