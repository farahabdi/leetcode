package searchInsert

func searchInsert(nums []int, target int) int {

	N := len(nums)


	if nums[N-1] < target {
		return N
	}
	
	for i, value := range nums {
		if value >= target {
			return i 
		}
	}
 
	return 0
}