package linked_lists

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	node := &ListNode{}
	if list1.Val <= list2.Val {
		node.Val = list1.Val
		node.Next = mergeTwoLists(list1.Next, list2)
	} else {
		node.Val = list2.Val
		node.Next = mergeTwoLists(list1, list2.Next)
	}
	return node
}
