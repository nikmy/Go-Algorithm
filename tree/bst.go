package tree

func (t *Node[T]) IsBST(cmp func(T, T) int) bool {
	if t == nil {
		return true
	}

	if t.Left != nil {
		if cmp(t.Left.Val, t.Val) > 0 || !t.Left.IsBST(cmp) {
			return false
		}
	}

	if t.Right != nil {
		if cmp(t.Val, t.Right.Val) >= 0 || !t.Right.IsBST(cmp) {
			return false
		}
	}

	return true
}

func (t *Node[T]) Find(cmp func(T, T) int, target T) *Node[T] {
	if t == nil {
		return nil
	}

	compare := cmp(t.Val, target)
	if compare == 0 {
		return t
	}
	if compare < 0 {
		return t.Left.Find(cmp, target)
	}
	return t.Right.Find(cmp, target)
}