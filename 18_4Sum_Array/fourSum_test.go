package fourSum

import "testing"

func TestFourSum(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	result := fourSum(nums, 4)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
