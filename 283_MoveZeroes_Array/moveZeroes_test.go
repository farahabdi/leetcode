package moveZero

import "testing"

/*           1
           /   \
          2     3
         /      / \
	    4      2   4
	    	  /
	         4
*/

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 0, 1}

	moveZeroes(nums)

}
