package dfs

import (
	"sort"
)

func findFarmland(land [][]int) [][]int {

	m := len(land)
	n := len(land[0])
	var ans [][]int
	var dfs func([][]int, int, int, *[]int)
	dfs = func(land [][]int, i int, j int, path *[]int) {
		if i < 0 || j < 0 || i >= len(land) || j >= len(land[0]) || land[i][j] != 1 {
			return
		}

		land[i][j] = 0
		*path = append(*path, i*len(land)+j)
		dfs(land, i+1, j, path) // 向下
		dfs(land, i-1, j, path) //向上
		dfs(land, i, j-1, path) // 向左
		dfs(land, i, j+1, path) // 向右
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			path := make([]int, 0)
			if land[i][j] == 1 {
				dfs(land, i, j, &path)
				cors := process(path, n)
				ans = append(ans, cors)
			}
		}
	}
	return ans
}

func process(path []int, n int) []int {
	// 处理
	sort.Ints(path)
	cors := make([]int, 4)
	min := path[0]
	startLeft := min / n
	cors[0] = startLeft
	startRight := min - startLeft*n
	cors[1] = startRight
	max := path[len(path)-1]
	endLeft := max / n
	cors[2] = endLeft
	endRight := max - endLeft*n
	cors[3] = endRight

	return cors
}
