package util

import (
	"golang.org/x/exp/constraints"
)

func IntegerValue[T constraints.Integer](ptr *T) T {
	if ptr == nil {
		return 0
	}
	return *ptr
}
