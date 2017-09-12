package plusOne

import "testing"

func TestPlusOne(t *testing.T) {

	nums := []int{4, 3, 2}

	result := plusOne(nums)

	if result != nil {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}

}
