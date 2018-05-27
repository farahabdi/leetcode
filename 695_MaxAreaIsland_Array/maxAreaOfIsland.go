package maxAreaOfIsland


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
func maxAreaOfIsland(grid [][]int) int {


	maxRegion := 0

	for i :=0; i < len(grid); i++ {
		for j :=0; j < len(grid[0]); j++  {

			if grid[i][j] == 1 {
				size := dfs(grid, i, j)
				maxRegion = max(maxRegion, size)
			}
		}
	}

	return maxRegion
    
}

func dfs(grid [][]int, row int, column int) int {

	if row < 0 || column < 0 || row >= len(grid) || column >= len(grid[0]) {
		return 0
	}

	if grid[row][column] == 0 {
		return 0
	} 

	grid[row][column] = 0

	return (1 + dfs(grid, row+1, column) + dfs(grid, row-1, column)+ dfs(grid, row, column-1) + dfs(grid, row, column+1))

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}