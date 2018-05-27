package numIslands


import "testing"

func TestNumIslands(t *testing.T) {
	
	

	grid1 := [][]byte{
		{'1','1','0','0','0'},
		{'1','1','0','0','0'},
		{'0','0','1','0','0'},
		{'0','0','0','1','1'},
	}


	grid2 := [][]byte{
		{'1','1','1','1','0'},
		{'1','1','0','1','0'},
		{'1','1','0','0','0'},
		{'0','0','0','0','0'},
	}

	result := numIslands(grid1)

	result2 := numIslands(grid2)

	if result != 3 && result2 != 3 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}

}
