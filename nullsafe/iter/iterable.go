package iter

import . "github.com/nikmy/algo/nullsafe/null"

type Iterable[E any] interface {
	Iterate() Iterator[E]
}

type ReverseIterable[E any] interface {
	Reverse() Iterator[E]
}

func ForEach[T any, S Iterable[T]](s S, callback func(*T)) {
	for i := s.Iterate(); i.Next(); {
		callback(i.Elem().Must())
	}
}

func Fold[T any, S Iterable[T], V any](s S, init V, f func(V, T) V) V {
	result := init
	ForEach(s, func(elem *T) { result = f(result, *elem) })
	return result
}

func Reduce[T any, S Iterable[T]](s S, reduceFunc func(T, T) T) Safe[T] {
	return Fold(s, Null[T](), func(result Safe[T], elem T) Safe[T] {
		if result.IsNull() {
			result.Set(elem)
		} else {
			result.Set(reduceFunc(*result.Must(), elem))
		}
		return result
	})
}
