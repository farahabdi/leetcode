package nonDecreasing

import "testing"

func TestNonDecreasing(t *testing.T) {
	coins := []int{3,4,2,13}

	result := checkPossibility(coins)

	if result != false {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
