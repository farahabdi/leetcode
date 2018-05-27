package largeGroupPositions


// nums := "abbxxxxzzy" // 
// nums := "aabxxxxzzy" // 1
// nums := "abcdddeeeeaabbbcd"
func largeGroupPositions(S string) [][]int {

	if len(S) < 3 {
		return [][]int{}
	}

	str := []rune(S)
	i := 0
	groups := [][]int{}
	N := len(str)

	for j :=0; j < N; j++ {

		if j == N-1 || str[j] != str[j+1] {
			if j-i+1 >= 3 {
				groups = append(groups, []int{i, j})
			}
			i = j + 1
		}
	}
    return groups
}