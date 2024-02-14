package null

func GetOrDefault[T any](maybe Nullable[T], defaultValue T) T {
	if maybe.IsNull() {
		return defaultValue
	}
	return *maybe.Must()
}

type Nullable[T any] interface {
	IsNull() bool
	Must() *T
	Set(value T)
}

func Null[T any]() Nullable[T] {
	return &nullable[T]{}
}

func Value[T any](value T) Nullable[T] {
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
