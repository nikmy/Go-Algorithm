package math

import (
	"math/cmplx"

	"golang.org/x/exp/constraints"
)

type Real interface {
	constraints.Integer | constraints.Float
}

type Complex interface {
	Real | constraints.Complex
}

func Conjugate[T constraints.Complex](x T) T {
	return T(cmplx.Conj(complex128(x)))
}
