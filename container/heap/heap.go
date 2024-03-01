package heap

import "cmp"

type Heap[T any] interface {
	Min() T
	Size() int
	Empty() bool

	Push(T)
	Pop() T
}

func Heapify[T cmp.Ordered](elems []T) Heap[T] {
	return Make(cmp.Compare[T], Init[T](elems))
}

func Init[T any](elems []T) func(*hImpl[T]) {
	return func(h *hImpl[T]) {
		h.data = elems
		h.heapify()
	}
}

func ToMax[T any](h *hImpl[T]) {
	h.comp = func(x, y T) int { return -h.comp(x, y) }
}

func Make[T any](comp func(T, T) int, opts ...func(*hImpl[T])) Heap[T] {
	h := &hImpl[T]{
		comp: comp,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

type hImpl[T any] struct {
	comp func(T, T) int
	data []T
}

func (h *hImpl[T]) heapify() {
	for i := len(h.data) / 2; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *hImpl[T]) Size() int {
	return len(h.data)
}

func (h *hImpl[T]) Empty() bool {
	return h.Size() == 0
}

func (h *hImpl[T]) Min() T {
	return h.data[0]
}

func (h *hImpl[T]) Push(x T) {
	h.data = append(h.data, x)
	h.siftUp(len(h.data) - 1)
}

func (h *hImpl[T]) Pop() T {
	m := h.data[0]
	last := len(h.data) - 1
	h.swap(0, last)
	h.data = h.data[:last]
	h.siftDown(0)
	return m
}

func (h *hImpl[T]) siftUp(i int) {
	for h.less(i, (i-1)/2) {
		h.swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}

func (h *hImpl[T]) siftDown(i int) {
	for 2*i+1 < h.Size() {
		left := 2*i + 1
		right := 2*i + 2
		j := left
		if right < h.Size() && h.less(right, left) {
			j = right
		}
		if h.lessOrEqual(i, j) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

func (h *hImpl[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *hImpl[T]) less(i, j int) bool {
	return h.comp(h.data[i], h.data[j]) < 0
}

func (h *hImpl[T]) lessOrEqual(i, j int) bool {
	return h.comp(h.data[i], h.data[j]) <= 0
}
