package isValidTree

import (
	"fmt"
	"testing"
)

/*           30
           /     \
          25      35
         /  \    /  \
	   20   26  31   37
	   / \		 \
	15    21 	  33
			\
			 22		  */


			 type Node struct {
				Value int
				Left  *Node
				Right *Node
			}
			
func TestIsValidTree(t *testing.T) {

  return isValidTree(root, 30)


}



