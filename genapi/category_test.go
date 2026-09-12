package genapi

import (
	"testing"
)

// parse categories from the GenApi 2.1.1 §2.8.2 example plus camera-style extras.
func TestCategoryTree(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <Category Name="Root" NameSpace="Standard">
    <pFeature>ScalarFeatures</pFeature>
    <pFeature>Trigger</pFeature>
  </Category>
  <Category Name="ScalarFeatures">
    <pFeature>Shutter</pFeature>
    <pFeature>Gain</pFeature>
    <pFeature>Offset</pFeature>
    <pFeature>WhiteBalance</pFeature>
  </Category>
  <Category Name="WhiteBalance">
    <pFeature>RedGain</pFeature>
    <pFeature>BlueGain</pFeature>
  </Category>
  <Category Name="Trigger">
    <pFeature>TriggerMode</pFeature>
    <pFeature>TriggerPolarity</pFeature>
  </Category>
  <IntReg Name="Shutter"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <IntReg Name="Gain"><Address>0x1004</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <IntReg Name="Offset"><Address>0x1008</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <IntReg Name="RedGain"><Address>0x100C</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <IntReg Name="BlueGain"><Address>0x100E</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <IntReg Name="TriggerMode"><Address>0x1010</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <IntReg Name="TriggerPolarity"><Address>0x1014</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
  <Category Name="Dangling">
    <pFeature>TriggerMode</pFeature>
    <pFeature>NoSuchFeature</pFeature>
  </Category>
</RegisterDescription>`
	nm, err := ParseNodeMap([]byte(xml), &memPort{})
	if err != nil {
		t.Fatal(err)
	}

	// Root -> ScalarFeatures, Trigger (document order).
	root, err := nm.CategoryTree("")
	if err != nil {
		t.Fatal(err)
	}
	if root.Name != "Root" {
		t.Fatalf("root name = %q", root.Name)
	}
	if len(root.Categories) != 2 || root.Categories[0].Name != "ScalarFeatures" || root.Categories[1].Name != "Trigger" {
		t.Fatalf("root categories = %+v", root.Categories)
	}
	if len(root.Features) != 0 {
		t.Fatalf("root leaf features = %v", root.Features)
	}

	scal := root.Categories[0]
	if len(scal.Categories) != 1 || scal.Categories[0].Name != "WhiteBalance" {
		t.Fatalf("ScalarFeatures sub = %+v", scal.Categories)
	}
	if len(scal.Features) != 3 || scal.Features[0] != "Shutter" || scal.Features[2] != "Offset" {
		t.Fatalf("ScalarFeatures leaves = %v", scal.Features)
	}
	if wb := scal.Categories[0]; len(wb.Features) != 2 || wb.Features[1] != "BlueGain" {
		t.Fatalf("WhiteBalance leaves = %v", wb.Features)
	}

	// RootCategories reflects the Root children.
	if rc := nm.RootCategories(); len(rc) != 2 || rc[0] != "ScalarFeatures" || rc[1] != "Trigger" {
		t.Fatalf("RootCategories = %v", rc)
	}
	// Categories() returns every Category, sorted.
	all := nm.Categories()
	if len(all) != 5 {
		t.Fatalf("Categories() = %v", all)
	}

	// Category(name) returns raw pFeature list, incl. sub-category names.
	raw, err := nm.Category("ScalarFeatures")
	if err != nil {
		t.Fatal(err)
	}
	if len(raw) != 4 || raw[3] != "WhiteBalance" {
		t.Fatalf("Category(ScalarFeatures) = %v", raw)
	}
	if _, err := nm.Category("Shutter"); err == nil {
		t.Fatal("non-category should error")
	}

	// Dangling references are skipped in the tree but listed raw.
	dang, err := nm.Category("Dangling")
	if err != nil {
		t.Fatal(err)
	}
	if len(dang) != 2 {
		t.Fatalf("dangling raw = %v", dang)
	}
	dtree, err := nm.CategoryTree("Dangling")
	if err != nil {
		t.Fatal(err)
	}
	if len(dtree.Features) != 1 || dtree.Features[0] != "TriggerMode" {
		t.Fatalf("dangling tree leaves = %v", dtree.Features)
	}
}

func TestCategoryCycle(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <Category Name="A"><pFeature>B</pFeature></Category>
  <Category Name="B"><pFeature>A</pFeature></Category>
</RegisterDescription>`
	nm, err := ParseNodeMap([]byte(xml), &memPort{})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := nm.CategoryTree("A"); err == nil || err.Error() != "gige: category cycle at \"A\"" {
		t.Fatalf("cycle: want cycle error, got %v", err)
	}
}

func TestCategoryNoRoot(t *testing.T) {
	const xml = `<?xml version="1.0"?>
<RegisterDescription>
  <Category Name="Z"><pFeature>X</pFeature></Category>
  <Category Name="M"><pFeature>X</pFeature></Category>
  <IntReg Name="X"><Address>0x1000</Address><Length>4</Length><AccessMode>RW</AccessMode></IntReg>
</RegisterDescription>`
	nm, err := ParseNodeMap([]byte(xml), &memPort{})
	if err != nil {
		t.Fatal(err)
	}
	// No Root node: RootCategories falls back to all categories, sorted.
	if rc := nm.RootCategories(); len(rc) != 2 || rc[0] != "M" || rc[1] != "Z" {
		t.Fatalf("RootCategories fallback = %v", rc)
	}
}
