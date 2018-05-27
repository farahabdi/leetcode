package flipAndInvertImage

import (
	"testing"
)
func TestFlipAndInvertImag(t *testing.T) {
	array := [][]int{
		{1, 1, 0, 0},
		{1, 0, 0, 1},
		{0, 1, 1, 1},
		{1, 0, 1, 0},
	}

	result := flipAndInvertImage(array)

	if result != nil {
		t.Errorf("Expected 3, but it was %d instead.", array)
	}
}
