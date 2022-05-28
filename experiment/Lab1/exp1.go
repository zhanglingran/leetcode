package Lab1

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

/**
凸包问题
*/

// 判断 一个点是不是在三角形内部; A[0],A[1] 表示一个点的横纵坐标
func checkPointInTriangle(A, B, C, D []int) bool {
	// 1、判断ABC是不是顺时针
	if calCrossProduct(B[0]-A[0], B[1]-A[1], C[0]-A[0], C[1]-A[1]) < 0 {
		// 如果顺时针，将BC交换，得到ABC逆时针
		B, C = C, B
	}

	// 2、判断D是不是在三角形内部
	if calCrossProduct(B[0]-A[0], B[1]-A[1], D[0]-A[0], D[1]-A[1]) >= 0 &&
		calCrossProduct(C[0]-B[0], C[1]-B[1], D[0]-B[0], D[1]-B[1]) >= 0 &&
		calCrossProduct(A[0]-C[0], A[1]-C[1], D[0]-C[0], D[1]-C[1]) >= 0 {
		return true
	}
	return false
}

func calCrossProduct(X1, Y1, X2, Y2 int) int {
	return X1*Y2 - X2*Y1
}

// ConvexHullForce 方案1、基于 枚举方法 的凸包求解算法
func ConvexHullForce(points [][]int) [][]int {

	isDelete := make([]bool, len(points))
	ans := make([][]int, 0, 0)

	for i := 0; i < len(points)-3; i++ {
		if isDelete[i] == true {
			continue
		}
		for j := i + 1; j < len(points)-2; j++ {
			if isDelete[j] == true {
				continue
			}
			for k := j + 1; k < len(points)-1; k++ {
				if isDelete[k] == true {
					continue
				}
				for m := k + 1; m < len(points); m++ {
					// 判断 得到的四个点 是不是存在一个点在三角形内部，若是则直接删除，因此要判断四次
					if isDelete[m] == true {
						continue
					}
					if checkPointInTriangle(points[i], points[j], points[k], points[m]) {
						isDelete[m] = true
						continue
					}
					if checkPointInTriangle(points[m], points[i], points[j], points[k]) {
						isDelete[k] = true
						continue
					}
					if checkPointInTriangle(points[k], points[m], points[i], points[j]) {
						isDelete[j] = true
						continue
					}
					if checkPointInTriangle(points[j], points[k], points[m], points[i]) {
						isDelete[i] = true
						continue
					}
				}
			}
		}
	}

	// 输出maps为false的节点
	for i := 0; i < len(points); i++ {
		if !isDelete[i] {
			ans = append(ans, points[i])
		}
	}

	return ans
}

// ConvexHullGrahamScan 方案2、基于 Graham-Scan 的凸包求解算法
func ConvexHullGrahamScan(points [][]int) [][]int {

	// 遍历取 (x,y) 最小的点, 作为原点
	origin, points := GetOriginNode(points)

	// 以 min 为极点将数据points按照角度排序
	SortByAngle(points, origin)

	// 排序后，将前三个点放入栈中
	stack := make([][]int, 0, 0)
	stack = append(stack, origin)
	stack = append(stack, points[0])
	stack = append(stack, points[1])

	for i := 2; i < len(points); i++ {
		for true {
			if len(stack) < 3 {
				break
			}
			top1 := stack[len(stack)-1]
			top2 := stack[len(stack)-2]
			// 判断 top\nextTop\pi 三个点是否产生了非左移动
			x1, y1 := top1[0]-top2[0], top1[1]-top2[1]
			x2, y2 := points[i][0]-top2[0], points[i][1]-top2[1]
			if calCrossProduct(x1, y1, x2, y2) <= 0 {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, points[i])
	}

	return stack
}

// ConvexHullDivideAndConquer 方案3、基于 分治思想 的凸包求解算法
func ConvexHullDivideAndConquer(points [][]int) [][]int {

	if len(points) < 3 {
		return points
	} else if len(points) == 3 {
		// 找到左下角的点
		origin, data := GetOriginNode(points)

		// 以 origin 为极点将数据 data 按照角度排序（此时数据将为逆时针序）
		SortByAngle(data, origin)

		// 重组排序后的节点
		res := make([][]int, len(data)+1)
		res[0] = origin
		for i := 0; i < len(data); i++ {
			res[i+1] = data[i]
		}

		// 将排序后的节点返回
		return res
	}
	// divide
	QL, QR := divide(points)
	QL = ConvexHullDivideAndConquer(QL)
	QR = ConvexHullDivideAndConquer(QR)
	// merge
	return ConvexHullGrahamScan(append(QL, QR...))

}

func divide(points [][]int) ([][]int, [][]int) {

	// 按照 横坐标 排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	// 按照横坐标，点集拆分成两组
	mid := len(points) >> 1
	return points[:mid], points[mid:]
}

// generatePointsData 生成测试点集数据，其范围在 (startX,startY)-(endX,endY)
func generatePointsData(startX, endX, startY, endY, count int) [][]int {

	rand.Seed(time.Now().Unix())
	data := make([][]int, count)

	for i := 0; i < count; i++ {
		x := rand.Intn(endX-startX) + startX
		y := rand.Intn(endY-startY) + startY
		data[i] = []int{x, y}
	}

	return data
}

// SortByAngle 按照 每个点与原点形成的角度 进行排序（输出序列为逆时针）
func SortByAngle(convex [][]int, origin []int) {
	// 取第一个点为 坐标原点
	x, y := origin[0], origin[1]

	sort.Slice(convex, func(i, j int) bool {
		xxA, yyA := convex[i][0]-x, convex[i][1]-y
		xxB, yyB := convex[j][0]-x, convex[j][1]-y

		// 若共线，将按照射线的长度进行排序
		if FastAtan(float64(yyA), float64(xxA)) == FastAtan(float64(yyB), float64(xxB)) {
			return xxA*xxA+yyA*yyA < xxB*xxB+yyB*yyB
		}

		return FastAtan(float64(yyA), float64(xxA)) < FastAtan(float64(yyB), float64(xxB))
	})
}

// GetOriginNode 获取点集左下角元素作为原点坐标
func GetOriginNode(points [][]int) ([]int, [][]int) {
	origin := points[0]
	pos := 0
	for i := 1; i < len(points); i++ {
		if points[i][1] < origin[1] {
			origin = points[i]
			pos = i
		}
		if points[i][1] == origin[1] && points[i][0] < origin[0] {
			origin = points[i]
			pos = i
		}
	}
	points[0], points[pos] = points[pos], points[0]
	origin = points[0]
	points = points[1:]
	return origin, points
}

// FastAtan 因为凸包上所有点都不可能与原点重合，那么也就不会出现 x,y = 0,0 的情况
func FastAtan(y, x float64) float64 {

	ax := math.Abs(x)
	ay := math.Abs(y) //首先不分象限，求得一个锐角角度
	var a float64
	if ax >= ay {
		a = math.Atan(ay / ax)
	} else {
		// 防止出现(0,1)点的情况
		a = 90 - math.Atan(ax/ay)
	}
	if x < 0 {
		a = 180.0 - a
	}
	if y < 0 {
		a = 360.0 - a
	}
	return a
}
