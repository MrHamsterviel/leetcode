package trees

import "math"

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt, math.MaxInt)
}

func validate(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
}
