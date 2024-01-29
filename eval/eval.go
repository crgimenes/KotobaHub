package eval

import (
	"errors"
)

type Kotoba struct{}

var (
	ERR_NOT_IMPLEMENTED = errors.New("Not implemented")
	ERR_INVALID_EXPR    = errors.New("Invalid expression")
)

func isInt(expr any) bool {
	switch expr.(type) {
	case int:
		return true
	case int8:
		return true
	case int16:
		return true
	case int32:
		return true
	case int64:
		return true
	case uint:
		return true
	case uint8:
		return true
	case uint16:
		return true
	case uint32:
		return true
	case uint64:
		return true
	default:
		return false
	}
}

func isString(expr any) bool {
	switch expr.(type) {
	case string:
		s := expr.(string)
		if len(s) == 2 {
			return false
		}
		if s[0] != '"' || s[len(s)-1] != '"' {
			return false
		}
		return true
	default:
		return false
	}
}

func (k *Kotoba) Eval(expr ...any) (any, error) {
	if len(expr) == 0 {
		return "", ERR_INVALID_EXPR
	}

	//fmt.Printf("expr: %T\n", expr)
	if len(expr) == 1 {
		if isInt(expr[0]) {
			return expr[0], nil
		}
		if isString(expr[0]) {
			s := expr[0].(string)
			return s[1 : len(s)-1], nil
		}
	}
	return "", ERR_NOT_IMPLEMENTED
}
