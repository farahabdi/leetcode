package containsDuplicate

import "testing"

func TestContainsDuplicates(t *testing.T) {

	nums1 := []int{1, 2, 3, 1}
	result := containsNearbyDuplicate(nums1, 3)

	if !result {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
