package maxSubarray

func maxSubArray(nums []int) int {
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum := nums[0]
	currentSum := 0
	for i := range nums {
		currentSum = currentSum + nums[i]

		if currentSum > maxSum {
			maxSum = currentSum
		}

		if currentSum < 0 {
			currentSum = 0
		}
	}

	return maxSum
}
