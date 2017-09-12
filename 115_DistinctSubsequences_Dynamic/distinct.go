package distinctSubsequence

// RecursiveSolution
func numDistinctRecursive(s string, t string) int {

	if len(t) == 0 {
		return 1
	} else if len(s) == 0 {
		return 0
	}

	if s[len(s)-1] == t[len(t)-1] {
		return numDistinctRecursive(s[:len(s)-1], t) + numDistinctRecursive(s[:len(s)-1], t[:len(t)-1])

	}
	return numDistinctRecursive(s[:len(s)-1], t)

}
func numDistinctDP(s string, t string) int {

	matrix := make([][]int, len(t)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(s)+1)
	}

	for i := 1; i <= len(t); i++ {
		matrix[i][0] = 0
	}

	for i := 0; i <= len(s); i++ {
		matrix[0][i] = 1
	}

	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {

			if s[j-1] != t[i-1] {
				matrix[i][j] = matrix[i][j-1]
			} else {
				matrix[i][j] = matrix[i][j-1] + matrix[i-1][j-1]
			}
		}
	}

	return matrix[len(t)][len(s)]

}
