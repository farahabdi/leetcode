package maxRectangle

import "testing"

func TestMaximalRectangle(t *testing.T) {
	array := [][]byte{
		{0, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 0, 0, 0},
	}
	result := maximalRectangle(array)

	if result == 3 {
		t.Errorf("Expected score of 24, but it was %d instead.", result)
	}
}
