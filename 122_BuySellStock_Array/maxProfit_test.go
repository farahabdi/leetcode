package buySellStock

import "testing"

func TestMaxProfit(t *testing.T) {

	nums := []int{2, 1}
	result := maxProfit(nums)

	if result != 3 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
