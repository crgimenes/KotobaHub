package eval

import (
	"errors"
	"kotoba/environment"
)

type Kotoba struct {
	Global *environment.Environment
}

var (
	ERR_NOT_IMPLEMENTED    = errors.New("Not implemented")
	ERR_INVALID_EXPR       = errors.New("Invalid expression")
	ERR_DIV_BY_ZERO        = errors.New("Division by zero")
	ERR_VARIABLE_NOT_FOUND = errors.New("Variable not found")
)

func New() *Kotoba {
	return &Kotoba{
		Global: environment.New(nil),
	}
}

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

func (k *Kotoba) Eval(parentEnv *environment.Environment, expr ...any) (any, error) {
	if parentEnv == nil {
		parentEnv = k.Global
	}

	env := environment.New(parentEnv)

	if len(expr) == 0 {
		return "", ERR_INVALID_EXPR
	}

	var err error
	for i := 0; i < len(expr); i++ {
		switch expr[i].(type) {
		case []any:
			expr[i], err = k.Eval(parentEnv, expr[i].([]any)...)
			if err != nil {
				return nil, err
			}
		}
	}

	if len(expr) == 1 {
		if isInt(expr[0]) {
			return expr[0], nil
		}
		if isString(expr[0]) {
			s := expr[0].(string)
			return s[1 : len(s)-1], nil
		}
	}

	if len(expr) == 2 {
		if expr[0] == `+` {
			if isInt(expr[1]) {
				v := expr[1].(int)
				return +v, nil
			}
			if isString(expr[1]) {
				s := expr[1].(string)
				return s[1 : len(s)-1], nil
			}
		}
		if expr[0] == "get" {
			v, ok := env.Get(expr[1].(string))
			if !ok {
				return nil, ERR_VARIABLE_NOT_FOUND
			}
			return v, nil
		}
	}

	if len(expr) == 3 {
		if expr[0] == "set" {
			env.Set(expr[1].(string), expr[2])
			return expr[2], nil
		}

		if expr[0] == "+" {
			if isInt(expr[1]) && isInt(expr[2]) {
				l := expr[1].(int)
				r := expr[2].(int)
				return l + r, nil
			}
			if isString(expr[1]) && isString(expr[2]) {
				return (expr[1].(string)[1:len(expr[1].(string))-1] + expr[2].(string)[1:len(expr[2].(string))-1]), nil
			}
			return nil, ERR_INVALID_EXPR
		}
		if expr[0] == "-" {
			if isInt(expr[1]) && isInt(expr[2]) {
				l := expr[1].(int)
				r := expr[2].(int)
				return l - r, nil
			}
			return nil, ERR_INVALID_EXPR
		}
		if expr[0] == "*" {
			if isInt(expr[1]) && isInt(expr[2]) {
				l := expr[1].(int)
				r := expr[2].(int)
				return l * r, nil
			}
			return nil, ERR_INVALID_EXPR
		}
		if expr[0] == "/" {
			if isInt(expr[1]) && isInt(expr[2]) {
				l := expr[1].(int)
				r := expr[2].(int)
				if r == 0 {
					return nil, ERR_DIV_BY_ZERO
				}
				return l / r, nil
			}
			return nil, ERR_INVALID_EXPR
		}
	}

	return nil, ERR_NOT_IMPLEMENTED
}
