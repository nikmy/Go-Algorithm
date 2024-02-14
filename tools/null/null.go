package null

type Nullable[T any] interface {
	IsNull() bool
	Must() *T
	Set(value T) Nullable[T]
	Or(defaultN Nullable[T]) Nullable[T]
	SetDefault(defaultValue T) Nullable[T]
}

func New[T any]() Nullable[T] {
	return &nullable[T]{}
}

func Value[T any](value T) Nullable[T] {
	return &nullable[T]{
		ptr:     &value,
		checked: true,
	}
}

func Safe[T any](ptr *T) Nullable[T] {
	if ptr == nil {
		return New[T]()
	}

	return &nullable[T]{
		ptr:     ptr,
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

func (n *nullable[T]) Set(value T) Nullable[T] {
	n.ptr = &value
	return n
}

func (n *nullable[T]) Or(defaultN Nullable[T]) Nullable[T] {
	if n.IsNull() {
		return defaultN
	}
	return n
}

func (n *nullable[T]) SetDefault(defaultValue T) Nullable[T] {
	if n.IsNull() {
		n.ptr = &defaultValue
	}
	return n
}
