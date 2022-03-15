package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList 反转链表
// 三种方案：
//	一、头插or尾插 构建新的链表；
func reverseList(head *ListNode) *ListNode {

	// 用 L 作为新链表的开始头节点
	var L *ListNode
	// cur 为游标指针，指向当前待反转的节点
	cur := head

	for cur != nil {
		tmp := cur.Next
		cur.Next = L
		L = cur
		cur = tmp
	}

	return L
}

// 二、递归实现链表反转
func reverseListRecurse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 三、双指针实现链表反转
func reverseListByTwoPoint(head *ListNode) *ListNode {

	var cur *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = cur
		cur = head
		head = tmp
	}

	return cur
}
