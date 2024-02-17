package heap

import (
	"slices"
)

func Sort[T any](comp func(T, T) int, slice []T) []T {
	s := slices.Clone(slice)
	SortInPlace(comp, s)
	return s
}

func SortInPlace[T any](comp func(T, T) int, slice []T) {
	h := Make(comp, Init(slice))

	i := 0
	for !h.Empty() {
		slice[i] = h.Pop()
	}
}

func MergeSorted[T any](comp func(T, T) int, lists ...[]T) []T {
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

	h := Make(
		func(i, j int) int { return comp(lists[i][pos[i]], lists[j][pos[j]]) },
		Init(index),
	)

	for !h.Empty() {
		next := h.Min()
		buffer = append(buffer, lists[next][pos[next]])
		if next < len(lists[next]) {
			pos[next]++
			h.Push(next)
		}
	}
	return buffer
}
