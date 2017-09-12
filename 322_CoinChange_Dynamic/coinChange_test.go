package coinChange

import "testing"

func TestCoinChange(t *testing.T) {
	coins := []int{5, 10, 30, 50, 100}
	amount := 200
	result := coinChange(coins, amount)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
