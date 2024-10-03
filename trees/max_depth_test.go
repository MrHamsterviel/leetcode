package trees

import (
	h "leetcode/helpers"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	arr := []*int{h.Ptr(1), h.Ptr(3), nil, h.Ptr(5), h.Ptr(6), h.Ptr(9), nil, nil, h.Ptr(12)}
	root := BuildBinTree(arr)
	actual := MaxDepth(root)
	if actual != 4 {
		t.Fatalf("max depth fail, expect %d, got %d", 4, actual)
	}
}
