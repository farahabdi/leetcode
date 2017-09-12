package maxSubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(nums)

	if result != 6 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
