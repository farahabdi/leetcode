package decodeWays

import "testing"

func TestDecodeWays(t *testing.T) {
	str := "123"

	result := numDecodings(str)

	if result != 0 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
