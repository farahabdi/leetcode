package unival

import (

	"testing"
)

/*           0
           /   \
          1      0
            \     \ 0
	         1  
	        / \
	       1   1
*/

func TestMaxBinaryTree(t *testing.T) {
	var root *TreeNode
	root = &TreeNode{Val: 0, Left: nil, Right: nil}

	root.Left = &TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 0,  Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 0, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right.Left.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right.Left.Right = &TreeNode{Val: 1, Left: nil, Right: nil}

	result, _ := unival(root)

	if result != 3 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}

}
