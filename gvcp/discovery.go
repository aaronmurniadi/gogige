package gvcp

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// DISCOVERY_ACK body mirrors GigE Vision ABRM from 0x0000. Minimum useful size
// covers through User-defined name (0x00E8 + 16).
const (
	discoveryAckMinSize = 0x00F8

	abrmName32 = 32
	abrmName16 = 16
)

// DiscoveredDevice is a peer that answered DISCOVERY_CMD.
type DiscoveredDevice struct {
	IP           string
	MAC          string
	Manufacturer string
	Model        string
	Serial       string
	UserName     string
}

// discoverIface is one IPv4 address to bind and its directed broadcast.
type discoverIface struct {
	src   net.IP
	bcast net.IP
}

// Discover broadcasts DISCOVERY_CMD on every up IPv4 interface and returns peers
// that answered. Multi-homed hosts must bind per interface: a single socket on
// 0.0.0.0 sends limited broadcast via the default route, which often misses the
// camera NIC (GigE Vision / Aravis practice).
func Discover(ctx context.Context, timeout time.Duration) ([]DiscoveredDevice, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if timeout <= 0 {
		timeout = 2 * time.Second
	}
	deadline, ok := ctx.Deadline()
	if !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
		deadline, _ = ctx.Deadline()
	}

	ifaces := listDiscoverIfaces()
	if len(ifaces) == 0 {
		ifaces = []discoverIface{{src: net.IPv4zero, bcast: net.IPv4bcast}}
	}

	pkt := encodeGVCPHeader(gvcpCmdDiscovery, 0, 1)
	limited := &net.UDPAddr{IP: net.IPv4bcast, Port: gvcpPort}

	type result struct {
		dev DiscoveredDevice
	}
	outCh := make(chan result, 16)

	var wg sync.WaitGroup
	for _, iface := range ifaces {
		conn, err := net.ListenUDP("udp4", &net.UDPAddr{IP: iface.src, Port: 0})
		if err != nil {
			continue
		}
		_ = conn.SetDeadline(deadline)
		_ = conn.SetReadBuffer(1 << 20)

		directed := &net.UDPAddr{IP: iface.bcast, Port: gvcpPort}
		_, _ = conn.WriteToUDP(pkt, limited)
		if !iface.bcast.Equal(net.IPv4bcast) {
			_, _ = conn.WriteToUDP(pkt, directed)
		}

		wg.Add(1)
		go func(conn *net.UDPConn) {
			defer wg.Done()
			defer conn.Close()
			buf := make([]byte, 2048)
			for {
				n, addr, err := conn.ReadFromUDP(buf)
				if err != nil {
					return
				}
				if n < gvcpHeaderSize || addr == nil || addr.IP == nil {
					continue
				}
				cmd := binary.BigEndian.Uint16(buf[2:4])
				if cmd != gvcpCmdDiscoveryAck {
					continue
				}
				ip := addr.IP.To4()
				if ip == nil {
					continue
				}
				dev := parseDiscoveryAck(buf[gvcpHeaderSize:n], ip.String())
				if dev.IP == "" {
					dev.IP = ip.String()
				}
				outCh <- result{dev: dev}
			}
		}(conn)
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	seen := map[string]struct{}{}
	var out []DiscoveredDevice
	for r := range outCh {
		key := r.dev.IP
		if key == "" {
			continue
		}
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, r.dev)
	}

	if len(out) == 0 && ctx.Err() != nil && !errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return nil, ctx.Err()
	}
	return out, nil
}

func listDiscoverIfaces() []discoverIface {
	netIfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	var out []discoverIface
	for _, ni := range netIfaces {
		if ni.Flags&net.FlagUp == 0 || ni.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := ni.Addrs()
		if err != nil {
			continue
		}
		for _, a := range addrs {
			ipnet, ok := a.(*net.IPNet)
			if !ok {
				continue
			}
			ip4 := ipnet.IP.To4()
			if ip4 == nil || !ip4.IsGlobalUnicast() {
				continue
			}
			out = append(out, discoverIface{
				src:   append(net.IP(nil), ip4...),
				bcast: ipv4Broadcast(ip4, ipnet.Mask),
			})
		}
	}
	return out
}

func ipv4Broadcast(ip net.IP, mask net.IPMask) net.IP {
	ip4 := ip.To4()
	m4 := mask
	if len(m4) == net.IPv6len {
		m4 = m4[12:]
	}
	if ip4 == nil || len(m4) != net.IPv4len {
		return net.IPv4bcast
	}
	bcast := make(net.IP, net.IPv4len)
	for i := 0; i < net.IPv4len; i++ {
		bcast[i] = ip4[i] | ^m4[i]
	}
	return bcast
}

// parseDiscoveryAck extracts ABRM identity fields from a DISCOVERY_ACK payload.
// sourceIP is the UDP peer address (fallback when Current IP is unset).
func parseDiscoveryAck(payload []byte, sourceIP string) DiscoveredDevice {
	d := DiscoveredDevice{IP: sourceIP}
	if len(payload) < discoveryAckMinSize {
		return d
	}
	hi := binary.BigEndian.Uint32(payload[gevMACHigh:])
	lo := binary.BigEndian.Uint32(payload[gevMACLow:])
	d.MAC = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		byte(hi>>8), byte(hi),
		byte(lo>>24), byte(lo>>16), byte(lo>>8), byte(lo))

	if ip4 := payload[gevCurrentIP : gevCurrentIP+4]; !(ip4[0] == 0 && ip4[1] == 0 && ip4[2] == 0 && ip4[3] == 0) {
		d.IP = net.IP(ip4).String()
	}
	d.Manufacturer = cString(payload[gevManufacturerName : gevManufacturerName+abrmName32])
	d.Model = cString(payload[gevModelName : gevModelName+abrmName32])
	d.Serial = cString(payload[gevSerialNumber : gevSerialNumber+abrmName16])
	d.UserName = cString(payload[gevUserDefinedName : gevUserDefinedName+abrmName16])
	return d
}

func cString(b []byte) string {
	n := 0
	for n < len(b) && b[n] != 0 {
		n++
	}
	return strings.TrimSpace(string(b[:n]))
}
