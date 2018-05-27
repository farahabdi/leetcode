package degreeArray
import ("testing")
func TestDegreeArray(t *testing.T) {
	//array := []int{3,3,4}
	array := []int{1,1}
	//array := []int{1,2,1,3,2}
	result := findShortestSubArray(array)

	if result != 2 {
		t.Errorf("Expected 3, but it was %d instead.", array)
	}
}
