package std

import "errors"

func Async[T any](callback func() (T, error)) Future[T] {
	ch := make(chan result[T])
	go func() { ch <- makeResult(callback()); close(ch) }()
	return func() result[T] { return <-ch }
}

type Future[T any] func() result[T]

func Await[T any, F Future[T]](future F) (T, error) {
	r := future()
	return r.val, r.err
}

func Chain[A, B any, F Future[A]](future F, then func(A) (B, error)) Future[B] {
	return func() result[B] {
		a := future()
		if a.err != nil {
			return result[B]{err: a.err}
		}

		return makeResult(then(a.val))
	}
}

func Catch[T any, F Future[T]](future F, catch func(error) (T, error)) Future[T] {
	return func() result[T] {
		r := future()
		if r.err != nil {
			return makeResult(catch(r.err))
		}

		return r
	}
}

func Finally[T any, F Future[T]](future F, finally func() error) Future[T] {
	return func() result[T] {
		r := future()
		r.err = errors.Join(r.err, finally())
		return r
	}
}

func makeResult[T any](val T, err error) result[T] {
	return result[T]{val: val, err: err}
}

type result[T any] struct {
	val T
	err error
}
