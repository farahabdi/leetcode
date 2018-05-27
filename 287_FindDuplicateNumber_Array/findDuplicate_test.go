package findDuplicate

import "testing"


func TestCombinationSum(t *testing.T) {

	nums := []int{1,3,4,2,2}
	result := findDuplicate(nums)

	if result != 0 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
