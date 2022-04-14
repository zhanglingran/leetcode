package linkedList

/**
合并两个有序链表
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{-1, nil}
	p := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
		p.Next = nil
	}

	for l1 != nil {
		p.Next = l1
		l1 = l1.Next
		p = p.Next
		p.Next = nil
	}
	for l2 != nil {
		p.Next = l2
		l2 = l2.Next
		p = p.Next
		p.Next = nil
	}

	return head.Next
}

// 使用递归完成
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head *ListNode
	if l2.Val > l1.Val {
		head = l1
		head.Next = mergeTwoLists1(l1.Next, l2)
	} else {
		head = l2
		head.Next = mergeTwoLists1(l1, l2.Next)
	}
	return head
}
