package medianSortedArrays

import "testing"

func TestMedianSortedArrays(t *testing.T) {

	A := []int{1, 2, 3, 6}
	B := []int{4, 6, 8, 10}
	result := findMedianSortedArrays(A, B)

	if result != 1.0 {
		t.Errorf("Expected 3, but it was %f instead.", result)
	}
}
