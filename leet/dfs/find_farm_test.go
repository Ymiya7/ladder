package dfs

import (
	"reflect"
	"testing"
)

func Test_findFarmland(t *testing.T) {
	type args struct {
		land [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				land: [][]int{
					{1, 0, 0},
					{0, 1, 1},
					{0, 1, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFarmland(tt.args.land); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFarmland() = %v, want %v", got, tt.want)
			}
		})
	}
}
