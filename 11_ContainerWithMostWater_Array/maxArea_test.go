package ContainerWithMostWater

import "testing"

func TestContainerWithMostWater(t *testing.T) {

	result := maxArea(coins, amount)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
