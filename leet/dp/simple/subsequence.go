package simple

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

func findSubsequence(s string, t string) bool {
	dp := make([][]bool, len(s)+1)
	// init
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, len(t)+1)
	}
	dp[0][0] = true
	//row
	for i := 1; i < len(s); i++ {
		dp[i][0] = false
	}
	for j := 1; j < len(t); j++ {
		dp[0][j] = true
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j < len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(t)]

}
