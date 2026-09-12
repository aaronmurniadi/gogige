package genapi

import (
	"fmt"
	"math"
	"testing"
)

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
	v, err = evalFormula("(SEL == 0)", map[string]int64{"SEL": 0})
	if err != nil {
		t.Fatal(err)
	}
	if v != 1 {
		t.Fatalf("== yielded %d", v)
	}
	v, err = evalFormula("2 ** 10", nil)
	if err != nil {
		t.Fatal(err)
	}
	if v != 1024 {
		t.Fatalf("2**10=%d", v)
	}
	v, err = evalFormula("2 ** 3 ** 2", nil)
	if err != nil {
		t.Fatal(err)
	}
	if v != 512 {
		t.Fatalf("2**3**2=%d", v)
	}
	v, err = evalFormula("-2 ** 2", nil)
	if err != nil {
		t.Fatal(err)
	}
	if v != -4 {
		t.Fatalf("-2**2=%d", v)
	}
	v, err = evalFormula("SQRT(144)", nil)
	if err != nil {
		t.Fatal(err)
	}
	if v != 12 {
		t.Fatalf("SQRT(144)=%d", v)
	}
	v, err = evalFormula("LG(1000)", nil)
	if err != nil {
		t.Fatal(err)
	}
	if v != 3 {
		t.Fatalf("LG(1000)=%d", v)
	}
	v, err = evalFormula("ROUND(250, -1)", nil)
	if err != nil {
		t.Fatal(err)
	}
	if v != 250 {
		t.Fatalf("ROUND(250,-1)=%d", v)
	}
	v, err = evalFormula("SGN(SEL)", map[string]int64{"SEL": -8})
	if err != nil {
		t.Fatal(err)
	}
	if v != -1 {
		t.Fatalf("SGN(-8)=%d", v)
	}
	v, err = evalFormula("PI + 1", nil)
	if err != nil {
		t.Fatal(err)
	}
	if v != 4 {
		t.Fatalf("PI+1=%d", v)
	}
}

// TestEvalFormulaNotEqual verifies the "<>" not-equal operator, which real
// GenICam XML uses for != (e.g. Huaray "VAR_LASERSELECTOR <> 0"). Previously
// parseRel swallowed the "<" as a less-than and the ">" was a bad token.
func TestEvalFormulaNotEqual(t *testing.T) {
	cases := []struct {
		expr string
		vars map[string]int64
		want int64
	}{
		{"SELECTOR <> 0", map[string]int64{"SELECTOR": 3}, 1},
		{"SELECTOR <> 3", map[string]int64{"SELECTOR": 3}, 0},
		{"SELECTOR < > 0", map[string]int64{"SELECTOR": 5}, 1},
		{"(SELECTOR = 0) || (SELECTOR <> 0)", map[string]int64{"SELECTOR": 0}, 1},
		{"(SELECTOR <> 2) && (SELECTOR <> 4)", map[string]int64{"SELECTOR": 2}, 0},
		{"(SELECTOR <> 2) && (SELECTOR <> 4)", map[string]int64{"SELECTOR": 3}, 1},
		// Relational operators must be unaffected.
		{"SELECTOR < 4", map[string]int64{"SELECTOR": 3}, 1},
		{"SELECTOR <= 4", map[string]int64{"SELECTOR": 4}, 1},
		{"SELECTOR >= 4", map[string]int64{"SELECTOR": 4}, 1},
		{"SELECTOR > 4", map[string]int64{"SELECTOR": 3}, 0},
	}
	for _, tc := range cases {
		v, err := evalFormula(tc.expr, tc.vars)
		if err != nil {
			t.Fatalf("%q: %v", tc.expr, err)
		}
		if v != tc.want {
			t.Fatalf("%q = %d, want %d", tc.expr, v, tc.want)
		}
	}
}

// TestEvalFormulaTernaryRegression locks the GevTimestampControlResetAvailExpr
// style formula, which nests a converter read ("unknown var TO" regression)
// behind a ternary that keeps working through the <>/&& expression chain.
func TestEvalFormulaTernaryRegression(t *testing.T) {
	const expr = "(VAR_GEVSUPPORTEDOPTIONALCOMMANDSIEEE1588SUPPORT) ? (VAR_GEVIEEE1588 = 0) : (1)"
	vars := map[string]int64{
		"VAR_GEVSUPPORTEDOPTIONALCOMMANDSIEEE1588SUPPORT": 1,
		"VAR_GEVIEEE1588": 0,
	}
	v, err := evalFormula(expr, vars)
	if err != nil {
		t.Fatal(err)
	}
	if v != 1 {
		t.Fatalf("ternary = %d, want 1", v)
	}
}

// TestEvalFormulaFloat verifies the float-domain SwissKnife evaluator
// (GenApi 2.1.1 §2.8.13): arithmetic stays float (no truncation), functions
// and E/PI constants resolve, and comparisons yield 1/0.
func TestEvalFormulaFloat(t *testing.T) {
	cases := []struct {
		expr string
		vars map[string]float64
		want float64
	}{
		{"VAR_X / 2.0", map[string]float64{"VAR_X": 23}, 11.5},
		{"1.0 + 2.5", nil, 3.5},
		{"-2.5 * 4", nil, -10},
		{"2.0 ** 3", nil, 8},
		{"ROUND(3.14159, 2)", nil, 3.14},
		{"FLOOR(2.9)", nil, 2},
		{"CEIL(2.1)", nil, 3},
		{"TRUNC(-2.9)", nil, -2},
		{"ABS(-3.5)", nil, 3.5},
		{"SQRT(144.0)", nil, 12},
		{"LG(1000)", nil, 3},
		{"LN(E)", nil, 1},
		{"MIN(3.0, 7.5)", nil, 3},
		{"MAX(3.0, 7.5)", nil, 7.5},
		{"POW(2.0, 10)", nil, 1024},
		{"ATAN2(0.0, 1.0)", nil, 0},
		{"PI > 3.0", nil, 1},
		{"PI < 3.0", nil, 0},
		{"(VAR_X > 10) ? VAR_X : 0.5", map[string]float64{"VAR_X": 23}, 23},
		{"(VAR_X > 10) ? VAR_X : 0.5", map[string]float64{"VAR_X": 5}, 0.5},
		{"VAR_X % 2.0", map[string]float64{"VAR_X": 7}, 1},
		{"0x10 * 1.5", nil, 24},
		{"1e3", nil, 1000},
	}
	for _, tc := range cases {
		v, err := evalFormulaFloat(tc.expr, tc.vars, nil)
		if err != nil {
			t.Fatalf("%q: %v", tc.expr, err)
		}
		if math.Abs(v-tc.want) > 1e-9 {
			t.Fatalf("%q = %v, want %v", tc.expr, v, tc.want)
		}
	}
	if _, err := evalFormulaFloat("1/0", nil, nil); err == nil {
		t.Fatal("division by zero should error")
	}
}

// TestEvalFormulaFloatSuffix verifies float-domain suffix resolution
// (".Value", ".Min", ".Max") returns the right domain value.
func TestEvalFormulaFloatSuffix(t *testing.T) {
	resolve := func(name, suffix string) (float64, bool, error) {
		switch name + "." + suffix {
		case "GAIN.Max", "GAIN.X":
			return 23, true, nil
		}
		return 0, true, fmt.Errorf("no such suffix %s.%s", name, suffix)
	}
	v, err := evalFormulaFloat("(GAIN.Max - GAIN.X) / 2.0", nil, resolve)
	if err != nil {
		t.Fatal(err)
	}
	if v != 0 {
		t.Fatalf("suffix formula = %v, want 0", v)
	}
}
