package isOneBitCharacter

import "testing"
func TestIsOneBitCharacter(t *testing.T) {
	array := []int{1, 1, 1, 0}

	result := isOneBitCharacter(array)

	if result != false {
		t.Errorf("Expected 3, but it was %d instead.", array)
	}
}
