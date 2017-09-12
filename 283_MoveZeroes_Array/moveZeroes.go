package moveZero

func moveZeroes(nums []int) {

	count := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}

	}

	for j := count; j < n; j++ {
		nums[count] = 0
		count++
	}

}
