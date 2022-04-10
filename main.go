package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	m, n := 0, 0
	fmt.Scan(&n, &m)

	maps := make([][]int, n)
	for i := 0; i < n; i++ {
		maps[i] = make([]int, n)
	}

	// 建图
	for i := 0; i < n; i++ {
		input := ""
		fmt.Scan(&input)
		nums := strings.Split(input, ",")
		for j := 1; j < len(nums); j++ {
			// num[j] 为 i 的前驱
			node, _ := strconv.Atoi(nums[j])
			maps[node][i] = 1
		}
	}

	// 深搜
	stack := make([]int, 0, 0)
	stack = append(stack, 0)

	path := make([]int, 0, 0)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		path = append(path, node)
		flag := false
		for j := 0; j < n; j++ {

			if maps[node][j] == 1 && count(maps, j, n) == 1 {
				if j == m {
					flag = true
					break
				}
				stack = append(stack, j)
			}
		}
		if flag {
			break
		}
	}

	for i := 0; i < len(path)-1; i++ {
		fmt.Printf("%d ", path[i])
	}
	fmt.Printf("%d", path[len(path)-1])

}

func count(maps [][]int, node, n int) int {
	counts := 0
	for i := 0; i < n; i++ {
		counts += maps[i][node]
	}
	return counts
}
