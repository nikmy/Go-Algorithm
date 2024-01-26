package dp

import "sort"

// LongestIncreasingSubsequence (LIS) returns length of LIS
// according to the passed order (strict-less func)
//
// O(n) space and O(n log n) time complexity
func LongestIncreasingSubsequence[E any, S ~[]E](order func(x, y E) bool, s S) int {
	lis := make([]E, 1, len(s))
	lis[0] = s[0]

	for _, x := range s {
		if order(lis[len(lis)-1], x) {
			lis = append(lis, x)
		} else {
			insertPos := sort.Search(len(lis), func(i int) bool { return !order(lis[i], x) })
			lis[insertPos] = x
		}
	}
	return len(lis)
}

// LongestCommonSubsequence (LCS) return length of LCS of 2 sequences
//
// O(n^2) space and time complexity
func LongestCommonSubsequence[E comparable, S1 ~[]E, S2 ~[]E](s1 S1, s2 S2) int {
	m, n := len(s1), len(s2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}
