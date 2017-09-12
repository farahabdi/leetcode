package plusOne

func plusOne(digits []int) []int {

	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		}

		digits[i] = 0
	}

	digits = append(digits, 0)

	digits[0] = 1

	return digits
}
