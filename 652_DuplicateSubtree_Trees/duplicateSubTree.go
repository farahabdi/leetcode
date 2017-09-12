package duplicateSubTree

/*           1
           /   \
          2     3
         /      / \
	    4      2   4
	    	  /
	         4
*/

func duplicateSubTree() {

}

func subtree(node *Node, root *Node) bool {
	if root == nil {
		return true
	} else if root == nil {
		return false
	} else if checkSubTree(node, root) == true {
		return true
	}

	return subtree(node.Left, root) || subtree(node.Right, root)
}

func checkSubTree(nodeA *Node, nodeB *Node) bool {

	if nodeA == nil {
		return true
	} else if nodeB == nil {
		return false
	} else if nodeA.Value == nodeB.Value &&
		checkSubTree(nodeA.Left, nodeB.Left) &&
		checkSubTree(nodeA.Right, nodeB.Right) {
		return true
	}
	return false
}
