package genapi

import (
	"testing"
)

// StructReg §2.8.6: entries expand into MaskedIntReg nodes that inherit the
// StructReg's shared elements. Uses the spec's Format 7 example (Address 0x14,
// pAddress VFormat7ModeCsrBase, Length 4, RO, BigEndian; per-entry Bit).
func TestStructRegExpansion(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="VFormat7ModeCsrBase"><Address>0x8000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <StructReg Comment="VFormat7 support inquiry registers">
    <pAddress>VFormat7ModeCsrBase</pAddress>
    <Address>0x14</Address>
    <Length>4</Length>
    <AccessMode>RO</AccessMode>
    <pPort>Device</pPort>
    <Endianess>BigEndian</Endianess>
    <StructEntry Name="VFormat7Mono8InqReg"><Bit>31</Bit></StructEntry>
    <StructEntry Name="VFormat7InqReg"><Bit>29</Bit></StructEntry>
    <StructEntry Name="VFormat7Mono16InqReg"><Bit>24</Bit></StructEntry>
  </StructReg>
</RegisterDescription>`

	port := &memPort{regs: map[uint32]uint32{0x8000: 0x8000}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}

	// The StructReg itself carries no Name and must not be registered.
	if nm.Has("VFormat7Mono8InqReg") == false {
		t.Fatal("expanded entry missing")
	}
	// Inherited elements visible on a single-bit entry.
	p := nm.nodes["VFormat7Mono8InqReg"]
	if p.Kind != "MaskedIntReg" {
		t.Fatalf("kind = %s", p.Kind)
	}
	if p.Address != 0x14 {
		t.Fatalf("inherited address = 0x%x", p.Address)
	}
	if len(p.PAddresses) != 1 || p.PAddresses[0] != "VFormat7ModeCsrBase" {
		t.Fatalf("inherited pAddress = %v", p.PAddresses)
	}
	if p.Length != 4 || p.Access != "RO" || p.Endianess != "BigEndian" {
		t.Fatalf("inherited len/access/endianess = %d/%s/%s", p.Length, p.Access, p.Endianess)
	}
	if !p.HasMask || p.LSB != 31 || p.MSB != 31 {
		t.Fatalf("bit mask = lsb %d msb %d has %v", p.LSB, p.MSB, p.HasMask)
	}
	if n := nm.nodes["VFormat7InqReg"]; n.LSB != 29 || !n.HasMask {
		t.Fatalf("entry VFormat7InqReg bits = %d-%d", n.LSB, n.MSB)
	}
	if n := nm.nodes["VFormat7Mono16InqReg"]; n.LSB != 24 || !n.HasMask {
		t.Fatalf("entry VFormat7Mono16InqReg bits = %d-%d", n.LSB, n.MSB)
	}

	// Functional read: effective address = 0x8000 + 0x14 = 0x8014.
	// BigEndian: byte 0 at 0x8014 holds bits 31..24.
	bin := map[string]uint32{
		"VFormat7Mono8InqReg":  0x80000000,
		"VFormat7InqReg":       0x20000000,
		"VFormat7Mono16InqReg": 0x01000000,
	}
	for entry, bit := range bin {
		port.mem = map[uint32]byte{}
		port.mem[0x8014] = byte(bit >> 24)
		port.mem[0x8015] = byte(bit >> 16)
		port.mem[0x8016] = byte(bit >> 8)
		port.mem[0x8017] = byte(bit)
		v, err := nm.ReadInteger(entry)
		if err != nil {
			t.Fatalf("ReadInteger(%s): %v", entry, err)
		}
		if v != 1 {
			t.Fatalf("ReadInteger(%s) = %d, want 1", entry, v)
		}
	}
	// Buttoned bit is zero.
	port.mem = map[uint32]byte{}
	if v, err := nm.ReadInteger("VFormat7Mono8InqReg"); err != nil || v != 0 {
		t.Fatalf("bit 31 unset read = %d, %v", v, err)
	}
}

// An entry that defines its own Address overrides the inherited parent Address
// rather than adding to it; likewise an entry-defined LSB/MSB replaces <Bit>.
func TestStructRegEntryOverride(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <StructReg Comment="override demo">
    <Address>0x10</Address>
    <StructEntry Name="SharedAddr"><Bit>7</Bit></StructEntry>
    <StructEntry Name="OwnAddr"><Address>0x20</Address><LSB>8</LSB><MSB>15</MSB></StructEntry>
  </StructReg>
</RegisterDescription>`
	port := &memPort{}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	if p := nm.nodes["SharedAddr"]; p.Address != 0x10 || p.LSB != 7 || p.MSB != 7 {
		t.Fatalf("shared = addr 0x%x lsb %d msb %d", p.Address, p.LSB, p.MSB)
	}
	p := nm.nodes["OwnAddr"]
	if p.Address != 0x20 {
		t.Fatalf("own address = 0x%x, want 0x20 (override, not sum 0x30)", p.Address)
	}
	if p.LSB != 8 || p.MSB != 15 {
		t.Fatalf("own lsm/msb = %d/%d", p.LSB, p.MSB)
	}
}
