package binaryTree

/**
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
B是A的子结构， 即 A中有出现和B相同的结构和节点值。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {

	ret := false

	if A != nil && B != nil {
		// 1、如果 相同则直接判断A是不是包含B
		if A.Val == B.Val {
			ret = doseTree1HasTree2(A, B)
		}

		// 如果两个节点不相同，判断A的左右节点
		if !ret {
			ret = isSubStructure(A.Left, B)
		}
		if !ret {
			ret = isSubStructure(A.Right, B)
		}
	}

	return ret
}

/**
判断 以root1为根  d的树 是不是包含 以root2为根的树
*/
func doseTree1HasTree2(root1, root2 *TreeNode) bool {

	// 如果 root2已经空了 就说明root1包含了root2
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}
	if root2.Val != root1.Val {
		return false
	}

	// 如果当前节点满足，还需要其左右子树满足条件
	return doseTree1HasTree2(root1.Left, root2.Left) && doseTree1HasTree2(root1.Right, root2.Right)
}
