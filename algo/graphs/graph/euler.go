package graph

import (
	"github.com/nikmy/algo/container/bitset"
)

func (g *Graph[Vertex, Edge]) IsEuler() bool {
	oddDeg := bitset.New(g.NVertices())
	for u, adj := range g.neighbors {
		if len(adj)%2 == 1 {
			oddDeg.FlipBit(u)
		}
		for _, v := range adj {
			oddDeg.FlipBit(int(v))
		}
	}
	return oddDeg.Count() == 0
}
