package genapi

import (
	"fmt"
	"testing"
)

type urlReaderMock struct {
	firstURL         string
	manifestTableURL string
	mem              map[uint32]byte
}

func (m *urlReaderMock) ReadReg(addr uint32) (uint32, error) { return 0, nil }
func (m *urlReaderMock) WriteReg(addr, value uint32) error   { return nil }
func (m *urlReaderMock) ReadMem(addr uint32, n int) ([]byte, error) {
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = m.mem[addr+uint32(i)]
	}
	return out, nil
}
func (m *urlReaderMock) WriteMem(addr uint32, data []byte) error { return nil }
func (m *urlReaderMock) FirstURL() (string, error)               { return m.firstURL, nil }
func (m *urlReaderMock) ManifestTableURL() (string, error)       { return m.manifestTableURL, nil }

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

func TestFetchXMLPrefersManifestTable(t *testing.T) {
	xmlBody := []byte(`<?xml version="1.0"?><RegisterDescription/>`)
	mock := &urlReaderMock{
		manifestTableURL: "local:manifest.xml;1000;" + fmt.Sprintf("%x", len(xmlBody)),
		mem:              map[uint32]byte{},
	}
	for i, b := range xmlBody {
		mock.mem[0x1000+uint32(i)] = b
	}
	out, err := FetchXML(mock)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != string(xmlBody) {
		t.Fatalf("got %q", out)
	}
}

func TestFetchXMLFallsBackToFirstURL(t *testing.T) {
	xmlBody := []byte(`<?xml version="1.0"?><RegisterDescription/>`)
	mock := &urlReaderMock{
		firstURL: "local:first.xml;2000;" + fmt.Sprintf("%x", len(xmlBody)),
		mem:      map[uint32]byte{},
	}
	for i, b := range xmlBody {
		mock.mem[0x2000+uint32(i)] = b
	}
	out, err := FetchXML(mock)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != string(xmlBody) {
		t.Fatalf("got %q", out)
	}
}

func TestFetchXMLManifestTableEmptyFallsBack(t *testing.T) {
	xmlBody := []byte(`<?xml version="1.0"?><RegisterDescription/>`)
	mock := &urlReaderMock{
		firstURL:         "local:first.xml;2000;" + fmt.Sprintf("%x", len(xmlBody)),
		manifestTableURL: "",
		mem:              map[uint32]byte{},
	}
	for i, b := range xmlBody {
		mock.mem[0x2000+uint32(i)] = b
	}
	out, err := FetchXML(mock)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != string(xmlBody) {
		t.Fatalf("got %q", out)
	}
}
