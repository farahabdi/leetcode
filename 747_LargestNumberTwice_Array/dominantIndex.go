
package dominantIndex

import ("math")
func dominantIndex(nums []int) int {
	var index int
	maxNumber := math.MinInt32

	for i, value := range nums {
		if value > maxNumber {
			maxNumber = value
			index = i
		}
	}

	minValue := maxNumber / 2

	for _, value := range nums {
		if value != maxNumber && value > minValue{
			return -1
		}
	}
	return index
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}