package monad

import . "github.com/nikmy/algo/lambda/function"

func (m Maybe[T]) make(v T) Maybe[T] {
	return Just(v)
}

func (m Maybe[T]) chain(to Lambda[T, Maybe[T]]) Maybe[T] {
	if m.raw == nil {
		return Nothing[T]()
	}
	return to(*m.raw)
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{}
}

func Just[T any](value T) Maybe[T] {
	return Maybe[T]{&value}
}

type Maybe[T any] struct {
	raw *T
}

func (m Maybe[T]) Ptr() *T {
	return m.raw
}
