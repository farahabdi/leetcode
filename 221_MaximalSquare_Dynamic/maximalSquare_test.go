package maximalSquare

import "testing"

func TestMaximalSquare(t *testing.T) {
	//	["00","00"]
	array1 := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	/*	array1 := [][]byte{
		{'0', '0'},
		{'1', '1'},
	}*/
	result := maximalSquare(array1)

	if result != 3 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
