package maxRectangle

func maximalRectangle(matrix [][]byte) int {

	/*	array := [][]byte{
			{1, 0, 1, 0, 0},
			{1, 0, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 0, 0, 1, 0},
		}
	*/

	temp := [5]int{}
	count := 0
	result := 0
	maxResult := 0

	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {

			if matrix[i][j] == 0 {
				temp[j] = 0
			} else {
				temp[j] = temp[j] + 1
			}

		}
		result = findMax(temp)

		if result > maxResult {
			maxResult = result
		}
		count++
	}

	return maxResult

}

func findMax(temp [5]int) int {

	lcd := 0
	count := 0
	start := 0
	result := 0
	maxResult := 0
	maxCount := 0
	for i := 0; i < 5; i++ {
		if temp[i] != 0 && lcd == 0 {
			lcd = temp[i]
			start = i
		}

		if temp[i] > 0 {
			count = count + 1

		}

		if i == 4 || temp[i] == 0 {

			for j := start; j < start+count; j++ {
				result = gcd(temp[j], temp[start])
			}
			start = i + 1
			if result > maxResult {
				maxResult = result
				maxCount = count
			}
			count = 0
			lcd = 0
		}

	}

	return maxResult * maxCount
}

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
