package isOneBitCharacter

func isOneBitCharacter(bits []int) bool {
	
	var i int
	N := len(bits)
	for i < N-1 {
		i = bits[i] + i + 1

	}
	return i == N-1
}