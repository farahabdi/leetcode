package threeSum

import "testing"

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, 4}
	result := threeSum(nums)

	if result != nil {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
