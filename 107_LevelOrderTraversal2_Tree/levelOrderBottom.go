package levelOrderBottom

type TreeNode struct {
	Val int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int  {

	matrix := [][]int{}

	if root == nil {
		return matrix
	}

	queue := []*TreeNode{root}
	
	for len(queue) != 0 {

		values := []int{}
		n := len(queue)

		for i:=0; i<n; i++{
			curr := queue[0]
			queue =  queue[1:]
			
			values = append(values, curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		matrix = append(matrix, values)
	
	}

	n := len(matrix)-1
	for i:=0; i<len(matrix)/2; i++ {
		matrix[i], matrix[n-i] = matrix[n-i], matrix[i]
	} 
	return matrix
}



func levelOrderBottoms(root *TreeNode) [][]int  {

	matrix := [][]int{}

	if root == nil {
		return matrix
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

	n := len(matrix)-1
	for i:=0; i<len(matrix)/2; i++ {
		matrix[i], matrix[n-i] = matrix[n-i], matrix[i]
	} 
	return matrix
}



