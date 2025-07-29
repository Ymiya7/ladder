package simple

func leastMinutes(n int) int {

	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 1; i <= n; i++ {
		dp[i] = dp[(i+1)/2] + 1
	}
	return dp[n]
}
