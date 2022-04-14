package binaryTree

/*
	将 NLR 与 NRL 的遍历结果进行对比，两者相同即为对称二叉树【这种思路错了,因为将null节点过滤了！】
	当数据为 ：[1,2,2,null,3,null,3] 的时候，两种遍历的结果是相同的，但是并非对称二叉树
*/

func isSymmetric(root *TreeNode) bool {

	// 1、先获取一个树的镜像树
	mirrorRoot := getMirrorTree(root)
	// 2、将两个树进行对比是否一致

	return isEqual(root, mirrorRoot)
}

// 获取镜像树
func getMirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	head := &TreeNode{Val: root.Val}
	head.Left = getMirrorTree(root.Right)
	head.Right = getMirrorTree(root.Left)
	return head
}

// 对比两个树是否一致
func isEqual(root1 *TreeNode, root2 *TreeNode) bool {

	if root2 == nil && root1 == nil {
		return true
	}
	if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) {
		return false
	}
	if root1.Val == root2.Val {
		return true
	}

	return isEqual(root1.Left, root2.Left) && isEqual(root1.Right, root2.Right)
}

/*
	剑指offer中的解法:
		判断 先序 与 对称先序 是否相同
*/
func isSymmetric1(root *TreeNode) bool {
	return isSymmetrical(root, root)
}

// 将 NLR 与 NRL 结合起来判断
func isSymmetrical(root1, root2 *TreeNode) bool {

	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return isSymmetrical(root1.Left, root2.Right) && isSymmetrical(root1.Right, root2.Left)
}
