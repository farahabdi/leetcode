package findMin

import "testing"

func TestFindMin(t *testing.T) {
	nums := []int{2, 1}
	result := findMin(nums)

	if result != 7 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
