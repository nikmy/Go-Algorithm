package container

import (
	"github.com/nikmy/algo/nullsafe/iter"
	. "github.com/nikmy/algo/nullsafe/null"
)

func NewLinkedList[T any]() List[T] {
	return makeList[T](&linkedList[T]{})
}

type linkedList[T any] struct {
	head llNode[T]
	tail llNode[T]
	size int
}

func (l *linkedList[T]) Size() int {
	return l.size
}

func (l *linkedList[T]) At(index int) Nullable[T] {
	return l.nodeAt(index).Elem()
}

func (l *linkedList[T]) add(elem T) {
	newTail := newNode[T](elem)
	l.tail.SetNext(newTail)
	newTail.SetPrev(l.tail)
	l.tail = newTail
	l.size++
}

func (l *linkedList[T]) RemoveAt(index int) {
	index = l.getIndex(index)
	switch index {
	case 0:
		l.head = l.head.Next()
		l.head.SetPrev(nullNode[T]())
		fallthrough
	case l.Size():
		return
	}

	after := l.nodeAt(index - 1)
	before := after.Next().Next()
	after.SetNext(before)
	before.SetPrev(after)
}

func (l *linkedList[T]) Sublist(from, to int) listImpl[T] {
	from, to = l.getIndex(from), l.getIndex(to)

	subList := &linkedList[T]{
		head: nullNode[T](),
		tail: nullNode[T](),
		size: to - from,
	}

	if from >= min(l.Size(), to) {
		return subList
	}

	head := l.head
	for i := 0; i < from; i++ {
		head = head.Next()
	}

	subList.head.Nullable.Set(unsafeLLNode[T]{
		Value: *head.Elem().Must(),
		next:  head.Next(),
		prev:  nullNode[T](),
	})

	tail := head
	for i := from; i < to; i++ {
		tail = tail.Next()
	}

	subList.tail.Nullable.Set(unsafeLLNode[T]{
		Value: *head.Elem().Must(),
		next:  nullNode[T](),
		prev:  tail.Prev(),
	})

	return subList
}

func (l *linkedList[T]) FillRange(from, to int, filler T) {
	from, to = l.getIndex(from), l.getIndex(to)

	if from >= min(l.Size(), to) {
		return
	}

	node := l.head
	for i := 0; i < from; i++ {
		node = node.Next()
	}

	for i := from; i < to; i++ {
		node.Elem().Set(filler)
		node = node.Next()
	}
}

func (l *linkedList[T]) nodeAt(index int) llNode[T] {
	current, target := 0, l.getIndex(index)

	if target >= l.Size() {
		return nullNode[T]()
	}

	node := l.head
	for current < target && !node.IsNull() {
		current, node = current+1, node.Next()
	}

	return node
}

func (l *linkedList[T]) getIndex(index int) int {
	return normalizeIndex[T](l, index)
}

func (l *linkedList[T]) Iterate() iter.Iterator[T] {
	return &llIterator[T]{curr: l.head}
}

func (l *linkedList[T]) Reverse() iter.Iterator[T] {
	return &llIterator[T]{curr: l.head, rev: true}
}

type llIterator[T any] struct {
	curr llNode[T]
	rev  bool
}

func (i *llIterator[T]) Next() bool {
	if i.rev {
		i.curr = i.curr.Prev()
	} else {
		i.curr = i.curr.Next()
	}

	return !i.curr.IsNull()
}

func (i *llIterator[T]) Elem() Nullable[T] {
	return i.curr.Elem()
}

func nullNode[T any]() llNode[T] {
	return llNode[T]{Null[unsafeLLNode[T]]()}
}

func newNode[T any](value any) llNode[T] {
	return llNode[T]{Value[unsafeLLNode[T]](unsafeLLNode[T]{Value: value})}
}

type unsafeLLNode[T any] struct {
	next  Nullable[unsafeLLNode[T]]
	prev  Nullable[unsafeLLNode[T]]
	Value T
}

type llNode[T any] struct {
	Nullable[unsafeLLNode[T]]
}

func (n llNode[T]) Elem() Nullable[T] {
	if n.IsNull() {
		return Null[T]()
	}

	return Value(n.Must().Value)
}

func (n llNode[T]) SetNext(next llNode[T]) {
	if !n.IsNull() {
		n.Must().next = next
	}
}

func (n llNode[T]) SetPrev(prev llNode[T]) {
	if !n.IsNull() {
		n.Must().prev = prev
	}
}

func (n llNode[T]) Next() llNode[T] {
	if n.IsNull() {
		return llNode[T]{Null[unsafeLLNode[T]]()}
	}
	return llNode[T]{n.Must().next}
}

func (n llNode[T]) Prev() llNode[T] {
	if n.IsNull() {
		return llNode[T]{Null[unsafeLLNode[T]]()}
	}
	return llNode[T]{n.Must().prev}
}

func (n llNode[T]) HasCycle() bool {
	slow, fast := n.Next(), n.Next().Next()
	for !fast.IsNull() && slow != fast {
		slow, fast = slow.Next(), fast.Next().Next()
	}
	return slow == fast
}
