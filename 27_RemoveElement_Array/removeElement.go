package removeElement

func removeElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		v := nums[i]
		if v == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}
