package perfectSquares

func numSquares(n int) int {

	size := n
	if n <= 3 {
		size = 3
	}
	dp := make([]int, size+1)

	for i := 0; i <= 3; i++ {
		dp[i] = i
	}

	for i := 4; i <= n; i++ {
		dp[i] = i

		for j := 1; j <= i; j++ {
			temp := j * j
			if temp > i {
				break
			} else {
				dp[i] = min(dp[i], dp[i-temp]+1)
			}
		}

	}

	return dp[n]

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
