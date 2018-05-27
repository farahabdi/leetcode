package maxAreaOfIsland

import (
	"testing"
)
func TestMaxAreaOfIsland(t *testing.T) {
/*
	nums := [][]int{
		{0,0,1,0,0,0,0,1,0,0,0,0,0},
 		{0,0,0,0,0,0,0,1,1,1,0,0,0},
 		{0,1,1,0,1,0,0,0,0,0,0,0,0},
 		{0,1,0,0,1,1,0,0,1,0,1,0,0},
 		{0,1,0,0,1,1,0,0,1,1,1,0,0},
 		{0,0,0,0,0,0,0,0,0,0,1,0,0},
 		{0,0,0,0,0,0,0,1,1,1,0,0,0},
 		{0,0,0,0,0,0,0,1,1,0,0,0,0},
	}
	*/
	nums := [][]int{
		{0,1},
 		{1,0},
	}

	result := maxAreaOfIsland(nums)

	if result == 3 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
