package drafts

import (
    "cmp"
    "math"
)

// Stack
/////////////////////////////////////////////////////////////////////////////////////////
type stackNode[T cmp.Ordered] struct {
    val, max int
}

type Stack[T cmp.Ordered] struct {
    data []stackNode[T]
    size uint64
}

func NewStack[T cmp.Ordered]() *Stack[T] {
    return &Stack[T]{make([]stackNode[T], 0), 0}
}

func (s *Stack[T]) GetMax() int {
    if s.size == 0 {
        return math.MinInt32
    }
    return s.data[s.size-1].max
}

func (s *Stack[T]) Push(x int) {
    s.data = append(s.data, stackNode[T]{x, max(x, s.GetMax())})
    s.size++
}

func (s *Stack[T]) Pop() int {
    s.size--
    top := s.data[s.size]
    s.data = s.data[:s.size]
    return top.val
}

func (s *Stack[T]) Clear() {
    s.data = s.data[:0]
    s.size = 0
}

/////////////////////////////////////////////////////////////////////////////////////////

// Queue
/////////////////////////////////////////////////////////////////////////////////////////
type Queue[T cmp.Ordered] struct {
    head *Stack[T]
    tail *Stack[T]
}

func NewQueue[T cmp.Ordered]() *Queue[T] {
    return &Queue[T]{NewStack[T](), NewStack[T]()}
}

func (q *Queue[T]) GetMax() int {
    return max(q.head.GetMax(), q.tail.GetMax())
}

func (q *Queue[T]) Push(x int) {
    q.tail.Push(x)
}

func (q *Queue[T]) Pop() int {
    if q.head.size == 0 {
        for q.tail.size > 0 {
            q.head.Push(q.tail.Pop())
        }
    }
    return q.head.Pop()
}

/////////////////////////////////////////////////////////////////////////////////////////

func MaxSlidingWindow(nums []int, k int) []int {
    q := NewQueue()
    for _, x := range nums[:k] {
        q.Push(x)
    }

    mx := make([]int, len(nums)-k+1)
    mx[0] = q.GetMax()

    for pos := k; pos < len(nums); pos++ {
        q.Pop()
        q.Push(nums[pos])
        mx[pos-k+1] = q.GetMax()
    }

    return mx
}
