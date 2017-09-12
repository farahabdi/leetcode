package interleavingString

func isInterleave(s1 string, s2 string, s3 string) bool {
	return is_interleave(s1, 0, s2, 0, "", s3)
}

func is_interleave(s1 string, i int, s2 string, j int, res string, s3 string) bool {

	if res == s3 && i == len(s1) && j == len(s2) {
		return true
	}

	ans := false
	if i < len(s1) {
		ans = ans || is_interleave(s1, i+1, s2, j, res+string(s1[i]), s3)
	}

	if j < len(s2) {
		ans = ans || is_interleave(s1, i, s2, j+1, res+string(s2[j]), s3)
	}

	return ans
}
