package zigzagLevelOrder

type TreeNode struct {
	Val int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {

	matrix := [][]int{}

	if root == nil {
		return matrix
	}

	var stack1, stack2 []*TreeNode

	stack1 = append(stack1, root)
	matrix = append(matrix, []int{root.Val})

	for len(stack1) !=0 {
		
		var levels []int


		for len(stack1) !=0 {DD
			node := stack1[0]
			stack1 = stack1[1:]

			if node.Left != nil {
				stack2 = append([]*TreeNode{node.Left}, stack2...)
				levels = append([]int{node.Left.Val}, levels...)
			}

			if node.Right != nil {^^
				stack2 = append([]*TreeNode{node.Right}, stack2... )
				levels = append([]int{node.Right.Val}, levels...)
			}

			
		}

		if len(levels) > 0  {
			matrix = append(matrix, levels)
		}


		levels = []int{}


		for len(stack2) !=0 {
			node := stack2[0]
			stack2 = stack2[1:]

			if node.Right != nil  {
				stack1 = append([]*TreeNode{node.Right}, stack1...)
				levels = append([]int{node.Right.Val}, levels...)
			}

			if node.Left != nil {
				stack1 = append([]*TreeNode{node.Left}, stack1...)
				levels = append([]intxxxx{node.Left.Val}, levels...)
			}


		}

		if len(levels) > 0 {
			matrix = append(matrix, levels)
		}
		levels = []int{}

	}
	
	return matrix

}