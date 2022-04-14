package divide_and_conquer

import "testing"

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
				arr1: []int{3, 5},
				arr2: []int{2, 4},
			},
		},
		{
			name: "case 2",
			args: args{
				arr1: []int{2, 3},
				arr2: []int{2, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Foo(tt.args.arr1, tt.args.arr2)
		})
	}
}

func Test_reversePairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{7, 5, 6, 4},
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{3, 2, 2, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePairs(tt.args.nums); got != tt.want {
				t.Errorf("reversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
