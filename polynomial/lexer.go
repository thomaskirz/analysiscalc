package polynomial

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

// Token represents a lexical token
type Token int

const (
	ILLEGAL Token = iota
	EOF
	WHITESPACE

	PLUS   // +
	MINUS  // -
	CARET  // ^
	EQUALS // =
	VAR    // x

	NAME // function name, like 'f(x)'

	INTEGER
	FLOAT
)

var eof = rune(0)

func isWhitespace(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func isLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{bufio.NewReader(r)}
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	char, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return char
}

// unread places the previously read rune back on the reader
func (s *Scanner) unread() {
	s.r.UnreadRune()
}

// Scan returns the next token and literal value
func (s *Scanner) Scan() (Token, string) {
	// Read the next rune
	char := s.read()

	// If see whitespace then consume all contiguous whitespace
	// If we see letter then consume as name
	// If we see digit then consume as number
	if isWhitespace(char) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(char) {
		s.unread()
		return s.scanLetters()
	} else if unicode.IsDigit(char) {
		s.unread()
		return s.scanNumber()
	}

	// Otherwise read the individual character.
	switch char {
	case eof:
		return EOF, ""
	case '+':
		return PLUS, string(char)
	case '-':
		return MINUS, string(char)
	case '^':
		return CARET, string(char)
	case '=':
		return EQUALS, string(char)
	}

	return ILLEGAL, string(char)
}
func (s *Scanner) scanWhitespace() (Token, string) {
	// Create a builder and read the current character into it
	var builder strings.Builder
	builder.WriteRune(s.read())

	// Read every subsequent whitespace character into the builder
	// Non-whitespace characters and EOF will cause the loop to exit
	for {
		if char := s.read(); char == eof {
			break
		} else if !isWhitespace(char) {
			s.unread()
			break
		} else {
			builder.WriteRune(char)
		}
	}

	return WHITESPACE, builder.String()
}
func (s *Scanner) scanLetters() (Token, string) {
	// Create a builder and read the current character into it
	var builder strings.Builder
	first := s.read()
	builder.WriteRune(first)

	// Check if token is either NAME (like 'f(x)', f being some letter) or VAR ('x').
	// Otherwise return ILLEGAL token
	if char := s.read(); char == '(' {
		builder.WriteRune(char)
	} else {
		s.unread()
		if first == 'x' {
			return VAR, builder.String()
		} else {
			return ILLEGAL, builder.String()
		}
	}

	if char := s.read(); char == 'x' {
		builder.WriteRune(char)
	} else {
		return ILLEGAL, builder.String()
	}

	if char := s.read(); char == ')' {
		builder.WriteRune(char)
	} else {
		return ILLEGAL, builder.String()
	}

	return NAME, builder.String()
}

func (s *Scanner) scanNumber() (Token, string) {
	// Create a builder and read the current character into it
	var builder strings.Builder
	builder.WriteRune(s.read())

	for {
		if char := s.read(); !unicode.IsDigit(char) {
			s.unread()
			break
		} else {
			builder.WriteRune(char)
		}
	}

	if char := s.read(); char == '.' {
		builder.WriteRune(char)
		for {
			if char = s.read(); char == eof {
				break
			} else if !unicode.IsDigit(char) {
				s.unread()
				break
			} else {
				builder.WriteRune(char)
			}
		}
		return FLOAT, builder.String()

	} else if char != eof {
		s.unread()
	}

	return INTEGER, builder.String()
}
