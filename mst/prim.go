package mst

import (
	"cmp"
	"math"
)

// PrimCost O(VE) version
func PrimCost(nVertices int, edges [][]int) int {
	w := make([]int, nVertices)
	for i := range w {
		w[i] = math.MaxInt
	}
	w[0] = 0

	queue := make([]int, nVertices)
	for i := range queue {
		queue[i] = i
	}

	inMST := make([]bool, nVertices)

	var cost int

	for len(queue) != 0 {
		next := extractMin(&queue, func(u, v int) int {
			return cmp.Compare(w[u], w[v])
		})

		inMST[next] = true
		cost += w[next]

		for _, edge := range edges {
			if inMST[edge[0]] && !inMST[edge[1]] && edge[2] < w[edge[1]] {
				w[edge[1]] = edge[2]
			}
		}
	}

	return cost
}

func extractMin(s *[]int, comp func(x, y int) int) int {
	var mi, mv int
	for i, v := range *s {
		if comp(v, mv) == -1 {
			mi, mv = i, v
		}
	}

	(*s)[mi], (*s)[len(*s)-1] = (*s)[len(*s)-1], (*s)[mi]
	*s = (*s)[:len(*s)-1]
	return mv
}
