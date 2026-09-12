package genapi

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/aaronmurniadi/gogige/gvcp"
)

// NodeMap is a minimal GenICam RegisterDescription feature map.
// It holds a parsed collection of GenICam nodes and provides access to their values
// through a bound gvcp.Port for register I/O.
type NodeMap struct {
	port  gvcp.Port
	nodes map[string]*gcNode
	pa    *portAdapter
}

// ParseNodeMap parses a RegisterDescription XML document and returns a NodeMap
// bound to the given gvcp.Port for register access.
func ParseNodeMap(xmlData []byte, port gvcp.Port) (*NodeMap, error) {
	nm := &NodeMap{
		port:  port,
		nodes: make(map[string]*gcNode),
		pa:    newPortAdapter(port),
	}

	err := parseNodeMapXML(xmlData, func(kind, name string, inner []byte) {
		if name == "" {
			return
		}
		gn := parseNodeXML(kind, name, inner)
		if gn.Name != "" {
			nm.nodes[gn.Name] = gn
		}
	})
	if err != nil {
		return nil, fmt.Errorf("gige: genapi xml: %w", err)
	}
	return nm, nil
}

// SetBoolean sets a Boolean feature. Honors <OnValue>/<OffValue> when the
// camera overrides the §2.8.7 default of 1/0.
func (nm *NodeMap) SetBoolean(name string, v bool) error {
	n, err := nm.lookup(name)
	if err != nil {
		return err
	}
	iv := int64(0)
	if v {
		iv = 1
	}
	if txt := n.OnValue; v {
		if p, perr := strconv.ParseInt(txt, 0, 64); perr == nil {
			iv = p
		}
	} else if txt := n.OffValue; txt != "" {
		if p, perr := strconv.ParseInt(txt, 0, 64); perr == nil {
			iv = p
		}
	}
	return nm.writeIntegerish(n, iv)
}

// ReadBoolean returns the current value of a Boolean feature. Per §2.8.7 a
// Boolean is true iff its register value equals <OnValue> (default 1).
func (nm *NodeMap) ReadBoolean(name string) (bool, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return false, err
	}
	if n.Kind != "Boolean" {
		return false, fmt.Errorf("gige: feature %s is %s, not Boolean", name, n.Kind)
	}
	if err := nm.checkReadAccess(n); err != nil {
		return false, err
	}
	if n.PValue == "" {
		return false, fmt.Errorf("gige: feature %s has no pValue", name)
	}
	v, err := nm.evalIntegerValue(n.PValue, 0)
	if err != nil {
		return false, err
	}
	on := int64(1)
	if txt := n.OnValue; txt != "" {
		if p, perr := strconv.ParseInt(txt, 0, 64); perr == nil {
			on = p
		}
	}
	return v == uint64(on), nil
}

// SetInteger sets an Integer feature.
func (nm *NodeMap) SetInteger(name string, v int64) error {
	n, err := nm.lookup(name)
	if err != nil {
		return err
	}
	return nm.writeIntegerish(n, v)
}

// SetFloat sets a Float feature.
func (nm *NodeMap) SetFloat(name string, v float64) error {
	n, err := nm.lookup(name)
	if err != nil {
		return err
	}
	if err := nm.checkWriteAccess(n); err != nil {
		return err
	}
	if n.Kind == "Float" && n.PValue != "" {
		reg, err := nm.lookup(n.PValue)
		if err != nil {
			return err
		}
		if err := nm.checkWriteAccess(reg); err != nil {
			return err
		}
		return nm.pa.writeFloatReg(reg, v, nm)
	}
	if n.Kind == "Converter" || n.Kind == "IntConverter" {
		return nm.writeIntegerish(n, int64(v))
	}
	return nm.pa.writeFloatReg(n, v, nm)
}

// SetString sets a String or Enumeration feature by name/value.
func (nm *NodeMap) SetString(name, val string) error {
	n, err := nm.lookup(name)
	if err != nil {
		return err
	}
	if err := nm.checkWriteAccess(n); err != nil {
		return err
	}
	if n.Kind == "Enumeration" {
		ev, ok := n.Entries[val]
		if !ok {
			return fmt.Errorf("gige: enum %s has no entry %q (have %v)", name, val, keysOf(n.Entries))
		}
		return nm.writeIntegerish(n, ev)
	}
	if n.Kind == "String" || n.Kind == "StringReg" {
		if n.PValue != "" {
			reg, err := nm.lookup(n.PValue)
			if err != nil {
				return err
			}
			if err := nm.checkWriteAccess(reg); err != nil {
				return err
			}
		}
		return nm.pa.writeStringReg(n, val, nm)
	}
	if n.PValue != "" {
		reg, err := nm.lookup(n.PValue)
		if err != nil {
			return err
		}
		if reg.Kind == "StringReg" {
			if err := nm.checkWriteAccess(reg); err != nil {
				return err
			}
			return nm.pa.writeStringReg(reg, val, nm)
		}
	}
	return fmt.Errorf("gige: feature %s is not string/enum", name)
}

// Execute runs a Command feature.
func (nm *NodeMap) Execute(name string) error {
	n, err := nm.lookup(name)
	if err != nil {
		return err
	}
	if n.Kind != "Command" {
		return fmt.Errorf("gige: %s is not a Command", name)
	}
	v := int64(1)
	if n.Value != "" {
		if parsed, err := strconv.ParseInt(n.Value, 0, 64); err == nil {
			v = parsed
		}
	}
	return nm.writeIntegerish(n, v)
}

func (nm *NodeMap) lookup(name string) (*gcNode, error) {
	if nm == nil {
		return nil, errors.New("gige: nil nodemap")
	}
	n := nm.nodes[name]
	if n == nil {
		return nil, fmt.Errorf("gige: unknown feature %q", name)
	}
	return n, nil
}

// accessRank ranks GenApi access modes: higher is more restrictive.
func accessRank(mode string) int {
	switch mode {
	case "RW":
		return 0
	case "RO", "WO":
		return 1
	default:
		return 2 // NI, NA, anything unrecognized
	}
}

// normalizeAccess canonicalizes an AccessMode string. Empty is the GenApi
// 2.1.1 §2.8.1 default RW; the §2.8.13 NI/NA meta-modes are surfaced as-is.
func normalizeAccess(mode string) string {
	switch mode {
	case "WO", "RO", "NI", "NA":
		return mode
	default:
		return "RW"
	}
}

// effectiveAccess computes the runtime access mode of a feature per GenApi
// 2.1.1 §2.8.1/§2.8.12: the declared AccessMode narrowed by any
// ImposedAccessMode, then downgraded by not-implemented/not-available/locked
// conditions. Returns one of "NI", "NA", "WO", "RO", "RW".
func (nm *NodeMap) effectiveAccess(n *gcNode) (string, error) {
	impl, err := nm.IsImplemented(n.Name)
	if err != nil {
		return "", fmt.Errorf("gige: access %s: %w", n.Name, err)
	}
	if !impl {
		return "NI", nil
	}
	avail, err := nm.IsAvailable(n.Name)
	if err != nil {
		return "", fmt.Errorf("gige: access %s: %w", n.Name, err)
	}
	if !avail {
		return "NA", nil
	}
	mode := normalizeAccess(n.Access)
	if im := normalizeAccess(n.ImposedAccess); accessRank(im) > accessRank(mode) {
		mode = im
	}
	locked, err := nm.IsLocked(n.Name)
	if err != nil {
		return "", fmt.Errorf("gige: access %s: %w", n.Name, err)
	}
	if locked {
		switch mode {
		case "RW":
			return "RO", nil
		case "WO":
			// Locked write-only: neither readable (WO) nor writable (locked).
			return "NA", nil
		}
	}
	return mode, nil
}

// checkWriteAccess rejects writes to features whose effective access mode is
// not RW/WO (GenApi 2.1.1 §2.8.1).
func (nm *NodeMap) checkWriteAccess(n *gcNode) error {
	mode, err := nm.effectiveAccess(n)
	if err != nil {
		return err
	}
	switch mode {
	case "RW", "WO":
		return nil
	case "RO":
		return fmt.Errorf("gige: feature %q is read-only; requires RW or WO to write", n.Name)
	case "NA":
		return fmt.Errorf("gige: feature %q is not available", n.Name)
	case "NI":
		return fmt.Errorf("gige: feature %q is not implemented", n.Name)
	}
	return fmt.Errorf("gige: feature %q has invalid access mode %q", n.Name, mode)
}

// checkReadAccess rejects reads of features whose effective access mode is not
// RW/RO (GenApi 2.1.1 §2.8.1).
func (nm *NodeMap) checkReadAccess(n *gcNode) error {
	mode, err := nm.effectiveAccess(n)
	if err != nil {
		return err
	}
	switch mode {
	case "RW", "RO":
		return nil
	case "WO":
		return fmt.Errorf("gige: feature %q is write-only", n.Name)
	case "NA":
		return fmt.Errorf("gige: feature %q is not available", n.Name)
	case "NI":
		return fmt.Errorf("gige: feature %q is not implemented", n.Name)
	}
	return fmt.Errorf("gige: feature %q has invalid access mode %q", n.Name, mode)
}

func (nm *NodeMap) writeIntegerish(n *gcNode, v int64) error {
	if err := nm.checkWriteAccess(n); err != nil {
		return err
	}
	if n.Kind == "Converter" || n.Kind == "IntConverter" {
		regVal, target, err := nm.convertUserToRegister(n, v)
		if err != nil {
			return err
		}
		if err := nm.checkWriteAccess(target); err != nil {
			return err
		}
		return nm.pa.writeIntReg(target, regVal, nm)
	}
	target, err := nm.pa.resolveIntegerTarget(n, nm)
	if err != nil {
		return err
	}
	if err := nm.checkWriteAccess(target); err != nil {
		return err
	}
	return nm.pa.writeIntReg(target, v, nm)
}

// convertUserToRegister converts a user-domain value into the register domain
// for a Converter/IntConverter node using FormulaTo (GenApi 2.1.1 §2.8.10),
// then resolves the register node to write to. A formula variable bound to the
// <pValue> register carries the incoming user value; a single-variable formula
// always binds the user value.
func (nm *NodeMap) convertUserToRegister(n *gcNode, v int64) (int64, *gcNode, error) {
	formula := converterFormula(n, false)
	if formula == "" {
		return 0, nil, fmt.Errorf("gige: converter %s has no FormulaTo", n.Name)
	}
	// FormulaTo exposes the incoming user value as the reserved variable FROM
	// (GenApi 2.1.1 §2.8.13); the pVariables then carry their current values,
	// giving the read-modify-write semantics bitfield converters rely on,
	// e.g. "(VAR_CFG & 0xFFF7FFFF) | (FROM << 19)". The legacy single-variable
	// convention (FormulaTo written purely in terms of the one pVariable)
	// instead binds that variable to the user value.
	modern := formulaUses(formula, "FROM")
	extra := map[string]int64{}
	if modern {
		extra["FROM"] = v
	} else {
		// Legacy: bind the user value onto the variable referencing the
		// <pValue> register (or the only variable when there is exactly one).
		for varName, feat := range n.Variables {
			if feat == n.PValue || len(n.Variables) == 1 {
				extra[varName] = v
			}
		}
	}
	rawVal, err := nm.pa.evaluateSwissKnife(formula, n.Variables, nm, extra)
	if err != nil {
		return 0, nil, fmt.Errorf("gige: converter %s FormulaTo: %w", n.Name, err)
	}
	regVal := int64(rawVal)
	target, err := nm.pa.resolveIntegerTarget(n, nm)
	if err != nil {
		return 0, nil, err
	}
	return regVal, target, nil
}

// evalIntegerValue returns a GenICam Integer-like node's numeric value (for pAddress).
func (nm *NodeMap) evalIntegerValue(name string, depth int) (uint64, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return 0, err
	}
	return nm.pa.resolveIntegerReference(n, nm, depth)
}

// evalBoolish evaluates a feature as a boolean (non-zero = true).
// Supports Integer, IntReg, MaskedIntReg, Boolean, SwissKnife, and Enumeration features.
func (nm *NodeMap) evalBoolish(name string) (bool, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return false, err
	}
	switch n.Kind {
	case "Boolean":
		v, err := nm.ReadBoolean(name)
		if err != nil {
			return false, err
		}
		return v, nil
	case "Enumeration":
		// Enumeration is considered "true" if it has any entries
		// (i.e., it's a valid feature). For boolean evaluation,
		// we check if the current value is non-zero.
		v, err := nm.evalIntegerValue(name, 0)
		if err != nil {
			return false, err
		}
		return v != 0, nil
	default:
		v, err := nm.evalIntegerValue(name, 0)
		if err != nil {
			return false, err
		}
		return v != 0, nil
	}
}

func (nm *NodeMap) writeIntReg(n *gcNode, v int64) error {
	return nm.pa.writeIntReg(n, v, nm)
}

func (nm *NodeMap) writeFloatReg(n *gcNode, v float64) error {
	return nm.pa.writeFloatReg(n, v, nm)
}

func (nm *NodeMap) writeStringReg(n *gcNode, val string) error {
	return nm.pa.writeStringReg(n, val, nm)
}

// Has reports whether a feature name exists.
func (nm *NodeMap) Has(name string) bool {
	_, err := nm.lookup(name)
	return err == nil
}

// Kind returns the GenApi node kind of a feature ("Enumeration", "Integer", …),
// or "" when the feature does not exist.
func (nm *NodeMap) Kind(name string) string {
	n, err := nm.lookup(name)
	if err != nil {
		return ""
	}
	return n.Kind
}

// CategoryNode is a node in the GenApi 2.1.1 §2.8.2 feature category tree.
// Categories nest other categories and reference leaf features via <pFeature>.
type CategoryNode struct {
	// Name is the Category node name.
	Name string
	// Features holds the non-Category <pFeature> children in document order.
	Features []string
	// Categories holds the nested child categories in document order.
	Categories []*CategoryNode
}

// Category returns the ordered <pFeature> names of a Category node
// (GenApi 2.1.1 §2.8.2). References to other categories are included verbatim.
func (nm *NodeMap) Category(name string) ([]string, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return nil, err
	}
	if n.Kind != "Category" {
		return nil, fmt.Errorf("gige: feature %s is %s, not Category", name, n.Kind)
	}
	return n.Features, nil
}

// Categories returns the ordered names of all Category nodes in the map.
func (nm *NodeMap) Categories() []string {
	var out []string
	for name, n := range nm.nodes {
		if n.Kind == "Category" {
			out = append(out, name)
		}
	}
	sort.Strings(out)
	return out
}

// RootCategories returns the ordered top-level category names: the <pFeature>
// children of the special "Root" category, or every Category in the map (sorted)
// when this camera description has no Root node.
func (nm *NodeMap) RootCategories() []string {
	if root, err := nm.lookup("Root"); err == nil && root.Kind == "Category" {
		return append([]string(nil), root.Features...)
	}
	return nm.Categories()
}

// CategoryTree builds the nested category tree (§2.8.2) rooted at the category
// named name (use "" for the standard "Root"). Leaf <pFeature> references are
// listed in document order; dangling references to unknown features are
// skipped. Cycles in the category graph return an error.
func (nm *NodeMap) CategoryTree(name string) (*CategoryNode, error) {
	if name == "" {
		name = "Root"
	}
	return nm.buildCategoryTree(name, map[string]bool{})
}

func (nm *NodeMap) buildCategoryTree(name string, active map[string]bool) (*CategoryNode, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return nil, err
	}
	if n.Kind != "Category" {
		return nil, fmt.Errorf("gige: feature %s is %s, not Category", name, n.Kind)
	}
	if active[name] {
		return nil, fmt.Errorf("gige: category cycle at %q", name)
	}
	active[name] = true
	defer delete(active, name)

	cn := &CategoryNode{Name: name}
	for _, f := range n.Features {
		child, err := nm.lookup(f)
		if err != nil {
			continue
		}
		if child.Kind == "Category" {
			sub, err := nm.buildCategoryTree(f, active)
			if err != nil {
				return nil, err
			}
			cn.Categories = append(cn.Categories, sub)
		} else {
			cn.Features = append(cn.Features, f)
		}
	}
	return cn, nil
}

// EnumEntries returns the sorted EnumEntry names of an Enumeration feature.
// It is useful for probing available values (e.g. PixelFormat, PayloadType).
func (nm *NodeMap) EnumEntries(name string) ([]string, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return nil, err
	}
	if n.Kind != "Enumeration" {
		return nil, fmt.Errorf("gige: feature %s is %s, not Enumeration", name, n.Kind)
	}
	keys := make([]string, 0, len(n.Entries))
	for k := range n.Entries {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys, nil
}

// CurrentEnum reads an Enumeration feature and returns the matching entry name,
// or "" when the current register value matches no declared entry.
func (nm *NodeMap) CurrentEnum(name string) (string, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return "", err
	}
	if n.Kind != "Enumeration" {
		return "", fmt.Errorf("gige: feature %s is %s, not Enumeration", name, n.Kind)
	}
	if err := nm.checkReadAccess(n); err != nil {
		return "", err
	}
	v, err := nm.evalIntegerValue(name, 0)
	if err != nil {
		return "", err
	}
	for k, ev := range n.Entries {
		if uint64(ev) == v {
			return k, nil
		}
	}
	return "", nil
}

// ReadInteger returns the current value of an Integer-like feature
// (Integer, IntReg, MaskedIntReg, SwissKnife/Converter or Enumeration).
func (nm *NodeMap) ReadInteger(name string) (int64, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return 0, err
	}
	if err := nm.checkReadAccess(n); err != nil {
		return 0, err
	}
	v, err := nm.evalIntegerValue(name, 0)
	return int64(v), err
}

// ReadFloat returns the current value of a Float feature.
func (nm *NodeMap) ReadFloat(name string) (float64, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return 0, err
	}
	if err := nm.checkReadAccess(n); err != nil {
		return 0, err
	}
	if n.Kind == "Converter" || n.Kind == "IntConverter" {
		v, err := nm.ReadInteger(name)
		if err != nil {
			return 0, err
		}
		return float64(v), nil
	}
	if n.Kind == "SwissKnife" {
		if n.Formula == "" {
			return 0, fmt.Errorf("gige: %s has empty Formula", name)
		}
		return nm.pa.evaluateSwissKnifeFloat(n.Formula, n.Variables, nm, nil)
	}
	target := n
	if n.PValue != "" {
		target, err = nm.lookup(n.PValue)
		if err != nil {
			return 0, err
		}
	}
	return nm.pa.readFloatReg(target, nm)
}

// ReadString returns the current value of a String or string-backed feature.
func (nm *NodeMap) ReadString(name string) (string, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return "", err
	}
	if err := nm.checkReadAccess(n); err != nil {
		return "", err
	}
	target := n
	if n.PValue != "" {
		reg, err := nm.lookup(n.PValue)
		if err != nil {
			return "", err
		}
		if reg.Kind == "StringReg" {
			target = reg
		}
	}
	return nm.pa.readStringReg(target, nm)
}

// GetMin returns the minimum constraint value for an Integer or Float feature.
// Returns (min, hasConstraint, error).
// If pMin is set, evaluates the referenced feature; otherwise returns static Min value.
func (nm *NodeMap) GetMin(name string) (int64, bool, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return 0, false, err
	}
	// Try pointer first
	if n.PMin != "" {
		v, err := nm.evalIntegerValue(n.PMin, 0)
		if err == nil {
			return int64(v), true, nil
		}
		// If pointer fails, don't fall back to static value; report error
		return 0, false, fmt.Errorf("gige: %s pMin %s: %w", name, n.PMin, err)
	}
	// Use static Min if set (check if it was explicitly parsed)
	if n.Min > math.MinInt64 {
		return n.Min, true, nil
	}
	return 0, false, nil
}

// GetMax returns the maximum constraint value for an Integer or Float feature.
// Returns (max, hasConstraint, error).
// If pMax is set, evaluates the referenced feature; otherwise returns static Max value.
func (nm *NodeMap) GetMax(name string) (int64, bool, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return 0, false, err
	}
	// Try pointer first
	if n.PMax != "" {
		v, err := nm.evalIntegerValue(n.PMax, 0)
		if err == nil {
			return int64(v), true, nil
		}
		// If pointer fails, don't fall back to static value; report error
		return 0, false, fmt.Errorf("gige: %s pMax %s: %w", name, n.PMax, err)
	}
	// Use static Max if set (check if it was explicitly parsed)
	if n.Max < math.MaxInt64 {
		return n.Max, true, nil
	}
	return 0, false, nil
}

// GetInc returns the increment constraint value for an Integer feature.
// Returns (inc, hasConstraint, error).
// If pInc is set, evaluates the referenced feature; otherwise returns static Inc value.
func (nm *NodeMap) GetInc(name string) (int64, bool, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return 0, false, err
	}
	// Try pointer first
	if n.PInc != "" {
		v, err := nm.evalIntegerValue(n.PInc, 0)
		if err == nil {
			return int64(v), true, nil
		}
		// If pointer fails, don't fall back to static value; report error
		return 0, false, fmt.Errorf("gige: %s pInc %s: %w", name, n.PInc, err)
	}
	// Use static Inc if set
	if n.Inc > 0 {
		return n.Inc, true, nil
	}
	return 0, false, nil
}

// GetConstraints returns all constraints for a feature.
// Returns (min, max, inc, hasMin, hasMax, hasInc, error).
func (nm *NodeMap) GetConstraints(name string) (int64, int64, int64, bool, bool, bool, error) {
	minVal, hasMin, errMin := nm.GetMin(name)
	maxVal, hasMax, errMax := nm.GetMax(name)
	incVal, hasInc, errInc := nm.GetInc(name)

	// Return first error encountered
	if errMin != nil {
		return 0, 0, 0, false, false, false, errMin
	}
	if errMax != nil {
		return 0, 0, 0, false, false, false, errMax
	}
	if errInc != nil {
		return 0, 0, 0, false, false, false, errInc
	}

	return minVal, maxVal, incVal, hasMin, hasMax, hasInc, nil
}

// IsImplemented reports whether a feature is implemented.
// If pIsImplemented is set, evaluates the referenced feature as a boolean (non-zero = true).
// Otherwise returns true.
func (nm *NodeMap) IsImplemented(name string) (bool, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return false, err
	}
	if n.PIsImplemented == "" {
		return true, nil
	}
	return nm.evalBoolish(n.PIsImplemented)
}

// IsAvailable reports whether a feature is currently available.
// If pIsAvailable is set, evaluates the referenced feature as a boolean (non-zero = true).
// Otherwise returns true.
func (nm *NodeMap) IsAvailable(name string) (bool, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return false, err
	}
	if n.PIsAvailable == "" {
		return true, nil
	}
	return nm.evalBoolish(n.PIsAvailable)
}

// IsLocked reports whether a feature is currently locked.
// If pIsLocked is set, evaluates the referenced feature as a boolean (non-zero = true).
// Otherwise returns false.
func (nm *NodeMap) IsLocked(name string) (bool, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return false, err
	}
	if n.PIsLocked == "" {
		return false, nil
	}
	return nm.evalBoolish(n.PIsLocked)
}

// GetInvalidator returns the pInvalidator feature name for a node,
// or "" if not set. When the invalidator feature changes, this node is invalidated.
func (nm *NodeMap) GetInvalidator(name string) (string, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return "", err
	}
	return n.PInvalidator, nil
}

func deviceOrder(port gvcp.Port) binary.ByteOrder {
	type orderer interface {
		DeviceByteOrder() binary.ByteOrder
	}
	if o, ok := port.(orderer); ok {
		if order := o.DeviceByteOrder(); order != nil {
			return order
		}
	}
	return binary.BigEndian
}
