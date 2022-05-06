package exp2

import "testing"

func Test_routingOneWay(t *testing.T) {
	type args struct {
		maps [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{maps: [][]byte{{'w', 'w', 'w'}, {'w', 'w', 'w'}, {'y', 'y', 'b'}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routingOneWay(tt.args.maps)
		})
	}
}

func Test_routingOneWaySimple(t *testing.T) {
	type args struct {
		maps [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{maps: [][]byte{{'S', '.', '.'}, {'.', '#', '.'}, {'.', '.', 'T'}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routingOneWaySimple(tt.args.maps)
		})
	}
}
