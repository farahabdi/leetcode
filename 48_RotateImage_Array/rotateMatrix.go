package rotateMatrix

/* 	array := [][]int{
	{5, 1, 9, 11},
	{2, 4, 8, 10},
	{13, 3, 6, 7},
	{15, 14, 12, 16},
} */
func rotate(matrix [][]int) {
	n := len(matrix[0])
	last := n - 1
	levels, totalLevels := 0, n/2

	for levels < totalLevels {
		for i := levels; i < last; i++ {
			matrix[levels][i], matrix[i][last] = matrix[i][last], matrix[levels][i]
			matrix[levels][i], matrix[last][last-i+levels] = matrix[last][last-i+levels], matrix[levels][i]
			matrix[levels][i], matrix[last-i+levels][levels] = matrix[last-i+levels][levels], matrix[levels][i]
		}

		levels++
		last--
	}

	n = n + 1 - 1

}
