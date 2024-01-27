package monad

import . "github.com/nikmy/algo/lambda/function"

func newResult[T any](val T, err error) Result[T] {
	return Result[T]{&val, err}
}

type Result[T any] struct {
	val *T
	err error
}

func (r Result[T]) make(v T) Result[T] {
	return newResult[T](v, nil)
}

func (r Result[T]) chain(to Lambda[T, Result[T]]) Result[T] {
	if r.err != nil {
		return newResult[T](nil, r.err)
	}
	return to(*r.val)
}
