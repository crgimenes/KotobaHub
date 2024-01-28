package eval

import "errors"

type Kotoba struct{}

var ERR_NOT_IMPLEMENTED = errors.New("Not implemented")

func isInt(expr string) bool {
	if len(expr) == 0 {
		return false
	}

	// Check if the first character is a sign
	c := expr[0]
	if c == '-' || c == '+' {
		expr = expr[1:]
		if len(expr) == 0 {
			return false
		}
	}

	for _, c := range expr {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func isString(expr string) bool {
	if len(expr) < 2 {
		return false
	}

	if expr[0] != '"' || expr[len(expr)-1] != '"' {
		return false
	}

	return true
}

func (k *Kotoba) Eval(expr string) (string, error) {
	if isInt(expr) {
		return expr, nil
	}
	if isString(expr) {
		expr = expr[1 : len(expr)-1]
		return expr, nil
	}
	return "", ERR_NOT_IMPLEMENTED
}
