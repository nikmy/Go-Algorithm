package drafts

func permute[T any](arr []T, l, r int, gen chan []T) {
	if l == r {
		gen <- arr
		return
	}

	for i := l; i <= r; i++ {
		arr[l], arr[i] = arr[i], arr[l]
		permute(arr, l+1, r, gen)
		arr[l], arr[i] = arr[i], arr[l]
	}
}

func NextPermutation[T any](arr []T) chan []T {
	gen := make(chan []T)
	go func() {
        defer close(gen)
        permute(arr, 0, len(arr)-1, gen)
    }()
	return gen
}
