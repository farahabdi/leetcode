package triangleNumber

func triangleNumberComb(nums []int) int {
	sofar := []int{}

	matrix := make([][]int, 0)
	matrix = combinations(sofar, nums, 3, matrix)
	result := validateSets(matrix)
	return result
}

func combinations(sofar []int, rest []int, n int, matrix [][]int) [][]int {
	if n == 0 {
			matrix = append(matrix, sofar)
			return matrix
	} else {
			for i := range rest[:len(rest)] {
					concat := sofar
					concat = append(concat, rest[i])
					matrix = combinations(concat, rest[i+1:], n-1, matrix)
			}
	}
	return matrix
}

func validateSets(matrix [][]int) int {
	count :=0
	valid := true
	for _, set := range matrix {
		if set[0] + set[1] <= set[2] {
			valid = false
		}

		if set[0] + set[2] <= set[1] {
			valid = false
		}

		if set[1] + set[2] <= set[0] {
			valid = false
		}

		if valid {
			count++
		}
		valid = true
	}

	return count
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}