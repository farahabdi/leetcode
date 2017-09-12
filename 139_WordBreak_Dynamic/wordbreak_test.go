package wordbreak

import "testing"

/*           1
           /   \
          2     3
         /      / \
	    4      2   4
	    	  /
	         4
*/

func TestWordBreak(t *testing.T) {
	str := "heyaman"

	result := wordbreak(str)

	if result == false {
		t.Errorf("Expected score of 24, but it was %d instead.", result)
	}
}
