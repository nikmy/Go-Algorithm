package null

import (
	"github.com/nikmy/algo/tools/generic"
)

func First[T any](args ...Safe[T]) Safe[T] {
	for _, maybe := range args {
		if !maybe.IsNull() {
			return maybe
		}
	}
	return Null[T]()
}

func GetOrDefault[T any](maybe Safe[T], defaultValue T) T {
	if maybe.IsNull() {
		return defaultValue
	}
	return *maybe.Must()
}

type Safe[T any] interface {
	IsNull() bool
	Must() *T
	Set(value T)
}

func Null[T any]() Safe[T] {
	return &nullable[T]{}
}

func Empty[T any]() Safe[T] {
	return Value(generic.Empty[T]())
}

func Value[T any](value T) Safe[T] {
	return &nullable[T]{
		ptr:     &value,
		checked: true,
	}
}

type nullable[T any] struct {
	ptr     *T
	checked bool
}

func (n *nullable[T]) IsNull() bool {
	n.checked = true
	return n.ptr == nil
}

func (n *nullable[T]) Must() *T {
	if n.checked {
		return n.ptr
	}
	panic("potential null dereference")
}

func (n *nullable[T]) Set(value T) {
	n.ptr = &value
	n.checked = true
}
