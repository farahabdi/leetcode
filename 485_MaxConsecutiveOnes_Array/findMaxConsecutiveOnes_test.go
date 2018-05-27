package findMaxConsecutiveOnes


import "testing"

func TestMaxConsecutiveOnes(t *testing.T) {

	nums := []int{0,1}
	result := findMaxConsecutiveOnes(nums)

	if result == 3 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
