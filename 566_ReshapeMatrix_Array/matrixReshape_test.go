package matrixReshape

import "testing"
func TestReshapeMatrix(t *testing.T) {
	array := [][]int{
		{1,2,3,4,5,6,7,8},
		{9,10,11,12,13,14,15,16},

	}
	array1 := [][]int{
		{1,2},
		{3,4},

	}
	result := matrixReshape(array1,4, 1)

	if result != nil {
		t.Errorf("Expected 3, but it was %d instead.", array)
	}
}
