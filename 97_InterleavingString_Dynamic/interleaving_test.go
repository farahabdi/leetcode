package interleavingString

import "testing"

func TestInterleavingString(t *testing.T) {

	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	result := isInterleave(s1, s2, s3)

	if result != true {
		t.Errorf("Expected 3, but it was %t instead.", result)
	}
}
