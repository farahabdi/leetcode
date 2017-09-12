package maxProduct

import "testing"

func TestMaxProduct(t *testing.T) {
	array := [5]int{-2, 4, 2, -3}
	result := findMaxProduct(array)

	if result == 1 {
		t.Errorf("Expected score of 24, but it was %d instead.", result)
	}
}
