package main

import (
	"fmt"
)

func main() {

	fmt.Println(cell(float64(2) / 3))
}

func cell(a float64) int {
	if a > float64(int(a)) {
		return int(a) + 1
	}
	return int(a)
}
