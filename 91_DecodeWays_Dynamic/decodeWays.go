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


	if n == 0 {
		return 1
	}

	sum := 0

	for headSize := 1; headSize <= len(s); headSize++ {
		head := s[0:headSize]
		tail := s[headSize:]

		result, _ := strconv.Atoi(head)

		if result > 26 {
			break
		}

		sum = numDecodings(tail) + sum
	}

	return sum
}
