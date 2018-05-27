package jump

func jumpHack(nums []int) int {

	n := len(nums)
	
	if n < 2 {
		return 0
	} 
	level := 0
	currentMax := 0
	i := 0
	nextMax := 0

	for currentMax-i+1 > 0 {	
		level++;
		for ;i<=currentMax; i++ {	
		 nextMax = max(nextMax, nums[i]+i)

		 if nextMax >= n-1 {
			 return level
			}

		}
		currentMax=nextMax;
	}
	return 0
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}