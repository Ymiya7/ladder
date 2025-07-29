package simple

func divisor(n int) bool {
	if n == 1 {
		return false
	}
	dp := make([]bool, n+1)
	dp[1] = false
	dp[2] = true
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !dp[i-j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]

}
