// Helper
func max(a, b int) int {
  if a > b { return a }
  return b
}

// Stack
/////////////////////////////////////////////////////////////////////////////////////////
type node struct {
  val, max int
}

type Stack struct {
  data []node
  size uint64
}

func NewStack() *Stack {
  return &Stack{make([]node, 0), 0}
}

func (this *Stack) GetMax() int {
  if this.size == 0 {
    return math.MinInt32
  }
  return this.data[this.size - 1].max
}

func (this *Stack) Push(x int) {
  this.data = append(this.data, node{x, max(x, this.GetMax())})
  this.size++
}

func (this *Stack) Pop() int {
  this.size--
  top := this.data[this.size]
  this.data = this.data[:this.size]
  return top.val
}

func (this *Stack) Clear() {
  this.data = this.data[:0]
  this.size = 0
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
  
  max := make([]int, len(nums) - k + 1)
  max[0] = q.GetMax()
  
  for pos := k; pos < len(nums); pos++ {
    q.Pop()
    q.Push(nums[pos])
    max[pos - k + 1] = q.GetMax()
  }
  
  return max
}
