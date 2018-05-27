package matrixReshape


func matrixReshape(nums [][]int, r int, c int) [][]int {

	row := len(nums)
	column := len(nums[0])

	if row *column < r*c {
		return nums
	}

	matrix := make([][]int, r)

	row = 0
	column = 0

	for i := 0; i< len(nums); i++ {
		for j:=0; j < len(nums[0]); j++ {
			matrix[row] = append(matrix[row], nums[i][j])
			column++

			if column == c {
				row++
				column =0
			}
		}
	}
	return matrix    
}

//Can also use modulus  res[count / c][count % c] = nums[i][j]; without if statement