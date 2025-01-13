package dfs

import "github.com/Ymiya7/ladder/leet/common"

func recoverTree(root *common.TreeNode) {

	var inorder func(node *common.TreeNode)
	var path []*common.TreeNode
	inorder = func(node *common.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		path = append(path, node)
		inorder(node.Right)
	}
	inorder(root)
	var (
		x *common.TreeNode
		y *common.TreeNode
	)
	for i := 0; i < len(path)-1; i++ {
		if path[i].Val > path[i+1].Val {
			y = path[i+1]
			if x == nil {
				x = path[i]
			}
		}
	}

	if x != nil && y != nil {
		x.Val, y.Val = y.Val, x.Val
	}

}
