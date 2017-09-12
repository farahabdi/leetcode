package maximalSquare

func maximalSquare(matrix [][]byte) int {

	m := len(matrix)

	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	value := 0
	if n == 1 && m == 1 {
		if matrix[0][0] == '1' {
			value = 1
		}
		return value
	}

	dp := make([][]int, m+1)

	for k := 0; k < m+1; k++ {
		dp[k] = make([]int, n+1)
	}

	max := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {

			if matrix[i-1][j-1] == '1' {
				cMin := min(dp[i][j-1], dp[i-1][j])
				dp[i][j] = min(dp[i-1][j-1], cMin) + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}

		}
	}

	return max * max
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
