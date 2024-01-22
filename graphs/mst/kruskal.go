package mst

import (
	"sort"

	"github.com/nikmy/algo/dsu"
)

// KruskalCost ~ O(ElogE)
func KruskalCost(nVertices int, edges [][]int) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	minCost, uf := 0, dsu.NewSetForest(nVertices)
	for _, nextEdge := range edges {
		if !uf.SameSet(nextEdge[0], nextEdge[1]) {
			uf.Union(nextEdge[0], nextEdge[1])
			minCost += nextEdge[2]
		}
	}
	return minCost
}
