package main

import (
	"fmt"
	"strings"
)

func main() {
	//ans := minNumber([]int{3, 30, 34, 5, 9})
	//fmt.Println(ans)

	ans := []string{"12", "34"}

	strings.Trim("  2  ", " ")

	fmt.Println(strings.Join(ans, " "))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
