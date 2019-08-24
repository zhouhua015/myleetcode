package leetcode

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = grid[m-1][n-1]

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			var min int
			if i+1 < m {
				min = dp[i+1][j]
			}
			if j+1 < n && (i+1 >= m || min > dp[i][j+1]) {
				min = dp[i][j+1]
			}

			dp[i][j] = min + grid[i][j]
		}
	}

	return dp[0][0]
}
