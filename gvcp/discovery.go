package gvcp

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

// DISCOVERY_ACK body mirrors ABRM from 0x0000 (GigE Vision). Minimum useful size
// covers through User-defined name (0x00E8 + 16).
const (
	discoveryAckMinSize = 0x00F8

	abrmMACHigh      = 0x0008
	abrmMACLow       = 0x000C
	abrmCurrentIP    = 0x0024
	abrmManufacturer = 0x0048
	abrmModel        = 0x0068
	abrmSerial       = 0x00D8
	abrmUserName     = 0x00E8

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

// Discover broadcasts DISCOVERY_CMD and returns peers that answered on the control port.
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

	conn, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4zero, Port: 0})
	if err != nil {
		return nil, fmt.Errorf("gige: discover listen: %w", err)
	}
	defer conn.Close()
	_ = conn.SetDeadline(deadline)
	_ = conn.SetReadBuffer(1 << 20)

	pkt := encodeGVCPHeader(gvcpCmdDiscovery, 0, 1)
	dst := &net.UDPAddr{IP: net.IPv4bcast, Port: gvcpPort}
	if _, err := conn.WriteToUDP(pkt, dst); err != nil {
		return nil, fmt.Errorf("gige: discover broadcast: %w", err)
	}

	seen := map[string]struct{}{}
	var out []DiscoveredDevice
	buf := make([]byte, 2048)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Timeout() {
				break
			}
			if errors.Is(err, context.DeadlineExceeded) || ctx.Err() != nil {
				break
			}
			break
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
		key := dev.IP
		if key == "" {
			key = ip.String()
			dev.IP = key
		}
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, dev)
	}
	return out, nil
}

// parseDiscoveryAck extracts ABRM identity fields from a DISCOVERY_ACK payload.
// sourceIP is the UDP peer address (fallback when Current IP is unset).
func parseDiscoveryAck(payload []byte, sourceIP string) DiscoveredDevice {
	d := DiscoveredDevice{IP: sourceIP}
	if len(payload) < discoveryAckMinSize {
		return d
	}
	hi := binary.BigEndian.Uint32(payload[abrmMACHigh:])
	lo := binary.BigEndian.Uint32(payload[abrmMACLow:])
	d.MAC = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		byte(hi>>8), byte(hi),
		byte(lo>>24), byte(lo>>16), byte(lo>>8), byte(lo))

	if ip4 := payload[abrmCurrentIP : abrmCurrentIP+4]; !(ip4[0] == 0 && ip4[1] == 0 && ip4[2] == 0 && ip4[3] == 0) {
		d.IP = net.IP(ip4).String()
	}
	d.Manufacturer = cString(payload[abrmManufacturer : abrmManufacturer+abrmName32])
	d.Model = cString(payload[abrmModel : abrmModel+abrmName32])
	d.Serial = cString(payload[abrmSerial : abrmSerial+abrmName16])
	d.UserName = cString(payload[abrmUserName : abrmUserName+abrmName16])
	return d
}

func cString(b []byte) string {
	n := 0
	for n < len(b) && b[n] != 0 {
		n++
	}
	return strings.TrimSpace(string(b[:n]))
}
