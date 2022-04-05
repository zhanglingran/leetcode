package contest287

func findWinners(matches [][]int) [][]int {

	// 维护两个map，记录win与lose的次数
	winMap := make(map[int]int)
	loseMap := make(map[int]int)

	answer0 := make([]int, 0, 0)
	answer1 := make([]int, 0, 0)

	// 记录最大编号
	maxNum := 0
	for i := 0; i < len(matches); i++ {
		win := matches[i][0]
		lose := matches[i][1]
		winMap[win]++
		loseMap[lose]++
		if win > maxNum {
			maxNum = win
		}
		if lose > maxNum {
			maxNum = lose
		}
	}

	for i := 1; i <= maxNum; i++ {
		if winMap[i] != 0 && loseMap[i] == 0 {
			answer0 = append(answer0, i)
		}
		if loseMap[i] == 1 {
			answer1 = append(answer1, i)
		}
	}

	return [][]int{answer0, answer1}
}
