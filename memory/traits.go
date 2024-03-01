package memory

import "unsafe"

type Allocator interface {
	Allocate(n int) []byte
	Free(p uintptr, n int)
}

type Ptr[T any] uintptr

func (p Ptr[T]) Elem() *T {
	return (*T)(unsafe.Pointer(p))
}

func New[T any](a Allocator) Ptr[T] {
	bytes := a.Allocate(alignOf[T]())
	return Ptr[T](unsafe.Pointer(&bytes[0]))
}

func NewArray[T any](a Allocator, n int) []T {
	bytes := a.Allocate(alignOf[T]() * n)
	return unsafe.Slice((*T)(unsafe.Pointer(&bytes[0])), n)
}

func Delete[T any](a Allocator, p Ptr[T]) {
	a.Free(uintptr(p), alignOf[T]())
}

func DeleteArray[T any](a Allocator, arr []T) {
	a.Free(uintptr(unsafe.Pointer(&arr[0])), len(arr) * alignOf[T]())
}

func alignOf[T any]() int {
	var stub T
	return int(unsafe.Alignof(stub))
}
