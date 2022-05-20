package algo

import "sort"

// Union-Find Forest
/////////////////////////////////////////////////////////////////////////////////////////

type DSUForest struct {
    root []int
    rank []int
    Size int
}

func NewSetForest(nComponents int) *DSUForest {
    f := &DSUForest{make([]int, nComponents), make([]int, nComponents), 0}
    for i := 0; i < nComponents; i++ {
        f.root[i] = i
        f.rank[i] = 1
    }
    f.Size = nComponents
    return f
}

func (f *DSUForest) Find(x int) int {
    if x != f.root[x] {
        f.root[x] = f.Find(f.root[x])
    }
    return f.root[x]
}

func (f *DSUForest) SameSet(x, y int) bool {
    return f.Find(x) == f.Find(y)
}

func (f *DSUForest) Union(x, y int) {
    x, y = f.Find(x), f.Find(y)
    if f.rank[y] < f.rank[x] {
        f.root[y] = x
    } else {
        f.root[x] = y
    }
    if f.rank[y] == f.rank[x] {
        f.rank[y]++
    }
    f.Size--
}

/////////////////////////////////////////////////////////////////////////////////////////

// MST: Kruskal ~ O(ElogE), Boruvka ~ O(ElogV)
/////////////////////////////////////////////////////////////////////////////////////////

// Edge: {from, to, cost}

func MSTCostKruskal(nVertices int, edges [][]int) int {
    sort.Slice(edges, func(i, j int) bool {
        return edges[i][2] < edges[j][2]
    })

    minCost, dsu := 0, NewSetForest(nVertices)
    for i := 0; dsu.Size > 1; i++ {
        nextEdge := edges[i]
        if !dsu.SameSet(nextEdge[0], nextEdge[1]) {
            dsu.Union(nextEdge[0], nextEdge[1])
            minCost += nextEdge[2]
        }
    }
    return minCost
}

// Edge: {from, to, cost}

func MSTCostBoruvka(nVertices int, edges [][]int) int {
    minCost, dsu := 0, NewSetForest(nVertices)
    used, minEdge := make([]bool, len(edges)), make([]int, nVertices)
    for dsu.Size > 1 {
        for i, _ := range minEdge {
            minEdge[i] = -1
        }

        for i, edge := range edges {
            src, dst := dsu.Find(edge[0]), dsu.Find(edge[1])
            if !dsu.SameSet(edge[0], edge[1]) {
                if minEdge[src] == -1 || edge[2] < edges[minEdge[src]][2] {
                    minEdge[src] = i
                }
                if minEdge[dst] == -1 || edge[2] < edges[minEdge[dst]][2] {
                    minEdge[dst] = i
                }
            }
        }

        for _, e := range minEdge {
            if e != -1 && !used[e] {
                dsu.Union(edges[e][0], edges[e][1])
                minCost += edges[e][2]
                used[e] = true
            }
        }
    }
    return minCost
}

/////////////////////////////////////////////////////////////////////////////////////////
