package dp

// BFS - generalized breadth-first search, applies callback for each state
// bfs level is passed into callback as first argument
// memo func checks whether state has been visited
func BFS[State any](
	callback func(level int, s State, next *[]State),
	init []State,
	memo func(State) bool,
) {
	queue, next := init, make([]State, 0, len(init))
	level := 0

	for len(queue) > 0 {
		for _, s := range queue {
			if memo(s) {
				continue
			}
			callback(level, s, &next)
		}
		queue, next = next, queue[:0]
		level++
	}
}
