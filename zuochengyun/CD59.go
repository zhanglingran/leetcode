package zuochengyun

import (
	"fmt"
)

type Node struct {
	X float64
	Y float64
}

func Main() {

	A := Node{
		X: 0,
		Y: 0,
	}
	X, Y := 0.0, 0.0
	fmt.Scan(&X, &Y)
	A.X = X
	A.Y = Y

	B := Node{
		X: 0,
		Y: 0,
	}
	fmt.Scan(&X, &Y)
	B.X = X
	B.Y = Y

	C := Node{
		X: 0,
		Y: 0,
	}
	fmt.Scan(&X, &Y)
	C.X = X
	C.Y = Y

	D := Node{
		X: 0,
		Y: 0,
	}
	fmt.Scan(&X, &Y)
	D.X = X
	D.Y = Y

	if checkPointInTriangle(A, B, C, D) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func checkPointInTriangle(A, B, C, D Node) bool {
	// 1、判断ABC是不是顺时针
	if calCrossProduct(B.X-A.X, B.Y-A.Y, C.X-A.X, C.Y-A.Y) < 0 {
		// 如果顺时针，将BC交换，得到ABC逆时针
		B, C = C, B
	}

	// 2、判断D是不是在三角形内部
	if calCrossProduct(B.X-A.X, B.Y-A.Y, D.X-A.X, D.Y-A.Y) >= 0 &&
		calCrossProduct(C.X-B.X, C.Y-B.Y, D.X-B.X, D.Y-B.Y) >= 0 &&
		calCrossProduct(A.X-C.X, A.Y-C.Y, D.X-C.X, D.Y-C.Y) >= 0 {
		return true
	}
	return false
}

func calCrossProduct(X1, Y1, X2, Y2 float64) float64 {
	return X1*Y2 - X2*Y1
}
