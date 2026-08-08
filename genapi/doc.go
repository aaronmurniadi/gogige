// Package genapi provides GenICam GenApi XML parser and node map implementation.
//
// GenICam Standard Version 2.1.1
// GenICam GenApi Standard Version 2.1.1
// GenICam SFNC Version 2.7
// Reference: GenICam_Standard_v2_1_1.pdf, GenICam_GenApi_Standard_v2_1_1.pdf, GenICam_SFNC_v2_7.pdf
//
// # Overview
//
// GenApi (GenICam Standard for Camera Control) defines a common interface for
// accessing camera features through XML descriptions. This package parses the
// XML, builds a node tree, and provides evaluation of expressions.
//
// # Package Content
//
// genapi/ provides:
//   - CameraDescription: Manifest parsing and XML decompression (ZIP/DEFLATE)
//   - NodeMap: Node tree structure for feature access
//   - Node: Core Node interface (Integer, Float, Enumeration, Category, etc.)
//   - Evaluator: SwissKnife expression parser and evaluation
//   - Port: GenApi node binding to GVCP client for register access
//
// # Node Types
//
// Supported node types:
//   - Category: Grouping node for organizing features
//   - Integer / IntReg: Integer value nodes with register binding
//   - MaskedIntReg: Integer with bit mask operations
//   - Float / FloatReg: Floating-point value nodes
//   - Enumeration / EnumEntry: Enumerated value choices
//   - Boolean: True/false value
//   - Command: Execute action nodes
//   - StringReg: String value with register binding
//   - SwissKnife / IntSwissKnife: Expression evaluation nodes
//   - Converter: Format conversion nodes
//   - Port: Register space access port
//
// # Expression Evaluation
//
// SwissKnife/IntSwissKnife nodes support:
//   - Operators: +, -, *, /, %, **, &, |, ^, <<, >>, &&, ||
//   - Comparisons: ==, !=, <, <=, >, >=
//   - Ternary: condition ? value1 : value2
//   - Functions: SQRT, FLOOR, CEIL, ABS
//
// # Port Binding
//
// Port nodes bind GenApi features to GVCP client register operations:
//   - Read(addr, len) -> ReadRegister / ReadMemory
//   - Write(addr, data) -> WriteRegister / WriteMemory
//   - Respects device endianness (BigEndian / LittleEndian)
//
// References
//
//   - GenICam Standard v2.1.1
//   - GenICam GenApi Standard v2.1.1
//   - GenICam Standard Feature Name Convention (SFNC) v2.7
//   - https://www.emva.org/standards-technical-documents/
package genapi
