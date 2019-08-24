package leetcode

func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}

	if m == 1 {
		return uniquePaths(m, n-1)
	} else if n == 1 {
		return uniquePaths(m-1, n)
	}
	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}

func uniquePathsDP(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i+1 < m {
				dp[i][j] = dp[i+1][j]
			}
			if j+1 < n {
				dp[i][j] += dp[i][j+1]
			}
		}
	}

	return dp[0][0]
}
