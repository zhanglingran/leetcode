package Lab4

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"log"
	"sort"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		A []int
		p int
		r int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{A: GenerateData(1e6, 0), p: 0, r: 1e6 - 1},
		},
		{
			name: "case2",
			args: args{A: GenerateData(1e6, 1), p: 0, r: 1e6 - 1},
		},
		{
			name: "case3",
			args: args{A: GenerateData(1e6, 2), p: 0, r: 1e6 - 1},
		},
		{
			name: "case4",
			args: args{A: GenerateData(1e6, 3), p: 0, r: 1e6 - 1},
		},
		{
			name: "case5",
			args: args{A: GenerateData(1e6, 4), p: 0, r: 1e6 - 1},
		},
		{
			name: "case6",
			args: args{A: GenerateData(1e6, 5), p: 0, r: 1e6 - 1},
		},
		{
			name: "case7",
			args: args{A: GenerateData(1e6, 6), p: 0, r: 1e6 - 1},
		},
		{
			name: "case8",
			args: args{A: GenerateData(1e6, 7), p: 0, r: 1e6 - 1},
		},
		{
			name: "case9",
			args: args{A: GenerateData(1e6, 8), p: 0, r: 1e6 - 1},
		},
		{
			name: "case10",
			args: args{A: GenerateData(1e6, 9), p: 0, r: 1e6 - 1},
		},
		{
			name: "case11",
			args: args{A: GenerateData(1e6, 10), p: 0, r: 1e6 - 1},
		},
	}

	timesBefore := make([]int64, len(tests))
	timesAfter := make([]int64, len(tests))
	timesLib := make([]int64, len(tests))

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			QuickSortBefore(tt.args.A, tt.args.p, tt.args.r)
			end := time.Now()
			timesBefore[i] = end.Sub(start).Milliseconds()

			start = time.Now()
			QuickSortAfter(tt.args.A, tt.args.p, tt.args.r)
			end = time.Now()
			timesAfter[i] = end.Sub(start).Milliseconds()

			start = time.Now()
			sort.Slice(tt.args.A, func(i, j int) bool {
				return tt.args.A[i] < tt.args.A[j]
			})
			end = time.Now()
			timesLib[i] = end.Sub(start).Milliseconds()

			fmt.Println(tt.name, "时间为", timesAfter[i], "ms")
			fmt.Println(tt.name, "时间为", timesLib[i], "ms")
		})
	}

	p := plot.New()
	p.Title.Text = "三种算法性能对比"
	p.X.Label.Text = "数据重复率(%)"
	p.Y.Label.Text = "性能(ms)"
	p.X.Label.Position = draw.PosRight
	p.Y.Label.Position = draw.PosTop

	plotutil.AddLinePoints(p,
		"Before", generatePointsForExp4(timesBefore),
		"After", generatePointsForExp4(timesAfter),
		"LibFunc", generatePointsForExp4(timesLib),
	)

	// 将散点图保存起来
	err := p.Save(20*vg.Centimeter, 20*vg.Centimeter, "performance_exp4_1.png")
	if err != nil {
		log.Fatalf("could not save plot: %+v", err)
	}
}

func generatePointsForExp4(data []int64) plotter.XYs {
	points := make(plotter.XYs, len(data))
	for i := 0; i < len(data); i++ {
		points[i].X = float64(i * 10)
		points[i].Y = float64(data[i])
	}
	return points
}
