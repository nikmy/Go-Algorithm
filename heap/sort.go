package heap

import (
	"slices"
)

func Sort[T any](lessFunc func(T, T) bool, slice []T) []T {
	s := slices.Clone(slice)
	SortInPlace(lessFunc, s)
	return s
}

func SortInPlace[T any](lessFunc func(T, T) bool, slice []T) {
	h := Make(func(a, b T) bool { return !lessFunc(a, b) }, slice)
	for h.Size() > 0 {
		h.ExtractMin()
	}
}

func MergeSorted[T any](lessFunc func(T, T) bool, lists ...[]T) []T {
	var bufferSize int
	for _, list := range lists {
		bufferSize += len(list)
	}
	buffer := make([]T, 0, bufferSize)

	k := len(lists)
	pos := make([]int, k)

	index := make([]int, k)
	for i := range index {
		index[i] = i
	}

	h := Make(func(i, j int) bool { return lessFunc(lists[i][pos[i]], lists[j][pos[j]]) }, index)
	for !h.Empty() {
		next := h.ExtractMin()
		buffer = append(buffer, lists[next][pos[next]])
		if next < len(lists[next]) {
			pos[next]++
			h.Insert(next)
		}
	}
	return buffer
}
