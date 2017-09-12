package rotateMatrix

import "testing"

func TestRotateMatrix(t *testing.T) {

	array := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(array)

	//	if result != ∂≈1.0 {
	//		t.Errorf("Expected 3, but it was %f instead.", result)
	//	}
}
