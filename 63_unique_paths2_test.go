package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}

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
