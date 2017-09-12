package findMin

//	nums := []int{4, 5, 6, 13, 14, 15, 16, 17, 18,19}
//	nums := []int{4, 5, 6, 7, 8, 15, 16, 17, 18,19}
//	nums := []int{4, 5, 6, 7, 8, 9, 15, 16, 17,18}

//	nums := []int{14, 15, 16, 17, 18, 19, 20, 1, 2, 3}
func findMin(nums []int) int {
	r := len(nums) - 1
	l := 0

	for l < r {

		if nums[l] < nums[r] {
			return nums[l]
		}
		mid := (l + r) / 2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}


