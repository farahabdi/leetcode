package searchInsert


import "testing"

func TestSearchInsert(t *testing.T) {

	nums := []int{1,3,5,6} // 1
	result := searchInsert(nums, 7)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}

