package merge

import "testing"

func TestMergeSortedArray(t *testing.T) {

	array1 := []int{1, 3, 4, 5}
	array2 := []int{2, 4, 6, 8}
	// result = {1, 2, 3, 4, 5, 6, 8}

	mergeArray(array1, len(array1), array2, len(array2))

}
