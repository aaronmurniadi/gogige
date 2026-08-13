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

// SetBoolean sets a Boolean feature.
func (nm *NodeMap) SetBoolean(name string, v bool) error {
	n, err := nm.lookup(name)
	if err != nil {
		return err
	}
	iv := int64(0)
	if v {
		iv = 1
	}
	return nm.writeIntegerish(n, iv)
}

// ReadBoolean returns the current value of a Boolean feature.
func (nm *NodeMap) ReadBoolean(name string) (bool, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return false, err
	}
	if n.Kind != "Boolean" {
		return false, fmt.Errorf("gige: feature %s is %s, not Boolean", name, n.Kind)
	}
	if n.PValue == "" {
		return false, fmt.Errorf("gige: feature %s has no pValue", name)
	}
	v, err := nm.evalIntegerValue(n.PValue, 0)
	if err != nil {
		return false, err
	}
	return v != 0, nil
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
	if n.Kind == "Float" && n.PValue != "" {
		reg, err := nm.lookup(n.PValue)
		if err != nil {
			return err
		}
		return nm.pa.writeFloatReg(reg, v, nm)
	}
	return nm.pa.writeFloatReg(n, v, nm)
}

// SetString sets a String or Enumeration feature by name/value.
func (nm *NodeMap) SetString(name, val string) error {
	n, err := nm.lookup(name)
	if err != nil {
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
		return nm.pa.writeStringReg(n, val, nm)
	}
	if n.PValue != "" {
		reg, err := nm.lookup(n.PValue)
		if err != nil {
			return err
		}
		if reg.Kind == "StringReg" {
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

func (nm *NodeMap) writeIntegerish(n *gcNode, v int64) error {
	target, err := nm.pa.resolveIntegerTarget(n, nm)
	if err != nil {
		return err
	}
	return nm.pa.writeIntReg(target, v, nm)
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
	v, err := nm.evalIntegerValue(name, 0)
	return int64(v), err
}

// ReadFloat returns the current value of a Float feature.
func (nm *NodeMap) ReadFloat(name string) (float64, error) {
	n, err := nm.lookup(name)
	if err != nil {
		return 0, err
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
