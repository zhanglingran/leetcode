package main

var (
	maps  map[byte]int
	queue []byte
)

func main() {
	if maps == nil {
		maps = make(map[byte]int)
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func isValidate(arr []int) bool {

	vis := make(map[int]int)
	union := make(map[int]int)

	//max, min := 0, math.MaxInt32
	// 将 arr 放到map中

	for i := 0; i < len(arr); i++ {
		vis[arr[i]]++
		if i > 0 {
			dif := abs(arr[i] - arr[i-1])
			union[dif]++
		}
	}

	// 计算 相邻数之差
	for i := 0; i < len(arr); i++ {
		vis[arr[i]]++
	}

	return false
}
