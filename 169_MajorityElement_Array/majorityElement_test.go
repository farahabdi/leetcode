package majorityElement

import "testing"

func TestMajorityElement(t *testing.T) {

	nums := []int{2, 2, 2, 2, 5, 5, 2, 3, 3}

	result := majorityElement(nums)

	if result != 2 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}

}
