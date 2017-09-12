package combinationSum

import "testing"

func TestCombinationSum(t *testing.T) {

	grid1 := [][]int{
		{1, 2},
		{1, 1},
	}

	result := combinationSum(grid1)

	if result != 0 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
