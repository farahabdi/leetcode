package powerset


func subsetsIterative(nums []int) [][]int {
	empty := []int{}
	sets := [][]int{empty}
	for _, n := range nums {
			newSets := make([][]int,0)
			for _, set := range sets {
					tmp := make([]int, len(set)+1)
					copy(tmp, append(set,n))
					newSets = append(newSets, tmp)
			}
			sets = append(sets, newSets...)
	}
	return sets
}


func subsetsRecursive(nums []int) [][]int {
	var matrix = make([][]int, 0)
	var tmp = make([]int, 0)
	backtrack(&matrix, tmp, nums, 0)

	return matrix
}

func backtrack(matrix *[][]int, tmp []int, nums []int, start int) {
	var subset = make([]int, len(tmp))
	copy(subset, tmp)
	*matrix = append(*matrix, subset)

	for i := start; i < len(nums); i++ {
			tmp = append(tmp, nums[i])
			backtrack(matrix, tmp, nums, i+1)
			tmp = tmp[:len(tmp)-1]
	}
}
