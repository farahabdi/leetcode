package unival

type TreeNode struct {
	Val int
	Left  *TreeNode
	Right *TreeNode
}

func unival(root *TreeNode) (int, bool) {

	if root == nil {
		return 0, true
	}

	leftCount, isLeftUnival  := unival(root.Left)
	rightCount, isRightUnival  := unival(root
		.Right)
	totalCount := leftCount + rightCount
	if isLeftUnival && isRightUnival {
		if root.Left !=nil && root.Val != root.Left.Val {
			return totalCount, false
		}
		if root.Right != nil && root.Val != root.Right.Val {
			return totalCount, false
		}
		return totalCount + 1, true
	}
return totalCount, false
}