package grayCode

import "testing"

func TestGrayCode(t *testing.T) {


	result := grayCode(3)

	if result != 2 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}

}
