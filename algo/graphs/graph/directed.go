package graph

func (g *Graph[Vertex, Edge]) hasCycle(start ID, color *[]byte) bool {
    if (*color)[start] == 2 {
        return false
    }
    if (*color)[start] == 1 {
        return true
    }

    (*color)[start] = 1

    for _, n := range g.GetNeighborsByID(start) {
        if g.hasCycle(n, color) {
            return true
        }
    }

    (*color)[start] = 2

    return false
}

func (g *Graph[Vertex, Edge]) HasCycle() bool {
    color := make([]byte, g.NVertices())
    for u := 0; u < g.NVertices(); u++ {
        if color[u] == 2 {
            continue
        }
        if g.hasCycle(ID(u), &color) {
            return true
        }
    }
    return false
}

func (g *Graph[Vertex, Edge]) topSort(v ID, visited []bool, sorted *[]ID) bool {
    if visited[v] {
        return false
    }
    visited[v] = true

    for _, n := range g.GetNeighborsByID(v) {
        if !g.topSort(n, visited, sorted) {
            return false
        }
    }

    *sorted = append(*sorted, v)
    return true
}

func (g *Graph[Vertex, Edge]) TopSort() []ID {
    n := g.NVertices()

    sorted := make([]ID, 0, n)
    visited := make([]bool, n)

    for u := 0; u < g.NVertices(); u++ {
        if visited[u] {
            continue
        }
        if !g.topSort(ID(u), visited, &sorted) {
            return nil
        }
    }

    return sorted
}
