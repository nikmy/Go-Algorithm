package tree

type Node[T any] struct {
	Left  *Node[T]
	Right *Node[T]
	Val   T
}
