package skiplist

import (
	"math"
	"strconv"
	"strings"
)

func New() *List {
	return &List{
		leftmost: &tower{
			elem: math.MinInt64,
			next: make([]*tower, maxLevel),
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
	base := l.leftmost.next[0]
	for node := base; node != nil; node = node.next[0] {
		if !yield(node.elem) {
			break
		}
	}
}

func (l *List) Find(x int64) bool {
	return l.leftmost.find(x)
}

func (l *List) Insert(x int64) (inserted bool) {
	linksToUpdate := l.leftmost.findLinks(x)

	target := linksToUpdate[0].next[0]
	if target != nil && target.elem == x {
		return false
	}

	return newTower(x).link(linksToUpdate)
}

func (l *List) Delete(x int64) (deleted bool) {
	linksToUpdate := l.leftmost.findLinks(x)

	target := linksToUpdate[0].next[0]
	if target == nil || target.elem != x {
		return false
	}

	return target.unlink(linksToUpdate)
}

func (l *List) String() string {
	elements := make([]string, 0)
	for elem := range l.Elements {
		elements = append(elements, strconv.FormatInt(elem, 10))
	}
	return "[" + strings.Join(elements, ", ") + "]"
}

// SDump is debug only. Dumps skip list in the following format:
//
//	[head] -------------------------------------> 5 --------------------------------> [end]
//	[head] -------------------> 3 --------------> 5 --------------> 7 --------------> [end]
//	[head] -> 1 --------------> 3 -----> 4 -----> 5 --------------> 7 --------------> [end]
//	[head] -> 1 --------------> 3 -----> 4 -----> 5 --------------> 7 --------------> [end]
//	[head] -> 1 -----> 2 -----> 3 -----> 4 -----> 5 -----> 6 -----> 7 -----> 8 -----> [end]
func (l *List) SDump() string {
	if l == nil || l.leftmost == nil {
		return "<nil>"
	}

	if l.leftmost.next[0] == nil {
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
		if l.leftmost.next[lvl] == nil {
			break
		}

		level := make([]int64, 0)
		for node := l.leftmost.next[lvl]; node != nil; node = node.next[lvl] {
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

/*
	Example

    L0  s_0 --------------------------------------------------> s_8   p=0
    L1  s_0 ----------------------> s_4 --------> s_6 ->------> s_8   p=0.25
    L2  s_0 --------> s_2 --------> s_4 --------> s_6 -> s_7 -> s_8   p=0.5
    L3  s_0 -> s_1 -> s_2 -> s_3 -> s_4 -> s_5 -> s_6 -> s_7 -> s_8   p=1
	    ===========================================================
	    T0     T1     T2     T3     T4     T5     T6     T7     T8

	List:
		leftmost: *T0

	Towers:
		T0: {
			elem: s_0
			next: [ *T8, *T4, *T2, *T1 ]
		}
		T1: {
			elem: s_1
			next: [ nil, nil, nil, *T2 ]
		}
		T2: {
			elem:s_2
			next: [ nil, nil, *T3, *T3 ]
		}
		T3: {
			elem: s_3
			next: [nil, nil, nil, *T4 ]
		}
		T4: {
			elem: s_4
			next: [ nil, *T6, *T6, *T5 ]
		}
		T5: {
			elem: s_5
			next: [ nil, nil, nil, *T6 ]
		}
		T6: {
			elem: s_6
			next: [ nil, *T8, *T7, *T7 ]
		}
		T7: {
			elem: s_7
			next: [ nil, nil, *T8, *T8 ]
		}
		T8: {
			elem: s_8
			next: [ nil, nil, nil, nil ]
		}


	Linking B between A and C

			link:
				C = __load__(A.next)

				if C == B:
					return false

				if C < A:      # A -> [Y] -> <B> -> C
					A = C
					goto link

				__store__(B.next, C)            # B -> C
				if not __cas__(A.next, C, B):   # A -> B
					goto link

				return true


*/
