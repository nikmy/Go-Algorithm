package functools

type MapLike[K comparable, V any] interface{ ~map[K]V }
type SliceLike[E any] interface{ ~[]E }

type iterable[K comparable, V any] interface{ MapLike[K, V] | SliceLike[V] }

func MapReduceFunc[K comparable, V any, M MapLike[K, V]](
	callbackFn func(accumulator V, currentKey K, currentValue V) V,
	initialValue V,
) func(M) V {
	return reduce[K, V, M](callbackFn, initialValue)
}

func SliceMapFunc[From, To any, S SliceLike[From]](mapFunc func(int, From) To) func(S) []To {
	return func(slice S) []To {
		result := make([]To, 0, len(slice))
		for i, v := range slice {
			result[i] = mapFunc(i, v)
		}
		return result
	}
}

func SliceReduceFunc[V any, S SliceLike[V]](
	callbackFn func(accumulator V, currentIndex int, currentValue V) V,
	initialValue V,
) func(S) V {
	return reduce[int, V, S](callbackFn, initialValue)
}

func reduce[K comparable, V any, S iterable[K, V]](
	callbackFn func(accumulator V, currentKey K, currentValue V) V,
	initialValue V,
) func(S) V {
	return func(s S) V {
		result := initialValue
		for k, v := range s {
			result = callbackFn(result, k, v)
		}
		return result
	}
}
