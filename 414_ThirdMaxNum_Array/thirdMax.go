package thirdMax


func thirdMax(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}


	dict := make(map[int]int, 0)

	var maxOne *int
	var maxTwo *int
	var maxThree *int


	for _, value := range nums {
		
		if dict[value] > 0 {
			continue
		}
		if maxOne == nil || value >= *maxOne {
			
			maxThree = maxTwo
			maxTwo = maxOne
			maxOne = new(int)
			*maxOne = value
			
		}else if maxTwo == nil || value >= *maxTwo{
			maxThree = maxTwo
			maxTwo = new(int)
			*maxTwo = value
		} else if maxThree == nil || value >= *maxThree {
			maxThree = new(int)
			*maxThree = value
		}

		dict[value] += 1
	}

	if len(dict) == 1 {
		return nums[0]
	}

	if maxThree == nil {
		return *maxOne
	}
	return *maxThree
    
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}