package array

import (
	"reflect"
	"testing"
)

func Test_constructArr(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{a: []int{1, 2, 3, 4, 5}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructArr(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_constructArr2(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{a: []int{1, 2, 3, 4, 5}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructArr2(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructArr2() = %v, want %v", got, tt.want)
			}
		})
	}
}
