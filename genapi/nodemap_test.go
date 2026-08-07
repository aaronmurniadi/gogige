package genapi

import "testing"

type memPort struct {
	regs map[uint32]uint32
	mem  map[uint32]byte
}

func (m *memPort) ReadReg(addr uint32) (uint32, error) {
	return m.regs[addr], nil
}
func (m *memPort) WriteReg(addr, value uint32) error {
	if m.regs == nil {
		m.regs = map[uint32]uint32{}
	}
	m.regs[addr] = value
	return nil
}
func (m *memPort) ReadMem(addr uint32, n int) ([]byte, error) {
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = m.mem[addr+uint32(i)]
	}
	return out, nil
}
func (m *memPort) WriteMem(addr uint32, data []byte) error {
	if m.mem == nil {
		m.mem = map[uint32]byte{}
	}
	for i, b := range data {
		m.mem[addr+uint32(i)] = b
	}
	if len(data) >= 4 {
		_ = m.WriteReg(addr, uint32(data[0])<<24|uint32(data[1])<<16|uint32(data[2])<<8|uint32(data[3]))
	}
	return nil
}

func TestParseNodeMapSetupFeatures(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="CommEnableReg"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Boolean Name="CommEnable"><pValue>CommEnableReg</pValue></Boolean>
  <IntReg Name="TransferWorkModeReg"><Address>0x1008</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Enumeration Name="TransferWorkMode"><EnumEntry Name="TCPServer"><Value>1</Value></EnumEntry><pValue>TransferWorkModeReg</pValue></Enumeration>
  <IntReg Name="TCPPortReg"><Address>0x100C</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Integer Name="TCPPort"><pValue>TCPPortReg</pValue></Integer>
  <IntReg Name="TriggerSoftwareReg"><Address>0x2008</Address><Length>4</Length><AccessMode>WO</AccessMode></IntReg>
  <Command Name="TriggerSoftware"><pValue>TriggerSoftwareReg</pValue><Value>1</Value></Command>
</RegisterDescription>`

	port := &memPort{regs: map[uint32]uint32{}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	if err := nm.SetBoolean("CommEnable", false); err != nil {
		t.Fatal(err)
	}
	if port.regs[0x1000] != 0 {
		t.Fatalf("CommEnable reg=%d", port.regs[0x1000])
	}
	if err := nm.SetString("TransferWorkMode", "TCPServer"); err != nil {
		t.Fatal(err)
	}
	if port.regs[0x1008] != 1 {
		t.Fatalf("TransferWorkMode=%d", port.regs[0x1008])
	}
	if err := nm.SetInteger("TCPPort", 3100); err != nil {
		t.Fatal(err)
	}
	if port.regs[0x100C] != 3100 {
		t.Fatalf("TCPPort=%d", port.regs[0x100C])
	}
	if err := nm.Execute("TriggerSoftware"); err != nil {
		t.Fatal(err)
	}
	if port.regs[0x2008] != 1 {
		t.Fatalf("TriggerSoftware=%d", port.regs[0x2008])
	}
}

func TestMaskedIntRegRMW(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <MaskedIntReg Name="CommEnableReg">
    <Address>0x2000</Address><Length>4</Length><AccessMode>RW</AccessMode>
    <LSB>1</LSB><MSB>1</MSB>
  </MaskedIntReg>
  <Boolean Name="CommEnable"><pValue>CommEnableReg</pValue></Boolean>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{0x2000: 0xffff0001}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	if err := nm.SetBoolean("CommEnable", true); err != nil {
		t.Fatal(err)
	}
	if port.regs[0x2000] != 0xffff0003 {
		t.Fatalf("reg=%x", port.regs[0x2000])
	}
	if err := nm.SetBoolean("CommEnable", false); err != nil {
		t.Fatal(err)
	}
	if port.regs[0x2000] != 0xffff0001 {
		t.Fatalf("reg after clear=%x", port.regs[0x2000])
	}
}

func TestRejectZeroAddress(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <Boolean Name="CommEnable"><pValue>MissingReg</pValue></Boolean>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	if err := nm.SetBoolean("CommEnable", false); err == nil {
		t.Fatal("expected error")
	}
}

func TestPAddressSum(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription xmlns="http://www.genapi.org/GenApi/Version_1_1">
  <Integer Name="RegBase"><Value>0x10000</Value></Integer>
  <IntReg Name="TriggerSourceReg">
    <Address>0x20</Address>
    <Length>4</Length>
    <AccessMode>RW</AccessMode>
    <pAddress>RegBase</pAddress>
  </IntReg>
  <Enumeration Name="TriggerSource">
    <EnumEntry Name="Software"><Value>3</Value></EnumEntry>
    <pValue>TriggerSourceReg</pValue>
  </Enumeration>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	if err := nm.SetString("TriggerSource", "Software"); err != nil {
		t.Fatal(err)
	}
	if port.regs[0x10020] != 3 {
		t.Fatalf("regs=%v", port.regs)
	}
}
