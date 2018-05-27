package dominantIndex

import "testing"

func TestLargestNumber(t *testing.T) {

	nums := []int{3, 6, 1, 0} // 1
	result := dominantIndex(nums)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}

