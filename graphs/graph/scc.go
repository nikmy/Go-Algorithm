package graph

import (
	"golang.org/x/exp/constraints"
	"slices"
)

func (g *Graph[Vertex, Edge]) BuildCondensation() *Graph[[]Vertex, Edge] {
	order := makeSCCVisitOrder(g.NVertices(), g.neighbors)
	scc := markSCC(g.NVertices(), g.neighbors, order)

	nSCC := slices.Max(scc)
	cNeigh := make([][]ID, nSCC)
	cEdges := make(map[edgeID]*Edge, len(g.edges))

	for from, toList := range g.neighbors {
		for _, to := range toList {
			fromComp, toComp := scc[from], scc[to]
			cNeigh[fromComp] = append(cNeigh[fromComp], toComp)
			cEdges[getEdgeID(fromComp, toComp)] = g.GetEdge(ID(from), to)
		}
	}

	cVert := make([][]Vertex, nSCC)
	for v, comp := range scc {
		cVert[comp] = append(cVert[comp], g.vertices[v])
	}

	return &Graph[[]Vertex, Edge]{
		neighbors: cNeigh,
		vertices:  cVert,
		edges:     cEdges,
	}
}

func markSCC[I constraints.Integer](nVertices int, adj [][]I, order []I) []I {
	scc := make([]I, 0, nVertices)
	rev := transpose(adj)
	currScc := I(0)

	visited := make([]bool, nVertices)
	var dfs func(I)
	dfs = func(u I) {
		if visited[u] {
			return
		}
		visited[u] = true

		for _, v := range rev[u] {
			dfs(v)
		}
		scc[u] = currScc
	}

	for i := 0; i < len(order); i++ {
		u := I(len(order) - i - 1)
		dfs(u)
		currScc++
	}
	return scc[:len(scc):len(scc)]
}

func makeSCCVisitOrder[I constraints.Integer](nVertices int, adj [][]I) []I {
	visited := make([]bool, nVertices)
	order := make([]I, 0, nVertices)

	var dfs func(I)
	dfs = func(u I) {
		if visited[u] {
			return
		}
		visited[u] = true

		for _, v := range adj[u] {
			dfs(v)
		}
		order = append(order, u)
	}

	for u := I(0); u < I(nVertices); u++ {
		dfs(u)
	}
	return order
}

func transpose[I constraints.Integer](adj [][]I) [][]I {
	r := make([][]I, len(adj))
	for to, fromList := range adj {
		for _, from := range fromList {
			r[from] = append(r[from], I(to))
		}
	}
	return r
}
