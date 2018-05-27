package maximumProduct

import "testing"

func TestMaximumProduct(t *testing.T) {


	//nums := []int{1,2,3,4}
	//nums := []int{2,3,-2,4}
	//nums := []int{-2,0,1}
	//nums := []int{3, -1, 4}
	nums := []int{-4,-3,-2,-1,60}
	//nums := []int{2,-1,-3,-3,1}
	result := maximumProduct(nums)

	if result != 2 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
