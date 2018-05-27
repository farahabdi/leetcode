package zigzagLevelOrder

import (

	"testing"
)

/*           3
           /   \
          9     20
                / \
	          15  7

*/


	
func TestZigzagLevelOrder(t *testing.T) {

	var root *TreeNode
	/*
	root = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 9, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 20, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 7, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 15, Left: nil, Right: nil} */

	root = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 5, Left: nil, Right: nil}




  zigzagLevelOrder(root)


}



