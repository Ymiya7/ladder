package simple

func maxSales(sales []int) int {

	dp := make([]int, len(sales))
	maxSale := sales[0]
	dp[0] = sales[0]
	for i := 1; i < len(sales); i++ {
		dp[i] = max(dp[i-1]+sales[i], sales[i])
		maxSale = max(maxSale, dp[i])
	}
	return maxSale
}
