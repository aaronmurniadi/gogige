package gvsp

import (
	"net"
	"testing"
)

func TestPacketSizeForMTU(t *testing.T) {
	if PacketSizeForMTU(0) != DefaultMTU {
		t.Fatal("fallback")
	}
	if PacketSizeForMTU(9000) != 9000 {
		t.Fatal("jumbo")
	}
	if PacketSizeForMTU(20000) != 16384 {
		t.Fatal("clamp high")
	}
}

func TestPathMTULocalSubnet(t *testing.T) {
	ifaces, err := net.Interfaces()
	if err != nil {
		t.Fatal(err)
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 || iface.MTU <= 0 {
			continue
		}
		addrs, _ := iface.Addrs()
		for _, a := range addrs {
			ipnet, ok := a.(*net.IPNet)
			if !ok || ipnet.IP.To4() == nil {
				continue
			}
			// Pick another address in the same subnet (network + 1 host octet flip-ish).
			dst := make(net.IP, len(ipnet.IP.To4()))
			copy(dst, ipnet.IP.To4())
			dst[3] ^= 1
			if !ipnet.Contains(dst) {
				continue
			}
			mtu := PathMTU(dst)
			if mtu != iface.MTU {
				t.Fatalf("PathMTU(%s)=%d want iface %s MTU %d", dst, mtu, iface.Name, iface.MTU)
			}
			return
		}
	}
	t.Skip("no usable IPv4 interface for PathMTU test")
}

func TestSetRecvBuffer(t *testing.T) {
	s, err := ListenStream(0)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	got := s.RecvBuffer()
	if got <= 0 {
		t.Fatalf("rcvbuf=%d", got)
	}
}
