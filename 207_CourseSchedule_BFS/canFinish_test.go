package canFinish

import "testing"

func TestCanFinish(t *testing.T) {
	nums := [][]int{{1,0}, {0, 1}}
	result := canFinish(2, nums)

	if result != false {
		t.Errorf("Expected 2, but it was %d instead.", result)
	}
}
