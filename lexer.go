package filter

import (
	"fmt"
	"strconv"
	"unicode"
)

const (
	EOF     = -1
	UNKNOWN = 0
)

var keywords = map[string]int{
	"not":   NOT,
	"and":   AND,
	"or":    OR,
	"pr":    PR,
	"eq":    EQ,
	"ne":    NE,
	"co":    CO,
	"sw":    SW,
	"ew":    EW,
	"gt":    GT,
	"ge":    GE,
	"lt":    LT,
	"le":    LE,
	"(":     LPAREN,
	")":     RPAREN,
	"[":     LBOXP,
	"]":     RBOXP,
	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
}

type Position struct {
	Line   int
	Column int
}

type Scanner struct {
	src      []rune
	offset   int
	lineHead int
	line     int
}

func (s *Scanner) Init(src string) {
	s.src = []rune(src)
}

func (s *Scanner) Scan() (tok int, lit interface{}, pos Position) {
	s.skipWhiteSpace()
	pos = s.position()

	switch ch := s.peek(); {
	case unicode.IsLetter(ch) || ch == '"':
		if ch == '"' {
			tok, lit = VALUE, s.scanAttrValue()
		} else {
			lit = s.scanIdentifier()
			if keyword, ok := keywords[lit.(string)]; ok {
				tok = keyword
			} else {
				tok = IDENT
			}
		}
	case isDigit(ch):
		i, _ := strconv.Atoi(s.scanNumber())
		tok, lit = NUMBER, i
	case ch == '(':
		tok = LPAREN
		lit = "("
		s.next()
	case ch == ')':
		tok = RPAREN
		lit = ")"
		s.next()
	case ch == '[':
		tok = LBOXP
		lit = "["
		s.next()
	case ch == ']':
		tok = RBOXP
		lit = "]"
		s.next()
	case ch == -1:
		tok = EOF
		s.next()
	default:
		panic("scan error, " + fmt.Sprintf("Line %d, Column %d: %q", pos.Line, pos.Column, ch))
	}
	return
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsPunct(ch) || unicode.IsSymbol(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func (s *Scanner) peek() rune {
	if !s.reachEOF() {
		return s.src[s.offset]
	} else {
		return -1
	}
}

func (s *Scanner) next() {
	if !s.reachEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

func (s *Scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

func (s *Scanner) position() Position {
	return Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

func (s *Scanner) skipWhiteSpace() {
	for isWhiteSpace(s.peek()) {
		s.next()
	}
}

func (s *Scanner) scanAttrValue() string {
	var ret []rune
	s.next()
	for isLetter(s.peek()) || isDigit(s.peek()) || s.peek() == ' ' {
		ret = append(ret, s.peek())
		s.next()
		if ret[len(ret)-1] != '\\' && s.peek() == '"' {
			s.next()
			break
		}

	}
	return string(ret)
}

func (s *Scanner) scanIdentifier() string {
	var ret []rune
	begin := s.peek()

	switch begin {
	default:
		for unicode.IsLetter(s.peek()) || isDigit(s.peek()) || s.peek() == '-' || s.peek() == '_' || s.peek() == '.' || s.peek() == ':' || s.peek() == '/' {
			ret = append(ret, s.peek())
			s.next()
		}
	}
	return string(ret)
}

func (s *Scanner) scanNumber() string {
	var ret []rune
	for isDigit(s.peek()) {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret)
}
