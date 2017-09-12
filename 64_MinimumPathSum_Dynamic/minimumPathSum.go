package minimumPathSum

func minPathSum(grid [][]int) int {

	if len(grid[0])-1 == 0 {
		return grid[0][0]
	}

	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
