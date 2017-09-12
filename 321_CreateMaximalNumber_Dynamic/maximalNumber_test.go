package maximalNumber

import "testing"

func TestMaximalNumber(t *testing.T) {
	nums1 := []int{5, 10, 30, 50, 100}
	nums2 := []int{5, 10, 30, 50, 100}
	k := 2
	result := maxNumber(nums1, nums2, k)

	if result != nil {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
