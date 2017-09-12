package productArray

import (
	"reflect"
	"testing"
)

func TestProductArray(t *testing.T) {

	nums := []int{2, 3, 4, 5}
	output := []int{24, 12, 8, 6}
	result := productExceptSelf(nums)

	if reflect.DeepEqual(nums, output) != true {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
