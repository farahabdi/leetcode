
package combinationSum2

import ("testing")
func TestCombinationSum2(t *testing.T) {
	combination := []int{2,5,2,1,2}
	//combination := []int{10,1,2,7,6,1,5}
	target := 5

	result := combinationSum2(combination, target)

	if result != nil {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}

