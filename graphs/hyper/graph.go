package hyper

import (
	"github.com/nikmy/algo/container/bitset"
)

func NewHypergraph(nVertices int, edges [][]int) Hypergraph {
	v2e := make([]*bitset.Bitset, nVertices)
	for i := range v2e {
		v2e[i] = bitset.New(len(edges))
	}

	e2v := make([]*bitset.Bitset, len(edges))
	for i, edge := range edges {
		e2v[i] = bitset.New(nVertices)
		for _, v := range edge {
			v2e[v].Fix(i)
			e2v[i].Fix(v)
		}
	}

	return Hypergraph{
		v2e: v2e,
		e2v: e2v,
	}
}

type Hypergraph struct {
	v2e []*bitset.Bitset
	e2v []*bitset.Bitset
}
