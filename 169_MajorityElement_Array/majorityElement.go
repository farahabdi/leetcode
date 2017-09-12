package majorityElement

//nums := []int{2, 2, 2, 2, 5, 5, 2, 3, 3}

func majorityElement(nums []int) int {

	cache := make(map[int]int, 10)

	for index, val := range nums {
		cache[nums[index]] = cache[nums[index]] + 1

		if cache[nums[index]] > len(nums)/2 {
			return val
		}

	}
	return -1
}
