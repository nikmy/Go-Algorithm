package graph

import "math"

type WeightFunc[Edge any] func(*Edge) int64

func (g *Graph[Vertex, Edge]) fordBellman(source, target ID, weight WeightFunc[Edge]) []ID {
	n := ID(g.NVertices())

	dist, prev := make([]int64, n), make([]ID, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[source] = 0

	relax := func(from, to ID) {
		if dist[from] == math.MaxInt64 {
			return
		}

		newDist := dist[from] + weight(g.GetEdge(from, to))
		if dist[to] > newDist {
			dist[to], prev[to] = newDist, from
		}
	}

	for k := ID(0); k < n; k++ {
		for i := ID(0); i < n; i++ {
			for _, j := range g.GetNeighborsByID(i) {
				relax(i, j)
			}
		}
	}

	path, iter := make([]ID, n), n-1
	for v := target; v != source; v = prev[v] {
		path[iter] = v
		iter--
	}
	path[iter] = source
	return path[iter:]
}

func (g *Graph[Vertex, Edge]) dijkstra(source, target ID, weight WeightFunc[Edge]) []ID {
	// TODO
	return nil
}

func (g *Graph[Vertex, Edge]) johnson(source, target ID, weight WeightFunc[Edge]) []ID {
	// TODO
	return nil
}

type SingleSourceAlgorithm uint8

const (
	FordBellman SingleSourceAlgorithm = iota
	Dijkstra
	Johnson
)

func (g *Graph[Vertex, Edge]) FindShortestPath(from, to ID, w WeightFunc[Edge], algorithm SingleSourceAlgorithm) []ID {
	switch algorithm {
	case FordBellman:
		return g.fordBellman(from, to, w)
	case Dijkstra:
		return g.dijkstra(from, to, w)
	case Johnson:
		return g.johnson(from, to, w)
	default:
		return nil
	}
}
