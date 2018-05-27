
package combinationSum2
import (

	"strconv"
	"strings"
	"sort"
)
func combinationSum2(candidates []int, target int) [][]int {

	subsets := make([][]int, 0)
	sort.Ints(candidates)
	subsets= subsetsRecursive(candidates)

	var tmp = make([][]int, 0)
	for i :=0; i < len(subsets); i++ {
		result := addsToTarget(subsets[i], target)
		if result {
			tmp = append(tmp, subsets[i])
		}
	}

	for i :=0; i < len(subsets); i++ {
		if i >
	}



	return tmp

}

func removeDuplicateSets(subsets [][]int) [][]int {
	encountered := make(map[string]bool, 0)
	result := [][]int{}

	for i :=0; i < len(subsets); i++ {
		subarray := subarrayToString(subsets[i])
		if !encountered[subarray] == true {

			encountered[subarray] = true
			result = append(result, subsets[i])
		}
	}
	return result
}

func addsToTarget(subset []int, target int) bool {
	sum := 0
	for i :=0; i < len(subset); i++ {
		sum += subset[i]
	}
	return sum == target
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