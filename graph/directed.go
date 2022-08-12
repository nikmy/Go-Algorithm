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

func (g *Graph[Vertex, Edge]) topSort(start ID, color *[]byte, sorted *[]ID, last *int) bool {
    if c := (*color)[start]; c == 1 || c == 2 {
        return false
    }

    (*color)[start] = 1

    for _, n := range g.GetNeighborsByID(start) {
        if !g.topSort(n, color, sorted, last) {
            return false
        }
    }

    (*color)[start] = 2

    *last--
    (*sorted)[*last] = start
    return true
}

func (g *Graph[Vertex, Edge]) TopSort() []ID {
    n := g.NVertices()

    sorted := make([]ID, 0, n)
    color := make([]byte, n)
    last := n

    for u := 0; u < g.NVertices(); u++ {
        if color[u] == 2 {
            continue
        }
        if !g.topSort(ID(u), &color, &sorted, &last) {
            return nil
        }
    }

    return sorted
}
