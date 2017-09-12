package burstBaloons

func maxCoins(nums []int) int {
	n := len(nums)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)
	return maxCoin(1, len(nums)-2, nums, dp)
}

func maxCoin(start int, end int, nums []int, dp [][]int) int {

	if start > end {
		return 0
	}

	if dp[start][end] != 0 {
		return dp[start][end]
	}

	if start == end {
		dp[start][end] = nums[start-1] * nums[start] * nums[start+1]
		return dp[start][end]
	}

	for i := start; i <= end; i++ {
		dp[start][end] = max(dp[start][end],
			nums[start-1]*nums[i]*nums[end+1]+maxCoin(start, i-1, nums, dp)+maxCoin(i+1, end, nums, dp))

	}
	return dp[start][end]

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
