package exp2

import (
	"fmt"
	"math"
	"strings"
)

// 迪杰斯特拉算法(单源最短路问题)
func f() {

	const INF = math.MaxInt32

	maps := make(map[byte]int)

	maps['a']++

	delete(maps, 'a')

	s := "12"
	s = strings.ToLower(s)
	data := []byte(s)
	fmt.Println(string(data))
}
