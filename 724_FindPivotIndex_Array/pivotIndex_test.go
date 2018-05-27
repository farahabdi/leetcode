package pivotIndex


import "testing"

func TestTriangleNumber(t *testing.T) {
	//coins := []int{1, 7, 3, 6, 5, 6}
	//coins := []int{2, 3, 5, 1, 5, 7,50,14,9}
	//coins := []int{-1,-1,-1,0,1,1}[]
	//coins := []int{-1,-1,0,1,-1,-1}

	//coins := []int{-1,-1,1,1,-1,0}
	//coins := []int{-1,-1,0,1,1,-1}// 0

	coins := []int{-1,-1,1,0,-1,-1} // 1
	result := pivotIndex(coins)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
