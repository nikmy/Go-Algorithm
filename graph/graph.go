package graph

type ID uint32

type edgeID struct {
    from, to ID
}

func NewGraph[Vertex, Edge any](nVertices int) *Graph[Vertex, Edge] {
    return &Graph[Vertex, Edge]{
        neighbors: make([][]ID, nVertices),
        vertices:  make([]Vertex, nVertices),
        edges:     make(map[edgeID]Edge),
    }
}

type Graph[Vertex, Edge any] struct {
    neighbors [][]ID
    vertices  []Vertex
    edges     map[edgeID]Edge
}

func (g *Graph[Vertex, Edge]) NVertices() int {
    return len(g.neighbors)
}

func (g *Graph[Vertex, Edge]) GetVertexByID(id ID) Vertex {
    return g.vertices[id]
}

func (g *Graph[Vertex, Edge]) GetEdge(from, to ID) Edge {
    return g.edges[edgeID{from, to}]
}

func (g *Graph[Vertex, Edge]) GetNeighborsByID(id ID) []ID {
    if id > ID(len(g.neighbors)) {
        return nil
    }
    return g.neighbors[id]
}
