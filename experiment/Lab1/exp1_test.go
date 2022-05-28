package Lab1

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"log"
	"testing"
	"time"
)

func TestConvexHullForce(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{points: generatePointsData(0, 100, 0, 100, 1000)},
		},
		{
			name: "case2",
			args: args{points: generatePointsData(0, 100, 0, 100, 2000)},
		},
		{
			name: "case3",
			args: args{points: generatePointsData(0, 100, 0, 100, 3000)},
		},
		{
			name: "case4",
			args: args{points: generatePointsData(0, 100, 0, 100, 4000)},
		},
		{
			name: "case5",
			args: args{points: generatePointsData(0, 100, 0, 100, 5000)},
		},
		{
			name: "case6",
			args: args{points: generatePointsData(0, 100, 0, 100, 6000)},
		},
		{
			name: "case7",
			args: args{points: generatePointsData(0, 100, 0, 100, 7000)},
		},
		{
			name: "case8",
			args: args{points: generatePointsData(0, 100, 0, 100, 8000)},
		},
		{
			name: "case9",
			args: args{points: generatePointsData(0, 100, 0, 100, 9000)},
		},
		{
			name: "case10",
			args: args{points: generatePointsData(0, 100, 0, 100, 10000)},
		},
	}

	timesForce := make([]int64, len(tests))
	timesGrahamScan := make([]int64, len(tests))
	timesDivideConquer := make([]int64, len(tests))

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 暴力算法耗时
			start := time.Now()
			_ = ConvexHullForce(tt.args.points)
			end := time.Now()
			tL := end.Sub(start).Milliseconds()
			fmt.Println("暴力 算法运行时长:", tL, "ms")
			//GeneratePicture(tt.name+".png", true, tt.args.points, got)
			timesForce[i] = tL

			// Graham-Scan算法耗时
			start = time.Now()
			_ = ConvexHullGrahamScan(tt.args.points)
			end = time.Now()
			tL = end.Sub(start).Milliseconds()
			fmt.Println("Graham-Scan 算法运行时长:", tL, "ms")
			timesGrahamScan[i] = tL

			// 分治算法耗时
			start = time.Now()
			_ = ConvexHullDivideAndConquer(tt.args.points)
			end = time.Now()
			tL = end.Sub(start).Milliseconds()
			fmt.Println("分治 算法运行时长:", tL, "ms")
			timesDivideConquer[i] = tL
		})
	}

	p := plot.New()
	p.Title.Text = "三种算法性能对比"
	p.X.Label.Text = "数据规模"
	p.Y.Label.Text = "性能(ms)"
	p.X.Label.Position = draw.PosRight
	p.Y.Label.Position = draw.PosTop

	plotutil.AddLinePoints(p,
		"Force", generatePointsForExp1(timesForce),
		"Graham-Scan", generatePointsForExp1(timesGrahamScan),
		"Divide-And-Conquer", generatePointsForExp1(timesDivideConquer),
	)

	// 将散点图保存起来
	err := p.Save(20*vg.Centimeter, 20*vg.Centimeter, "performance_exp1.png")
	if err != nil {
		log.Fatalf("could not save plot: %+v", err)
	}
}

func generatePointsForExp1(data []int64) plotter.XYs {
	points := make(plotter.XYs, len(data))
	for i := 0; i < len(data); i++ {
		points[i].X = float64((i + 1) * 1000)
		points[i].Y = float64(data[i])
	}
	return points
}
