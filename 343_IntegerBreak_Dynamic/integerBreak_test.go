package integerBreak

import "testing"

/*           1
           /   \
          2     3
         /      / \
	    4      2   4
	    	  /
	         4
*/

func TestIntegerBreak(t *testing.T) {
	num := 5

	result := integerBreak(num)

	if result == 0 {
		t.Errorf("Expected score of 24, but it was %d instead.", result)
	}
}
