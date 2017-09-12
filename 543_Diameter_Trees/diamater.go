package Tree_543

/*           30
           /     \
          25      35
         /  \    /  \
	   20   26  31   37
	   / \		 \
	15    21 	  33
			\
			 22		  */

func Diamater(root *Node) int {
	if root == nil {
		return 0
	}

	left := Diamater(root.Left) + 1
	right := Diamater(root.Right) + 1

	return max(1, 2)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
