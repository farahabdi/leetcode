package findLength

import (

	"strconv"
	"strings"
)

func findLength(A []int, B []int) int {

	lo := 0
	hi := min(len(A), len(B)) + 1

	for lo < hi {
			mi := (lo + hi) / 2;
			if check(mi, A, B) {
				lo = mi + 1;
			} else {
				hi = mi
			}
	}
	return lo - 1

}

func check(length int, A []int, B []int) bool  {

	seen := make(map[string]bool, 0)

	for i := 0; i + length <= len(A); i++ {

		subarray := subarrayToString(A[i:i+length])
		seen[subarray] = true

	}
	for j := 0; j + length <= len(B); j++ {
			subarray := subarrayToString(B[j:j+length])
			if seen[subarray] {
					return true
			}
	}
	return false
}

func subarrayToString(array []int) string {
	valuesText := []string{}

	for i := range array {
		number := array[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}
	result := strings.Join(valuesText, "")

	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}