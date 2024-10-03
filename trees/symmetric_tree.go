package trees

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// isMirror checks if two subtrees are mirror images of each other
func isMirror(left *TreeNode, right *TreeNode) bool {
	// Base case: both nodes are nil, meaning the subtrees are mirrors
	if left == nil && right == nil {
		return true
	}

	// If only one of them is nil, the subtrees are not mirrors
	if left == nil || right == nil {
		return false
	}

	// Check if the values of the two nodes are the same and recursively
	// check if their subtrees are mirrors
	return (left.Val == right.Val) &&
		isMirror(left.Left, right.Right) &&
		isMirror(left.Right, right.Left)
}
