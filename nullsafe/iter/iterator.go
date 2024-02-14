package iter

import (
	"github.com/nikmy/algo/nullsafe/container"
	. "github.com/nikmy/algo/nullsafe/null"
)

type Iterator[T any] interface {
	Next() bool
	Elem() Nullable[T]
}

func Advance[T any, Iter Iterator[T]](i Iter, distance int) {
	for distance > 0 && i.Next() {
		distance--
	}
}

func SkipUntil[T any, Iter Iterator[T]](i Iter, cond func(T) bool) Iter {
	for i.Next() {
		if cond(*i.Elem().Must()) {
			break
		}
	}
	return i
}

func TakeWhile[T any, Iter Iterator[T], L container.List[T]](from Iter, to L, cond func(T) bool) Iter {
	return SkipUntil(from, func(elem T) bool {
		if cond(elem) {
			to.Add(elem)
			return true
		}
		return false
	})
}

func Range[T any, Iter Iterator[T]](start Iter, callback func(T) bool) {
	SkipUntil(start, func(elem T) bool { return !callback(elem) })
}
