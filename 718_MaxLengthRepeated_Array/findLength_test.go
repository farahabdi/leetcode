package findLength

import ("testing")
func TestDegreeArray(t *testing.T) {

	A := []int{1,2,3,2,1}
	B := []int{3,2,1,4,7}
	result := findLengthDP(A, B)

	if result != 2 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}

