package rightSideView

type TreeNode struct {
	Val int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
    
	matrix := [][]int{}

	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}

	for len(stack) != 0 {
		nodes := []*TreeNode{}
		values := []int{}

		for i:= range stack {
			values = append(values, stack[i].Val)

			if stack[i].Left != nil {
				nodes = append(nodes, stack[i].Left)
			}
			if stack[i].Right != nil {
				nodes = append(nodes, stack[i].Right)
			}
		}

		matrix = append(matrix, values)
		stack = nodes
	}

	var values []int
	for i:= range matrix {
		values = append(values, matrix[i][len(matrix[i])-1])
	}


	return values
}