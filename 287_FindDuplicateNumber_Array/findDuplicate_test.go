package findDuplicate

import "testing"

func TestCombinationSum(t *testing.T) {

	nums := []int{1, 6, 4, 7, 3, 5, 0, 8, 3, 9}
	result := findDuplicate(nums)

	if result != 0 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
