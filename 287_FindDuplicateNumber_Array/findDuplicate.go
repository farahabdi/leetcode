package findDuplicate

import ("sort")
func findDuplicate(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	
	sort.Ints(nums)

	for i, value := range nums {
		if value == nums[i+1] {
			return i
		}
	}

	return -1
}
