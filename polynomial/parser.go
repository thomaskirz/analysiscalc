package polynomial

import (
	"fmt"
	"io"
	"strconv"
)

type InputStatement struct {
	Name     string
	Function Polynomial
}

// Parser represents a parser
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max = 1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead
func (p *Parser) scan() (tok Token, lit string) {
	// If the buffer holds a token, the return it
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan it later
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back into the buffer
func (p *Parser) unscan() {
	p.buf.n = 1
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (Token, string) {
	tok, lit := p.scan()
	if tok == WHITESPACE {
		tok, lit = p.scan()
	}
	return tok, lit
}

func (p *Parser) Parse() (*InputStatement, error) {
	stmt := &InputStatement{}

	if tok, lit := p.scanIgnoreWhitespace(); tok == NAME {
		stmt.Name = string([]rune(lit)[0])
	} else {
		return nil, fmt.Errorf("found %q, expected function name", lit)
	}

	// Expecting EOF or function definition
	if tok, lit := p.scanIgnoreWhitespace(); tok == EOF {
		return stmt, nil
	} else if tok != EQUALS {
		return nil, fmt.Errorf("found, %q, expected =", lit)
	}

	stmt.Function = make(Polynomial)

	sign := true // true for positive, false for negative
	if tok, lit := p.scanIgnoreWhitespace(); tok == MINUS {
		sign = false
	} else if tok != PLUS && tok != INTEGER && tok != FLOAT {
		return nil, fmt.Errorf("found %q, expected function declaration", lit)
	} else {
		p.unscan()
	}

	for {
		var (
			coeff float64
			exp   uint
		)

		// Expecting coefficient or 'x'
		if tok, lit := p.scanIgnoreWhitespace(); tok == FLOAT || tok == INTEGER {
			if flt, err := strconv.ParseFloat(lit, 64); err != nil {
				return nil, fmt.Errorf("problem parsing float %q", lit)
			} else {
				coeff = flt
				if !sign {
					coeff *= -1
				}
			}
		} else if tok == VAR {
			coeff = 1
			p.unscan()
		} else {
			return nil, fmt.Errorf("found %q, expected polynomial term", lit)
		}

		if tok, _ := p.scanIgnoreWhitespace(); tok == VAR {
			// Expecting either exponent or next term
			if tok, lit := p.scanIgnoreWhitespace(); tok != CARET {
				exp = 1
				p.unscan()
			} else {
				if tok, lit = p.scanIgnoreWhitespace(); tok == INTEGER {
					if i, err := strconv.Atoi(lit); err != nil {
						return nil, fmt.Errorf("problem parsing integer %q", lit)
					} else {
						exp = uint(i)
					}
				} else {
					return nil, fmt.Errorf("found %q, expected integer", lit)
				}
			}
		} else {
			exp = 0
			p.unscan()
		}

		stmt.Function[exp] += coeff

		if tok, lit := p.scanIgnoreWhitespace(); tok == PLUS {
			sign = true
		} else if tok == MINUS {
			sign = false
		} else if tok == EOF {
			break
		} else {
			return nil, fmt.Errorf("found %q, expected polynomial term", lit)
		}
	}

	return stmt, nil
}
