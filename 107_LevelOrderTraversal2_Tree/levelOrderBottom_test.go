package levelOrderBottom

import (

	"testing"
)

/*           3
           /   \
          9     20
                / \
	          15  7

*/


	
func TestLevelOrderBottom(t *testing.T) {

	var root *TreeNode
	root = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 9, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 20, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 7, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 15, Left: nil, Right: nil}


  levelOrderBottom(root)


}



