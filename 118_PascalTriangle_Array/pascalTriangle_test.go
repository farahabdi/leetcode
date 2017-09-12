package pascalTriangle

import "testing"

func TestPascalTriangle(t *testing.T) {
	result := generate(5)

	if result != nil {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
