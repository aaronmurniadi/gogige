package genapi

// Node is the core GenICam Node interface for all feature types.
// Concrete node kinds (Integer, IntReg, MaskedIntReg, etc.) implement this interface
// by embedding gcNode and providing type-specific accessors.
type Node interface {
	// GetName returns the node's Name attribute.
	GetName() string

	// GetKind returns the node's GenApi kind string
	// (e.g., "Integer", "IntReg", "Enumeration", "Boolean", "Command", "Float", "String", etc.).
	GetKind() string

	// GetAccess returns the AccessMode for this node ("RO", "RW", "WO", etc.).
	GetAccess() string

	// GetConstraints returns constraint pointers (pMin, pMax, pInc) and static values.
	// pMin/pMax/pInc are feature names to evaluate; Min/Max/Inc are static values.
	// For Integer/Float nodes, at least one of these may be populated.
	GetConstraints() (pMin, pMax, pInc string, minVal, maxVal, incVal int64)

	// GetInvalidator returns the pInvalidator feature name for this node,
	// or "" if not set. When the referenced feature changes, this node is invalidated.
	GetInvalidator() string
}

// gcNode is the internal representation of a GenICam node.
// It holds raw parsed data that may be resolved later via pValue/pAddress pointers.
type gcNode struct {
	// Name is the node's Name attribute from XML.
	Name string

	// Kind is the GenApi node kind: "Integer", "IntReg", "Boolean", "Enumeration", etc.
	Kind string

	// Address is the sum of all <Address> children (constant offsets).
	Address uint64

	// Addresses stores each individual <Address> for debug/inspection.
	Addresses []uint64

	// Length is the register size in bytes (typically 1, 2, 4, 8).
	Length int

	// Access is the AccessMode: "RO", "RW", "WO", etc.
	Access string

	// PValue is the feature name this node points to (via <pValue>).
	// Used for Boolean, Enumeration, Integer nodes that wrap a register.
	PValue string

	// PAddresses are feature names to evaluate and sum into the address (via <pAddress>).
	PAddresses []string

	// Value is a constant <Value> for this node (used by Integer, Enumeration entries).
	Value string

	// Entries maps enum entry names to their integer values (for Enumeration nodes).
	Entries map[string]int64

	// Variables maps SwissKnife variable names (e.g., "X") to feature names
	// (used by SwissKnife, IntSwissKnife, Converter, IntConverter nodes).
	Variables map[string]string

	// Formula is the SwissKnife expression for this node
	// (used by SwissKnife, IntSwissKnife, Converter, IntConverter nodes).
	Formula string

	// LSB and MSB define the bit range for MaskedIntReg nodes.
	// LSB=MSB when <Bit> is used; both set when <LSB>+<MSB> are present.
	LSB int
	MSB int

	// HasMask is true when LSB and MSB define a valid bit mask (both >= 0).
	HasMask bool

	// Constraint pointers for Integer and Float nodes.
	// These reference other feature nodes that compute min/max/inc values.
	// GenApi: typically SwissKnife nodes that evaluate to numeric bounds.
	PMin string
	PMax string
	PInc string

	// Availability and lock pointers.
	// These reference other features that determine whether this node is
	// implemented, available, or locked.
	PIsImplemented string
	PIsAvailable   string
	PIsLocked      string

	// PInvalidator is a feature name that, when changed, invalidates this node.
	// The node must be re-read after the invalidator changes.
	PInvalidator string

	// Static constraint values for Integer and Float nodes.
	// Used when Min/Max/Inc are constant rather than computed via pointers.
	Min int64
	Max int64
	Inc int64
}

// GetName implements Node.
func (n *gcNode) GetName() string {
	return n.Name
}

// GetKind implements Node.
func (n *gcNode) GetKind() string {
	return n.Kind
}

// GetAccess implements Node.
func (n *gcNode) GetAccess() string {
	return n.Access
}

// GetConstraints implements Node.
func (n *gcNode) GetConstraints() (pMin, pMax, pInc string, minVal, maxVal, incVal int64) {
	return n.PMin, n.PMax, n.PInc, n.Min, n.Max, n.Inc
}

// GetInvalidator implements Node.
func (n *gcNode) GetInvalidator() string {
	return n.PInvalidator
}
