package maxProfitf

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)

	dp := make([][]int, 3)
	for k := range dp {
		dp[k] = make([]int, n)
	}

	for i := 1; i < 3; i++ {

		for j := 1; j < n; j++ {
			maxSum := -10000
			for m := 0; m <= j-1; m++ {
				if dp[i-1][m]-prices[m] > maxSum {
					maxSum = dp[i-1][m] - prices[m]
				}

			}

			dp[i][j] = max(dp[i][j-1], prices[j]+maxSum)

		}
	}

	return dp[2][n-1]

}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
