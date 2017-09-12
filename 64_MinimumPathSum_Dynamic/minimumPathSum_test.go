package minimumPathSum

import "testing"

func TestMinimumPathSum(t *testing.T) {
	/*grid := [][]int{
		{1, 3, 5, 8},
		{4, 2, 1, 7},
		{4, 3, 2, 3},
	}*/
	grid1 := [][]int{
		{1, 2},
		{1, 1},
	}

	result := minPathSum(grid1)

	if result != 0 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
