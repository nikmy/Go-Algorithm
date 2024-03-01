package container

import (
	"github.com/nikmy/algo/nullsafe/iter"
	. "github.com/nikmy/algo/nullsafe/null"
)

func NewSlice[T any](opts SliceOptions[T]) List[T] {
	var s slice[T]

	sz := GetOrDefault(opts.Size, 0)
	cp := GetOrDefault(opts.Cap, sz)

	if cp == 0 {
		return makeList[T](&s)
	}

	s.s = make([]T, sz, cp)
	return makeList[T](&s)
}

type SliceOptions[T any] struct {
	Size Safe[int]
	Cap  Safe[int]
}

type slice[T any] struct {
	s []T
}

func (s *slice[T]) Size() int {
	return len(s.s)
}

func (s *slice[T]) At(index int) Safe[T] {
	index = s.getIndex(index)
	if index >= s.Size() {
		return Null[T]()
	}

	return Value(s.s[index])
}

func (s *slice[T]) RemoveAt(index int) {
	index = s.getIndex(index)
	if index >= s.Size() {
		return
	}

	if index < s.Size()-1 {
		copy(s.s[index:], s.s[index+1:])
	}

	s.s = s.s[:s.Size()-1]
}

func (s *slice[T]) Sublist(from, to int) listImpl[T] {
	from, to = s.getIndex(from), s.getIndex(to)

	if from >= min(s.Size(), to) {
		return &slice[T]{s: nil}
	}

	return &slice[T]{s: s.s[from:to]}
}

func (s *slice[T]) FillRange(from, to int, filler T) {
	from, to = s.getIndex(from), s.getIndex(to)

	if from >= min(s.Size(), to) {
		return
	}

	for i := from; i < to; i++ {
		s.s[i] = filler
	}
}

func (s *slice[T]) getIndex(i int) int {
	return normalizeIndex(s, i)
}

func (s *slice[T]) add(elem T) {
	s.s = append(s.s, elem)
}

func (s *slice[T]) Iterate() iter.Iterator[T] {
	return &sliceIterator[T]{s: s}
}

func (s *slice[T]) Reverse() iter.Iterator[T] {
	return &sliceIterator[T]{
		s:   s,
		pos: s.Size() - 1,
		rev: true,
	}
}

type sliceIterator[T any] struct {
	s   *slice[T]
	pos int
	rev bool
}

func (i *sliceIterator[T]) Next() bool {
	if i.rev {
		i.pos--
		return i.pos >= 0
	}
	i.pos++
	return i.pos < i.s.Size()
}

func (i *sliceIterator[T]) Elem() Safe[T] {
	return i.s.At(i.pos)
}
