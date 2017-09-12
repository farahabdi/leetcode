package jumpGame

import "testing"

func TestJumpGame(t *testing.T) {

	nums := []int{-1, 0, 1, 2, -1, 4}
	result := canJump(nums)

	if result {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
