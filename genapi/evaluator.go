package genapi

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

var formulaFuncs = map[string]func(int64) int64{
	"ABS": func(v int64) int64 {
		if v < 0 {
			return -v
		}
		return v
	},
	"FLOOR": func(v int64) int64 { return v },
	"CEIL":  func(v int64) int64 { return v },
	"SQRT": func(v int64) int64 {
		if v < 0 {
			return 0
		}
		return int64(math.Sqrt(float64(v)))
	},
}

// evalFormula evaluates a GenICam SwissKnife formula with integer variables.
// Supports + - * / % & | ^ << >> ~ ( ) ? : = <> != < > <= >= && || and hex/decimal literals.
// Supports functions: ABS, FLOOR, CEIL, SQRT.
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
		case p.accept("<>"), p.accept("!="):
			op = "!="
		case p.accept("="):
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
	return p.parsePrimary()
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
			arg, err := p.parseExpr()
			if err != nil {
				return 0, err
			}
			if !p.accept(")") {
				return 0, fmt.Errorf("gige: formula missing ) in %s call", name)
			}
			return fn(arg), nil
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
