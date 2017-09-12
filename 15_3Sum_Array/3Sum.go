package threeSum

/*
For example, given array S = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/
func threeSum(nums []int) [][]int {

	var result [][]int
	var l int

	for i := 0; i < len(nums)-2; i++ {

		for j := i + 1; j < len(nums)-1; j++ {

			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}

		}
	}

	//result = append(result, []int{4, 2, 3})

	return result

}
