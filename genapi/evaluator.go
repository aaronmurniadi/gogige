package genapi

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

var formulaFuncs = map[string]func([]int64) (int64, error){
	// Evaluated in the float domain; results are truncated back to the integer
	// domain this evaluator operates on (GenApi 2.1.1 §2.8.13).
	"SGN": func(a []int64) (int64, error) {
		switch {
		case a[0] > 0:
			return 1, nil
		case a[0] < 0:
			return -1, nil
		}
		return 0, nil
	},
	"NEG":   func(a []int64) (int64, error) { return -a[0], nil },
	"ABS":   func(a []int64) (int64, error) { return ffn(math.Abs, a) },
	"FLOOR": func(a []int64) (int64, error) { return ffn(math.Floor, a) },
	"CEIL":  func(a []int64) (int64, error) { return ffn(math.Ceil, a) },
	"TRUNC": func(a []int64) (int64, error) { return ffn(math.Trunc, a) },
	"ROUND": func(a []int64) (int64, error) {
		prec := 0
		if len(a) > 1 {
			prec = int(a[1])
		}
		scale := math.Pow(10, float64(prec))
		return int64(math.Round(float64(a[0])*scale) / scale), nil
	},
	"SQRT": func(a []int64) (int64, error) { return ffn(math.Sqrt, a) },
	"EXP":  func(a []int64) (int64, error) { return ffn(math.Exp, a) },
	"LN":   func(a []int64) (int64, error) { return ffn(math.Log, a) },
	"LG":   func(a []int64) (int64, error) { return ffn(math.Log10, a) },
	"SIN":  func(a []int64) (int64, error) { return ffn(math.Sin, a) },
	"COS":  func(a []int64) (int64, error) { return ffn(math.Cos, a) },
	"TAN":  func(a []int64) (int64, error) { return ffn(math.Tan, a) },
	"ATAN": func(a []int64) (int64, error) { return ffn(math.Atan, a) },
	"ASIN": func(a []int64) (int64, error) { return ffn(math.Asin, a) },
	"ACOS": func(a []int64) (int64, error) { return ffn(math.Acos, a) },
}

// ffn applies a float->float math function to the first argument and truncates
// the result back to the integer domain.
func ffn(f func(float64) float64, a []int64) (int64, error) {
	return int64(f(float64(a[0]))), nil
}

// formulaConsts maps SwissKnife symbolic constants (GenApi 2.1.1 §2.8.13) into
// the integer domain used by this evaluator.
var formulaConsts = map[string]int64{
	"E":  floatToInt(math.E),
	"PI": floatToInt(math.Pi),
}

// floatToInt truncates a float64 to the integer domain. Kept as a function so
// the truncation happens at run time (constant folding would reject it).
func floatToInt(f float64) int64 { return int64(f) }

// evalFormula evaluates a GenICam SwissKnife formula with integer variables.
// Supports + - * / % ** & | ^ << >> ~ ( ) ? : = == <> != < > <= >= && || and
// hex/decimal literals.
// Supports functions: SGN NEG ABS FLOOR CEIL TRUNC ROUND SQRT EXP LN LG
// SIN COS TAN ATAN ASIN ACOS and constants E, PI.
func evalFormula(expr string, vars map[string]int64) (int64, error) {
	p := &formParser{s: strings.TrimSpace(expr), vars: vars}
	v, err := p.parseExpr()
	if err != nil {
		return 0, err
	}
	if strings.TrimSpace(p.s[p.i:]) != "" {
		return 0, fmt.Errorf("gige: formula trailing junk %q", p.s[p.i:])
	}
	return v, nil
}

type formParser struct {
	s    string
	i    int
	vars map[string]int64
}

func (p *formParser) peek() byte {
	p.skipSpace()
	if p.i >= len(p.s) {
		return 0
	}
	return p.s[p.i]
}

func (p *formParser) skipSpace() {
	for p.i < len(p.s) && unicode.IsSpace(rune(p.s[p.i])) {
		p.i++
	}
}

func (p *formParser) accept(prefix string) bool {
	p.skipSpace()
	if strings.HasPrefix(p.s[p.i:], prefix) {
		p.i += len(prefix)
		return true
	}
	return false
}

// acceptNE consumes the "<>" not-equal operator, tolerating whitespace between
// the two characters, e.g. "A < > B".
func (p *formParser) acceptNE() bool {
	p.skipSpace()
	if p.i >= len(p.s) || p.s[p.i] != '<' {
		return false
	}
	j := p.i + 1
	for j < len(p.s) && unicode.IsSpace(rune(p.s[j])) {
		j++
	}
	if j >= len(p.s) || p.s[j] != '>' {
		return false
	}
	p.i = j + 1
	return true
}

// isNE reports whether the next token is the "<>" not-equal operator (with
// optional whitespace between the characters), without consuming it.
func (p *formParser) isNE() bool {
	p.skipSpace()
	if p.i >= len(p.s) || p.s[p.i] != '<' {
		return false
	}
	j := p.i + 1
	for j < len(p.s) && unicode.IsSpace(rune(p.s[j])) {
		j++
	}
	return j < len(p.s) && p.s[j] == '>'
}

// formulaUses reports whether the formula text references the given identifier
// as a standalone token. Used to detect the GenICam converter reserved
// variables FROM and TO (GenApi 2.1.1 §2.8.13).
func formulaUses(formula, tok string) bool {
	for i := 0; i+len(tok) <= len(formula); i++ {
		if formula[i] != tok[0] || formula[i:i+len(tok)] != tok {
			continue
		}
		before := i == 0 || !isIdentByte(formula[i-1])
		after := i+len(tok) == len(formula) || !isIdentByte(formula[i+len(tok)])
		if before && after {
			return true
		}
	}
	return false
}

func isIdentByte(b byte) bool {
	return unicode.IsLetter(rune(b)) || unicode.IsDigit(rune(b)) || b == '_'
}

func (p *formParser) parseExpr() (int64, error) {
	return p.parseTernary()
}

func (p *formParser) parseTernary() (int64, error) {
	cond, err := p.parseLogOr()
	if err != nil {
		return 0, err
	}
	if !p.accept("?") {
		return cond, nil
	}
	thenV, err := p.parseTernary()
	if err != nil {
		return 0, err
	}
	if !p.accept(":") {
		return 0, fmt.Errorf("gige: formula missing : in ternary")
	}
	elseV, err := p.parseTernary()
	if err != nil {
		return 0, err
	}
	if cond != 0 {
		return thenV, nil
	}
	return elseV, nil
}

func (p *formParser) parseLogOr() (int64, error) {
	v, err := p.parseLogAnd()
	if err != nil {
		return 0, err
	}
	for p.accept("||") {
		r, err := p.parseLogAnd()
		if err != nil {
			return 0, err
		}
		if v != 0 || r != 0 {
			v = 1
		} else {
			v = 0
		}
	}
	return v, nil
}

func (p *formParser) parseLogAnd() (int64, error) {
	v, err := p.parseOr()
	if err != nil {
		return 0, err
	}
	for p.accept("&&") {
		r, err := p.parseOr()
		if err != nil {
			return 0, err
		}
		if v != 0 && r != 0 {
			v = 1
		} else {
			v = 0
		}
	}
	return v, nil
}

func (p *formParser) parseOr() (int64, error) {
	v, err := p.parseXor()
	if err != nil {
		return 0, err
	}
	for {
		p.skipSpace()
		if p.i < len(p.s) && p.s[p.i] == '|' && (p.i+1 >= len(p.s) || p.s[p.i+1] != '|') {
			p.i++
			r, err := p.parseXor()
			if err != nil {
				return 0, err
			}
			v |= r
			continue
		}
		break
	}
	return v, nil
}

func (p *formParser) parseXor() (int64, error) {
	v, err := p.parseAnd()
	if err != nil {
		return 0, err
	}
	for p.accept("^") {
		r, err := p.parseAnd()
		if err != nil {
			return 0, err
		}
		v ^= r
	}
	return v, nil
}

func (p *formParser) parseAnd() (int64, error) {
	v, err := p.parseEquality()
	if err != nil {
		return 0, err
	}
	for {
		p.skipSpace()
		if p.i < len(p.s) && p.s[p.i] == '&' && (p.i+1 >= len(p.s) || p.s[p.i+1] != '&') {
			p.i++
			r, err := p.parseEquality()
			if err != nil {
				return 0, err
			}
			v &= r
			continue
		}
		break
	}
	return v, nil
}

func (p *formParser) parseEquality() (int64, error) {
	v, err := p.parseRel()
	if err != nil {
		return 0, err
	}
	for {
		var op string
		switch {
		case p.acceptNE(), p.accept("!="):
			op = "!="
		case p.accept("=="), p.accept("="):
			op = "="
		default:
			return v, nil
		}
		r, err := p.parseRel()
		if err != nil {
			return 0, err
		}
		switch op {
		case "=":
			if v == r {
				v = 1
			} else {
				v = 0
			}
		case "!=":
			if v != r {
				v = 1
			} else {
				v = 0
			}
		}
	}
}

func (p *formParser) parseRel() (int64, error) {
	v, err := p.parseShift()
	if err != nil {
		return 0, err
	}
	for {
		var op string
		switch {
		case p.accept("<="):
			op = "<="
		case p.accept(">="):
			op = ">="
		case p.isNE():
			// "<>" is the GenICam SwissKnife not-equal operator. It belongs to
			// the equality tier, so leave it for parseEquality to consume
			// instead of swallowing its "<" as a relational operator here.
			return v, nil
		case p.accept("<"):
			op = "<"
		case p.accept(">"):
			op = ">"
		default:
			return v, nil
		}
		r, err := p.parseShift()
		if err != nil {
			return 0, err
		}
		switch op {
		case "<=":
			if v <= r {
				v = 1
			} else {
				v = 0
			}
		case ">=":
			if v >= r {
				v = 1
			} else {
				v = 0
			}
		case "<":
			if v < r {
				v = 1
			} else {
				v = 0
			}
		case ">":
			if v > r {
				v = 1
			} else {
				v = 0
			}
		}
	}
}

func (p *formParser) parseShift() (int64, error) {
	v, err := p.parseAdd()
	if err != nil {
		return 0, err
	}
	for {
		if p.accept("<<") {
			r, err := p.parseAdd()
			if err != nil {
				return 0, err
			}
			v <<= uint(r)
		} else if p.accept(">>") {
			r, err := p.parseAdd()
			if err != nil {
				return 0, err
			}
			v >>= uint(r)
		} else {
			break
		}
	}
	return v, nil
}

func (p *formParser) parseAdd() (int64, error) {
	v, err := p.parseMul()
	if err != nil {
		return 0, err
	}
	for {
		if p.accept("+") {
			r, err := p.parseMul()
			if err != nil {
				return 0, err
			}
			v += r
		} else if p.accept("-") {
			r, err := p.parseMul()
			if err != nil {
				return 0, err
			}
			v -= r
		} else {
			break
		}
	}
	return v, nil
}

func (p *formParser) parseMul() (int64, error) {
	v, err := p.parseUnary()
	if err != nil {
		return 0, err
	}
	for {
		if p.accept("*") {
			r, err := p.parseUnary()
			if err != nil {
				return 0, err
			}
			v *= r
		} else if p.accept("/") {
			r, err := p.parseUnary()
			if err != nil {
				return 0, err
			}
			if r == 0 {
				return 0, fmt.Errorf("gige: formula divide by zero")
			}
			v /= r
		} else if p.accept("%") {
			r, err := p.parseUnary()
			if err != nil {
				return 0, err
			}
			if r == 0 {
				return 0, fmt.Errorf("gige: formula mod by zero")
			}
			v %= r
		} else {
			break
		}
	}
	return v, nil
}

func (p *formParser) parseUnary() (int64, error) {
	if p.accept("+") {
		return p.parseUnary()
	}
	if p.accept("-") {
		v, err := p.parseUnary()
		return -v, err
	}
	if p.accept("~") {
		v, err := p.parseUnary()
		return ^v, err
	}
	if p.accept("!") {
		v, err := p.parseUnary()
		if err != nil {
			return 0, err
		}
		if v == 0 {
			return 1, nil
		}
		return 0, nil
	}
	return p.parsePow()
}

// parsePow handles the ** power operator, which binds tighter than unary
// operators and is right-associative.
func (p *formParser) parsePow() (int64, error) {
	v, err := p.parsePrimary()
	if err != nil {
		return 0, err
	}
	if p.accept("**") {
		r, err := p.parseUnary()
		if err != nil {
			return 0, err
		}
		if r < 0 {
			return 0, fmt.Errorf("gige: formula negative exponent %d", r)
		}
		return power(v, r), nil
	}
	return v, nil
}

// power computes base**exp via exponentiation by squaring.
func power(base, exp int64) int64 {
	if exp == 0 {
		return 1
	}
	result := int64(1)
	for exp > 0 {
		if exp&1 == 1 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

func (p *formParser) parsePrimary() (int64, error) {
	p.skipSpace()
	if p.accept("(") {
		v, err := p.parseExpr()
		if err != nil {
			return 0, err
		}
		if !p.accept(")") {
			return 0, fmt.Errorf("gige: formula missing )")
		}
		return v, nil
	}
	if p.i >= len(p.s) {
		return 0, fmt.Errorf("gige: formula unexpected eof")
	}
	// number
	if unicode.IsDigit(rune(p.s[p.i])) || (p.s[p.i] == '0' && p.i+1 < len(p.s) && (p.s[p.i+1] == 'x' || p.s[p.i+1] == 'X')) {
		start := p.i
		if strings.HasPrefix(strings.ToLower(p.s[p.i:]), "0x") {
			p.i += 2
			for p.i < len(p.s) && isHex(p.s[p.i]) {
				p.i++
			}
		} else {
			for p.i < len(p.s) && unicode.IsDigit(rune(p.s[p.i])) {
				p.i++
			}
		}
		return strconv.ParseInt(p.s[start:p.i], 0, 64)
	}
	// identifier or function call
	if unicode.IsLetter(rune(p.s[p.i])) || p.s[p.i] == '_' {
		start := p.i
		for p.i < len(p.s) && (unicode.IsLetter(rune(p.s[p.i])) || unicode.IsDigit(rune(p.s[p.i])) || p.s[p.i] == '_') {
			p.i++
		}
		name := p.s[start:p.i]
		if fn, ok := formulaFuncs[name]; ok {
			if !p.accept("(") {
				return 0, fmt.Errorf("gige: formula missing ( after %s", name)
			}
			var args []int64
			for {
				arg, err := p.parseExpr()
				if err != nil {
					return 0, err
				}
				args = append(args, arg)
				if !p.accept(",") {
					break
				}
			}
			if !p.accept(")") {
				return 0, fmt.Errorf("gige: formula missing ) in %s call", name)
			}
			return fn(args)
		}
		if c, ok := formulaConsts[name]; ok {
			return c, nil
		}
		v, ok := p.vars[name]
		if !ok {
			return 0, fmt.Errorf("gige: formula unknown var %q", name)
		}
		return v, nil
	}
	return 0, fmt.Errorf("gige: formula bad token at %q", p.s[p.i:])
}

func isHex(b byte) bool {
	return unicode.IsDigit(rune(b)) || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
}
