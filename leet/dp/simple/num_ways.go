package simple

func numWays(n int, relation [][]int, k int) int {
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for i := 1; i <= k; i++ {
		for _, r := range relation {
			start := r[0]
			end := r[1]
			dp[i][end] += dp[i-1][start]
		}
	}

	return dp[k][n-1]

}

// 使用 bfs 求解 从编号 0 到编号 n-1 走 k步 统计队列中编号为 n-1的节点的出现次数
func numWays2(n int, relation [][]int, k int) int {
	relationMap := make(map[int][]int)
	for _, r := range relation {
		relationMap[r[0]] = append(relationMap[r[0]], r[1])
	}
	queue := make([]int, 0)
	queue = append(queue, 0) // 编号 0
	for len(queue) > 0 && k > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			//获取它的边
			nodes := relationMap[node]
			if len(nodes) == 0 {
				continue
			}
			for _, nd := range nodes {
				queue = append(queue, nd)
			}
		}
		k--
	}
	ans := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == n-1 {
			ans++
		}
	}
	return ans
}

// dfs 求解
var ans int

func numWays3(n int, relation [][]int, k int) int {
	dfs(0, 0, k, n, relation)
	return ans
}

func dfs(node, step, n, k int, relation [][]int) {
	if step == k {
		if node == n-1 {
			ans++
		}
		return
	}
	for _, r := range relation {
		if r[0] == node {
			dfs(r[1], step+1, n, k, relation)
		}
	}
}
