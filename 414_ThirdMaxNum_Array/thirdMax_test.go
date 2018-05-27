package thirdMax
import ("testing")
func TestThirdMax(t *testing.T) {
	//array := []int{2, 3, -1, -9, 11, 4, 3, 0, -100}
	array := []int{1, 1, 2}
	//array := []int{1,2,-2147483648}
	//array := []int{3,2,1}
	result := thirdMax(array)

	if result != 2 {
		t.Errorf("Expected 3, but it was %d instead.", array)
	}
}
