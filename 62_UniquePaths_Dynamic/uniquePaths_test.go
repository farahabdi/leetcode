package uniquePaths

import "testing"

func TestPerfectSquares(t *testing.T) {

	result := uniquePaths(3, 7)

	if result != 28 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
