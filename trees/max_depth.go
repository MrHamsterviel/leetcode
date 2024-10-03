package trees

func MaxDepth(root *TreeNode) int {
	res := maxDepth(root)
	return res
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// Return the larger depth plus one for the current node
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func BuildBinTree(arr []*int) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *arr[0]}
	queue := []*TreeNode{root} // Queue to hold the tree nodes
	i := 1

	for i < len(arr) {
		current := queue[0] // Dequeue the current node
		queue = queue[1:]   // Remove the first element from the queue

		// Left child
		if i < len(arr) {
			if arr[i] != nil {
				current.Left = &TreeNode{Val: *arr[i]}
				queue = append(queue, current.Left)
			}
			i++
		}

		// Right child
		if i < len(arr) {
			if arr[i] != nil {
				current.Right = &TreeNode{Val: *arr[i]}
				queue = append(queue, current.Right)
			}
			i++
		}
	}

	return root
}
