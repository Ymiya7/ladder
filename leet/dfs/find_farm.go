package dfs

func findFarmland(land [][]int) [][]int {

	m := len(land)
	n := len(land[0])
	var ans [][]int
	var dfs func([][]int, int, int, []int)
	dfs = func(land [][]int, i int, j int, path []int) {
		if i < 0 || j < 0 || i >= len(land) || j >= len(land[0]) || land[i][j] != 1 {
			return
		}

		land[i][j] = 0
		if i < path[0] || j < path[1] {
			path[0] = i
			path[1] = j
		}

		if i > path[2] || j > path[3] {
			path[2] = i
			path[3] = j
		}
		dfs(land, i+1, j, path) // 向下
		dfs(land, i-1, j, path) //向上
		dfs(land, i, j-1, path) // 向左
		dfs(land, i, j+1, path) // 向右
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			path := []int{i, j, i, j}
			if land[i][j] == 1 {
				dfs(land, i, j, path)
				ans = append(ans, path)
			}
		}
	}
	return ans
}
