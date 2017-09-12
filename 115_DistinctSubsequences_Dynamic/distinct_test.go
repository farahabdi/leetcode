package distinctSubsequence

import (
	"testing"
)

func TestDistinctSubsequence(t *testing.T) {
	S := "banana"
	T := "ban"

	num := numDistinctRecursive(S, T)
	num = numDistinctDP(S, T)

	if num != 3 {
		t.Errorf("Expected coutn of 3, got %d instead.", num)
	}
}
