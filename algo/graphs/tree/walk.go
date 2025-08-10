package tree

func (t *Node[T]) Preorder(callback func(T)) {
	if t == nil {
		return
	}

	callback(t.Val)
	t.Left.Preorder(callback)
	t.Right.Preorder(callback)
}

func (t *Node[T]) Inorder(callback func(T)) {
	if t == nil {
		return
	}

	t.Left.Inorder(callback)
	callback(t.Val)
	t.Right.Inorder(callback)
}

func (t *Node[T]) Postorder(callback func(T)) {
	if t == nil {
		return
	}

	t.Left.Postorder(callback)
	t.Right.Postorder(callback)
	callback(t.Val)
}

func (t *Node[T]) LevelOrder(callback func(int, T)) {
	if t == nil {
		return
	}

	var queue, next []*Node[T]
	queue = append(queue, t)

	depth := 0
	for len(queue) > 0 {
		for _, node := range queue {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
			callback(depth, t.Val)
		}
		queue, next = next, queue[:0]
		depth++
	}
}
