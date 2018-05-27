package mySqrt

import "testing"


func TestMySqrt(t *testing.T) {

	result := mySqrt(8)

	if result != 2 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
