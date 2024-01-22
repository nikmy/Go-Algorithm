package draft

// Generalized DFS
//
//	@param base - checks for base case and return pair (baseCaseValue, isBaseCase)
//	@param memo - storage for results for visited states
//	@param calc - function that calculates result using abstract dfs function and current state
//
//	@return dfs - recursive DFS function
func DFS[State comparable, Value any](
	base func(State) (Value, bool),
	memo map[State]Value,
	calc func(func(State) Value, State) Value,
) (dfs func(State) Value) {
	dfs = func(s State) Value {
		if v, ok := base(s); ok {
			return v
		}
		if v, ok := memo[s]; ok {
			return v
		}
		memo[s] = calc(dfs, s)
		return memo[s]
	}
	return dfs
}

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
