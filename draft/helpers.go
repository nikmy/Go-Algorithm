package draft

import (
	"golang.org/x/exp/constraints"
	"sort"
)

type number interface {
	constraints.Integer | constraints.Float
}

func Distance[T number](x, y T) T {
	if x < y {
		return y - x
	}
	return x - y
}

func LowerBound[T any](arr []T, x T, less func(T, T) bool) int {
	return sort.Search(len(arr), func(i int) bool { return less(x, arr[i]) })
}

func UpperBound[T any](arr []T, x T, less func(T, T) bool) int {
	return sort.Search(len(arr), func(i int) bool {
		if i == len(arr)-1 {
			return false
		}
		return less(arr[i], x) && !less(arr[i+1], x)
	})
}
