package binaryTree

/*
使用递归解决问题：
	需要创建新节点？
*/
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	head := &TreeNode{Val: root.Val}

	head.Left = mirrorTree(root.Right)
	head.Right = mirrorTree(root.Left)

	return head
}

/*
不创建节点的方法试一试
*/
func mirrorTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	leftSubTree := root.Left
	rightSubTree := root.Right

	root.Left = mirrorTree(rightSubTree)
	root.Right = mirrorTree(leftSubTree)

	return root
}

/*
错误解决方案：
	如果不将 left 与 right 临时存储，那么 root.Left = mirrorTree(root.Right) 执行结束之后，就会出现右子树覆盖左子树的情况，造成数据丢失
*/

func mirrorTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(root.Left)

	return root
}

/*
	非递归的方式 : 使用队列解决问题
*/
func mirrorTreeWithoutRecursion(root *TreeNode) *TreeNode {

	// 特判特殊条件
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	queue := make([]*TreeNode, 0, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		top := queue[0]
		left := top.Left
		right := top.Right
		if left != nil {
			queue = append(queue, left)
		}
		if right != nil {
			queue = append(queue, right)
		}
		queue = queue[1:]
		top.Left = right
		top.Right = left
	}

	return root
}
