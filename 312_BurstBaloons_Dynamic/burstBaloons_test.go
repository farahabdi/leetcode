package burstBaloons

import "testing"

func TestBurstBaloons(t *testing.T) {
	nums := []int{3, 1, 5, 8}
	result := maxCoins(nums)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
