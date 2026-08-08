package genapi

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/aaronmurniadi/gogige/gvcp"
)

// NodeMap is a minimal GenICam RegisterDescription feature map.
type NodeMap struct {
	port  gvcp.Port
	nodes map[string]*gcNode
}

type gcNode struct {
	Name       string
	Kind       string
	Address    uint64   // sum of constant <Address> children
	Addresses  []uint64 // each <Address> (for debug)
	Length     int
	Access     string
	PValue     string
	PAddresses []string // each <pAddress> — values are summed with Address (Aravis GenApi)
	Value      string
	Entries    map[string]int64
	Variables  map[string]string // SwissKnife: formula var name → feature name
	Formula    string
	LSB        int
	MSB        int
	HasMask    bool
}

// ParseNodeMap parses a RegisterDescription XML document.
func ParseNodeMap(xmlData []byte, port gvcp.Port) (*NodeMap, error) {
	dec := xml.NewDecoder(bytes.NewReader(xmlData))
	nm := &NodeMap{port: port, nodes: make(map[string]*gcNode)}

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("gige: genapi xml: %w", err)
		}
		se, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		kind := se.Name.Local
		switch kind {
		case "Integer", "Boolean", "Float", "String", "Enumeration", "Command",
			"IntReg", "FloatReg", "StringReg", "MaskedIntReg",
			"IntSwissKnife", "SwissKnife", "IntConverter", "Converter":
			name := attrLocal(se, "Name")
			var holder struct {
				Inner []byte `xml:",innerxml"`
			}
			if err := dec.DecodeElement(&holder, &se); err != nil {
				return nil, err
			}
			fields := parseNodeFields(holder.Inner)
			gn := &gcNode{
				Name:       name,
				Kind:       kind,
				Access:     fields.Access,
				PValue:     fields.PValue,
				PAddresses: fields.PAddresses,
				Value:      fields.Value,
				Address:    fields.AddressSum,
				Addresses:  fields.Addresses,
				Length:     fields.Length,
				LSB:        fields.LSB,
				MSB:        fields.MSB,
				HasMask:    fields.HasMask,
				Variables:  fields.Variables,
				Formula:    fields.Formula,
			}
			// IntConverter uses FormulaTo for forward mapping when reading value
			if gn.Formula == "" && fields.FormulaTo != "" {
				gn.Formula = fields.FormulaTo
			}
			if kind == "Enumeration" {
				gn.Entries = parseEnumEntries(holder.Inner)
			}
			if gn.Name != "" {
				nm.nodes[gn.Name] = gn
			}
		case "RegisterDescription", "Group", "Category", "StructReg",
			"Port", "Node", "XMLDescription":
			continue
		default:
			_ = dec.Skip()
		}
	}
	return nm, nil
}

func attrLocal(se xml.StartElement, local string) string {
	for _, a := range se.Attr {
		if a.Name.Local == local {
			return a.Value
		}
	}
	return ""
}

type nodeFields struct {
	AddressSum uint64
	Addresses  []uint64
	Length     int
	Access     string
	PValue     string
	PAddresses []string
	Value      string
	Variables  map[string]string
	Formula    string
	FormulaTo  string
	LSB, MSB   int
	HasMask    bool
}

func parseNodeFields(inner []byte) nodeFields {
	f := nodeFields{LSB: -1, MSB: -1, Variables: map[string]string{}}
	dec := xml.NewDecoder(bytes.NewReader(inner))
	for {
		tok, err := dec.Token()
		if err != nil {
			break
		}
		se, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		local := se.Name.Local
		if local == "pVariable" {
			varName := attrLocal(se, "Name")
			var v struct {
				Chardata string `xml:",chardata"`
			}
			if err := dec.DecodeElement(&v, &se); err != nil {
				continue
			}
			feat := strings.TrimSpace(v.Chardata)
			if varName != "" && feat != "" {
				f.Variables[varName] = feat
			}
			continue
		}
		var v struct {
			Chardata string `xml:",chardata"`
		}
		if err := dec.DecodeElement(&v, &se); err != nil {
			continue
		}
		text := strings.TrimSpace(v.Chardata)
		switch local {
		case "Address":
			if a, err := parseUint(text); err == nil {
				f.AddressSum += a
				f.Addresses = append(f.Addresses, a)
			}
		case "pAddress":
			if text != "" {
				f.PAddresses = append(f.PAddresses, text)
			}
		case "Length":
			if n, err := strconv.Atoi(text); err == nil {
				f.Length = n
			}
		case "AccessMode":
			f.Access = text
		case "pValue":
			f.PValue = text
		case "Value":
			if f.Value == "" {
				f.Value = text
			}
		case "Formula":
			f.Formula = text
		case "FormulaTo":
			f.FormulaTo = text
		case "Bit":
			if n, err := strconv.Atoi(text); err == nil {
				f.LSB, f.MSB, f.HasMask = n, n, true
			}
		case "LSB":
			if n, err := strconv.Atoi(text); err == nil {
				f.LSB = n
				f.HasMask = f.MSB >= 0
			}
		case "MSB":
			if n, err := strconv.Atoi(text); err == nil {
				f.MSB = n
				f.HasMask = f.LSB >= 0
			}
		}
	}
	if f.LSB >= 0 && f.MSB >= 0 {
		f.HasMask = true
	}
	return f
}

func maskFromBits(lsb, msb int) (mask uint32, shift uint) {
	lo, hi := lsb, msb
	if lo > hi {
		lo, hi = hi, lo
	}
	shift = uint(lo)
	for b := lo; b <= hi; b++ {
		mask |= 1 << uint(b)
	}
	return mask, shift
}

func parseEnumEntries(inner []byte) map[string]int64 {
	out := make(map[string]int64)
	dec := xml.NewDecoder(bytes.NewReader(inner))
	for {
		tok, err := dec.Token()
		if err != nil {
			break
		}
		se, ok := tok.(xml.StartElement)
		if !ok || se.Name.Local != "EnumEntry" {
			continue
		}
		name := attrLocal(se, "Name")
		fields := nodeFields{}
		var holder struct {
			Inner []byte `xml:",innerxml"`
		}
		if err := dec.DecodeElement(&holder, &se); err != nil {
			continue
		}
		fields = parseNodeFields(holder.Inner)
		if name == "" || fields.Value == "" {
			continue
		}
		v, err := strconv.ParseInt(fields.Value, 0, 64)
		if err != nil {
			continue
		}
		out[name] = v
	}
	return out
}

func parseUint(s string) (uint64, error) {
	s = strings.TrimSpace(s)
	return strconv.ParseUint(s, 0, 64)
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
		return nm.writeFloatReg(reg, v)
	}
	return nm.writeFloatReg(n, v)
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
		return nm.writeStringReg(n, val)
	}
	if n.PValue != "" {
		reg, err := nm.lookup(n.PValue)
		if err != nil {
			return err
		}
		if reg.Kind == "StringReg" {
			return nm.writeStringReg(reg, val)
		}
	}
	return fmt.Errorf("gige: feature %s is not string/enum", name)
}

func keysOf(m map[string]int64) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
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
	target := n
	if n.PValue != "" {
		reg, err := nm.lookup(n.PValue)
		if err != nil {
			return err
		}
		target = reg
		// Enumeration/Integer may point at another Integer that points at IntReg.
		if target.PValue != "" && (target.Kind == "Integer" || target.Kind == "Enumeration") {
			reg2, err := nm.lookup(target.PValue)
			if err != nil {
				return err
			}
			target = reg2
		}
	}
	return nm.writeIntReg(target, v)
}

// evalIntegerValue returns a GenICam Integer-like node's numeric value (for pAddress).
func (nm *NodeMap) evalIntegerValue(name string, depth int) (uint64, error) {
	if depth > 12 {
		return 0, fmt.Errorf("gige: pAddress recursion on %q", name)
	}
	n, err := nm.lookup(name)
	if err != nil {
		return 0, err
	}
	switch n.Kind {
	case "Integer":
		if n.Value != "" {
			return parseUint(n.Value)
		}
		if n.PValue != "" {
			return nm.evalIntegerValue(n.PValue, depth+1)
		}
		if n.Address != 0 {
			return n.Address, nil
		}
	case "IntSwissKnife", "SwissKnife", "IntConverter", "Converter":
		vars := make(map[string]int64, len(n.Variables))
		for varName, feat := range n.Variables {
			v, err := nm.evalIntegerValue(feat, depth+1)
			if err != nil {
				return 0, fmt.Errorf("gige: %s var %s→%s: %w", name, varName, feat, err)
			}
			vars[varName] = int64(v)
		}
		if n.Formula == "" {
			return 0, fmt.Errorf("gige: %s has empty Formula", name)
		}
		v, err := evalFormula(n.Formula, vars)
		if err != nil {
			return 0, fmt.Errorf("gige: %s formula %q: %w", name, n.Formula, err)
		}
		return uint64(v), nil
	case "IntReg", "MaskedIntReg":
		addr, length, err := nm.resolveAddr(n)
		if err != nil {
			return 0, err
		}
		if length == 4 {
			v, err := nm.port.ReadReg(addr)
			return uint64(v), err
		}
	case "Enumeration":
		if n.PValue != "" {
			return nm.evalIntegerValue(n.PValue, depth+1)
		}
	}
	return 0, fmt.Errorf("gige: cannot evaluate integer %q (kind=%s value=%q pValue=%q addr=0x%x pAddr=%v)",
		name, n.Kind, n.Value, n.PValue, n.Address, n.PAddresses)
}

func (nm *NodeMap) resolveAddr(n *gcNode) (uint32, int, error) {
	// GenApi: address = sum(Address*) + sum(value(pAddress*))
	addr := n.Address
	for _, pName := range n.PAddresses {
		v, err := nm.evalIntegerValue(pName, 0)
		if err != nil {
			return 0, 0, fmt.Errorf("gige: %s pAddress %s: %w", n.Name, pName, err)
		}
		addr += v
	}
	length := n.Length
	if length <= 0 {
		length = 4
	}
	if addr == 0 {
		return 0, 0, fmt.Errorf("gige: feature %q has no register address (kind=%s addresses=%v pAddress=%v pValue=%s)",
			n.Name, n.Kind, n.Addresses, n.PAddresses, n.PValue)
	}
	return uint32(addr), length, nil
}

func (nm *NodeMap) writeIntReg(n *gcNode, v int64) error {
	addr, length, err := nm.resolveAddr(n)
	if err != nil {
		return err
	}
	if n.Kind == "MaskedIntReg" && n.HasMask && length == 4 {
		cur, err := nm.port.ReadReg(addr)
		if err != nil {
			return fmt.Errorf("gige: read %s @0x%x for mask: %w", n.Name, addr, err)
		}
		mask, shift := maskFromBits(n.LSB, n.MSB)
		nv := (cur &^ mask) | ((uint32(v) << shift) & mask)
		if err := nm.port.WriteReg(addr, nv); err != nil {
			return fmt.Errorf("gige: write MaskedIntReg %s @0x%x (val=%d mask=0x%x): %w", n.Name, addr, v, mask, err)
		}
		return nil
	}
	switch length {
	case 4:
		if err := nm.port.WriteReg(addr, uint32(v)); err != nil {
			return fmt.Errorf("gige: write %s @0x%x = %d: %w", n.Name, addr, v, err)
		}
		return nil
	case 8:
		var b [8]byte
		deviceOrder(nm.port).PutUint64(b[:], uint64(v))
		return nm.port.WriteMem(addr, b[:])
	case 2:
		var b [2]byte
		deviceOrder(nm.port).PutUint16(b[:], uint16(v))
		return nm.port.WriteMem(addr, b[:])
	case 1:
		return nm.port.WriteMem(addr, []byte{byte(v)})
	default:
		if err := nm.port.WriteReg(addr, uint32(v)); err != nil {
			return fmt.Errorf("gige: write %s @0x%x = %d: %w", n.Name, addr, v, err)
		}
		return nil
	}
}

func (nm *NodeMap) writeFloatReg(n *gcNode, v float64) error {
	addr, length, err := nm.resolveAddr(n)
	if err != nil {
		return err
	}
	order := deviceOrder(nm.port)
	if length >= 8 {
		var b [8]byte
		order.PutUint64(b[:], math.Float64bits(v))
		return nm.port.WriteMem(addr, b[:])
	}
	var b [4]byte
	order.PutUint32(b[:], math.Float32bits(float32(v)))
	return nm.port.WriteMem(addr, b[:])
}

func (nm *NodeMap) writeStringReg(n *gcNode, val string) error {
	target := n
	if n.PValue != "" {
		reg, err := nm.lookup(n.PValue)
		if err != nil {
			return err
		}
		target = reg
	}
	addr, length, err := nm.resolveAddr(target)
	if err != nil {
		return err
	}
	if length <= 0 {
		length = len(val) + 1
	}
	buf := make([]byte, length)
	copy(buf, val)
	return nm.port.WriteMem(addr, buf)
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
