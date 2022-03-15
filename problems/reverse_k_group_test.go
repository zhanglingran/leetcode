// Package Leetcode K 个一组翻转链表
package problems

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReverseKGroup(t *testing.T) {

	Convey("测试k个一组进行链表反转", t, func() {
		node1 := &ListNode{Next: nil, Val: 1}
		node2 := &ListNode{Next: nil, Val: 2}
		node3 := &ListNode{Next: nil, Val: 3}
		node4 := &ListNode{Next: nil, Val: 4}
		node5 := &ListNode{Next: nil, Val: 5}
		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5

		ret := ReverseKGroup(node1, 2)
		var arr [5]int
		i := 0
		for ret != nil {
			arr[i] = ret.Val
			ret = ret.Next
			i++
		}
		So(arr, ShouldEqual, [5]int{2, 1, 4, 3, 5})
	})

}
