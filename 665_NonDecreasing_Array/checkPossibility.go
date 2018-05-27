package nonDecreasing

/*
Detect the decreasing pair in the whole array.

If the decreasing pair happened more than one time, return false.
If no decreasing pair, then it is already the non-decreasing array, return true.
If the decreasing pair happened just one time, there are two ways to do that. Make the bigger one smaller, make the smaller one bigger. Please see the return statement with comment.

*/

func checkPossibility(nums []int) bool {

	index := -1;  // index is where decreasing pair happened
	hasHappened := false;
	
	for i := 1; i < len(nums); i++ {
			if nums[i - 1] > nums[i] {
					if hasHappened { 
						return false  
						}  // need to modify more than one time
					hasHappened = true
					index = i
			}
	}

	alreadyDecreasing := index == -1
	makeSmallerOrBigger := index == len(nums) - 1 || nums[index + 1] >= nums[index - 1]
	makeSmallerminus := index <= 1 || nums[index] >= nums[index - 2]

	result := alreadyDecreasing ||  makeSmallerOrBigger || makeSmallerminus
	
	/*index == -1 // it's already non-decreasing
	|| index == 0 || index == nums.length - 1 || nums[index + 1] >= nums[index - 1]  // make index (smaller one) bigger
	|| index <= 1 || nums[index] >= nums[index - 2];  // make index - 1 (bigger one) smaller */

	return result
    

}