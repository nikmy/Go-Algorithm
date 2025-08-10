package dp

func MapMemo[T comparable]() func(T) bool {
	cache := make(map[T]struct{})
	return func(x T) bool {
		if _, ok := cache[x]; ok {
			return true
		}

		cache[x] = struct{}{}
		return false
	}
}

func NoMemo[T any](_ T) bool { return false }
