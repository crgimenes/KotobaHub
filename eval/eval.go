package eval

import "errors"

type Kotoba struct{}

var ERR_NOT_IMPLEMENTED = errors.New("Not implemented")

func (k *Kotoba) Eval(expr string) (string, error) {
	return "", ERR_NOT_IMPLEMENTED
}
