package draft

/*
HasCycle implementation

	color:
	  0: not completely processed, not visited
	  1: not completely processed, visited
	  2: processed
*/
func hasCycleDFS(vertex int, neighbors [][]int, color []int) bool {
	if color[vertex] == 2 {
		return false
	}
	if color[vertex] == 1 {
		return true
	}
	color[vertex] = 1
	for _, n := range neighbors[vertex] {
		if hasCycleDFS(n, neighbors, color) {
			return true
		}
	}
	color[vertex] = 2
	return false
}

func HasCycle(neighbors [][]int) bool {
	color := make([]int, len(neighbors))
	for v := 0; v < len(neighbors); v++ {
		if color[v] == 2 {
			continue
		}
		if hasCycleDFS(v, neighbors, color) {
			return true
		}
	}
	return false
}

func NeighborsFromEdgesList(numCourses int, prerequisites [][]int) [][]int {
	adj := make([][]int, numCourses)
	for _, edge := range prerequisites {
		from, to := edge[1], edge[0]
		adj[from] = append(adj[from], to)
	}
	return adj
}
