package filter

import (
	"errors"
	"fmt"
	"labix.org/v2/mgo/bson"
	"time"
)

type Env map[string]string

func Evaluate(statement Statement, env Env) (interface{}, error) {
	notOp := false
	if env["operator"] == "not" {
		env["operator"] = ""
		notOp = true
	}

	switch stmt := statement.(type) {
	case *LogStatement:
		if notOp {
			env["operator"] = "not"
		}
		lv, err := Evaluate(stmt.LHS, env)
		if err != nil {
			return nil, err
		}
		if notOp {
			env["operator"] = "not"
		}
		rv, err := Evaluate(stmt.RHS, env)
		if err != nil {
			return nil, err
		}

		var array []interface{}
		array = append(array, lv, rv)
		q := make(map[string]interface{})

		switch stmt.Operator {
		case "and":
			q["$and"] = array
			return q, nil
		case "or":
			q["$or"] = array
			return q, nil
		default:
			return "", err
		}
	case *AttrStatement:
		q := make(map[string]interface{})

		attr, err := evaluateExpr(stmt.Attr, env)
		if err != nil {
			return nil, err
		}

		switch stmt.Operator {
		case "pr":
			if notOp {
				m := make(map[string]interface{})
				m1 := map[string]bool{"$exists": true}
				m["$not"] = m1
				q[attr.(string)] = m
			} else {
				m := map[string]bool{"$exists": true}
				q[attr.(string)] = m
			}
			return q, nil
		default:
			return "", nil
		}
	case *CompStatement:
		var r interface{}
		l, err := evaluateExpr(stmt.LHE, env)
		if err != nil {
			return nil, err
		}
		if env["parentAttr"] != "" {
			l = env["parentAttr"] + "." + l.(string)
		}

		rhs, err := evaluateExpr(stmt.RHE, env)
		if err != nil {
			return nil, err
		}

		t := make(map[string]interface{})
		if rhs == "null" {
			t["$type"] = 10
			r = t
		} else {
			r = rhs
		}

		m := make(map[string]interface{})
		q := make(map[string]interface{})

		switch stmt.Operator {
		case "eq":
			if notOp {
				m["$ne"] = r
				q[l.(string)] = m
				return q, nil
			} else {
				q[l.(string)] = r
				return q, nil
			}
		case "ne":
			if notOp {
				q[l.(string)] = r
				return q, nil
			} else {
				m["$ne"] = r
				q[l.(string)] = m
				return q, nil
			}
		case "gt":
			s, ok := r.(string)
			if ok {
				r = timeOrString(s)
			}

			if notOp {
				m1 := make(map[string]interface{})
				m1["$gt"] = r
				m["$not"] = m1
			} else {
				m["$gt"] = r
			}
			q[l.(string)] = m
			return q, nil
		case "ge":
			s, ok := rhs.(string)
			if ok {
				r = timeOrString(s)
			}

			if notOp {
				m1 := make(map[string]interface{})
				m1["$gte"] = r
				m["$not"] = m1
			} else {
				m["$gte"] = r
			}
			q[l.(string)] = m
			return q, nil
		case "lt":
			s, ok := rhs.(string)
			if ok {
				r = timeOrString(s)
			}

			if notOp {
				m1 := make(map[string]interface{})
				m1["$lt"] = r
				m["$not"] = m1
			} else {
				m["$lt"] = r
			}
			q[l.(string)] = m
			return q, nil
		case "le":
			s, ok := rhs.(string)
			if ok {
				r = timeOrString(s)
			}

			if notOp {
				m1 := make(map[string]interface{})
				m1["$lte"] = r
				m["$not"] = m1
			} else {
				m["$lte"] = r
			}
			q[l.(string)] = m
			return q, nil
		default:
			return "", nil
		}
	case *RegexStatement:
		var r interface{}
		l, err := evaluateExpr(stmt.LHE, env)
		if err != nil {
			return nil, err
		}
		if env["parentAttr"] != "" {
			l = env["parentAttr"] + "." + l.(string)
		}

		r = stmt.Value

		m := make(map[string]interface{})
		q := make(map[string]interface{})

		switch stmt.Operator {
		case "co":
			if notOp {
				m1 := make(map[string]interface{})
				m1["$regex"] = bson.RegEx{Pattern: r.(string)}
				m["$not"] = m1
			} else {
				m["$regex"] = bson.RegEx{Pattern: r.(string)}
			}
			q[l.(string)] = m
			return q, nil
		case "sw":
			if notOp {
				m1 := make(map[string]interface{})
				m1["$regex"] = bson.RegEx{Pattern: fmt.Sprintf("^%s", r)}
				m["$not"] = m1
			} else {
				m["$regex"] = bson.RegEx{Pattern: fmt.Sprintf("^%s", r)}
			}
			q[l.(string)] = m
			return q, nil
		case "ew":
			if notOp {
				m1 := make(map[string]interface{})
				m1["$regex"] = bson.RegEx{Pattern: fmt.Sprintf("%s$", r)}
				m["$not"] = m1
			} else {
				m["$regex"] = bson.RegEx{Pattern: fmt.Sprintf("%s$", r)}
			}
			q[l.(string)] = m
			return q, nil
		default:
			return "", nil
		}
	case *ParenStatement:
		env["operator"] = stmt.Operator
		v, err := Evaluate(stmt.SubStatement, env)
		if err != nil {
			return "", err
		}
		return v, nil
	case *GroupStatement:
		parentAttr, err := evaluateExpr(stmt.ParentAttr, Env{})
		if err != nil {
			return nil, err
		}
		env["parentAttr"] = parentAttr.(string)
		v, err := Evaluate(stmt.SubStatement, env)
		if err != nil {
			return "", err
		}
		return v, nil

		return nil, nil
	default:
		return nil, errors.New("Unknown Statement type")
	}
}

func evaluateExpr(expr Expression, env Env) (interface{}, error) {
	switch e := expr.(type) {
	case *NumberExpression:
		v := e.Lit
		return v, nil
	case *IdentifierExpression:
		if v, ok := env[e.Lit]; ok {
			return v, nil
		} else {
			return e.Lit, nil
		}
	case *AttrValueExpression:
		return e.Lit, nil
	case *BoolExpression:
		return e.Lit, nil

	default:
		return nil, errors.New("Unknown Expression type")
	}
}

func timeOrString(str string) interface{} {
	var r interface{}
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		r = str
	} else {
		r = t
	}
	return r
}
