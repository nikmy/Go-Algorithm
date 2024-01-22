package heap

func Make[T any](lessFunc func(T, T) bool, elems []T) *Heap[T] {
	h := &Heap[T]{
		comp: lessFunc,
		data: elems,
	}
	h.heapify()
	return h
}

type Heap[T any] struct {
	comp func(T, T) bool
	data []T
	size int
}

func (h *Heap[T]) heapify() {
	for i := len(h.data) / 2; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *Heap[T]) Size() int {
	return h.size
}

func (h *Heap[T]) Empty() bool {
	return h.Size() == 0
}

func (h *Heap[T]) Insert(x T) {
	h.data = append(h.data, x)
	h.size++
	h.siftUp(len(h.data) - 1)
}

func (h *Heap[T]) GetMin() T {
	return h.data[0]
}

func (h *Heap[T]) ExtractMin() T {
	m := h.data[0]
	h.swap(0, len(h.data)-1)
	h.size--
	h.siftDown(0)
	return m
}

func (h *Heap[T]) Modify(i int, val T) {
	prev := h.data[i]
	h.data[i] = val
	if h.comp(prev, val) {
		h.siftDown(i)
	} else {
		h.siftUp(i)
	}
}

func (h *Heap[T]) siftUp(i int) {
	for h.less(i, (i-1)/2) {
		h.swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}

func (h *Heap[T]) siftDown(i int) {
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

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) less(i, j int) bool {
	return h.comp(h.data[i], h.data[j])
}

func (h *Heap[T]) lessOrEqual(i, j int) bool {
	if h.less(i, j) {
		return true
	}
	return !h.less(j, i)
}
