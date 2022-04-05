package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
一定要考虑代码的鲁棒性
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {

	// 判断head是空指针的情况 和 k = 0 的情况
	if head == nil || k == 0 {
		return nil
	}

	// 两个指针，right先移动k个，然后left从head开始与right一起移动
	left, right := head, head
	for i := 0; i < k; i++ {
		right = right.Next

		// 判断k比数字大的情况
		if right == nil {
			return nil
		}
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}
	return left
}
