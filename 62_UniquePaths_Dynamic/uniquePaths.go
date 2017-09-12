package uniquePaths

func uniquePaths(m int, n int) int {

	paths := make([][]int, m)
	for i := range paths {
		paths[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		paths[i][0] = 1
	}

	for i := 0; i < n; i++ {
		paths[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			paths[i][j] = paths[i-1][j] + paths[i][j-1]
		}
	}
	return paths[m-1][n-1]

}
