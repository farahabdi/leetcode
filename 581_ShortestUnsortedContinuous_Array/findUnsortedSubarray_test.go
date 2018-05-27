package findUnsortedSubarray

import (
	"testing"
)
func TestMaxAreaOfIsland(t *testing.T) {

	nums := []int{2, 6, 4, 8, 10, 9, 15}
	//nums := []int{1,2,3,4}
	//nums := []int{5,4,3,2,1}
	//nums := []int{1,3,2}
	 //nums := []int{1,3,2,2,2} //4
	//nums := []int{1,2}
	//nums := []int{1,2,3,3,3}
	
	result := findUnsortedSubarray(nums)

	if result == 3 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}