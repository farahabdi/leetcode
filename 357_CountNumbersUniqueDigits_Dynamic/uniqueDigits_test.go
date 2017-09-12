package uniqueDigits

import "testing"

func TestCoinChange(t *testing.T) {

	result := countNumbersWithUniqueDigits(11)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
