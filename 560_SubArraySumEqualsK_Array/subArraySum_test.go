package subArraySum

import "testing"

func TestSubArraySum(t *testing.T) {
	coins := []int{5, 10, 30, 50, 100}

	result := subarraySum(coins, 2)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
