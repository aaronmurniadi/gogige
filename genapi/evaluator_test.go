package genapi

import "testing"

func TestEvalFormula(t *testing.T) {
	v, err := evalFormula("BASE + 0x20", map[string]int64{"BASE": 0x10000})
	if err != nil {
		t.Fatal(err)
	}
	if v != 0x10020 {
		t.Fatalf("got %x", v)
	}
	v, err = evalFormula("WIDTH * HEIGHT * ((PIXELFORMAT>>16)&0xFF) / 8", map[string]int64{
		"WIDTH": 16, "HEIGHT": 8, "PIXELFORMAT": 0x01080001,
	})
	if err != nil {
		t.Fatal(err)
	}
	if v != 16*8*1 {
		t.Fatalf("payload=%d", v)
	}
	const trig = "(SEL = 0) ? 0x4E05C180 : ((SEL = 1) ? 0x4E05C184 : (0xFFFFFFFF))"
	v, err = evalFormula(trig, map[string]int64{"SEL": 0})
	if err != nil {
		t.Fatal(err)
	}
	if v != 0x4E05C180 {
		t.Fatalf("sel0=%x", v)
	}
	v, err = evalFormula(trig, map[string]int64{"SEL": 1})
	if err != nil {
		t.Fatal(err)
	}
	if v != 0x4E05C184 {
		t.Fatalf("sel1=%x", v)
	}
}

func TestIntSwissKnifeAddress(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <Integer Name="RegBase"><Value>0x10000</Value></Integer>
  <IntSwissKnife Name="TriggerSourceAddrCalc">
    <pVariable Name="BASE">RegBase</pVariable>
    <Formula>BASE + 0x20</Formula>
  </IntSwissKnife>
  <IntReg Name="TriggerSourceReg">
    <Length>4</Length>
    <AccessMode>RW</AccessMode>
    <pAddress>TriggerSourceAddrCalc</pAddress>
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

func TestEvalFormulaFunctions(t *testing.T) {
	v, err := evalFormula("ABS(X)", map[string]int64{"X": -5})
	if err != nil {
		t.Fatal(err)
	}
	if v != 5 {
		t.Fatalf("ABS(-5)=%d", v)
	}
	v, err = evalFormula("ABS(X)", map[string]int64{"X": 7})
	if err != nil {
		t.Fatal(err)
	}
	if v != 7 {
		t.Fatalf("ABS(7)=%d", v)
	}
	v, err = evalFormula("SQRT(16)", nil)
	if err != nil {
		t.Fatal(err)
	}
	if v != 4 {
		t.Fatalf("SQRT(16)=%d", v)
	}
	v, err = evalFormula("SQRT(0)", nil)
	if err != nil {
		t.Fatal(err)
	}
	if v != 0 {
		t.Fatalf("SQRT(0)=%d", v)
	}
	v, err = evalFormula("FLOOR(X)", map[string]int64{"X": 7})
	if err != nil {
		t.Fatal(err)
	}
	if v != 7 {
		t.Fatalf("FLOOR(7)=%d", v)
	}
	v, err = evalFormula("CEIL(X)", map[string]int64{"X": 7})
	if err != nil {
		t.Fatal(err)
	}
	if v != 7 {
		t.Fatalf("CEIL(7)=%d", v)
	}
	v, err = evalFormula("SQRT(ABS(X))", map[string]int64{"X": -9})
	if err != nil {
		t.Fatal(err)
	}
	if v != 3 {
		t.Fatalf("SQRT(ABS(-9))=%d", v)
	}
}
