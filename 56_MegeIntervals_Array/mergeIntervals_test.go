package mergeIntervals

import "testing"

type Interval struct {
	Start int
	End   int
}

func TestMergeIntervals(t *testing.T) {

	intervals := []Interval{Interval{1, 4}, Interval{0, 0}}
	result := merge(intervals)

	if result != nil {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
