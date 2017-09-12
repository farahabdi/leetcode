package summaryRanges

import "testing"

func TestSummaryRanges(t *testing.T) {

	nums := []int{0, 2, 3, 4, 6, 8, 9}
	result := summaryRanges(nums)

	if result != nil {
		t.Errorf("Expected 3, but it was %s instead.", result)
	}
}
