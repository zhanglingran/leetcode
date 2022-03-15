package problems

// 二叉树的层次遍历

// 使用slice实现一个队列

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	ret := [][]int{}

	if root == nil {
		return ret
	}

	// 实现队列
	queue1 := []*TreeNode{}
	queue2 := []*TreeNode{root}

	for len(queue2) != 0 {
		//levelRet := []int{}
		levelRet := make([]int, 0, len(queue2))
		for len(queue2) != 0 {
			tmp := queue2[0]
			levelRet = append(levelRet, tmp.Val)
			queue2 = queue2[1:]
			if tmp.Left != nil {
				queue1 = append(queue1, tmp.Left)
			}
			if tmp.Right != nil {
				queue1 = append(queue1, tmp.Right)
			}
		}
		queue2 = queue1
		queue1 = nil

		if len(levelRet) != 0 {
			ret = append(ret, levelRet)
		}

	}

	return ret
}
