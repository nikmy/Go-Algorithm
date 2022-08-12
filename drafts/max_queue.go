package drafts

import "math"

// Helper
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Stack
/////////////////////////////////////////////////////////////////////////////////////////
type stackNode struct {
    val, max int
}

type Stack struct {
    data []stackNode
    size uint64
}

func NewStack() *Stack {
    return &Stack{make([]stackNode, 0), 0}
}

func (s *Stack) GetMax() int {
    if s.size == 0 {
        return math.MinInt32
    }
    return s.data[s.size-1].max
}

func (s *Stack) Push(x int) {
    s.data = append(s.data, stackNode{x, max(x, s.GetMax())})
    s.size++
}

func (s *Stack) Pop() int {
    s.size--
    top := s.data[s.size]
    s.data = s.data[:s.size]
    return top.val
}

func (s *Stack) Clear() {
    s.data = s.data[:0]
    s.size = 0
}

/////////////////////////////////////////////////////////////////////////////////////////

// Queue
/////////////////////////////////////////////////////////////////////////////////////////
type Queue struct {
    head *Stack
    tail *Stack
}

func NewQueue() *Queue {
    return &Queue{NewStack(), NewStack()}
}

func (this *Queue) GetMax() int {
    return max(this.head.GetMax(), this.tail.GetMax())
}

func (this *Queue) Push(x int) {
    this.tail.Push(x)
}

func (this *Queue) Pop() int {
    if this.head.size == 0 {
        for this.tail.size > 0 {
            this.head.Push(this.tail.Pop())
        }
    }
    return this.head.Pop()
}

/////////////////////////////////////////////////////////////////////////////////////////

func MaxSlidingWindow(nums []int, k int) []int {
    q := NewQueue()
    for _, x := range nums[:k] {
        q.Push(x)
    }

    max := make([]int, len(nums)-k+1)
    max[0] = q.GetMax()

    for pos := k; pos < len(nums); pos++ {
        q.Pop()
        q.Push(nums[pos])
        max[pos-k+1] = q.GetMax()
    }

    return max
}
