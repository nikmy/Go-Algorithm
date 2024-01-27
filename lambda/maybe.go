package lambda

func newResult[Value any](val Value, err error) Result[Value] {
	return Result[Value]{val, err}
}

type Result[Value any] struct {
	val Value
	err error
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{}
}

func Just[T any](value T) Maybe[T] {
	return Maybe[T]{&value}
}

type Maybe[V any] struct {
	raw *V
}

func Pass[V, R any](fn Lambda[V, Maybe[R]]) Lambda[Maybe[V], Maybe[R]] {
	return func(in Maybe[V]) Maybe[R] {
		if in.raw == nil {
			return Nothing[R]()
		}
		return fn(*in.raw)
	}
}

func (m Maybe[T]) Ptr() *T {
	return m.raw
}
