package genapi

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"

	"github.com/aaronmurniadi/gogige/gvcp"
)

// portAdapter wraps a gvcp.Port and provides GenApi-specific register resolution and I/O.
type portAdapter struct {
	port gvcp.Port
}

// newPortAdapter creates a portAdapter for a given gvcp.Port.
func newPortAdapter(port gvcp.Port) *portAdapter {
	return &portAdapter{port: port}
}

// addressable reports whether a node declares a register address of its own,
// i.e. at least one <Address> or <pAddress> element. A node without any is a
// Constant (see §2.8.13) and must not be read as a register.
func addressable(n *gcNode) bool {
	return len(n.Addresses) > 0 || len(n.PAddresses) > 0
}

// resolveAddr computes the effective device address for a node by summing
// its constant Address with evaluated pAddress pointers.
// Returns (address, length, error).
func (pa *portAdapter) resolveAddr(n *gcNode, nm *NodeMap) (uint32, int, error) {
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
	// An address of 0 is legitimate (e.g. GevVersionReg at ABRM 0x0000); only
	// reject nodes that declare no address at all (Constants), and never
	// silently truncate 64-bit pAddress sums (GVCP registers are 32-bit).
	if !addressable(n) {
		return 0, 0, fmt.Errorf("gige: feature %q has no register address (kind=%s addresses=%v pAddress=%v pValue=%s)",
			n.Name, n.Kind, n.Addresses, n.PAddresses, n.PValue)
	}
	if addr > math.MaxUint32 {
		return 0, 0, fmt.Errorf("gige: feature %q address 0x%x exceeds 32-bit register space",
			n.Name, addr)
	}
	return uint32(addr), length, nil
}

// readIntReg reads an integer-like register from device memory.
func (pa *portAdapter) readIntReg(n *gcNode, nm *NodeMap) (uint64, error) {
	addr, length, err := pa.resolveAddr(n, nm)
	if err != nil {
		return 0, err
	}
	raw, err := pa.readRawValue(n, addr, length)
	if err != nil {
		return 0, err
	}
	// MaskedIntReg: extract the LSB..MSB bit slice (GenApi 2.1.1 §2.8.5).
	if n.Kind == "MaskedIntReg" && n.HasMask {
		mask, shift := maskFromBits(n.LSB, n.MSB)
		raw = (raw & uint64(mask)) >> shift
	}
	// Sign-bit extension for <Sign>Signed</Sign> nodes narrower than 64 bits.
	if n.Sign == "Signed" {
		if bits := intBitWidth(n); bits > 0 && bits < 64 {
			if raw&(uint64(1)<<uint(bits-1)) != 0 {
				raw |= ^uint64(0) << uint(bits)
			}
		}
	}
	return raw, nil
}

// readRawValue reads the raw register value, honoring an explicit per-register
// <Endianess> element (a LittleEndian 4-byte register is byte-swapped).
func (pa *portAdapter) readRawValue(n *gcNode, addr uint32, length int) (uint64, error) {
	if n.Endianess != "" {
		data, err := pa.port.ReadMem(addr, length)
		if err != nil {
			return 0, err
		}
		return decodeRegUint(pa.registerByteOrder(n), data, length), nil
	}
	if length == 4 {
		v, err := pa.port.ReadReg(addr)
		if err != nil {
			return 0, err
		}
		return uint64(v), nil
	}
	data, err := pa.port.ReadMem(addr, length)
	if err != nil {
		return 0, err
	}
	return decodeRegUint(pa.deviceByteOrder(), data, length), nil
}

// registerByteOrder returns the byte order declared by an explicit <Endianess>
// element on a register node, falling back to the device byte order.
func (pa *portAdapter) registerByteOrder(n *gcNode) binary.ByteOrder {
	switch n.Endianess {
	case "LittleEndian":
		return binary.LittleEndian
	case "BigEndian":
		return binary.BigEndian
	}
	return pa.deviceByteOrder()
}

// intBitWidth returns the number of meaningful bits of an integer-like register
// node, used for sign extension. MaskedIntReg spans LSB..MSB.
func intBitWidth(n *gcNode) int {
	if n.Kind == "MaskedIntReg" && n.HasMask {
		if w := n.MSB - n.LSB + 1; w > 0 {
			return w
		}
	}
	switch n.Length {
	case 1:
		return 8
	case 2:
		return 16
	case 4:
		return 32
	case 8:
		return 64
	}
	return 32
}

// decodeRegUint decodes a register value read as raw bytes.
func decodeRegUint(order binary.ByteOrder, data []byte, length int) uint64 {
	switch length {
	case 1:
		return uint64(data[0])
	case 2:
		return uint64(order.Uint16(data))
	case 8:
		return order.Uint64(data)
	default:
		return uint64(order.Uint32(data[:4]))
	}
}

// encodeRegUint encodes a register value into length bytes in the given order.
func encodeRegUint(order binary.ByteOrder, v uint64, length int) []byte {
	buf := make([]byte, length)
	switch length {
	case 1:
		buf[0] = byte(v)
	case 2:
		order.PutUint16(buf, uint16(v))
	case 8:
		order.PutUint64(buf, v)
	default:
		order.PutUint32(buf[:4], uint32(v))
	}
	return buf
}

// writeIntReg writes an integer-like register to device memory.
// For MaskedIntReg nodes, it performs a read-modify-write using LSB/MSB mask bits.
func (pa *portAdapter) writeIntReg(n *gcNode, v int64, nm *NodeMap) error {
	addr, length, err := pa.resolveAddr(n, nm)
	if err != nil {
		return err
	}
	if n.Kind == "MaskedIntReg" && n.HasMask && length <= 4 {
		cur, err := pa.readRawValue(n, addr, length)
		if err != nil {
			return fmt.Errorf("gige: read %s @0x%x for mask: %w", n.Name, addr, err)
		}
		mask, shift := maskFromBits(n.LSB, n.MSB)
		v = int64((cur &^ uint64(mask)) | ((uint64(uint32(v)) << shift) & uint64(mask)))
	}
	if n.Endianess != "" {
		order := pa.registerByteOrder(n)
		return pa.port.WriteMem(addr, encodeRegUint(order, uint64(v), length))
	}
	switch length {
	case 4:
		if err := pa.port.WriteReg(addr, uint32(v)); err != nil {
			return fmt.Errorf("gige: write %s @0x%x = %d: %w", n.Name, addr, v, err)
		}
		return nil
	case 8:
		var b [8]byte
		pa.deviceByteOrder().PutUint64(b[:], uint64(v))
		return pa.port.WriteMem(addr, b[:])
	case 2:
		var b [2]byte
		pa.deviceByteOrder().PutUint16(b[:], uint16(v))
		return pa.port.WriteMem(addr, b[:])
	case 1:
		return pa.port.WriteMem(addr, []byte{byte(v)})
	default:
		if err := pa.port.WriteReg(addr, uint32(v)); err != nil {
			return fmt.Errorf("gige: write %s @0x%x = %d: %w", n.Name, addr, v, err)
		}
		return nil
	}
}

// writeFloatReg writes a floating-point register to device memory.
func (pa *portAdapter) writeFloatReg(n *gcNode, v float64, nm *NodeMap) error {
	addr, length, err := pa.resolveAddr(n, nm)
	if err != nil {
		return err
	}
	order := pa.registerByteOrder(n)
	if length >= 8 {
		var b [8]byte
		order.PutUint64(b[:], math.Float64bits(v))
		return pa.port.WriteMem(addr, b[:])
	}
	var b [4]byte
	order.PutUint32(b[:], math.Float32bits(float32(v)))
	return pa.port.WriteMem(addr, b[:])
}

// readFloatReg reads a floating-point register from device memory. A constant
// Float node (a <Value> with no regsiter address, §2.8.13) yields its parsed
// constant instead of a register access.
func (pa *portAdapter) readFloatReg(n *gcNode, nm *NodeMap) (float64, error) {
	if !addressable(n) && n.Value != "" {
		f, err := strconv.ParseFloat(n.Value, 64)
		if err != nil {
			return 0, fmt.Errorf("gige: feature %q constant Value %q: %w", n.Name, n.Value, err)
		}
		return f, nil
	}
	addr, length, err := pa.resolveAddr(n, nm)
	if err != nil {
		return 0, err
	}
	order := pa.registerByteOrder(n)
	if n.Endianess != "" || length >= 8 {
		data, err := pa.port.ReadMem(addr, length)
		if err != nil {
			return 0, err
		}
		if length >= 8 {
			return math.Float64frombits(order.Uint64(data)), nil
		}
		return float64(math.Float32frombits(order.Uint32(data[:4]))), nil
	}
	v, err := pa.port.ReadReg(addr)
	if err != nil {
		return 0, err
	}
	return float64(math.Float32frombits(v)), nil
}

// readStringReg reads a string register from device memory. A constant String
// node (a <Value> with no register address) yields its constant.
func (pa *portAdapter) readStringReg(n *gcNode, nm *NodeMap) (string, error) {
	if !addressable(n) && n.Value != "" {
		return n.Value, nil
	}
	addr, length, err := pa.resolveAddr(n, nm)
	if err != nil {
		return "", err
	}
	if length <= 0 {
		length = 256
	}
	data, err := pa.port.ReadMem(addr, length)
	if err != nil {
		return "", err
	}
	if i := bytes.IndexByte(data, 0); i >= 0 {
		data = data[:i]
	}
	return string(data), nil
}

// writeStringReg writes a string register to device memory.
func (pa *portAdapter) writeStringReg(n *gcNode, val string, nm *NodeMap) error {
	target := n
	if n.PValue != "" {
		reg, err := nm.lookup(n.PValue)
		if err != nil {
			return err
		}
		target = reg
	}
	addr, length, err := pa.resolveAddr(target, nm)
	if err != nil {
		return err
	}
	if length <= 0 {
		length = len(val) + 1
	}
	buf := make([]byte, length)
	copy(buf, val)
	return pa.port.WriteMem(addr, buf)
}

// deviceByteOrder returns the byte order for device register access.
// Checks if the gvcp.Port implements a DeviceByteOrder method; otherwise defaults to BigEndian.
// This respects the ImplementationEndianness register (0x020C) if present.
func (pa *portAdapter) deviceByteOrder() binary.ByteOrder {
	type orderer interface {
		DeviceByteOrder() binary.ByteOrder
	}
	if o, ok := pa.port.(orderer); ok {
		if order := o.DeviceByteOrder(); order != nil {
			return order
		}
	}
	return binary.BigEndian
}

// readFormulaVariable resolves a SwissKnife variable name to its integer value
// by evaluating the referenced feature's value through the NodeMap.
func (pa *portAdapter) readFormulaVariable(varName, featName string, nm *NodeMap, depth int) (int64, error) {
	v, err := nm.evalIntegerValue(featName, depth)
	if err != nil {
		return 0, fmt.Errorf("gige: var %s→%s: %w", varName, featName, err)
	}
	return int64(v), nil
}

// formulaContext builds the eager variable values plus a suffix resolver that
// maps "Var.Min/.Max/.Inc/.Value/.Entry" (GenApi 2.1.1 §2.8.13) onto the
// referenced feature's constraints. Variables already bound in extra take the
// caller-provided value without being evaluated from the feature map.
func (pa *portAdapter) formulaContext(variables map[string]string, nm *NodeMap, extra map[string]int64) (map[string]int64, suffixResolver, error) {
	vars := make(map[string]int64, len(variables))
	for varName, feat := range variables {
		if _, bound := extra[varName]; bound {
			vars[varName] = extra[varName]
			continue
		}
		v, err := pa.readFormulaVariable(varName, feat, nm, 0)
		if err != nil {
			return nil, nil, err
		}
		vars[varName] = v
	}
	for k, v := range extra {
		if _, isVar := variables[k]; !isVar {
			vars[k] = v
		}
	}
	resolve := func(name, suffix string) (int64, bool, error) {
		feat, ok := variables[name]
		if !ok {
			return 0, false, nil
		}
		if _, err := nm.lookup(feat); err != nil {
			return 0, false, fmt.Errorf("gige: var %s suffix %s: %w", name, suffix, err)
		}
		switch suffix {
		case "Value", "Entry":
			v, err := nm.evalIntegerValue(feat, 0)
			return int64(v), true, err
		case "Min", "Max", "Inc":
			var v int64
			var has bool
			var err error
			switch suffix {
			case "Min":
				v, has, err = nm.GetMin(feat)
			case "Max":
				v, has, err = nm.GetMax(feat)
			case "Inc":
				v, has, err = nm.GetInc(feat)
			}
			if err != nil {
				return 0, true, fmt.Errorf("gige: var %s.%s: %w", name, suffix, err)
			}
			if !has {
				return 0, true, fmt.Errorf("gige: var %s: feature %q has no %s constraint", name, feat, suffix)
			}
			return v, true, nil
		}
		return 0, true, fmt.Errorf("gige: var %s has no such suffix %q", name, suffix)
	}
	return vars, resolve, nil
}

// evaluateSwissKnife computes a SwissKnife formula where pVariables may use
// .Value/.Min/.Max/.Inc/.Entry suffixes (GenApi 2.1.1 §2.8.13).
func (pa *portAdapter) evaluateSwissKnife(formula string, variables map[string]string, nm *NodeMap, extra map[string]int64) (uint64, error) {
	vars, resolve, err := pa.formulaContext(variables, nm, extra)
	if err != nil {
		return 0, err
	}
	v, err := evalFormulaRefs(formula, vars, resolve)
	return uint64(v), err
}

// converterFormula returns the SwissKnife formula to apply in the given
// direction for Converter/IntConverter nodes (GenApi 2.1.1 §2.8.10):
//   - forward (register -> user, read): FormulaFrom, else legacy Formula,
//     else FormulaTo.
//   - backward (user -> register, write): FormulaTo, else legacy Formula,
//     else FormulaFrom.
func converterFormula(n *gcNode, forward bool) string {
	if forward {
		if n.FormulaFrom != "" {
			return n.FormulaFrom
		}
		if n.Formula != "" {
			return n.Formula
		}
		return n.FormulaTo
	}
	if n.FormulaTo != "" {
		return n.FormulaTo
	}
	if n.Formula != "" {
		return n.Formula
	}
	return n.FormulaFrom
}

// resolveIntegerReference follows pValue and formula chains to compute an integer value
// for nodes that represent integers indirectly (Integer, Enumeration, SwissKnife, etc.).
func (pa *portAdapter) resolveIntegerReference(n *gcNode, nm *NodeMap, depth int) (uint64, error) {
	if depth > 12 {
		return 0, fmt.Errorf("gige: pAddress recursion on %q", n.Name)
	}
	switch n.Kind {
	case "Integer":
		if n.Value != "" {
			return parseUint(n.Value)
		}
		if n.PValue != "" {
			reg, err := nm.lookup(n.PValue)
			if err != nil {
				return 0, err
			}
			return pa.resolveIntegerReference(reg, nm, depth+1)
		}
		if n.Address != 0 {
			return n.Address, nil
		}
	case "IntSwissKnife", "SwissKnife":
		if n.Formula == "" {
			return 0, fmt.Errorf("gige: %s has empty Formula", n.Name)
		}
		return pa.evaluateSwissKnife(n.Formula, n.Variables, nm, nil)
	case "IntConverter", "Converter":
		formula := converterFormula(n, true)
		if formula == "" {
			return 0, fmt.Errorf("gige: %s has no FormulaFrom", n.Name)
		}
		if formulaUses(formula, "TO") {
			// FormulaFrom exposes the current register value as the reserved
			// variable TO (GenApi 2.1.1 §2.8.13), e.g. "(TO & 0x00080000) >> 19".
			reg, err := nm.lookup(n.PValue)
			if err != nil {
				return 0, fmt.Errorf("gige: %s FormulaFrom uses TO but pValue %q: %w", n.Name, n.PValue, err)
			}
			rv, err := pa.resolveIntegerReference(reg, nm, depth+1)
			if err != nil {
				return 0, fmt.Errorf("gige: %s FormulaFrom TO: %w", n.Name, err)
			}
			return pa.evaluateSwissKnife(formula, n.Variables, nm, map[string]int64{"TO": int64(rv)})
		}
		return pa.evaluateSwissKnife(formula, n.Variables, nm, nil)
	case "IntReg", "MaskedIntReg":
		return pa.readIntReg(n, nm)
	case "Enumeration":
		if n.PValue != "" {
			reg, err := nm.lookup(n.PValue)
			if err != nil {
				return 0, err
			}
			return pa.resolveIntegerReference(reg, nm, depth+1)
		}
	}
	return 0, fmt.Errorf("gige: cannot evaluate integer %q (kind=%s value=%q pValue=%q addr=0x%x pAddr=%v)",
		n.Name, n.Kind, n.Value, n.PValue, n.Address, n.PAddresses)
}

// resolveIntegerTarget follows pValue chains to find the actual register node to write to.
// For Boolean, Enumeration, Integer nodes that wrap a register, this resolves to the register.
func (pa *portAdapter) resolveIntegerTarget(n *gcNode, nm *NodeMap) (*gcNode, error) {
	target := n
	if n.PValue != "" {
		reg, err := nm.lookup(n.PValue)
		if err != nil {
			return nil, err
		}
		target = reg
		// Enumeration/Integer may point at another Integer that points at IntReg.
		if target.PValue != "" && (target.Kind == "Integer" || target.Kind == "Enumeration") {
			reg2, err := nm.lookup(target.PValue)
			if err != nil {
				return nil, err
			}
			target = reg2
		}
	}
	return target, nil
}
