package genapi

import (
	"bytes"
	"encoding/xml"
	"io"
	"math"
	"strconv"
	"strings"
)

// nodeFields holds intermediate parsed fields during XML parsing before constructing a gcNode.
type nodeFields struct {
	AddressSum     uint64
	Addresses      []uint64
	Length         int
	Access         string
	PValue         string
	PAddresses     []string
	Value          string
	Variables      map[string]string
	Formula        string
	FormulaTo      string
	LSB, MSB       int
	HasMask        bool
	PMin           string
	PMax           string
	PInc           string
	PIsImplemented string
	PIsAvailable   string
	PIsLocked      string
	PInvalidator   string
	Min            int64
	Max            int64
	Inc            int64
}

// parseNodeFields parses the inner XML of a GenICam node element
// and extracts all relevant fields (Address, pAddress, Length, AccessMode, pValue, Value, Formula, etc.).
func parseNodeFields(inner []byte) nodeFields {
	f := nodeFields{LSB: -1, MSB: -1, Min: math.MinInt64, Max: math.MaxInt64, Variables: map[string]string{}}
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
		case "pMin":
			f.PMin = text
		case "pMax":
			f.PMax = text
		case "pInc":
			f.PInc = text
		case "pIsImplemented":
			f.PIsImplemented = text
		case "pIsAvailable":
			f.PIsAvailable = text
		case "pIsLocked":
			f.PIsLocked = text
		case "pInvalidator":
			f.PInvalidator = text
		case "Min":
			if n, err := strconv.ParseInt(text, 0, 64); err == nil {
				f.Min = n
			}
		case "Max":
			if n, err := strconv.ParseInt(text, 0, 64); err == nil {
				f.Max = n
			}
		case "Inc":
			if n, err := strconv.ParseInt(text, 0, 64); err == nil {
				f.Inc = n
			}
		}
	}
	if f.LSB >= 0 && f.MSB >= 0 {
		f.HasMask = true
	}
	return f
}

// parseEnumEntries extracts EnumEntry name→value mappings from enumeration node XML.
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

// attrLocal returns the value of an XML attribute with the given local name,
// or "" if not found.
func attrLocal(se xml.StartElement, local string) string {
	for _, a := range se.Attr {
		if a.Name.Local == local {
			return a.Value
		}
	}
	return ""
}

// parseUint parses a uint64 from a string, handling hexadecimal (0x prefix) and decimal.
func parseUint(s string) (uint64, error) {
	s = strings.TrimSpace(s)
	return strconv.ParseUint(s, 0, 64)
}

// maskFromBits computes a bit mask and shift amount from LSB and MSB bit positions.
// For example, LSB=1, MSB=3 produces mask=0b1110 and shift=1.
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

// keysOf extracts keys from a map[string]int64 and returns them as a string slice.
// Used for error messages and debugging.
func keysOf(m map[string]int64) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

// parseNodeXML parses a single GenICam node element (identified by its start element and inner XML)
// and returns the corresponding gcNode or error. This is the main entry point for node construction.
func parseNodeXML(kind string, name string, inner []byte) *gcNode {
	fields := parseNodeFields(inner)
	gn := &gcNode{
		Name:           name,
		Kind:           kind,
		Access:         fields.Access,
		PValue:         fields.PValue,
		PAddresses:     fields.PAddresses,
		Value:          fields.Value,
		Address:        fields.AddressSum,
		Addresses:      fields.Addresses,
		Length:         fields.Length,
		LSB:            fields.LSB,
		MSB:            fields.MSB,
		HasMask:        fields.HasMask,
		Variables:      fields.Variables,
		Formula:        fields.Formula,
		PMin:           fields.PMin,
		PMax:           fields.PMax,
		PInc:           fields.PInc,
		PIsImplemented: fields.PIsImplemented,
		PIsAvailable:   fields.PIsAvailable,
		PIsLocked:      fields.PIsLocked,
		PInvalidator:   fields.PInvalidator,
		Min:            fields.Min,
		Max:            fields.Max,
		Inc:            fields.Inc,
	}
	// IntConverter uses FormulaTo for forward mapping when reading value.
	if gn.Formula == "" && fields.FormulaTo != "" {
		gn.Formula = fields.FormulaTo
	}
	// For Enumeration nodes, extract enum entries.
	if kind == "Enumeration" {
		gn.Entries = parseEnumEntries(inner)
	}
	return gn
}

// parseNodeMapXML is the low-level XML stream parser that yields GenICam node elements.
// It iterates over a RegisterDescription XML document and yields (kind, name, innerXML) tuples
// for recognized node types, skipping structural nodes like Category and Group.
//
// Yields recognized node types:
//   - Integer, Boolean, Float, String, Enumeration, Command
//   - IntReg, FloatReg, StringReg, MaskedIntReg
//   - IntSwissKnife, SwissKnife, IntConverter, Converter
//
// Skips structural containers:
//   - RegisterDescription, Group, Category, StructReg, Port, Node, XMLDescription
func parseNodeMapXML(xmlData []byte, onNode func(kind, name string, inner []byte)) error {
	dec := xml.NewDecoder(bytes.NewReader(xmlData))

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
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
				return err
			}
			onNode(kind, name, holder.Inner)
		case "RegisterDescription", "Group", "Category", "StructReg",
			"Port", "Node", "XMLDescription":
			// Structural containers; skip their content.
			continue
		default:
			_ = dec.Skip()
		}
	}
	return nil
}
