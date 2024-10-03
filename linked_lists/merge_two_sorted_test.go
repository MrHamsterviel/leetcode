package linked_lists

import "testing"

func TestMergeTwoSorted(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	res := mergeTwoLists(list1, list2)
	expected := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 6}
	for _, e := range expected {
		if res.Val != e {
			t.Fatalf("incorrect list val %d, expected %d", res.Val, e)
		}
		if res.Next != nil {
			res = res.Next
		} else {
			break
		}
	}
}
