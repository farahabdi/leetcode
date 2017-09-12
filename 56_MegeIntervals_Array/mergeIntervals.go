package mergeIntervals

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
// intervals[interval+1].Start < intervals[interval].Start &&
// intervals[interval+1].End < intervals[interval].End
// {1, 4}..{0, 1} == {0, 4}
// {1, 4}..{0, 0} == {0, 0} { 1, 4}

func merge(intervals []Interval) []Interval {

	for interval := 0; interval < len(intervals)-1; interval++ {

	}

	return intervals

}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
