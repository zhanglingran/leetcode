package problems

type node struct {
	index int
	value int
}

func DailyTemperatures(temperatures []int) []int {
	var stack = make([]*node, 10e5)
	var res = make([]int, 10e5)

	cur := 0

	for idx, val := range temperatures {
		if cur == 0 {
			stack[cur] = &node{index: idx, value: val}
			cur++
		} else {
			for cur != 0 && val > stack[cur-1].value {
				res[stack[cur-1].index] = idx - stack[cur-1].index
				cur--
			}
			stack[cur] = &node{index: idx, value: val}
			cur++
		}
	}
	return res[:len(temperatures)]
}

// DailyTemperatures2 效率最高的版本; 节省了大量的内存空间
func DailyTemperatures2(temperatures []int) []int {

	stackIdx := make([]int, 0)
	res := make([]int, len(temperatures))

	for idx, val := range temperatures {
		for len(stackIdx) > 0 && val > temperatures[stackIdx[len(stackIdx)-1]] {
			curIdx := stackIdx[len(stackIdx)-1]
			res[curIdx] = idx - curIdx
			stackIdx = stackIdx[:len(stackIdx)-1]
		}
		stackIdx = append(stackIdx, idx)
	}
	return res
}

// DailyTemperatures3 使用for 和 foreach 两种遍历方式，并没有什么效率上的差别
func DailyTemperatures3(temperatures []int) []int {
	length := len(temperatures)
	stackIdx := make([]int, 0)
	res := make([]int, length)

	for idx := 0; idx < length; idx++ {
		for len(stackIdx) > 0 && temperatures[idx] > temperatures[stackIdx[len(stackIdx)-1]] {
			curIdx := stackIdx[len(stackIdx)-1]
			res[curIdx] = idx - curIdx
			stackIdx = stackIdx[:len(stackIdx)-1]
		}
		stackIdx = append(stackIdx, idx)
	}
	return res
}

/*

73,74,75,71,69,72,76,73
 1, 1, 4, 2, 1, 1, 0, 0

 单调栈
73(0),74(1),75(2),71(3),69(4),72(5),76(6),73(7)

push 73(0)
	74(1) > top; res[top.idx] = cur.idx-top.idx
pop 73(0)
push 74(1)
	75(2) > top; res[top.idx] = cur.idx-top.idx
pop 74(1)
push 75(2)
	71(3) < top;
push 71(3)
	69(4) < top;
push 69(4)
	72(5) > top; res[top.idx] = cur.idx-top.idx
pop 64(4)
	72(5) > top; res[top.idx] = cur.idx-top.idx
pop 71(3)
	72(5) < top;
*/
