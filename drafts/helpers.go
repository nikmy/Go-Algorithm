import "sort"

type ordered interface {
    ~int   | ~uint   | ~int8    | ~uint8  |
    ~int16 | ~uint16 | ~int32   | ~uint32 |
    ~int64 | ~uint64 | ~float32 | ~float64
}

func Min[T ordered](args ...T) T {
    m := args[0]
    for _, x := range args[1:] {
        if x < m {
            m = x
        }
    }
    return m
}

func Max[T ordered](args ...T) T {
    m := args[0]
    for _, x := range args[1:] {
        if x > m {
            m = x
        }
    }
    return m
}

func Abs[T ordered](x, y T) T {
    if x < y {
        return y - x
    }
    return x - y
}

func LowerBound[T any](arr []T, x T, func less(x, y T) bool) {
    return sort.Search(len(arr), func (i int) bool { return less(x, arr[i]) })
}

func UpperBound[T any](arr []T, func less(x, y T) bool) {
    return sort.Search(len(arr), func (i int) bool {
        if i == len(arr) - 1 {
            return len(arr)
        }
        return less(arr[i], x) && !less(arr[i+1], x)
    })
}
