package skiplist

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"strings"
	"sync/atomic"
)

func NewGeneric[T cmp.Ordered]() *GenericList[T] {
	stub := &genericTower[T]{
		next: make([]atomic.Pointer[genericTower[T]], maxLevel),
	}
	stub.state.Store(towerStateDeleting)
	return &GenericList[T]{
		leftmost: stub,
	}
}

func MakeGeneric[T cmp.Ordered](values ...T) *GenericList[T] {
	l := NewGeneric[T]()
	for _, v := range values {
		l.Insert(v)
	}
	return l
}

// GenericList is generalized skip list for ordered types.
type GenericList[T cmp.Ordered] struct {
	leftmost *genericTower[T]
}

// Lookup returns whether element is in the list or not.
func (l *GenericList[T]) Lookup(x T) bool {
	return l.leftmost.find(x)
}

// Insert inserts element with value x to the list, if it does not exist.
// Returns true, if element has been deleted by current goroutine.
func (l *GenericList[T]) Insert(x T) bool {
	var linksToUpdate [maxLevel]*genericTower[T]
	n, found := l.leftmost.findLinks(linksToUpdate[:], x)
	if found != nil {
		return false
	}
	return newGenericTower[T](x).link(linksToUpdate[:n])
}

// Delete removes element with value x from the list, if one exists.
// Returns true, if element has been deleted by current goroutine.
func (l *GenericList[T]) Delete(x T) bool {
	var linksToUpdate [maxLevel]*genericTower[T]
	n, target := l.leftmost.findLinks(linksToUpdate[:], x)
	if target == nil {
		return false
	}
	return target.unlink(linksToUpdate[:n])
}

func (l *GenericList[T]) Elements(yield func(T) bool) {
	base := l.leftmost.next[0].Load()
	for node := base; node != nil; node = node.next[0].Load() {
		if !yield(node.elem) {
			break
		}
	}
}

func (l *GenericList[T]) IsEmpty() bool {
	return l.leftmost.next[0].Load() == nil
}

// String formats elements like a slice.
func (l *GenericList[T]) String() string {
	if l == nil || l.leftmost == nil {
		return "<nil>"
	}
	if l.leftmost.next[0].Load() == nil {
		return "[]"
	}
	elements := make([]string, 0)
	for elem := range l.Elements {
		elements = append(elements, fmt.Sprintf("%v", elem))
	}
	return "[" + strings.Join(elements, ", ") + "]"
}

// SDump is debug only. Dumps skip list in the following format:
//
//	[head] -> 1 --------------> 3 --------------> [end]
//	[head] -> 1 --------------> 3 -----> 4 -----> [end]
//	[head] -> 1 -----> 2 -----> 3 -----> 4 -----> [end]
func (l *GenericList[T]) SDump() string {
	if l == nil || l.leftmost == nil {
		return "<nil>"
	}

	if l.leftmost.next[0].Load() == nil {
		return "[]"
	}

	levels := make([][]T, 0)
	maxWidth := 0
	for lvl := range maxLevel {
		if l.leftmost.next[lvl].Load() == nil {
			break
		}

		level := make([]T, 0)
		for node := l.leftmost.next[lvl].Load(); node != nil; node = node.next[lvl].Load() {
			level = append(level, node.elem)
			maxWidth = max(maxWidth, len(fmt.Sprintf("%v", node.elem)))
		}
		levels = append(levels, level)
	}

	var sb strings.Builder
	for k := len(levels) - 1; k >= 0; k-- {
		level := levels[k]

		sb.WriteString("[head] -")
		i := 0
		for _, elem := range levels[0] {
			if i < len(level) && elem == level[i] {
				s := fmt.Sprintf("%v", elem)
				sb.WriteString("> [")
				sb.WriteString(s)
				sb.WriteString("] ")
				sb.WriteString(strings.Repeat("-", maxWidth-len(s)))
				i++
			} else {
				sb.WriteString(strings.Repeat("-", maxWidth+5))
			}
			sb.WriteString("-")
		}
		sb.WriteString("> [end]\n")
	}
	return sb.String()
}

func newGenericTower[T cmp.Ordered](x T) *genericTower[T] {
	levels := 1
	for levels < maxLevel && rand.Int()%4 == 0 {
		levels++
	}

	return &genericTower[T]{
		elem: x,
		next: make([]atomic.Pointer[genericTower[T]], levels),
	}
}

type genericTower[T cmp.Ordered] struct {
	elem T
	next []atomic.Pointer[genericTower[T]]

	state atomic.Int32
}

func (t *genericTower[T]) find(x T) bool {
	node := t

	for level := len(t.next) - 1; level >= 0; level-- {
		for node != nil && (node.state.Load() == towerStateDeleting || node.elem < x) {
			next := node.next[level].Load()
			if next == nil || next.elem > x {
				break
			}
			node = next
		}
		if node == nil {
			return false
		}
		if node.state.Load() != towerStateDeleting && node.elem == x {
			return true
		}
	}

	return false
}

func (t *genericTower[T]) findLinks(links []*genericTower[T], x T) (int, *genericTower[T]) {
	var (
		node = t
		next *genericTower[T]
	)
	for level := len(t.next) - 1; level >= 0; level-- {
		next = node.next[level].Load()
		for next != nil && next.elem < x {
			if next.state.Load() == towerStateDeleting {
				next = next.next[level].Load()
				continue
			}
			node = next
			next = next.next[level].Load()
		}
		links[level] = node
	}

	if next != nil && next.state.Load() != towerStateDeleting && next.elem == x {
		return len(t.next), next
	}
	return len(t.next), nil
}

func (t *genericTower[T]) link(links []*genericTower[T]) bool {
	for level := 0; level < len(t.next); level++ {
		left := links[level]
		for {
			right := left.next[level].Load()
			for right != nil && right.elem < t.elem {
				left = right
				right = right.next[level].Load()
			}
			if level == 0 && right != nil && right.elem == t.elem {
				return false
			}

			t.next[level].Store(right)
			if left.next[level].CompareAndSwap(right, t) {
				break
			}
		}
	}

	t.state.Store(towerStateCreated)

	return true
}

func (t *genericTower[T]) unlink(links []*genericTower[T]) bool {
	if !t.state.CompareAndSwap(towerStateCreated, towerStateDeleting) {
		return false
	}

	for level := 0; level < len(t.next); level++ {
		/*
			Unlinking B from A -> B -> C

			1. Make loop:

				A -> (B)

			2. Switch forward link:

				A   (B)   C
			    |---------^

			3. Make reverse link:

				A <--- B    C
				|-----------^
		*/

		// Step 1: announce deletion (make a loop link)
		var right *genericTower[T]
		for {
			right = t.next[level].Load()
			if t.next[level].CompareAndSwap(right, t) {
				break
			}
		}

		// Step 2: switch forward link
		left := links[level]
		for {
			next := left.next[level].Load()
			for next != nil && next.elem < t.elem {
				left = next
				next = left.next[level].Load()
			}
			if next == nil || next.elem > t.elem {
				break
			}
			if left.next[level].CompareAndSwap(t, right) {
				break
			}
		}

		// Step 3: make reverse link
		if t.next[level].CompareAndSwap(t, left) {
			if t.next[level].Load() != left {
				panic("genericTower[T] unlink fail")
			}
		}
	}

	return true
}
