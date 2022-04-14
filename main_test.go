package main

import "testing"

func Test_getWordsCount(t *testing.T) {
	type args struct {
		words []string
		chars string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				words: []string{"cat", "bt", "hat", "tree"},
				chars: "atacht",
			},
			want: 6,
		},
		{
			name: "case2",
			args: args{
				words: []string{"hello", "world", "leetcode"},
				chars: "welldonehoneyrl",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWordsCount(tt.args.words, tt.args.chars); got != tt.want {
				t.Errorf("getWordsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
