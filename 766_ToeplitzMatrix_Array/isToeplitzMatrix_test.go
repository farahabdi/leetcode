package toePlitz
func TestPowerSet(t *testing.T) {
	array := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}

	result := isToeplitzMatrix(array)

	if result != nil {
		t.Errorf("Expected 3, but it was %d instead.", array)
	}
}
