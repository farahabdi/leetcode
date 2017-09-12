package findPeakElement

import "testing"

func TestFindPeakElement(t *testing.T) {
	coins := []int{5, 10, 30, 50, 100}

	result := findPeakElement(coins)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
