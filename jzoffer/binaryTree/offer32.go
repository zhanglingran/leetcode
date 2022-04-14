package binaryTree

func levelOrder(root *TreeNode) []int {
	// 层次遍历，使用队列完成
	res := make([]int, 0, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		res = append(res, top.Val)
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	return res
}
