package filter

import (
	"reflect"
	"testing"
)

func parseStmt(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init(src)
	statements := Parse(s)
	if len(statements) != 1 {
		t.Errorf("Expect %q to be parsed", src)
		return
	}
	if !reflect.DeepEqual(statements[0], expect) {
		t.Errorf("Expect %+#v to be %+#v", statements[0], expect)
		return
	}
}

func parseErrorStmt(t *testing.T, src string) {
	defer func() {
		if error := recover(); error != nil {
			return
		} else {
			t.Errorf("Expect %q to be panic", src)
			return
		}
	}()

	s := new(Scanner)
	s.Init(src)
	Parse(s)
}

func TestParseStatement(t *testing.T) {
	//equal
	parseStmt(t, "ham eq \"spam\"", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "eq",
		RHE:      &AttrValueExpression{"spam"},
	})
	parseStmt(t, "ham eq \"spam@example.com\"", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "eq",
		RHE:      &AttrValueExpression{"spam@example.com"},
	})
	parseStmt(t, "ham eq \"山田 太郎\"", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "eq",
		RHE:      &AttrValueExpression{"山田 太郎"},
	})
	parseStmt(t, "ham eq null", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "eq",
		RHE:      &IdentifierExpression{"null"},
	})
	parseStmt(t, "ham123 eq \"spam\"", &CompStatement{
		LHE:      &IdentifierExpression{"ham123"},
		Operator: "eq",
		RHE:      &AttrValueExpression{"spam"},
	})
	parseStmt(t, "ham-123 eq \"spam\"", &CompStatement{
		LHE:      &IdentifierExpression{"ham-123"},
		Operator: "eq",
		RHE:      &AttrValueExpression{"spam"},
	})
	parseStmt(t, "ham_123 eq \"spam\"", &CompStatement{
		LHE:      &IdentifierExpression{"ham_123"},
		Operator: "eq",
		RHE:      &AttrValueExpression{"spam"},
	})
	parseStmt(t, "ham eq 123", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "eq",
		RHE:      &NumberExpression{123},
	})
	parseStmt(t, "ham eq true", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "eq",
		RHE:      &BoolExpression{true},
	})
	parseStmt(t, "ham eq false", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "eq",
		RHE:      &BoolExpression{false},
	})

	//not-equal
	parseStmt(t, "ham ne \"spam\"", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "ne",
		RHE:      &AttrValueExpression{"spam"},
	})
	parseStmt(t, "ham ne 123", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "ne",
		RHE:      &NumberExpression{123},
	})
	parseStmt(t, "ham ne true", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "ne",
		RHE:      &BoolExpression{true},
	})
	parseStmt(t, "ham ne false", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "ne",
		RHE:      &BoolExpression{false},
	})
	parseStmt(t, "ham ne null", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "ne",
		RHE:      &IdentifierExpression{"null"},
	})

	// contains
	parseStmt(t, "ham co \"spam\"", &RegexStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "co",
		Value:    "spam"},
	)

	// starts with
	parseStmt(t, "ham sw \"spam\"", &RegexStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "sw",
		Value:    "spam"},
	)

	// ends with
	parseStmt(t, "ham ew \"spam\"", &RegexStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "ew",
		Value:    "spam"},
	)

	// present (has value)
	parseStmt(t, "ham pr", &AttrStatement{
		Attr:     &IdentifierExpression{"ham"},
		Operator: "pr",
	})

	// greater than
	parseStmt(t, "ham gt \"spam\"", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "gt",
		RHE:      &AttrValueExpression{"spam"},
	})
	parseStmt(t, "ham gt 123", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "gt",
		RHE:      &NumberExpression{123},
	})
	parseStmt(t, "ham gt true", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "gt",
		RHE:      &BoolExpression{true},
	})
	parseStmt(t, "ham gt false", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "gt",
		RHE:      &BoolExpression{false},
	})
	parseStmt(t, "ham gt null", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "gt",
		RHE:      &IdentifierExpression{"null"},
	})

	// reater than equal
	parseStmt(t, "ham ge \"spam\"", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "ge",
		RHE:      &AttrValueExpression{"spam"},
	})
	parseStmt(t, "ham ge 123", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "ge",
		RHE:      &NumberExpression{123},
	})
	parseStmt(t, "ham ge true", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "ge",
		RHE:      &BoolExpression{true},
	})
	parseStmt(t, "ham ge false", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "ge",
		RHE:      &BoolExpression{false},
	})
	parseStmt(t, "ham ge null", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "ge",
		RHE:      &IdentifierExpression{"null"},
	})

	// less than
	parseStmt(t, "ham lt \"spam\"", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "lt",
		RHE:      &AttrValueExpression{"spam"},
	})
	parseStmt(t, "ham lt 123", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "lt",
		RHE:      &NumberExpression{123},
	})
	parseStmt(t, "ham lt true", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "lt",
		RHE:      &BoolExpression{true},
	})
	parseStmt(t, "ham lt false", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "lt",
		RHE:      &BoolExpression{false},
	})
	parseStmt(t, "ham lt null", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "lt",
		RHE:      &IdentifierExpression{"null"},
	})

	// less than equal
	parseStmt(t, "ham le \"spam\"", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "le",
		RHE:      &AttrValueExpression{"spam"},
	})
	parseStmt(t, "ham le 123", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "le",
		RHE:      &NumberExpression{123},
	})
	parseStmt(t, "ham le true", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "le",
		RHE:      &BoolExpression{true},
	})
	parseStmt(t, "ham le false", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "le",
		RHE:      &BoolExpression{false},
	})
	parseStmt(t, "ham le null", &CompStatement{
		LHE:      &IdentifierExpression{"ham"},
		Operator: "le",
		RHE:      &IdentifierExpression{"null"},
	})

	// Logical And
	parseStmt(t, "ham pr and ham eq \"spam\"", &LogStatement{
		LHS: &AttrStatement{
			Attr:     &IdentifierExpression{"ham"},
			Operator: "pr",
		},
		Operator: "and",
		RHS: &CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"spam"},
		},
	})
	parseStmt(t, "ham eq \"spam\" and foo eq \"bar\"", &LogStatement{
		LHS: &CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"spam"},
		},
		Operator: "and",
		RHS: &CompStatement{
			LHE:      &IdentifierExpression{"foo"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"bar"},
		},
	})
	parseStmt(t, "ham co \"spam\" and foo eq \"bar\"", &LogStatement{
		LHS: &RegexStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "co",
			Value:    "spam",
		},
		Operator: "and",
		RHS: &CompStatement{
			LHE:      &IdentifierExpression{"foo"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"bar"},
		},
	})

	// Logical or
	parseStmt(t, "ham pr or ham eq \"spam\"", &LogStatement{
		LHS: &AttrStatement{
			Attr:     &IdentifierExpression{"ham"},
			Operator: "pr",
		},
		Operator: "or",
		RHS: &CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"spam"},
		},
	})
	parseStmt(t, "ham eq \"spam\" or foo eq \"bar\"", &LogStatement{
		LHS: &CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"spam"},
		},
		Operator: "or",
		RHS: &CompStatement{
			LHE:      &IdentifierExpression{"foo"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"bar"},
		},
	})
	parseStmt(t, "ham co \"spam\" or foo eq \"bar\"", &LogStatement{
		LHS: &RegexStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "co",
			Value:    "spam",
		},
		Operator: "or",
		RHS: &CompStatement{
			LHE:      &IdentifierExpression{"foo"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"bar"},
		},
	})

	// Precedence grouping
	parseStmt(t, "(ham eq \"spam\")", &ParenStatement{
		Operator: "",
		SubStatement: &CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"spam"},
		},
	})
	parseStmt(t, "( ham co \"spam\" )", &ParenStatement{
		Operator: "",
		SubStatement: &RegexStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "co",
			Value:    "spam",
		},
	})

	parseStmt(t, "(ham eq \"spam\") and (foo co \"bar\")", &LogStatement{
		LHS: &ParenStatement{
			Operator: "",
			SubStatement: &CompStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "eq",
				RHE:      &AttrValueExpression{"spam"},
			},
		},
		Operator: "and",
		RHS: &ParenStatement{
			Operator: "",
			SubStatement: &RegexStatement{
				LHE:      &IdentifierExpression{"foo"},
				Operator: "co",
				Value:    "bar",
			},
		},
	})

	// Complex  attribute filter grouping
	parseStmt(t, "foo[ham eq \"spam\"]", &GroupStatement{
		ParentAttr: &IdentifierExpression{"foo"},
		SubStatement: &CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"spam"},
		},
	})

	// not function
	parseStmt(t, "not (ham eq \"spam\")", &ParenStatement{
		Operator: "not",
		SubStatement: &CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"spam"},
		},
	})
	parseStmt(t, "not ( ham co \"spam\")", &ParenStatement{
		Operator: "not",
		SubStatement: &RegexStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "co",
			Value:    "spam",
		},
	})
	parseStmt(t, "not ( foo[ham eq \"spam\"] )", &ParenStatement{
		Operator: "not",
		SubStatement: &GroupStatement{
			ParentAttr: &IdentifierExpression{"foo"},
			SubStatement: &CompStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "eq",
				RHE:      &AttrValueExpression{"spam"},
			},
		},
	})

	// associative
	parseStmt(t, "ham eq \"spam\" and foo eq \"bar\" and baz eq \"qux\"", &LogStatement{
		LHS: &LogStatement{
			LHS: &CompStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "eq",
				RHE:      &AttrValueExpression{"spam"},
			},
			Operator: "and",
			RHS: &CompStatement{
				LHE:      &IdentifierExpression{"foo"},
				Operator: "eq",
				RHE:      &AttrValueExpression{"bar"},
			},
		},
		Operator: "and",
		RHS: &CompStatement{
			LHE:      &IdentifierExpression{"baz"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"qux"},
		},
	})
	parseStmt(t, "ham eq \"spam\" or foo eq \"bar\" and baz eq \"qux\"", &LogStatement{
		LHS: &LogStatement{
			LHS: &CompStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "eq",
				RHE:      &AttrValueExpression{"spam"},
			},
			Operator: "or",
			RHS: &CompStatement{
				LHE:      &IdentifierExpression{"foo"},
				Operator: "eq",
				RHE:      &AttrValueExpression{"bar"},
			},
		},
		Operator: "and",
		RHS: &CompStatement{
			LHE:      &IdentifierExpression{"baz"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"qux"},
		},
	})

	parseErrorStmt(t, "foo bar baz")
	parseErrorStmt(t, "ham eq spam")
	parseErrorStmt(t, "123 eq \"spam\"")
	parseErrorStmt(t, "123ham eq \"spam\"")
	parseErrorStmt(t, "ham@123 eq \"spam\"")
	parseErrorStmt(t, "ham co 123")
	parseErrorStmt(t, "ham co true")
	parseErrorStmt(t, "ham co false")
	parseErrorStmt(t, "123 co 123")
	parseErrorStmt(t, "ham sw 123")
	parseErrorStmt(t, "ham sw true")
	parseErrorStmt(t, "ham sw false")
	parseErrorStmt(t, "123 sw 123")
	parseErrorStmt(t, "ham ew 123")
	parseErrorStmt(t, "ham ew true")
	parseErrorStmt(t, "ham ew false")
	parseErrorStmt(t, "123 ew 123")
	parseErrorStmt(t, "123 gt 123")
	parseErrorStmt(t, "123 ge 123")
	parseErrorStmt(t, "123 lt 123")
	parseErrorStmt(t, "123 le 123")
	parseErrorStmt(t, "ham and 123")
	parseErrorStmt(t, "ham and (foo eq \"spam\")")
	parseErrorStmt(t, "ham or 123")
	parseErrorStmt(t, "ham or (foo eq \"spam\")")
	parseErrorStmt(t, "(foo)")
	parseErrorStmt(t, "(foo bar baz)")
	parseErrorStmt(t, "(foo bar baz]")

}
