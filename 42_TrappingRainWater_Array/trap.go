package trap

func trap(height []int) int {
	left, right := 0, len(height) - 1
	ans, left_max, right_max := 0, 0, 0

	for left < right {
		if  height[right] > height[left] {
			if height[left] >= left_max {
				left_max = height[left] 
			} else {
				ans += left_max - height[left]
			}
			left++
		} else {
			if height[right] >= right_max {
			   right_max = height[right]
			} else {
			   ans += right_max - height[right]
			}
			right--
		}
	}
	return ans
}

func trapBrute(height []int) int {
    ans := 0;
    size := len(height)
    for i := 1; i < size - 1; i++ {
		maxLeft := 0
		maxRight := 0
        for j := i; j >= 0; j-- { //Search the left part for max bar size
            maxLeft = max(maxLeft, height[j])
        }
        for j := i; j < size; j++ { //Search the right part for max bar size
            maxRight = max(maxRight, height[j])
        }
        ans += min(maxLeft, maxRight) - height[i]
    }
    return ans
}

func trapDynamic(height []int) int {

	if height == nil {
		return 0
	}
	
    ans := 0;
    size := len(height)
	leftMax := []int{}
	rightMax := []int{}
	leftMax[0] = height[0]
	
    for i := 1; i < size; i++ {
        leftMax[i] = max(height[i], leftMax[i - 1])
    }
	rightMax[size - 1] = height[size - 1]
	
    for i := size - 2; i >= 0; i-- {
        rightMax[i] = max(height[i], rightMax[i + 1])
    }
    for i := 1; i < size - 1; i++ {
        ans += min(leftMax[i], rightMax[i]) - height[i]
    }
    return ans

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