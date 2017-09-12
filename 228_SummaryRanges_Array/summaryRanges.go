package summaryRanges

import "strconv"

func summaryRanges(nums []int) []string {

	//	nums := []int{0, 2, 3, 4, 6, 8, 9}
	var result []string

	for start, end := 0, 0; end < len(nums); end++ {

		if end+1 < len(nums) && nums[end]+1 == nums[end+1] {
			continue
		}

		if start == end {
			result = append(result, strconv.Itoa(nums[start]))
		} else {
			result = append(result, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[end]))
		}

		start = end + 1

	}

	//fmt.Println(result)

	return result
}
