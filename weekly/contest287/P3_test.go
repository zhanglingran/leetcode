package contest287

import "testing"

func Test_maximumCandies(t *testing.T) {
	type args struct {
		candies []int
		k       int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				candies: []int{1, 2, 1},
				k:       4,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCandies(tt.args.candies, tt.args.k); got != tt.want {
				t.Errorf("maximumCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
