package dp

// DFS - generalized depth-first search, returns recursive DFS function
//	base - checks for base case and return pair (baseCaseValue, isBaseCase)
//	memo - storage for results for visited states
//	calc - function that calculates result using abstract dfs function and current state
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
