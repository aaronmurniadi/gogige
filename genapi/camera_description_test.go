package genapi

import (
	"fmt"
	"testing"
)

func TestParseLocalGenICamURL(t *testing.T) {
	name, addr, length, err := parseLocalGenICamURL("S5_BS.zip;f0000000;cc5c")
	if err != nil {
		t.Fatal(err)
	}
	if name != "S5_BS.zip" {
		t.Fatalf("name=%q", name)
	}
	if addr != 0xf0000000 {
		t.Fatalf("addr=%x", addr)
	}
	if length != 0xcc5c {
		t.Fatalf("length=%x", length)
	}
	_, _, _, err = parseLocalGenICamURL("file.xml;0x1000;0x200")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFetchDeviceXMLLocalCaseInsensitive(t *testing.T) {
	xmlBody := []byte(`<?xml version="1.0"?><RegisterDescription/>`)
	port := &memPort{mem: map[uint32]byte{}}
	for i, b := range xmlBody {
		port.mem[0xf0000000+uint32(i)] = b
	}
	out, err := fetchDeviceXML(port, "local:S5_BS.xml;f0000000;"+fmt.Sprintf("%x", len(xmlBody)))
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != string(xmlBody) {
		t.Fatalf("got %q", out)
	}
}
