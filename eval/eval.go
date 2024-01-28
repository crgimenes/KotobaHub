package eval

import "errors"

type Kotoba struct{}

var ERR_NOT_IMPLEMENTED = errors.New("Not implemented")

func isNumber(expr string) bool {
	for _, c := range expr {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func (k *Kotoba) Eval(expr string) (string, error) {
	if isNumber(expr) {
		return expr, nil
	}
	return "", ERR_NOT_IMPLEMENTED
}
