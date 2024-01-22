package slices

import (
	"cmp"
	"golang.org/x/exp/constraints"
)

func Repeat[T any](elem T, n int) []T {
	s := make([]T, n)
	for i := range s {
		s[i] = elem
	}
	return s
}

func Compare[E constraints.Ordered, S1, S2 ~[]E](a S1, b S2) int {
	if lCmp := cmp.Compare(len(a), len(b)); lCmp != 0 {
		return lCmp
	}

	for i := range a {
		if c := cmp.Compare(a[i], b[i]); c != 0 {
			return c
		}
	}
	return 0
}

type addable interface {
	constraints.Ordered | constraints.Complex
}

func Sum[E addable, S ~[]E](s S) E {
	var sum E
	for _, x := range s {
		sum += x
	}
	return sum
}
