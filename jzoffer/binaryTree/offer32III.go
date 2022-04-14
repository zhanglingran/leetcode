package binaryTree

func levelOrderIII(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	kth := 0
	ret := make([][]int, 1001)
	queue := make([]*TreeNode, 0, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		lenLevel := len(queue)
		ret[kth] = make([]int, 0, 0)

		// 如果使用两个循环的话，那么倒序遍历的结果中，入队的顺序也是倒叙，那么将会影响下次的结果
		for i := 0; i < lenLevel; i++ {
			if kth%2 == 0 {
				ret[kth] = append(ret[kth], queue[i].Val)
			} else {
				ret[kth] = append(ret[kth], queue[lenLevel-1-i].Val)
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[lenLevel:]
		kth++
	}

	return ret[:kth]
}
