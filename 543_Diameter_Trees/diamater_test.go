package Tree_543

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

func TestDiamater(t *testing.T) {
	var root *Node
	root = Insert(root, 30)
	Insert(root, 25)
	Insert(root, 35)
	Insert(root, 31)
	Insert(root, 33)
	Insert(root, 37)
	Insert(root, 25)
	Insert(root, 20)
	Insert(root, 26)
	Insert(root, 15)
	Insert(root, 21)
	Insert(root, 22)
	printTree(root)
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
