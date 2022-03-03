// Union-Find Forest
/////////////////////////////////////////////////////////////////////////////////////////
type SetForest struct {
  root []int
  rank []int
  Size int
}

func NewSetForest(n_components int) *SetForest {
  f := &SetForest{make([]int, n_components), make([]int, n_components), 0}
  for i := 0; i < n_components; i++ {
    f.root[i] = i
    f.rank[i] = 1
  }
  f.Size = n_components
  return f
}

func (f *SetForest) Find(x int) int {
  if x != f.root[x] {
    f.root[x] = f.Find(f.root[x])
  }
  return f.root[x]
}

func (f *SetForest) SameSet(x, y int) bool {
  return f.Find(x) == f.Find(y)
}

func (f *SetForest) Union(x, y int) {
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
func MSTCostKruskal(n_vertices int, edges [][]int) int {
  sort.Slice(edges, func (i, j int) bool {
    return edges[i][2] < edges[j][2]
  })
  
  min_cost, dsu := 0, NewSetForest(n_vertices)
  for i := 0; dsu.Size > 1; i++ {
    next_edge := edges[i]
    if !dsu.SameSet(next_edge[0], next_edge[1]) {
      dsu.Union(next_edge[0], next_edge[1])
      min_cost += next_edge[2]
    }
  }
  return min_cost
}

// Edge: {from, to, cost}
func MSTCostBoruvka(n_vertices int, edges [][]int) int {
  min_cost, dsu := 0, NewSetForest(n_vertices)
  used, min_edge := make([]bool, len(edges)), make([]int, n_vertices)
  for dsu.Size > 1 {
    for i, _ := range min_edge {
      min_edge[i] = -1
    }
    
    for i, edge := range edges {
      src, dst := dsu.Find(edge[0]), dsu.Find(edge[1])
      if !dsu.SameSet(edge[0], edge[1]) {
        if min_edge[src] == -1 || edge[2] < edges[min_edge[src]][2] {
          min_edge[src] = i
        }
        if min_edge[dst] == -1 || edge[2] < edges[min_edge[dst]][2] {
          min_edge[dst] = i
        }
      }
    }
    
    for _, e := range min_edge {
      if e != -1 && !used[e] {
        dsu.Union(edges[e][0], edges[e][1])
        min_cost += edges[e][2]
        used[e] = true
      }
    }
  }
  return min_cost
}
/////////////////////////////////////////////////////////////////////////////////////////
