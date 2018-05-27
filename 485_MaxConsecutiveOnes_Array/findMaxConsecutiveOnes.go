package findMaxConsecutiveOnes



func findMaxConsecutiveOnes(nums []int) int {

	if len(nums) < 2 {
		return nums[0]
	}
	count := 0
	maxCount := 0

	for _, value := range nums {
		if value == 1 {
			count++
		} else {
			count = 0
		}
		maxCount = max(maxCount, count)
	}
	return maxCount
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}