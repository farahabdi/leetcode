package maxBinaryTree

/*           6
           /   \
          3      5
            \    /
	         2  0
	    	  \
	           1
*/

func maxBinaryTree() {
	num := []int{3, 2, 1, 6, 0, 5}
	construct(num, 0, len(num))
}

func construct(num []int, l int, r int) *Node {
	if l == r {
		return nil
	}

	max_i := max(num, l, r)
	root := &Node{Value: num[max_i], Left: nil, Right: nil}
	root.Left = construct(num, l, max_i)
	root.Right = construct(num, max_i+1, r)

	return root
}

func max(num []int, l int, r int) int {

	max_i := l
	for i := max_i; i < r; i++ {
		if num[max_i] < num[i] {
			max_i = i
		}

	}
	return max_i
}
