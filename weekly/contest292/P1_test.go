package contest292

import "testing"

func Test_largestGoodInteger(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{num: "6777133339"},
			want: "777",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestGoodInteger(tt.args.num); got != tt.want {
				t.Errorf("largestGoodInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
