package triangleNumber

import "testing"

func TestTriangleNumber(t *testing.T) {
	coins := []int{1,2,3,4}

	result := triangleNumber(coins)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
