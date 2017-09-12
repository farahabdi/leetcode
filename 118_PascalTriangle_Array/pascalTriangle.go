package pascalTriangle

import "fmt"

/*
numRows = 5
[
[1],
[1,1],
[1,2,1],
[1,3,3,1],
[1,4,6,4,1]
]
*/

func generate(numRows int) [][]int {

	matrix := make([][]int, numRows)

	for i := 1; i <= numRows; i++ {
		matrix[i-1] = make([]int, i)
	}

	for i := 0; i < numRows; i++ {
		matrix[i][0] = 1
		matrix[i][i] = 1
	}

	for i := 2; i <= numRows-1; i++ {
		for j := 1; j < i; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i-1][j-1]
			print2Array(matrix)
		}
	}

	return matrix
}

func print2Array(array [][]int) {
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d", array[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
