package pairSum

import "testing"

func TestArrayPartition(t *testing.T) {

	nums := []int{1,4,3,2}
	result := arrayPairSum(nums)

	if result == 3 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
