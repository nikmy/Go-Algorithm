package slices

import (
	"golang.org/x/exp/constraints"
)

func Count[E comparable, S ~[]E](slice S, elem E) int {
	var cnt int
	for _, x := range slice {
		if x == elem {
			cnt++
		}
	}
	return cnt
}

func Repeat[T any](elem T, n int) []T {
	s := make([]T, n)
	for i := range s {
		s[i] = elem
	}
	return s
}

type addable interface {
	constraints.Ordered | constraints.Complex | ~string
}

func Sum[E addable, S ~[]E](s S) E {
	var sum E
	for _, x := range s {
		sum += x
	}
	return sum
}

func Mean[E interface {
	constraints.Integer | constraints.Float
}, S ~[]E](s S) E {
	return Sum(s) / E(len(s))
}

func Prod[E interface {
	constraints.Integer | constraints.Float | constraints.Complex
}, S ~[]E](s S) E {
	var p E
	for _, x := range s {
		p *= x
	}
	return p
}
