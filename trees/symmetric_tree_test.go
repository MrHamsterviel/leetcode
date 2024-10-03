package trees

import "testing"

func TestSymmetric(t *testing.T) {
	sym := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}

	if !isSymmetric(sym) {
		t.Fatalf("wrong symmetric check")
	}

}

func TestNotSymmetric(t *testing.T) {
	notSym := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
	}

	if isSymmetric(notSym) {
		t.Fatalf("wrong notSymmetric check")
	}

}
