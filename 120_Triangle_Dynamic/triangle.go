package triangle

/* [
     [2],      [2]
    [3,4],	   [3,4]
   [6,5,7],    [6,5,7]
  [4,1,8,3]    [4,1,8,3]
]              [4,1,8,3,5]
			   [4,1,8,3,5,4]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
*/

func minimumTotalDP(matrix [][]int) int {

	for i := len(matrix) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			matrix[i][j] = min(matrix[i+1][j], matrix[i+1][j+1]) + matrix[i][j]
		}
	}

	return matrix[0][0]
}

func minimumTotalRecursive(matrix [][]int, m int, n int) int {

	if m >= len(matrix)-1 {
		return matrix[m][n]
	}

	return matrix[m][n] + min(minimumTotalRecursive(matrix, m+1, n), minimumTotalRecursive(matrix, m+1, n+1))
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
