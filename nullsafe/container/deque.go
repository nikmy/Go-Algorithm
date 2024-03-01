package container

import (
	"github.com/nikmy/algo/nullsafe/iter"
	. "github.com/nikmy/algo/nullsafe/null"
)

type deque[T any] struct {
	buffers []circularBuffer[T]
	bufSize int
}

func (d *deque[T]) Iterate() iter.Iterator[T] {
	iters := make([]iter.Iterator[T], 0, len(d.buffers))
	for _, buf := range d.buffers {
		iters = append(iters, buf.Iterate())
	}
	return iter.Chain[T](iters...)
}

func (d *deque[T]) Reverse() iter.Iterator[T] {
	iters := make([]iter.Iterator[T], 0, len(d.buffers))
	for i := len(d.buffers) - 1; i >= 0; i-- {
		iters = append(iters, d.buffers[i].Reverse())
	}
	return iter.Chain[T](iters...)
}

func (d *deque[T]) Size() int {
	if len(d.buffers) == 0 {
		return 0
	}

	size := d.buffers[0].Size()
	if len(d.buffers) == 1 {
		return size
	}

	size += d.lastBuf().Size()
	if len(d.buffers) == 2 {
		return size
	}

	return size + (len(d.buffers)-2)*d.bufSize
}

func (d *deque[T]) At(index int) Safe[T] {
	index = normalizeIndex(d, index)
	bucket, index := d.makeIndex(index)
	if bucket == -1 {
		return Null[T]()
	}

	return d.buffers[bucket].At(index)
}

func (d *deque[T]) add(elem T) {
	if len(d.buffers) == 0 || d.lastBuf().Size() == d.bufSize {
		d.buffers = append(d.buffers, circularBuffer[T]{
			data: make([]T, d.bufSize),
		})
	}

	d.lastBuf().add(elem)
}

func (d *deque[T]) RemoveAt(index int) {
	bucket, index := d.makeIndex(index)
	if bucket != -1 {
		d.buffers[bucket].RemoveAt(index)
	}
}

func (d *deque[T]) Sublist(from, to int) listImpl[T] {
	subDeque := &deque[T]{
		buffers: nil,
		bufSize: d.bufSize,
	}

	from = normalizeIndex(d, from)

	if from == d.Size() {
		return subDeque
	}

	to = normalizeIndex(d, to)
	if to <= from {
		return subDeque
	}

	fromBuf, from := d.makeIndex(from)
	toBuf, to := d.makeIndex(to)

	subDeque.buffers = make([]circularBuffer[T], len(d.buffers)-fromBuf)
	copy(subDeque.buffers, d.buffers[fromBuf:])

	subDeque.buffers[0].size -= from

	subDeque.buffers[0].head += from
	subDeque.buffers[0].head %= d.bufSize

	if toBuf != -1 {
		subDeque.buffers = subDeque.buffers[:toBuf-fromBuf]
		subDeque.lastBuf().size = to
	}

	return subDeque
}

func (d *deque[T]) FillRange(from, to int, filler T) {

}

func (d *deque[T]) makeIndex(index int) (int, int) {
	if index == d.Size() {
		return -1, d.bufSize
	}

	cur := 0
	for index >= d.buffers[cur].Size() {
		index -= d.buffers[cur].Size()
		cur++
	}

	return cur, index
}

func (d *deque[T]) lastBuf() *circularBuffer[T] {
	return &d.buffers[len(d.buffers)-1]
}

type circularBuffer[T any] struct {
	data []T
	size int
	head int
}

func (c *circularBuffer[T]) Size() int {
	return c.size
}

func (c *circularBuffer[T]) At(index int) Safe[T] {
	return Value(c.data[c.getIndex(index)])
}

func (c *circularBuffer[T]) add(elem T) {
	insertIndex := (c.head + c.size) % len(c.data)
	c.data[insertIndex] = elem
	c.size++
}

func (c *circularBuffer[T]) RemoveAt(index int) {
	// omg
}

func (c *circularBuffer[T]) Sublist(from, to int) listImpl[T] {
	return &circularBuffer[T]{
		data: c.data,
		head: c.getIndex(from),
		size: to - from,
	}
}

func (c *circularBuffer[T]) FillRange(from, to int, filler T) {
	from = c.getIndex(from)
	if from >= c.Size() {
		return
	}

	to = normalizeIndex(c, to)
	if to > len(c.data) {
		to -= len(c.data)
		for i := 0; i < to; i++ {
			c.data[i] = filler
		}

		to = len(c.data)
	}

	for i := from; i < to; i++ {
		c.data[from] = filler
	}
}

func (c *circularBuffer[T]) getIndex(index int) int {
	index = normalizeIndex(c, index)
	return (c.head + index) % len(c.data)
}

func (c *circularBuffer[T]) Iterate() iter.Iterator[T] { panic("not implemented") }
func (c *circularBuffer[T]) Reverse() iter.Iterator[T] { panic("not implemented") }
