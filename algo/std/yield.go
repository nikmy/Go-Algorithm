package std

type Iterator[T any] interface {
	Next() bool
	Elem() *T
}

func ForEach[T any, I Iterator[T]](iter I, callback func(*T)) {
	for iter.Next() {
		callback(iter.Elem())
	}
}

// RangeFunc constantly iterates through values yielded from generate func
//
// example:
//
// type tree struct {
//	 left, right *tree
//	 value       int
// }
//
// func (t *tree) Leaves() Iterator[int] {
//	 return RangeFunc[int](t.leaves)
// }
//
// func (t *tree) leaves(Yield func(int), Return func()) {
//	 var gen func(*tree)
//	 gen = func(t *tree) {
//		 if t == nil {
//			 return
//		 }
//
//		 if t.left == nil && t.right == nil {
//			 Yield(t.value)
//		 }
//
//		 gen(t.left)
//		 gen(t.right)
//	 }
//	 gen(t)
//	 Return()
// }
//
func RangeFunc[T any](generate func(Yield func(T), Return func())) Iterator[T] {
	i := &yIterator[T]{buf: make(chan T)}
	go generate(
		func(x T) { i.buf <- x },
		func() { i.end = true; close(i.buf) },
	)
	return i
}

type Generator[T any] interface {
	Generate(Yield func(T), Return func())
}

// Range iterates through values from Generator
func Range[T any, G Generator[T]](generator G) Iterator[T] {
	return RangeFunc(generator.Generate)
}

type yIterator[T any] struct {
	buf chan T
	end bool
}

func (i *yIterator[T]) Next() bool {
	return !i.end
}

func (i *yIterator[T]) Elem() *T {
	t := <-i.buf
	return &t
}
