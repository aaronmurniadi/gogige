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

func TestIntrospection(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription xmlns="http://www.genapi.org/GenApi/Version_1_1">
  <IntReg Name="PixelFormatReg"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Enumeration Name="PixelFormat">
    <EnumEntry Name="Mono8"><Value>0x1080001</Value></EnumEntry>
    <EnumEntry Name="RGB8"><Value>0x2180014</Value></EnumEntry>
    <EnumEntry Name="YUV422_8"><Value>0x2100032</Value></EnumEntry>
    <pValue>PixelFormatReg</pValue>
  </Enumeration>
  <Integer Name="Width"><pValue>WidthReg</pValue></Integer>
  <IntReg Name="WidthReg"><Address>0x2000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{0x1000: 0x02180014, 0x2000: 1920}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	if got := nm.Has("PixelFormat"); !got {
		t.Fatal("PixelFormat should exist")
	}
	if got := nm.Has("Nope"); got {
		t.Fatal("Nope should not exist")
	}
	if got := nm.Kind("PixelFormat"); got != "Enumeration" {
		t.Fatalf("Kind=%q", got)
	}
	if got := nm.Kind("Nope"); got != "" {
		t.Fatalf("Kind(Nope)=%q", got)
	}
	entries, err := nm.EnumEntries("PixelFormat")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"Mono8", "RGB8", "YUV422_8"}
	if len(entries) != len(want) || entries[0] != want[0] || entries[1] != want[1] || entries[2] != want[2] {
		t.Fatalf("entries=%v want=%v", entries, want)
	}
	if _, err := nm.EnumEntries("Width"); err == nil {
		t.Fatal("EnumEntries on non-enum should fail")
	}
	cur, err := nm.CurrentEnum("PixelFormat")
	if err != nil {
		t.Fatal(err)
	}
	if cur != "RGB8" {
		t.Fatalf("CurrentEnum=%q", cur)
	}
	w, err := nm.ReadInteger("Width")
	if err != nil {
		t.Fatal(err)
	}
	if w != 1920 {
		t.Fatalf("Width=%d", w)
	}
	if _, err := nm.ReadInteger("PixelFormat"); err != nil {
		t.Fatal(err)
	}
}

func TestConstraintPointers(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="FrameRateReg"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <IntReg Name="FrameRateMinReg"><Address>0x1004</Address><Length>4</Length><AccessMode>RO</AccessMode></IntReg>
  <IntReg Name="FrameRateMaxReg"><Address>0x1008</Address><Length>4</Length><AccessMode>RO</AccessMode></IntReg>
  <Integer Name="FrameRateMin"><pValue>FrameRateMinReg</pValue></Integer>
  <Integer Name="FrameRateMax"><pValue>FrameRateMaxReg</pValue></Integer>
  <Integer Name="FrameRate">
    <pValue>FrameRateReg</pValue>
    <pMin>FrameRateMin</pMin>
    <pMax>FrameRateMax</pMax>
  </Integer>
  <Integer Name="StaticConstraints">
    <Address>0x2000</Address>
    <Length>4</Length>
    <AccessMode>RW</AccessMode>
    <Min>10</Min>
    <Max>100</Max>
    <Inc>5</Inc>
  </Integer>
</RegisterDescription>`

	port := &memPort{regs: map[uint32]uint32{
		0x1004: 1,   // FrameRateMin
		0x1008: 120, // FrameRateMax
	}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}

	// Test pMin pointer
	minVal, hasMin, err := nm.GetMin("FrameRate")
	if err != nil {
		t.Fatalf("GetMin error: %v", err)
	}
	if !hasMin {
		t.Fatal("FrameRate should have min constraint")
	}
	if minVal != 1 {
		t.Fatalf("FrameRate min=%d, want 1", minVal)
	}

	// Test pMax pointer
	maxVal, hasMax, err := nm.GetMax("FrameRate")
	if err != nil {
		t.Fatalf("GetMax error: %v", err)
	}
	if !hasMax {
		t.Fatal("FrameRate should have max constraint")
	}
	if maxVal != 120 {
		t.Fatalf("FrameRate max=%d, want 120", maxVal)
	}

	// Test pInc (should not have)
	incVal, hasInc, err := nm.GetInc("FrameRate")
	if err != nil {
		t.Fatalf("GetInc error: %v", err)
	}
	if hasInc {
		t.Fatalf("FrameRate should not have inc constraint, got %d", incVal)
	}

	// Test static constraints
	staticMin, hasStaticMin, err := nm.GetMin("StaticConstraints")
	if err != nil {
		t.Fatalf("GetMin(StaticConstraints) error: %v", err)
	}
	if !hasStaticMin {
		t.Fatal("StaticConstraints should have min")
	}
	if staticMin != 10 {
		t.Fatalf("StaticConstraints min=%d, want 10", staticMin)
	}

	staticMax, hasStaticMax, err := nm.GetMax("StaticConstraints")
	if err != nil {
		t.Fatalf("GetMax(StaticConstraints) error: %v", err)
	}
	if !hasStaticMax {
		t.Fatal("StaticConstraints should have max")
	}
	if staticMax != 100 {
		t.Fatalf("StaticConstraints max=%d, want 100", staticMax)
	}

	staticInc, hasStaticInc, err := nm.GetInc("StaticConstraints")
	if err != nil {
		t.Fatalf("GetInc(StaticConstraints) error: %v", err)
	}
	if !hasStaticInc {
		t.Fatal("StaticConstraints should have inc")
	}
	if staticInc != 5 {
		t.Fatalf("StaticConstraints inc=%d, want 5", staticInc)
	}

	// Test GetConstraints combined
	min, max, inc, hasMin, hasMax, hasInc, err := nm.GetConstraints("FrameRate")
	if err != nil {
		t.Fatalf("GetConstraints error: %v", err)
	}
	if !hasMin || !hasMax || hasInc {
		t.Fatalf("FrameRate constraints: hasMin=%v hasMax=%v hasInc=%v", hasMin, hasMax, hasInc)
	}
	if min != 1 || max != 120 {
		t.Fatalf("FrameRate: min=%d max=%d", min, max)
	}
	_ = inc // unused; FrameRate doesn't have Inc constraint
}

func TestAvailabilityPointers(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="ImplReg"><Address>0x1000</Address><Length>4</Length><AccessMode>RO</AccessMode></IntReg>
  <IntReg Name="AvailReg"><Address>0x1004</Address><Length>4</Length><AccessMode>RO</AccessMode></IntReg>
  <IntReg Name="LockReg"><Address>0x1008</Address><Length>4</Length><AccessMode>RO</AccessMode></IntReg>
  <IntReg Name="FeatureReg"><Address>0x100C</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Integer Name="Impl"><pValue>ImplReg</pValue></Integer>
  <Integer Name="Avail"><pValue>AvailReg</pValue></Integer>
  <Integer Name="Lock"><pValue>LockReg</pValue></Integer>
  <Integer Name="Feature">
    <pValue>FeatureReg</pValue>
    <pIsImplemented>Impl</pIsImplemented>
    <pIsAvailable>Avail</pIsAvailable>
    <pIsLocked>Lock</pIsLocked>
  </Integer>
</RegisterDescription>`

	port := &memPort{regs: map[uint32]uint32{
		0x1000: 1, // ImplReg: implemented
		0x1004: 1, // AvailReg: available
		0x1008: 0, // LockReg: not locked
	}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}

	impl, err := nm.IsImplemented("Feature")
	if err != nil {
		t.Fatalf("IsImplemented error: %v", err)
	}
	if !impl {
		t.Fatal("Feature should be implemented")
	}

	avail, err := nm.IsAvailable("Feature")
	if err != nil {
		t.Fatalf("IsAvailable error: %v", err)
	}
	if !avail {
		t.Fatal("Feature should be available")
	}

	locked, err := nm.IsLocked("Feature")
	if err != nil {
		t.Fatalf("IsLocked error: %v", err)
	}
	if locked {
		t.Fatal("Feature should not be locked")
	}
}

func TestAvailabilityPointersDefaults(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="FeatureReg"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Integer Name="Feature">
    <pValue>FeatureReg</pValue>
  </Integer>
</RegisterDescription>`

	port := &memPort{regs: map[uint32]uint32{}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}

	impl, err := nm.IsImplemented("Feature")
	if err != nil {
		t.Fatalf("IsImplemented error: %v", err)
	}
	if !impl {
		t.Fatal("Feature without pIsImplemented should default to true")
	}

	avail, err := nm.IsAvailable("Feature")
	if err != nil {
		t.Fatalf("IsAvailable error: %v", err)
	}
	if !avail {
		t.Fatal("Feature without pIsAvailable should default to true")
	}

	locked, err := nm.IsLocked("Feature")
	if err != nil {
		t.Fatalf("IsLocked error: %v", err)
	}
	if locked {
		t.Fatal("Feature without pIsLocked should default to false")
	}
}

func TestInvalidatorPointer(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="InvalidatorReg"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <IntReg Name="FeatureReg"><Address>0x1004</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Integer Name="Invalidator"><pValue>InvalidatorReg</pValue></Integer>
  <Integer Name="Feature">
    <pValue>FeatureReg</pValue>
    <pInvalidator>Invalidator</pInvalidator>
  </Integer>
</RegisterDescription>`

	port := &memPort{regs: map[uint32]uint32{}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}

	inv, err := nm.GetInvalidator("Feature")
	if err != nil {
		t.Fatalf("GetInvalidator error: %v", err)
	}
	if inv != "Invalidator" {
		t.Fatalf("GetInvalidator=%q, want Invalidator", inv)
	}

	inv, err = nm.GetInvalidator("Invalidator")
	if err != nil {
		t.Fatalf("GetInvalidator error: %v", err)
	}
	if inv != "" {
		t.Fatalf("GetInvalidator for node without invalidator should be empty, got %q", inv)
	}
}

func TestReadFloatAndStringRoundTrip(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <FloatReg Name="ExposureTimeReg"><Address>0x4000</Address><Length>4</Length><AccessMode>RW</AccessMode></FloatReg>
  <Float Name="ExposureTime"><pValue>ExposureTimeReg</pValue></Float>
  <StringReg Name="DeviceUserIDReg"><Address>0x4100</Address><Length>64</Length><AccessMode>RW</AccessMode></StringReg>
  <String Name="DeviceUserID"><pValue>DeviceUserIDReg</pValue></String>
</RegisterDescription>`

	port := &memPort{}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}

	// Float write → read round trip.
	_ = nm.SetFloat("ExposureTime", 12.5)
	got, err := nm.ReadFloat("ExposureTime")
	if err != nil {
		t.Fatalf("ReadFloat: %v", err)
	}
	if got != 12.5 {
		t.Fatalf("ReadFloat=%v, want 12.5", got)
	}

	// String write → read round trip.
	_ = nm.SetString("DeviceUserID", "hello")
	s, err := nm.ReadString("DeviceUserID")
	if err != nil {
		t.Fatalf("ReadString: %v", err)
	}
	if s != "hello" {
		t.Fatalf("ReadString=%q, want \"hello\"", s)
	}
}
