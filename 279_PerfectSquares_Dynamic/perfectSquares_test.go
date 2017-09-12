package perfectSquares

import "testing"

func TestPerfectSquares(t *testing.T) {

	result := numSquares(4)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
