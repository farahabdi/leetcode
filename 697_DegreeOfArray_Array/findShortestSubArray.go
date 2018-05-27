package degreeArray

import ("math")
func findShortestSubArray(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	if len(nums) == 1 {
		return 1
	}

	dict := make(map[int]int, 0)
	maxValue := math.MinInt32
	maxDegree := math.MinInt32

	for _, value := range nums {
		dict[value] += 1
		maxDegree = max(maxDegree, dict[value])
		maxValue = max(maxValue, value)
	}

	if maxDegree == 1 {
		return 1
	}

	matrix := make([][]int, maxValue+1)

	for i, value := range nums {
		if dict[value] == maxDegree {
			matrix[value] = append(matrix[value], i)
		}
	}

	minLength := math.MaxInt32



	for i :=0; i < len(matrix); i++  {
		if len(matrix[i]) > 1 {
			minLength = min(minLength, matrix[i][len(matrix[i])-1]-matrix[i][0])
		}
	}



	return minLength+1
    
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}