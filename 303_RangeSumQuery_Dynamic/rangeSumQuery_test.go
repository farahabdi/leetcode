package rangeSumQuery

import "testing"

func TestRangeSumQuery(t *testing.T) {

	nums := []int{-2, 0, 3, -5, 2, -1}

	obj := Constructor(nums)
	result := obj.SumRange(0, 5)

	if result != -3 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
