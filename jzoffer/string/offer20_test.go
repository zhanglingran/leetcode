package string

import "testing"

func Test_isNumber1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case2",
			args: args{s: "0e"},
			want: false,
		},
		{
			name: "case3",
			args: args{s: "-1E-16"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber1(tt.args.s); got != tt.want {
				t.Errorf("isNumber1() = %v, want %v", got, tt.want)
			}
		})
	}
}
