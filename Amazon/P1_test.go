package Amazon

import "testing"

func Test_subString(t *testing.T) {
	type args struct {
		search string
		result string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "case1",
			args: args{
				search: "amarz",
				result: "amarzon",
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				search: "aa",
				result: "aba",
			},
			want: 2,
		},
		{
			name: "case3",
			args: args{
				search: "aaaaa",
				result: "aaaabbbbaaaa",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subString(tt.args.search, tt.args.result); got != tt.want {
				t.Errorf("subString() = %v, want %v", got, tt.want)
			}
		})
	}
}
