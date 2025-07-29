package simple

import "testing"

func Test_numWays3(t *testing.T) {
	type args struct {
		n        int
		relation [][]int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				n:        3,
				relation: [][]int{{0, 2}, {2, 1}},
				k:        2,
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				n:        5,
				relation: [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}},
				k:        3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays3(tt.args.n, tt.args.relation, tt.args.k); got != tt.want {
				t.Errorf("numWays3() = %v, want %v", got, tt.want)
			}
		})
	}
}
