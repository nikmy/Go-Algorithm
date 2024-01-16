package tree

func FindLCA[T any](root *Node[T], c1, c2 *Node[T]) *Node[T] {
	parent := make(map[*Node[T]]*Node[T])

	var findNode func(root *Node[T], target *Node[T]) bool
	findNode = func(root *Node[T], target *Node[T]) bool {
		if root == nil {
			return false
		}

		if root == target {
			return true
		}

		for _, child := range []*Node[T]{root.Left, root.Right} {
			if child == nil {
				continue
			}
			parent[child] = root
			if findNode(child, target) {
				return true
			}
		}

		return false
	}

	if !findNode(root, c1) || !findNode(root, c2) {
		return nil
	}

	for c1 != c2 {
		c1, c2 = parent[c1], parent[c2]
	}

	return c1
}
