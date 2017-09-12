package searchRange



import "testing"

func TestSearchRange(t *testing.T) {
	coins := []int{5, 10, 30, 50, 100}

	result := searchRange(coins)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
