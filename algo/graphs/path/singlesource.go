package path

import (
	"github.com/nikmy/algo/graphs/graph"
	"math"
)

type WeightFunc[Edge any] func(*Edge) int64

func fordBellman[Vertex, Edge any](
	g *graph.Graph[Vertex, Edge],
	source, target graph.ID,
	weight WeightFunc[Edge],
) []graph.ID {
	n := graph.ID(g.NVertices())
	

	dist, prev := make([]int64, n), make([]graph.ID, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[source] = 0

	relax := func(from, to graph.ID) {
		if dist[from] == math.MaxInt64 {
			return
		}

		newDist := dist[from] + weight(g.GetEdge(from, to))
		if dist[to] > newDist {
			dist[to], prev[to] = newDist, from
		}
	}

	for k := graph.ID(0); k < n; k++ {
		for i := graph.ID(0); i < n; i++ {
			for _, j := range g.GetNeighborsByID(i) {
				relax(i, j)
			}
		}
	}

	path, iter := make([]graph.ID, n), n-1
	for v := target; v != source; v = prev[v] {
		path[iter] = v
		iter--
	}
	path[iter] = source
	return path[iter:]
}

func dijkstra[Vertex, Edge any](
	g *graph.Graph[Vertex, Edge],
	source, target graph.ID,
	weight WeightFunc[Edge],
) []graph.ID {
	// TODO
	return nil
}

type SingleSourceAlgorithm uint8

const (
	FordBellman SingleSourceAlgorithm = iota
	Dijkstra
)

func FindShortestPath[Vertex, Edge any](
	g *graph.Graph[Vertex, Edge],
	from, to graph.ID,
	w WeightFunc[Edge],
	algorithm SingleSourceAlgorithm,
) []graph.ID {
	switch algorithm {
	case FordBellman:
		return fordBellman(g, from, to, w)
	case Dijkstra:
		return dijkstra(g, from, to, w)
	default:
		return nil
	}
}
