package maximumSwap

import "testing"

func TestMaximumSwap(t *testing.T) {
	coins := []int{5, 10, 30, 50, 100}

	result := maximumSwap(44)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
