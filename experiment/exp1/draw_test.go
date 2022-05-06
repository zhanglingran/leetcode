package exp1

import (
	"testing"
)

func TestPlotPicture(t *testing.T) {
	type args struct {
		count    int
		name     string
		f        func([][]int) [][]int
		needSort bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				count:    100,
				name:     "picture/case1.png",
				f:        ConvexHullForce,
				needSort: true,
			},
		},
		{
			name: "case2",
			args: args{
				count:    100,
				name:     "picture/case2.png",
				f:        ConvexHullGrahamScan,
				needSort: false,
			},
		},
		{
			name: "case3",
			args: args{
				count:    100,
				name:     "picture/case3.png",
				f:        ConvexHullDivideAndConquer,
				needSort: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PlotPicture(tt.args.count, tt.args.name, tt.args.f, tt.args.needSort)
		})
	}
}
