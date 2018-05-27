package decodeWays

import "testing"

func TestDecodeWays(t *testing.T) {
	str := "12321"

	result := numDecodings(str)

	if result != 3 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
