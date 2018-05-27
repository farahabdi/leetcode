package jumpGame

import "testing"

func TestJumpGame(t *testing.T) {

	nums := []int{2,3,1,1,4}
	result := canJumpBacktrack(nums)

	if result {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
