package container

import (
	"github.com/nikmy/algo/math"
	"github.com/nikmy/algo/nullsafe/iter"
	. "github.com/nikmy/algo/nullsafe/null"
)

type List[T any] interface {
	listImpl[T]

	IsEmpty() bool

	First() Nullable[T]
	Last() Nullable[T]

	Fill(filler T)
	Add(elems ...T)

	PopFirst()
	PopLast()
}

type listImpl[T any] interface {
	Iterate() iter.Iterator[T]
	Reverse() iter.Iterator[T]

	Size() int
	At(index int) Nullable[T]

	add(elem T)

	RemoveAt(index int)

	Sublist(from, to int) listImpl[T]

	FillRange(from, to int, filler T)
}

func makeList[E any](impl listImpl[E]) List[E] {
	return &listWrapper[E]{impl}
}

type listWrapper[T any] struct {
	listImpl[T]
}

func (l *listWrapper[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *listWrapper[T]) First() Nullable[T] {
	return l.At(0)
}

func (l *listWrapper[T]) Fill(filler T) {
	l.FillRange(0, l.Size(), filler)
}

func (l *listWrapper[T]) Add(elems ...T) {
	for _, elem := range elems {
		l.add(elem)
	}
}

func (l *listWrapper[T]) Last() Nullable[T] {
	return l.At(-1)
}

func (l *listWrapper[T]) PopFirst() {
	l.listImpl = l.listImpl.Sublist(1, l.Size())
}

func (l *listWrapper[T]) PopLast() {
	l.listImpl = l.listImpl.Sublist(0, l.Size()-1)
}

func normalizeIndex[S interface { Size() int }](s S, i int) int {
	if math.Abs(i) >= s.Size() {
		return s.Size()
	}

	if i < 0 {
		return i + s.Size()
	}

	return i
}
