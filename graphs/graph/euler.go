package graph

import (
	"github.com/nikmy/algo/container/bitset"
)

func (g *Graph[Vertex, Edge]) IsEuler() bool {
	oddDeg := bitset.New(g.NVertices())
	for u, adj := range g.neighbors {
		if len(adj) % 2 == 1 {
			oddDeg.Xor(u)
		}
		for _, v := range adj {
			oddDeg.Xor(int(v))
		}
	}
	return oddDeg.Count() == 0
}
