package maxBinaryTree

import (
	"fmt"
	"testing"
)

/*           6
           /   \
          3      5
            \    /
	         2  0
	    	  \
	           1
*/

func TestMaxBinaryTree(t *testing.T) {
	maxBinaryTree()
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
