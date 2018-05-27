package pivotIndex

import (
	"math"
)

func pivotIndex(nums []int) int {

	if len(nums) < 1  {
		return -1
	}

	left := make(map[int]int, len(nums))
	right := make(map[int]int, len(nums))
	j := len(nums)-1

	left[0] = nums[0]
	right[0] = nums[j]
	j--

	for index :=1; index < len(nums)-1; index++ {
		left[index] = left[index-1] + nums[index]
		right[index] = right[index-1] + nums[j]
		j--
	}

	j = 0
	n := len(nums)-1

	if right[n-1] == 0 {
		return 0
	}
	minKey := math.MaxInt32
	for key, value := range left {
		if right[n-key-2] == value {
			minKey = min(minKey, key+1)
		}
	}


	if (minKey == math.MaxInt32) {
		minKey = -1
	}
	return minKey
}

func min(a int, b int) int {
	 if a < b {
		 return a
	 }
	 return b
}