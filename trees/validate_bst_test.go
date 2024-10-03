package trees

import (
	"testing"
)

func TestCorrectBST(t *testing.T) {
	correct := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	if !isValidBST(correct) {
		t.Fatalf("correct tree is not valid")
	}

}

func TestIncorrectBST(t *testing.T) {
	incorrect := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	if isValidBST(incorrect) {
		t.Fatalf("incorrect tree is valid")
	}

}
