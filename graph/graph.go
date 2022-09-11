package graph

type ID uint32

type edgeID uint64

func getEdgeID(from, to ID) edgeID {
    return (edgeID(from) << 32) + edgeID(to)
}

func NewGraph[Vertex, Edge any](nVertices int) *Graph[Vertex, Edge] {
    return &Graph[Vertex, Edge]{
        neighbors: make([][]ID, nVertices),
        vertices:  make([]Vertex, nVertices),
        edges:     make(map[edgeID]*Edge),
    }
}

type Graph[Vertex, Edge any] struct {
    neighbors [][]ID
    vertices  []Vertex
    edges     map[edgeID]*Edge
}

func (g *Graph[Vertex, Edge]) NVertices() int {
    return len(g.neighbors)
}

func (g *Graph[Vertex, Edge]) GetVertexByID(id ID) Vertex {
    return g.vertices[id]
}

func (g *Graph[Vertex, Edge]) GetEdge(from, to ID) *Edge {
    return g.edges[getEdgeID(from, to)]
}

func (g *Graph[Vertex, Edge]) GetNeighborsByID(id ID) []ID {
    if id > ID(len(g.neighbors)) {
        return nil
    }
    return g.neighbors[id]
}

func (g *Graph[Vertex, Edge]) HasEdge(from, to ID) bool {
    _, has := g.edges[getEdgeID(from, to)]
    return has
}

func (g *Graph[Vertex, Edge]) AddEdge(from, to ID, e *Edge) {
    g.edges[getEdgeID(from, to)] = e
}
