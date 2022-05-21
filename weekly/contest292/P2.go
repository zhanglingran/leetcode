package contest292

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 其实就是在后续遍历的时候，回溯的同时判断即可
func averageOfSubtree(root *TreeNode) int {
	var ans = 0
	post(root, &ans)
	return ans
}

// 计算以root为根节点的子树的均值，是否等于root节点的val值
func post(root *TreeNode, ans *int) (int, int) {
	if root == nil {
		return 0, 0
	}

	lcount, lval := post(root.Left, ans)
	rcount, rval := post(root.Right, ans)

	count, val := lcount+rcount+1, lval+rval+root.Val

	if val/count == root.Val {
		*ans++
	}

	return count, val
}
