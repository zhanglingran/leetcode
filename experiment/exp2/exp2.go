package exp2

/*
 * 单向A*算法
 */

var (
	// 八个方向
	direct = [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
)

/*
 * 单向A*: 不考虑地形代价，只考虑移动代价。
 * #表示障碍; .表示可通行
 * 输入地图，输出代价
 */
func routingOneWaySimple(maps [][]byte) int {
	// 遍历节点，计算代价，找到最短的路径

	return 0
}

/*
 * 单向A*
 */
func routingOneWay(maps [][]byte) {

	// 遍历节点，计算代价，找到最短的路径

}

// 双向A*算法
func routingTwoWay() {

}
