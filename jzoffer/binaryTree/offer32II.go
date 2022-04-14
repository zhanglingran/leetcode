package binaryTree

func levelOrderII(root *TreeNode) [][]int {

	queue := make([]*TreeNode, 0, 0)
	res := make([][]int, 1001)
	kth := 0

	if root == nil {
		return res[:kth]
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		lenLevel := len(queue)
		res[kth] = make([]int, 0, lenLevel)
		for i := 0; i < lenLevel; i++ {
			res[kth] = append(res[kth], queue[i].Val)
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
	return res[:kth]
}
