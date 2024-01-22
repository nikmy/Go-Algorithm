package dp

// BFS - generalized breadth-first search, applies callback for each state
// bfs level is passed into callback as first argument
func BFS[State comparable](
	callback func(level int, s State, next *[]State) bool,
	init []State,
) {
	queue, next := init, make([]State, 0, len(init))
	visited := make(map[State]struct{}, len(init))
	level := 0

	for len(queue) > 0 {
		for _, s := range queue {
			if _, ok := visited[s]; ok {
				continue
			}
			visited[s] = struct{}{}
			callback(level, s, &next)
		}
		queue, next = next, queue[:0]
	}
}
