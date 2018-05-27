package maximumProduct

import ("sort")

func maximumProduct(nums []int) int {
	N := len(nums)
	sort.Ints(nums)

	firstThree := nums[0] * nums[1] * nums[N - 1]
	lastThree := nums[N - 1] * nums[N - 2] * nums[N - 3]

	return max(firstThree, lastThree)
}



func maxProducts(nums []int) int {
	maxProduct  := nums[0]

	var imin int

	N := len(nums)
	// imax/imin stores the max/min product of
	// subarray that ends with the current number nums[i]
	for i, imax := 1, maxProduct; i < N; imin, i = maxProduct, i+1 {
			// multiplied by a negative makes big number smaller, small number bigger
			// so we redefine the extremums by swapping them
			
			if nums[i] < 0 {
				temp := imax
				imax = temp
				imin = imax
			}
					

			// max/min product for the current number is either the current number itself
			// or the max/min by the previous number times the current one
			imax = max(nums[i], imax * nums[i]);
			imin = min(nums[i], imin * nums[i]);

			// the newly computed max value is a candidate for our global result
			maxProduct = max(maxProduct, imax);
	}
	return maxProduct;

		
}


func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

