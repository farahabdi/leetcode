package countingBits

import "testing"

/*           1
           /   \
          2     3
         /      / \
	    4      2   4
	    	  /
	         4
*/

func TestCountingBits(t *testing.T) {
	num := 5

	result := countingBits(num)

	if result == 0 {
		t.Errorf("Expected score of 24, but it was %d instead.", result)
	}
}
