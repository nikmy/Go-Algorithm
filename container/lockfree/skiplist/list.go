package skiplist

import (
	"math"
	"strconv"
	"strings"
	"sync/atomic"
)

func New() *List {
	return &List{
		leftmost: &tower{
			elem: math.MinInt64,
			next: make([]atomic.Pointer[tower], maxLevel),
		},
	}
}

func Make(values ...int64) *List {
	l := New()
	for _, v := range values {
		l.Insert(v)
	}
	return l
}

type List struct {
	leftmost *tower
}

func (l *List) Elements(yield func(int64) bool) {
	base := l.leftmost.next[0].Load()
	for node := base; node != nil; node = node.next[0].Load() {
		if !yield(node.elem) {
			break
		}
	}
}

func (l *List) IsEmpty() bool {
	return l.leftmost.next[0].Load() == nil
}

// Find checks whether x is present in the list or not.
func (l *List) Find(x int64) bool {
	return l.leftmost.find(x)
}

// Insert inserts element with value x to the list, if it does not exist.
// Returns true if element has been inserted by current goroutine.
func (l *List) Insert(x int64) (inserted bool) {
	var linksToUpdate [maxLevel]*tower
	n, found := l.leftmost.findLinks(linksToUpdate[:], x)
	if found != nil {
		// optimize tower creation
		return false
	}

	return newTower(x).link(linksToUpdate[:n])
}

// Delete removes element with value x from the list, if one exists.
// Returns true, if element has been deleted by current goroutine.
//
// Delete that returns false provides eventual consistency guarantee.
// If Delete(x) return true, any call of
// Find(x) (that is sequenced just after Delete via happens-before)
// will return true. So, if you need to ensure that element is completely
// deleted after Delete(x) returned false, you may use for-loop with Find(x).
func (l *List) Delete(x int64) (deleted bool) {
	var linksToUpdate [maxLevel]*tower
	n, target := l.leftmost.findLinks(linksToUpdate[:], x)
	if target == nil {
		return false
	}

	return target.unlink(linksToUpdate[:n])
}

// String formats elements like a slice.
func (l *List) String() string {
	if l == nil || l.leftmost == nil {
		return "<nil>"
	}
	if l.leftmost.next[0].Load() == nil {
		return "[]"
	}
	elements := make([]string, 0)
	for elem := range l.Elements {
		elements = append(elements, strconv.FormatInt(elem, 10))
	}
	return "[" + strings.Join(elements, ", ") + "]"
}

// SDump is debug only. Dumps skip list in the following format:
//
//	[head] -> 1 --------------> 3 --------------> [end]
//	[head] -> 1 --------------> 3 -----> 4 -----> [end]
//	[head] -> 1 -----> 2 -----> 3 -----> 4 -----> [end]
func (l *List) SDump() string {
	if l == nil || l.leftmost == nil {
		return "<nil>"
	}

	if l.leftmost.next[0].Load() == nil {
		return "[]"
	}

	getWidth := func(n int64) int {
		k := 1
		if n < 0 {
			k++
			n = -n
		}
		for n >= 10 {
			n /= 10
			k++
		}
		return k
	}

	levels := make([][]int64, 0)
	maxWidth := 0
	for lvl := range maxLevel {
		if l.leftmost.next[lvl].Load() == nil {
			break
		}

		level := make([]int64, 0)
		for node := l.leftmost.next[lvl].Load(); node != nil; node = node.next[lvl].Load() {
			level = append(level, node.elem)
			maxWidth = max(maxWidth, getWidth(node.elem))
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
				s := strconv.FormatInt(elem, 10)
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
