package trap

import "testing"

func TestTrapRainWater(t *testing.T) {
	nums := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	result := trap(nums)

	if result != 6 {
		t.Errorf("Expected 6s, but it was %d instead.", result)
	}
}
