package mst

import "github.com/nikmy/algo/container/dsu"

// BoruvkaCost - O(ElogV)
func BoruvkaCost(nVertices int, edges [][]int) int {
	minCost, uf := 0, dsu.NewSetForest(nVertices)
	used, minEdge := make([]bool, len(edges)), make([]int, nVertices)
	for uf.Size > 1 {
		for i, _ := range minEdge {
			minEdge[i] = -1
		}

		for i, edge := range edges {
			src, dst := uf.Find(edge[0]), uf.Find(edge[1])
			if !uf.SameSet(edge[0], edge[1]) {
				if minEdge[src] == -1 || edge[2] < edges[minEdge[src]][2] {
					minEdge[src] = i
				}
				if minEdge[dst] == -1 || edge[2] < edges[minEdge[dst]][2] {
					minEdge[dst] = i
				}
			}
		}

		for _, e := range minEdge {
			if e != -1 && !used[e] {
				uf.Union(edges[e][0], edges[e][1])
				minCost += edges[e][2]
				used[e] = true
			}
		}
	}
	return minCost
}
