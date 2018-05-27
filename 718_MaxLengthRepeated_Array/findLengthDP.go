package findLength

func findLengthDP(A []int, B []int)int {
	
	ans := 0;
	r := len(A)+1
	c := len(B)+1

	memo := make([][]int, r)

	for i := 0; i < r; i++ {
		memo[i] = make([]int, c)
	}

	for i := len(A) - 1; i >= 0; i-- {
			for j := len(B) - 1; j >= 0; j-- {
					if A[i] == B[j] {
							memo[i][j] = memo[i+1][j+1] + 1

							if ans < memo[i][j] {
								ans = memo[i][j]
							}
					}
			}
	}
	return ans
}