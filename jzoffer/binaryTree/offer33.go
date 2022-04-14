package binaryTree

// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。 如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同

/*
	可以递归的进行判断：最后一个元素是根节点，其大于左子树任意元素且小于右子树任意节点
*/
func verifyPostorder(postorder []int) bool {

	// 首先对于输入数据长度进行判断
	length := len(postorder)
	// 如果长度为0 返回true
	if length == 0 {
		return true
	}

	// 找到最后一个元素 【为根节点】
	last := postorder[length-1]
	// 游标
	cur := 0

	// 作为拆分左右子树的中间点
	split := 0

	// 找到比 根节点小的元素，此为左子树
	for postorder[cur] < last {
		cur++
	}

	// 记录中间位置，将左右子树进行拆分
	split = cur

	// 然后判断右子树是不是都大于根节点
	for cur < length-1 && postorder[cur] > last {
		cur++
	}

	// 如果右子树都大于根节点那么以 last 为根的树就满足了后续的条件；
	var flag bool
	if cur == length-1 {
		flag = true
	} else {
		flag = false
	}

	// 然后递归的判断左子树和右子树
	return flag && verifyPostorder(postorder[:split]) && verifyPostorder(postorder[split:length-1])
}
