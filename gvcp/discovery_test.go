package gvcp

import (
	"encoding/binary"
	"net"
	"testing"
	"time"
)

func TestParseDiscoveryAck(t *testing.T) {
	payload := make([]byte, discoveryAckMinSize)
	// MAC 00:11:22:33:44:55 → high=0x0011, low=0x22334455
	binary.BigEndian.PutUint32(payload[abrmMACHigh:], 0x00000011)
	binary.BigEndian.PutUint32(payload[abrmMACLow:], 0x22334455)
	copy(payload[abrmCurrentIP:], net.IPv4(192, 168, 1, 42).To4())
	copy(payload[abrmManufacturer:], []byte("AcmeCam\x00"))
	copy(payload[abrmModel:], []byte("Model-X\x00"))
	copy(payload[abrmSerial:], []byte("SN123\x00"))
	copy(payload[abrmUserName:], []byte("line1\x00"))

	d := parseDiscoveryAck(payload, "10.0.0.1")
	if d.IP != "192.168.1.42" {
		t.Fatalf("IP: got %q", d.IP)
	}
	if d.MAC != "00:11:22:33:44:55" {
		t.Fatalf("MAC: got %q", d.MAC)
	}
	if d.Manufacturer != "AcmeCam" || d.Model != "Model-X" {
		t.Fatalf("names: %q %q", d.Manufacturer, d.Model)
	}
	if d.Serial != "SN123" || d.UserName != "line1" {
		t.Fatalf("serial/user: %q %q", d.Serial, d.UserName)
	}
}

func TestParseDiscoveryAckShortFallsBackToSourceIP(t *testing.T) {
	d := parseDiscoveryAck([]byte{1, 2, 3}, "10.1.2.3")
	if d.IP != "10.1.2.3" || d.MAC != "" {
		t.Fatalf("got %+v", d)
	}
}

func TestHeartbeatStop(t *testing.T) {
	// Dial a black-hole UDP port; StartHeartbeat must still be stoppable.
	g, err := DialGVCP("127.0.0.1", 50*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	defer g.Close()
	h := g.StartHeartbeat()
	if h == nil {
		t.Fatal("nil heartbeat")
	}
	done := make(chan struct{})
	go func() {
		h.Stop()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("Stop hung")
	}
	h.Stop() // idempotent
}
