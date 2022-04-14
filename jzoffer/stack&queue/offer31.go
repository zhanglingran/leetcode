package stack_queue

func validateStackSequences(pushed []int, popped []int) bool {

	// 借助一个辅助栈，如果popped中第一个数是栈顶元素则弹出，否则pushed入栈，直到栈顶元素与popped相同

	stack := make([]int, 0, 0)
	curPop, curPush := 0, 0
	length := len(popped)

	if length == 0 {
		return true
	}

	for curPush < length {
		// 如果 popped[curPop] 与栈顶元素不同，则入栈，直到栈顶元素与 popped[curPop] 相同
		for curPush < length && (len(stack) == 0 || stack[len(stack)-1] != popped[curPop]) {
			stack = append(stack, pushed[curPush])
			curPush++
		}

		// 找到 popped[curPop] 与栈顶元素相同之后，不断出栈，直到栈顶元素 与 popped[curPop] 不同
		for curPop < length && (len(stack) > 0 && stack[len(stack)-1] == popped[curPop]) {
			stack = stack[:len(stack)-1]
			curPop++
		}
	}
	return curPush == curPop
}
