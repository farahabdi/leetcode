package pairSum
import ("sort")
func arrayPairSum(nums []int) int {

	//pairs := len(nums)/2
	sum := 0
	start := 0
	sort.Ints(nums)
	for i := start; i <= len(nums)-2; i +=2 {
		sum += min(nums[i], nums[i+1])
	
	}

	return sum
    
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}