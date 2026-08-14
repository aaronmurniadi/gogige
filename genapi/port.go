package genapi

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"

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
	if addr == 0 {
		return 0, 0, fmt.Errorf("gige: feature %q has no register address (kind=%s addresses=%v pAddress=%v pValue=%s)",
			n.Name, n.Kind, n.Addresses, n.PAddresses, n.PValue)
	}
	return uint32(addr), length, nil
}

// readIntReg reads an integer-like register from device memory.
func (pa *portAdapter) readIntReg(n *gcNode, nm *NodeMap) (uint64, error) {
	addr, length, err := pa.resolveAddr(n, nm)
	if err != nil {
		return 0, err
	}
	if length == 4 {
		v, err := pa.port.ReadReg(addr)
		return uint64(v), err
	}
	// For non-4-byte reads, use ReadMem.
	data, err := pa.port.ReadMem(addr, length)
	if err != nil {
		return 0, err
	}
	order := pa.deviceByteOrder()
	switch length {
	case 1:
		return uint64(data[0]), nil
	case 2:
		return uint64(order.Uint16(data)), nil
	case 8:
		return order.Uint64(data), nil
	default:
		return uint64(order.Uint32(data[:4])), nil
	}
}

// writeIntReg writes an integer-like register to device memory.
// For MaskedIntReg nodes, it performs a read-modify-write using LSB/MSB mask bits.
func (pa *portAdapter) writeIntReg(n *gcNode, v int64, nm *NodeMap) error {
	addr, length, err := pa.resolveAddr(n, nm)
	if err != nil {
		return err
	}
	if n.Kind == "MaskedIntReg" && n.HasMask && length == 4 {
		cur, err := pa.port.ReadReg(addr)
		if err != nil {
			return fmt.Errorf("gige: read %s @0x%x for mask: %w", n.Name, addr, err)
		}
		mask, shift := maskFromBits(n.LSB, n.MSB)
		nv := (cur &^ mask) | ((uint32(v) << shift) & mask)
		if err := pa.port.WriteReg(addr, nv); err != nil {
			return fmt.Errorf("gige: write MaskedIntReg %s @0x%x (val=%d mask=0x%x): %w", n.Name, addr, v, mask, err)
		}
		return nil
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
	order := pa.deviceByteOrder()
	if length >= 8 {
		var b [8]byte
		order.PutUint64(b[:], math.Float64bits(v))
		return pa.port.WriteMem(addr, b[:])
	}
	var b [4]byte
	order.PutUint32(b[:], math.Float32bits(float32(v)))
	return pa.port.WriteMem(addr, b[:])
}

// readFloatReg reads a floating-point register from device memory.
func (pa *portAdapter) readFloatReg(n *gcNode, nm *NodeMap) (float64, error) {
	addr, length, err := pa.resolveAddr(n, nm)
	if err != nil {
		return 0, err
	}
	order := pa.deviceByteOrder()
	if length >= 8 {
		data, err := pa.port.ReadMem(addr, 8)
		if err != nil {
			return 0, err
		}
		return math.Float64frombits(order.Uint64(data)), nil
	}
	v, err := pa.port.ReadReg(addr)
	if err != nil {
		return 0, err
	}
	return float64(math.Float32frombits(v)), nil
}

// readStringReg reads a string register from device memory.
func (pa *portAdapter) readStringReg(n *gcNode, nm *NodeMap) (string, error) {
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

// evaluateFormulaVariables builds a map of variable names to their computed integer values
// for SwissKnife expression evaluation.
func (pa *portAdapter) evaluateFormulaVariables(variables map[string]string, nm *NodeMap) (map[string]int64, error) {
	vars := make(map[string]int64, len(variables))
	for varName, feat := range variables {
		v, err := pa.readFormulaVariable(varName, feat, nm, 0)
		if err != nil {
			return nil, err
		}
		vars[varName] = v
	}
	return vars, nil
}

// evaluateIntegerFormula computes a SwissKnife formula's value given a variable map.
// This is a wrapper around the evaluator that returns uint64 for address/value computation.
func (pa *portAdapter) evaluateIntegerFormula(formula string, vars map[string]int64) (uint64, error) {
	v, err := evalFormula(formula, vars)
	return uint64(v), err
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
	case "IntSwissKnife", "SwissKnife", "IntConverter", "Converter":
		vars, err := pa.evaluateFormulaVariables(n.Variables, nm)
		if err != nil {
			return 0, fmt.Errorf("gige: %s vars: %w", n.Name, err)
		}
		if n.Formula == "" {
			return 0, fmt.Errorf("gige: %s has empty Formula", n.Name)
		}
		return pa.evaluateIntegerFormula(n.Formula, vars)
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
