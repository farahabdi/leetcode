package subsetsII

import "testing"

func TestSubsetsWithDup(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	result := subsetsWithDup(nums)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
