package containsDuplicate

func containsNearbyDuplicate(nums []int, k int) bool {

	for i := 0; i < len(nums); i++ {
		for j := max(i-k, 0); j < i; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
