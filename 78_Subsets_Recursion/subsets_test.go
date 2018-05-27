package powerset

import "testing"

func TestPowerSet(t *testing.T) {
	nums := []int{1,2,3}


	//result := subsetsRecursive(nums)
	result := subsetsIterative(nums)
	if result != nil {
		t.Errorf("Expected 3, but it was %d instead.", 3)
	}
}
