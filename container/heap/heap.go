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

func Init[T any](elems []T) func (*hImpl[T]) {
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

func (pq *hImpl[T]) heapify() {
	for i := len(pq.data) / 2; i >= 0; i-- {
		pq.siftDown(i)
	}
}

func (pq *hImpl[T]) Size() int {
	return len(pq.data)
}

func (pq *hImpl[T]) Empty() bool {
	return pq.Size() == 0
}

func (pq *hImpl[T]) Min() T {
	return pq.data[0]
}

func (pq *hImpl[T]) Push(x T) {
	pq.data = append(pq.data, x)
	pq.siftUp(len(pq.data) - 1)
}

func (pq *hImpl[T]) Pop() T {
	m := pq.data[0]
	last := len(pq.data) - 1
	pq.swap(0, last)
	pq.data = pq.data[:last]
	pq.siftDown(0)
	return m
}

func (pq *hImpl[T]) siftUp(i int) {
	for pq.less(i, (i-1)/2) {
		pq.swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}

func (pq *hImpl[T]) siftDown(i int) {
	for 2*i+1 < pq.Size() {
		left := 2*i + 1
		right := 2*i + 2
		j := left
		if right < pq.Size() && pq.less(right, left) {
			j = right
		}
		if pq.lessOrEqual(i, j) {
			break
		}
		pq.swap(i, j)
		i = j
	}
}

func (pq *hImpl[T]) swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *hImpl[T]) less(i, j int) bool {
	return pq.comp(pq.data[i], pq.data[j]) < 0
}

func (pq *hImpl[T]) lessOrEqual(i, j int) bool {
	return pq.comp(pq.data[i], pq.data[j]) <= 0
}
