package containsDuplicate

func containsDuplicate(nums []int) bool {

	m := make(map[int]int)

	for i := range nums {
		m[nums[i]]++
	}

	for j := range m {
		if m[j] >= 2 {
			return true
		}
	}

	return false

}
