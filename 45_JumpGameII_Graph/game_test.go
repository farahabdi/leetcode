package jump

import "testing"

func TestJumpGame(t *testing.T) {

	nums := []int{2,3,0,1,4}
	result := jump(nums)

	if result !=2 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
