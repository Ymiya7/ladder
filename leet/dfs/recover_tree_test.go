package dfs

import (
	"github.com/Ymiya7/ladder/leet/common"
	"testing"
)

func Test_recoverTree(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{

			name: "case 1",
			args: args{
				root: &common.TreeNode{
					Val: 1,
					Left: &common.TreeNode{
						Val: 3,
						Right: &common.TreeNode{
							Val: 2,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.args.root)
		})
	}
}
