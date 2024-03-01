package generic

func Empty[T any]() T {
	var x T
	return x
}
