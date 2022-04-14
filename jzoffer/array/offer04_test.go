package array

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 3, 5, 7, 9},
				k:    7,
			},
			want:  3,
			want1: true,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 3, 5, 7, 9},
				k:    8,
			},
			want:  -4,
			want1: false,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 3, 5, 7, 9},
				k:    6,
			},
			want:  -3,
			want1: false,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{1, 3, 5, 7, 9},
				k:    4,
			},
			want:  -2,
			want1: false,
		},
		{
			name: "case 5",
			args: args{
				nums: []int{1, 3, 5, 7, 9},
				k:    2,
			},
			want:  -1,
			want1: false,
		},
		{
			name: "case 6",
			args: args{
				nums: []int{1, 3, 5, 7, 9},
				k:    -2,
			},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := binarySearch(tt.args.nums, tt.args.k)
			if got != tt.want {
				t.Errorf("binarySearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("binarySearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_findNumberIn2DArray(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 5,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 16,
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
				},
				target: 16,
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				matrix: [][]int{
					{1},
				},
				target: 1,
			},
			want: true,
		},
		{
			name: "case 5",
			args: args{
				matrix: [][]int{
					{},
				},
				target: 1,
			},
			want: false,
		},
		{
			name: "case 6",
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
					{21, 22, 23, 24, 25},
				},
				target: 19,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberIn2DArray(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("findNumberIn2DArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNumberIn2DArray1(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 5,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 16,
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
				},
				target: 16,
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				matrix: [][]int{
					{1},
				},
				target: 1,
			},
			want: true,
		},
		{
			name: "case 5",
			args: args{
				matrix: [][]int{
					{},
				},
				target: 1,
			},
			want: false,
		},
		{
			name: "case 6",
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
					{21, 22, 23, 24, 25},
				},
				target: 19,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberIn2DArray1(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("findNumberIn2DArray1() = %v, want %v", got, tt.want)
			}
		})
	}
}
