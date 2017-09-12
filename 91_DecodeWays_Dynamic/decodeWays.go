package decodeWays

import "strconv"

func numDecodings(s string) int {
	n := len(s)

	if n == 0 {
		return 1
	}

	if s[0] == '0' {
		return 0
	}

	upperLimit := 26

	if n == 0 {
		return 1
	}

	sum := 0

	for headSize := 1; headSize <= len(s); headSize++ {
		head := s[0:headSize]
		tail := s[headSize:]

		result, err := strconv.Atoi(head)

		if err != nil {

		}
		if result > upperLimit {
			break
		}

		sum = numDecodings(tail) + sum
	}

	return sum
}
