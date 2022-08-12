package drafts

func permute[T any](arr []T, l, r int, gen chan []T, init bool) {
    if l == r {
        gen <- arr
        return
    }

    for i := l; i <= r; i++ {
        arr[l], arr[i] = arr[i], arr[l]
        permute(arr, l+1, r, gen, false)
        arr[l], arr[i] = arr[i], arr[l]
    }

    if init {
        close(gen)
    }
}

func NextPermutation[T any](arr []T) chan []T {
    gen := make(chan []T)
    go permute(arr, 0, len(arr)-1, gen, true)
    return gen
}
