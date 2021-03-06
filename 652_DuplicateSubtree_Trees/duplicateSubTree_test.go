package duplicateSubTree

import (
	"fmt"
	"testing"
)

/*           1
           /   \
          2     3
         /      / \
	    4      2   4
	    	  /
	         4
*/

func TestDuplicateSubTree(t *testing.T) {
	var root *Node
	root = Insert(root, 1)
	root.Left = &Node{Value: 2, Left: nil, Right: nil}
	root.Left.Left = &Node{Value: 4, Left: nil, Right: nil}
	root.Right = &Node{Value: 3, Left: nil, Right: nil}
	root.Right.Right = &Node{Value: 4, Left: nil, Right: nil}
	root.Right.Left = &Node{Value: 2, Left: nil, Right: nil}
	root.Right.Left.Left = &Node{Value: 4, Left: nil, Right: nil}
	//hey := checkSubTree(root.Left, root.Right)

	m := subtree(root, root.Right.Left)

	fmt.Println(m)
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Insert(root *Node, key int) *Node {

	if root == nil {
		root = &Node{Value: key, Left: nil, Right: nil}
	} else if key < root.Value {
		root.Left = Insert(root.Left, key)
	} else if key > root.Value {
		root.Right = Insert(root.Right, key)
	}

	return root

}

func printTree(root *Node) *Node {
	if root == nil {
		return root
	}

	printTree(root.Left)

	fmt.Println(root.Value)

	printTree(root.Right)

	return root
}
