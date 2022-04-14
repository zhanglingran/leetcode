package array

import (
	"reflect"
	"testing"
)

func Test_fakeMain(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "case1",
			want: []int{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fakeMain(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fakeMain() = %v, want %v", got, tt.want)
			}
		})
	}
}
