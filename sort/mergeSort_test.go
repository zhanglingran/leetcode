package sort

import (
	"reflect"
	"testing"
)

func TestFoo(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case 1",
			args: args{
				arr1: []int{1, 3, 5, 7, 9},
				arr2: []int{2, 4, 6, 8, 10},
			},
		},
		{
			name: "case 2",
			args: args{
				arr1: []int{1, 1, 5, 6, 9},
				arr2: []int{2, 2, 4, 5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Foo(tt.args.arr1, tt.args.arr2)
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{3, 5, 1, 4},
			},
			want: []int{1, 3, 4, 5},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{3, 5, 1, 4, 2},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
