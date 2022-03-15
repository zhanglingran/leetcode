// Package Leetcode K 个一组翻转链表
package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ReverseKGroup 用start-end表示待逆转的子链表
// 用dummy表示头节点
func ReverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	dummy.Next = head

	end := dummy
	pre := dummy

	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		// end 后边的指针，记录一下
		next := end.Next
		// 将start-end摘出来
		start := pre.Next
		end.Next = nil

		pre.Next = reverse(start)

		// 注意：revers函数返回后，start指针指向哪里？
		start.Next = next
		pre = start
		end = pre
	}

	// 用pre来作为游动指针
	return dummy.Next
}

// 链表逆转 按照头尾进行逆转
func reverse(head *ListNode) *ListNode {

	var pre *ListNode

	for head != nil {
		t := head.Next
		head.Next = pre
		pre = head
		head = t
	}

	return pre
}
