package Lab1

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"log"
)

// PlotPicture (startX, startY),(endX, endY) 分别表示点集对角线构成的矩形
func PlotPicture(count int, name string, f func([][]int) [][]int, needSort bool) {

	startX, startY, endX, endY := 0, 0, 100, 100

	p := plot.New()
	p.Title.Text = "Convex"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.X.Label.Position = draw.PosRight
	p.Y.Label.Position = draw.PosTop
	// 散点图横走坐标的范围
	p.X.Min = float64(startX - 5)
	p.X.Max = float64(endX + 5)
	p.Y.Min = float64(startY - 5)
	p.Y.Max = float64(endY + 5)

	// 生成的点集数据 data
	data := generatePointsData(startX, endX, startY, endY, count)
	// 计算得到凸包点集
	convex := f(data)

	if needSort {
		// 对 凸包点集 按照顺时针顺序进行排序
		sumX, sumY := 0, 0
		// 得到重心 作为 原点坐标
		for i := 0; i < len(convex); i++ {
			sumX += convex[i][0]
			sumY += convex[i][1]
		}
		originX, originY := sumX/len(convex), sumY/len(convex)
		SortByAngle(convex, []int{originX, originY})
	}

	// 生成散点图
	plotutil.AddScatters(p,
		"", insertData(len(data), data),
		"", plotter.XYs{},
		"", plotter.XYs{},
		"", plotter.XYs{},
		"", plotter.XYs{},
		"", insertData(len(convex), convex),
	)

	// 连接凸包点集, 生成边界线
	plotutil.AddLines(p, "", insertData(len(convex), convex))

	// 将最后一个点和第一个点连接起来
	plotutil.AddLines(p, "", insertData(2, [][]int{convex[0], convex[len(convex)-1]}))

	// 将散点图保存起来
	err := p.Save(20*vg.Centimeter, 20*vg.Centimeter, name)
	if err != nil {
		log.Fatalf("could not save plot: %+v", err)
	}
}

// GeneratePicture 根据传入数据生成
func GeneratePicture(name string, needSort bool, points, convex [][]int) {

	startX, startY, endX, endY := 0, 0, 100, 100

	p := plot.New()
	p.Title.Text = "Convex"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.X.Label.Position = draw.PosRight
	p.Y.Label.Position = draw.PosTop
	// 散点图横走坐标的范围
	p.X.Min = float64(startX - 5)
	p.X.Max = float64(endX + 5)
	p.Y.Min = float64(startY - 5)
	p.Y.Max = float64(endY + 5)

	if needSort {
		// 对 凸包点集 按照顺时针顺序进行排序
		sumX, sumY := 0, 0
		// 得到重心 作为 原点坐标
		for i := 0; i < len(convex); i++ {
			sumX += convex[i][0]
			sumY += convex[i][1]
		}
		originX, originY := sumX/len(convex), sumY/len(convex)
		SortByAngle(convex, []int{originX, originY})
	}

	// 生成散点图
	plotutil.AddScatters(p,
		"", insertData(len(points), points),
		"", plotter.XYs{},
		"", plotter.XYs{},
		"", plotter.XYs{},
		"", plotter.XYs{},
		"", insertData(len(convex), convex),
	)

	// 连接凸包点集, 生成边界线
	plotutil.AddLines(p, "", insertData(len(convex), convex))

	// 将最后一个点和第一个点连接起来
	plotutil.AddLines(p, "", insertData(2, [][]int{convex[0], convex[len(convex)-1]}))

	// 将散点图保存起来
	err := p.Save(20*vg.Centimeter, 20*vg.Centimeter, name)
	if err != nil {
		log.Fatalf("could not save plot: %+v", err)
	}
}

// 为画布构建点集
func insertData(n int, data [][]int) plotter.XYs {

	points := make(plotter.XYs, n)
	for i := 0; i < n; i++ {
		points[i].X = float64(data[i][0])
		points[i].Y = float64(data[i][1])
	}
	return points
}
