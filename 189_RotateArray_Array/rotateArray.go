package rotateArray

// 	nums := []int{1, 2, 3, 4, 5, 6, 7}

func rotateReverse(nums []int, k int) {

	n := len(nums)
	k %= n
	if n < k {
		return
	}
	rotateMoudlus(nums, 3)
	//	reverse(nums, 0, n-1)
	//	reverse(nums, 0, k-1)
	//	reverse(nums, k, n-1)

}

func rotateMoudlus(nums []int, k int) {

	n := len(nums)

	for i := 0; i < n; i++ {
		nums[(i+k)%n], nums[i] = nums[i], nums[(i+k)%n]
	}

	n = n + 1

}

func reverse(nums []int, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
