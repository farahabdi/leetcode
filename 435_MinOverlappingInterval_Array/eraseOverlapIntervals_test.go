package eraseOverlapIntervals

import "testing"

func TestOverlapIntervals(t *testing.T) {

	nums := [][]int{{1, 6}, {4, 7}, {3, 5}}
	result := eraseOverlapIntervals(nums)

	if result != 0 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
