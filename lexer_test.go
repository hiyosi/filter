package filter

import (
	"strconv"
	"strings"
	"testing"
)

func testScanner(t *testing.T, src string, expectTok int) {
	s := new(Scanner)
	s.Init(src)
	tok, lit, _ := s.Scan()
	if tok != expectTok {
		t.Errorf("Expect Scanner{%q}.Scan() = %#v, _ want %#v", src, tok, expectTok)
	}

	switch lit.(type) {
	case string:
		if lit != strings.Replace(src, "\"", "", -1) {
			t.Errorf("Expect Scanner{%q}.Scan() = _, %#v want %#v", src, lit, src)
		}
	case int:
		s, _ := strconv.Atoi(src)
		if lit != s {
			t.Errorf("Expect Scanner{%q}.Scan() = _, %#v want %#v", src, lit, src)
		}
	case bool:
		s, _ := strconv.ParseBool(src)
		if lit != s {
			t.Errorf("Expect Scanner{%q}.Scan() = _, %#v want %#v", src, lit, src)
		}
	default:
		t.Errorf("Expect Scanner{%q}.Scan() = _, %#v want %#v", src, lit, src)
	}

	tok, lit, _ = s.Scan()
	if tok != EOF {
		t.Errorf("Expect Scanner{%q}.Scan() = %#v, _ want %#v", src, tok, EOF)
	}
}

func TestScanner(t *testing.T) {
	testScanner(t, "ham", IDENT)
	testScanner(t, "ham123", IDENT)
	testScanner(t, "ham-123", IDENT)
	testScanner(t, "ham_123", IDENT)
	testScanner(t, "ham.123", IDENT)
	testScanner(t, "ham:123", IDENT)
	testScanner(t, "ham//123", IDENT)
	testScanner(t, "\"ham\"", VALUE)
	testScanner(t, "\"ham spam\"", VALUE)
	testScanner(t, "\"ham@spam\"", VALUE)
	testScanner(t, "\"ham (spam)\"", VALUE)
	testScanner(t, "123", NUMBER)
	testScanner(t, "123", NUMBER)
	testScanner(t, "true", TRUE)
	testScanner(t, "false", FALSE)
	testScanner(t, "(", LPAREN)
	testScanner(t, ")", RPAREN)
	testScanner(t, "[", LBOXP)
	testScanner(t, "]", RBOXP)

}
