package eraseOverlapIntervals

import (
	"sort"
)

type Interval struct {
	  Start int
	  End   int
}
	
func eraseOverlapIntervals(intervals []Interval) int {

	count := 0
	if len(intervals) == 0 {
		return 0
	} else if len(intervals) < 2 {
		return 1
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	newMin := intervals[0].End

	for i, interval := range intervals[:len(intervals)] {
		currentStart, currentEnd := interval.Start, interval.End
		nextStart, nextEnd := intervals[i+1].Start, intervals[i+1].End

		// Check if start of current is more than end of last interval
		if currentStart > intervals[i+1].Start {

			if intervals[i+1].End < currentEnd {
				newMin = currentEnd
				continue
			} else if intervals[i+1].End > currentEnd && currentStart > newMin{
				newMin++
				newMin--
				count++
				continue
			}

			newMin = intervals[i+1].Start

		} 








	}

    return 0
}