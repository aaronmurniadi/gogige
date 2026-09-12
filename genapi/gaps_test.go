package genapi

import (
	"strings"
	"testing"
)

// Regression tests for FINDINGS §3 items 1-4.

func TestZeroAddressPermitted(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="GevVersionReg"><Address>0x0</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Integer Name="GevVersion"><pValue>GevVersionReg</pValue></Integer>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{0: 1}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	v, err := nm.ReadInteger("GevVersion")
	if err != nil {
		t.Fatalf("ABRM address 0 read: %v", err)
	}
	if v != 1 {
		t.Fatalf("got %d, want 1", v)
	}
}

func TestAddressOverflowRejected(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <Integer Name="RegBase"><Value>0x100000000</Value></Integer>
  <IntReg Name="OverflowReg"><pAddress>RegBase</pAddress><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Integer Name="Overflow"><pValue>OverflowReg</pValue></Integer>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	_, err = nm.ReadInteger("Overflow")
	if err == nil || !strings.Contains(err.Error(), "exceeds 32-bit") {
		t.Fatalf("want 32-bit overflow error, got %v", err)
	}
}

func TestConstantFloatRead(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <Float Name="Saturation"><Value>2.5</Value></Float>
</RegisterDescription>`
	port := &memPort{}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	v, err := nm.ReadFloat("Saturation")
	if err != nil {
		t.Fatalf("constant Float read: %v", err)
	}
	if v != 2.5 {
		t.Fatalf("got %v, want 2.5", v)
	}
}

func TestConstantStringRead(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <String Name="DeviceVendorName"><Value>Huaray</Value></String>
</RegisterDescription>`
	port := &memPort{}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	v, err := nm.ReadString("DeviceVendorName")
	if err != nil {
		t.Fatalf("constant String read: %v", err)
	}
	if v != "Huaray" {
		t.Fatalf("got %q, want %q", v, "Huaray")
	}
}

func TestReadOnlyWriteRejected(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="GainReg"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Integer Name="GainRO"><AccessMode>RO</AccessMode><pValue>GainReg</pValue></Integer>
  <Integer Name="GainRW"><AccessMode>RW</AccessMode><pValue>GainReg</pValue></Integer>
  <IntReg Name="GainRegRO"><Address>0x1004</Address><Length>4</Length><AccessMode>RO</AccessMode></IntReg>
  <Integer Name="GainViaRegRO"><AccessMode>RW</AccessMode><pValue>GainRegRO</pValue></Integer>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	if err := nm.SetInteger("GainRO", 3); err == nil || !strings.Contains(err.Error(), "read-only") {
		t.Fatalf("feature RO write: want read-only error, got %v", err)
	}
	if err := nm.SetInteger("GainRW", 3); err != nil {
		t.Fatalf("RW write: %v", err)
	}
	// Indirect RO: feature is RW but its target register is RO.
	if err := nm.SetInteger("GainViaRegRO", 3); err == nil || !strings.Contains(err.Error(), "read-only") {
		t.Fatalf("target RO write: want read-only error, got %v", err)
	}
}

func TestWriteOnlyRejected(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="AcquisitionStartReg"><Address>0x4000</Address><Length>4</Length><AccessMode>WO</AccessMode></IntReg>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := nm.ReadInteger("AcquisitionStartReg"); err == nil || !strings.Contains(err.Error(), "write-only") {
		t.Fatalf("WO read: want write-only error, got %v", err)
	}
	if err := nm.SetInteger("AcquisitionStartReg", 1); err != nil {
		t.Fatalf("WO write: %v", err)
	}
	if port.regs[0x4000] != 1 {
		t.Fatalf("regs=%v", port.regs)
	}
}

func TestNotImplemented(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="GainReg"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <IntReg Name="ImplReg"><Address>0x1004</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Integer Name="Gain"><pValue>GainReg</pValue><pIsImplemented>ImplReg</pIsImplemented></Integer>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := nm.ReadInteger("Gain"); err == nil || !strings.Contains(err.Error(), "not implemented") {
		t.Fatalf("NI read: want not-implemented error, got %v", err)
	}
	if err := nm.SetInteger("Gain", 1); err == nil || !strings.Contains(err.Error(), "not implemented") {
		t.Fatalf("NI write: want not-implemented error, got %v", err)
	}
}

func TestNotAvailable(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="GainReg"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <IntReg Name="AvailReg"><Address>0x1004</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Integer Name="Gain"><pValue>GainReg</pValue><pIsAvailable>AvailReg</pIsAvailable></Integer>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := nm.ReadInteger("Gain"); err == nil || !strings.Contains(err.Error(), "not available") {
		t.Fatalf("NA read: want not-available error, got %v", err)
	}
	if err := nm.SetInteger("Gain", 1); err == nil || !strings.Contains(err.Error(), "not available") {
		t.Fatalf("NA write: want not-available error, got %v", err)
	}
}

func TestLockedDowngrade(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="GainReg"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <IntReg Name="LockReg"><Address>0x1004</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Integer Name="Gain"><pValue>GainReg</pValue><pIsLocked>LockReg</pIsLocked></Integer>
  <IntReg Name="CmdReg"><Address>0x2000</Address><Length>4</Length><AccessMode>WO</AccessMode></IntReg>
  <IntReg Name="CmdLockReg"><Address>0x2004</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <IntReg Name="CmdRegLocked"><Address>0x2008</Address><Length>4</Length><AccessMode>WO</AccessMode><pIsLocked>CmdLockReg</pIsLocked></IntReg>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{0x1004: 1, 0x2004: 1}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	// RW + locked => RO: write rejected, read allowed.
	if err := nm.SetInteger("Gain", 3); err == nil || !strings.Contains(err.Error(), "read-only") {
		t.Fatalf("locked RW write: want read-only error, got %v", err)
	}
	if v, err := nm.ReadInteger("Gain"); err != nil || v != 0 {
		t.Fatalf("locked RW read: %v %d", err, v)
	}
	// WO + locked => NA: even register-write is not available.
	if err := nm.SetInteger("CmdRegLocked", 1); err == nil || !strings.Contains(err.Error(), "not available") {
		t.Fatalf("locked WO write: want not-available error, got %v", err)
	}
	if _, err := nm.ReadInteger("CmdRegLocked"); err == nil || !strings.Contains(err.Error(), "not available") {
		t.Fatalf("locked WO read: want not-available error, got %v", err)
	}
}

func TestImposedAccessMode(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="GainReg"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Integer Name="Gain"><AccessMode>RW</AccessMode><ImposedAccessMode>RO</ImposedAccessMode><pValue>GainReg</pValue></Integer>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	if err := nm.SetInteger("Gain", 3); err == nil || !strings.Contains(err.Error(), "read-only") {
		t.Fatalf("ImposedAccessMode RO write: want read-only error, got %v", err)
	}
	if v, err := nm.ReadInteger("Gain"); err != nil || v != 0 {
		t.Fatalf("ImposedAccessMode read: %v %d", err, v)
	}
}

func TestSwissKnifeSuffixes(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <IntReg Name="GainReg"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Integer Name="Gain"><pValue>GainReg</pValue><Min>1</Min><Max>23</Max><Inc>1</Inc></Integer>
  <IntSwissKnife Name="ConstSum">
    <pVariable Name="VAR_GAIN">Gain</pVariable>
    <Formula>VAR_GAIN.Min + VAR_GAIN.Max + VAR_GAIN.Inc + VAR_GAIN.Value</Formula>
  </IntSwissKnife>
  <IntSwissKnife Name="ConstEntry">
    <pVariable Name="VAR_GAIN">Gain</pVariable>
    <Formula>VAR_GAIN.Entry</Formula>
  </IntSwissKnife>
  <IntSwissKnife Name="BadSuffix">
    <pVariable Name="VAR_GAIN">Gain</pVariable>
    <Formula>VAR_GAIN.Foo</Formula>
  </IntSwissKnife>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{0x1000: 5}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	// Register reads 5: Min(1) + Max(23) + Inc(1) + Value(5) = 30.
	v, err := nm.ReadInteger("ConstSum")
	if err != nil {
		t.Fatal(err)
	}
	if v != 30 {
		t.Fatalf("suffix sum got %d, want 30", v)
	}
	if v, err := nm.ReadInteger("ConstEntry"); err != nil || v != 5 {
		t.Fatalf("Entry suffix got %d (%v), want 5", v, err)
	}
	if _, err := nm.ReadInteger("BadSuffix"); err == nil || !strings.Contains(err.Error(), "such suffix") {
		t.Fatalf("unknown suffix: want error, got %v", err)
	}
}

// TestSwissKnifeFloatDomain mirrors the real Huaray camera XML: a Float
// feature read from a FloatReg, referenced by float-domain SwissKnife nodes
// used as <pMax>. Arithmetic must stay in the float domain, and integer
// reads of the SwissKnife must truncate rather than error.
func TestSwissKnifeFloatDomain(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <FloatReg Name="GainMaxReg"><Address>0x4E05D720</Address><Length>4</Length><AccessMode>RW</AccessMode></FloatReg>
  <IntReg Name="GainReg"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Float Name="GainMax"><AccessMode>RW</AccessMode><pValue>GainMaxReg</pValue><Min>1</Min><Max>23</Max></Float>
  <SwissKnife Name="GainRangeExpr">
    <pVariable Name="VAR_GAINMAX">GainMax</pVariable>
    <Formula>VAR_GAINMAX.Max / 2.0</Formula>
  </SwissKnife>
  <SwissKnife Name="GainIdentityExpr">
    <pVariable Name="VAR_GAINMAX">GainMax</pVariable>
    <Formula>VAR_GAINMAX</Formula>
  </SwissKnife>
  <Integer Name="GainRangeMax"><AccessMode>RW</AccessMode><pValue>GainReg</pValue><pMax>GainRangeExpr</pMax></Integer>
</RegisterDescription>`
	port := &memPort{regs: map[uint32]uint32{}}
	nm, err := ParseNodeMap([]byte(xml), port)
	if err != nil {
		t.Fatal(err)
	}
	// .Max suffix resolves in float domain: 23 / 2.0 = 11.5, no truncation.
	if v, err := nm.ReadFloat("GainRangeExpr"); err != nil || v != 11.5 {
		t.Fatalf("float suffix eval got %v (%v), want 11.5", v, err)
	}
	// Integer read of the float SwissKnife truncates instead of erroring.
	if v, err := nm.ReadInteger("GainRangeExpr"); err != nil || v != 11 {
		t.Fatalf("truncated integer eval got %d (%v), want 11", v, err)
	}
	// A Float feature bound by pMax -> SwissKnife inherits the truncated max.
	if v, has, err := nm.GetMax("GainRangeMax"); err != nil || !has || v != 11 {
		t.Fatalf("GetMax via SwissKnife got %d has=%v (%v), want 11", v, has, err)
	}
	// Identity formula carries the live Float register value in float domain.
	if err := nm.SetFloat("GainMax", 4); err != nil {
		t.Fatal(err)
	}
	if v, err := nm.ReadFloat("GainIdentityExpr"); err != nil || v != 4 {
		t.Fatalf("identity float eval got %v (%v), want 4", v, err)
	}
	if v, err := nm.ReadInteger("GainIdentityExpr"); err != nil || v != 4 {
		t.Fatalf("identity integer eval got %d (%v), want 4", v, err)
	}
}
