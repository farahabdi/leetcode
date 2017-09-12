package containsDuplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {

	nums := []int{2, 1}
	result := containsDuplicate(nums)

	if result != false {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
