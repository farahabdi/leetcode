package removeElement

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	result := removeElement(nums, 3)

	if result != 2 {
		t.Errorf("Expected 2, but it was %d instead.", result)
	}
}
