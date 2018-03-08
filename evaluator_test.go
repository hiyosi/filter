package filter

import (
	"labix.org/v2/mgo/bson"
	"reflect"
	"testing"
)

func TestEvaluate(t *testing.T) {
	{ // ham eq "spam"
		actual, _ := Evaluate(&CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "eq",
			RHE:      &AttrValueExpression{"spam"},
		}, Env{})

		expect := make(map[string]interface{})
		expect["ham"] = "spam"

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}
	{ // ham eq 123
		actual, _ := Evaluate(&CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "eq",
			RHE:      &NumberExpression{123},
		}, Env{})

		expect := make(map[string]interface{})
		expect["ham"] = 123

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}

		av, _ := actual.(map[string]interface{})
		at := reflect.TypeOf(av["ham"]).Kind()
		et := reflect.Int

		if at != et {
			t.Errorf("actual=%s, expect=%s", at, et)
		}
	}
	{ // ham eq true
		actual, _ := Evaluate(&CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "eq",
			RHE:      &BoolExpression{true},
		}, Env{})

		expect := make(map[string]interface{})
		expect["ham"] = true

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}

		av, _ := actual.(map[string]interface{})
		at := reflect.TypeOf(av["ham"]).Kind()
		et := reflect.Bool

		if at != et {
			t.Errorf("actual=%s, expect=%s", at, et)
		}
	}
	{ // ham eq null
		actual, _ := Evaluate(&CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "eq",
			RHE:      &IdentifierExpression{"null"},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$type"] = 10
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}

		av, _ := actual.(map[string]interface{})
		tv, _ := av["ham"].(map[string]interface{})

		if tv["$type"] != 10 {
			t.Errorf("actual=%v, expect=%v", tv["$type"], 10)
		}
	}

	{ // ham ne "spam"
		actual, _ := Evaluate(&CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "ne",
			RHE:      &AttrValueExpression{"spam"},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$ne"] = "spam"
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}
	{ // ham ne 123
		actual, _ := Evaluate(&CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "ne",
			RHE:      &NumberExpression{123},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$ne"] = 123
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}

		av, _ := actual.(map[string]interface{})
		sv, _ := av["ham"].(map[string]interface{})
		at := reflect.TypeOf(sv["$ne"]).Kind()
		et := reflect.Int

		if at != et {
			t.Errorf("actual=%s, expect=%s", at, et)
		}
	}
	{ // ham ne true
		actual, _ := Evaluate(&CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "ne",
			RHE:      &BoolExpression{true},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$ne"] = true
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}

		av, _ := actual.(map[string]interface{})
		sv, _ := av["ham"].(map[string]interface{})
		at := reflect.TypeOf(sv["$ne"]).Kind()
		et := reflect.Bool

		if at != et {
			t.Errorf("actual=%s, expect=%s", at, et)
		}
	}
	{ // ham ne null
		actual, _ := Evaluate(&CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "ne",
			RHE:      &IdentifierExpression{"null"},
		}, Env{})

		s := make(map[string]interface{})
		s["$type"] = 10
		sub := make(map[string]interface{})
		sub["$ne"] = s
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}

		av, _ := actual.(map[string]interface{})
		sv, _ := av["ham"].(map[string]interface{})
		tv, _ := sv["$ne"].(map[string]interface{})

		if tv["$type"] != 10 {
			t.Errorf("actual=%v, expect=%v", tv["$type"], 10)
		}
	}

	{ //ham co "spam"
		actual, _ := Evaluate(&RegexStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "co",
			Value:    "spam",
		}, Env{})

		sub := make(map[string]interface{})
		sub["$regex"] = bson.RegEx{Pattern: "spam"}
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //ham sw "spam"
		actual, _ := Evaluate(&RegexStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "sw",
			Value:    "spam",
		}, Env{})

		sub := make(map[string]interface{})
		sub["$regex"] = bson.RegEx{Pattern: "^spam"}
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //ham sw "spam"
		actual, _ := Evaluate(&RegexStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "ew",
			Value:    "spam",
		}, Env{})

		sub := make(map[string]interface{})
		sub["$regex"] = bson.RegEx{Pattern: "spam$"}
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ // ham pr
		actual, _ := Evaluate(&AttrStatement{
			Attr:     &IdentifierExpression{"ham"},
			Operator: "pr",
		}, Env{})

		sub := map[string]bool{"$exists": true}
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //ham gt 123
		actual, _ := Evaluate(&CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "gt",
			RHE:      &NumberExpression{123},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$gt"] = 123
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //ham ge 123
		actual, _ := Evaluate(&CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "ge",
			RHE:      &NumberExpression{123},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$gte"] = 123
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //ham lt 123
		actual, _ := Evaluate(&CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "lt",
			RHE:      &NumberExpression{123},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$lt"] = 123
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //ham le 123
		actual, _ := Evaluate(&CompStatement{
			LHE:      &IdentifierExpression{"ham"},
			Operator: "le",
			RHE:      &NumberExpression{123},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$lte"] = 123
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ // ham eq "spam" and abc eq 123
		actual, _ := Evaluate(&LogStatement{
			LHS: &CompStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "eq",
				RHE:      &AttrValueExpression{"spam"},
			},
			Operator: "and",
			RHS: &CompStatement{
				LHE:      &IdentifierExpression{"abc"},
				Operator: "eq",
				RHE:      &NumberExpression{123},
			},
		}, Env{})

		l := make(map[string]interface{})
		l["ham"] = "spam"
		r := make(map[string]interface{})
		r["abc"] = 123
		expect := make(map[string]interface{})
		var sub []interface{}
		sub = append(sub, l, r)
		expect["$and"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ // ham eq "spam" or abc eq 123
		actual, _ := Evaluate(&LogStatement{
			LHS: &CompStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "eq",
				RHE:      &AttrValueExpression{"spam"},
			},
			Operator: "or",
			RHS: &CompStatement{
				LHE:      &IdentifierExpression{"abc"},
				Operator: "eq",
				RHE:      &NumberExpression{123},
			},
		}, Env{})

		l := make(map[string]interface{})
		l["ham"] = "spam"
		r := make(map[string]interface{})
		r["abc"] = 123
		expect := make(map[string]interface{})
		var sub []interface{}
		sub = append(sub, l, r)
		expect["$or"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ // (ham eq "spam")
		actual, _ := Evaluate(&ParenStatement{
			Operator: "",
			SubStatement: &CompStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "eq",
				RHE:      &AttrValueExpression{"spam"},
			},
		}, Env{})

		expect := make(map[string]interface{})
		expect["ham"] = "spam"

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}

	}

	{ // ham[ foo eq "spam"]
		actual, _ := Evaluate(&GroupStatement{
			ParentAttr: &IdentifierExpression{"ham"},
			SubStatement: &CompStatement{
				LHE:      &IdentifierExpression{"foo"},
				Operator: "eq",
				RHE:      &AttrValueExpression{"spam"},
			},
		}, Env{})

		expect := make(map[string]interface{})
		expect["ham.foo"] = "spam"

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ // (ham eq "spam") and (foo co "bar")
		actual, _ := Evaluate(&LogStatement{
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
		}, Env{})

		l := make(map[string]interface{})
		l["ham"] = "spam"
		rsub := make(map[string]interface{})
		rsub["$regex"] = bson.RegEx{Pattern: "bar"}
		r := make(map[string]interface{})
		r["foo"] = rsub
		expect := make(map[string]interface{})
		var sub []interface{}
		sub = append(sub, l, r)
		expect["$and"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}

	}

	{ // not (ham eq "spam")  * this statement is evaluate to 'not equal, (equal <-> not equal)'
		actual, _ := Evaluate(&ParenStatement{
			Operator: "not",
			SubStatement: &CompStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "eq",
				RHE:      &AttrValueExpression{"spam"},
			},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$ne"] = "spam"
		expect := make(map[string]interface{})
		expect["ham"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //not (ham ne "spam") * this statement is evaluate to 'equal', (not equal <-> equal)"
		actual, _ := Evaluate(&ParenStatement{
			Operator: "not",
			SubStatement: &CompStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "ne",
				RHE:      &AttrValueExpression{"spam"},
			},
		}, Env{})

		expect := make(map[string]interface{})
		expect["ham"] = "spam"

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //not (ham co "spam")
		actual, _ := Evaluate(&ParenStatement{
			Operator: "not",
			SubStatement: &RegexStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "co",
				Value:    "spam",
			},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$regex"] = bson.RegEx{Pattern: "spam"}
		n := make(map[string]interface{})
		n["$not"] = sub
		expect := make(map[string]interface{})
		expect["ham"] = n

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //not (ham sw "spam")
		actual, _ := Evaluate(&ParenStatement{
			Operator: "not",
			SubStatement: &RegexStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "sw",
				Value:    "spam",
			},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$regex"] = bson.RegEx{Pattern: "^spam"}
		n := make(map[string]interface{})
		n["$not"] = sub
		expect := make(map[string]interface{})
		expect["ham"] = n

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //not (ham ew "spam")
		actual, _ := Evaluate(&ParenStatement{
			Operator: "not",
			SubStatement: &RegexStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "ew",
				Value:    "spam",
			},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$regex"] = bson.RegEx{Pattern: "spam$"}
		n := make(map[string]interface{})
		n["$not"] = sub
		expect := make(map[string]interface{})
		expect["ham"] = n

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //not (ham gt 123)
		actual, _ := Evaluate(&ParenStatement{
			Operator: "not",
			SubStatement: &CompStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "gt",
				RHE:      &NumberExpression{123},
			},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$gt"] = 123
		n := make(map[string]interface{})
		n["$not"] = sub
		expect := make(map[string]interface{})
		expect["ham"] = n

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //not (ham ge 123)
		actual, _ := Evaluate(&ParenStatement{
			Operator: "not",
			SubStatement: &CompStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "ge",
				RHE:      &NumberExpression{123},
			},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$gte"] = 123
		n := make(map[string]interface{})
		n["$not"] = sub
		expect := make(map[string]interface{})
		expect["ham"] = n

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //not (ham lt 123)
		actual, _ := Evaluate(&ParenStatement{
			Operator: "not",
			SubStatement: &CompStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "lt",
				RHE:      &NumberExpression{123},
			},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$lt"] = 123
		n := make(map[string]interface{})
		n["$not"] = sub
		expect := make(map[string]interface{})
		expect["ham"] = n

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ //not (ham le 123)
		actual, _ := Evaluate(&ParenStatement{
			Operator: "not",
			SubStatement: &CompStatement{
				LHE:      &IdentifierExpression{"ham"},
				Operator: "le",
				RHE:      &NumberExpression{123},
			},
		}, Env{})

		sub := make(map[string]interface{})
		sub["$lte"] = 123
		n := make(map[string]interface{})
		n["$not"] = sub
		expect := make(map[string]interface{})
		expect["ham"] = n

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

	{ // (ham eq "spam") and not (foo co "bar")
		actual, _ := Evaluate(&LogStatement{
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
				Operator: "not",
				SubStatement: &RegexStatement{
					LHE:      &IdentifierExpression{"foo"},
					Operator: "co",
					Value:    "bar",
				},
			},
		}, Env{})

		l := make(map[string]interface{})
		l["ham"] = "spam"
		rsub := make(map[string]interface{})
		rsub["$regex"] = bson.RegEx{Pattern: "bar"}
		n := make(map[string]interface{})
		n["$not"] = rsub
		r := make(map[string]interface{})
		r["foo"] = n
		expect := make(map[string]interface{})
		var sub []interface{}
		sub = append(sub, l, r)
		expect["$and"] = sub

		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual=%v, expect=%v", actual, expect)
		}
	}

}
