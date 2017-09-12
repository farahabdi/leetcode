package arrayNesting

import "testing"

func TestArrayNesting(t *testing.T) {
	coins := []int{5, 4, 0, 3, 1, 6, 2}

	result := arrayNesting(coins)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
