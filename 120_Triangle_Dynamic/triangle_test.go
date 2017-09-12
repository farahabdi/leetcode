package triangle

import "testing"

func TestPerfectSquares(t *testing.T) {

	array := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	result := minimumTotal(array, 0, 0)

	if result != 11 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
