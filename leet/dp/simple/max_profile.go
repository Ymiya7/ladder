package simple

import (
	"math"
)

// 121 买卖股票的最佳时机  常规做法 超时
func maxProfit(prices []int) int {
	ans := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			ans = maxInt(ans, prices[j]-prices[i])
		}
	}
	return ans
}

// 优化第一步 j 开始找 后续的最大的节点 超时
func maxProfit2(prices []int) int {

	ans := 0
	for i := 0; i < len(prices)-1; i++ {
		postMax := prices[i+1]
		for j := i + 2; j < len(prices); j++ {
			if postMax < prices[j] {
				postMax = prices[j]
			}
		}
		ans = maxInt(ans, postMax-prices[i])
	}
	return ans
}

func maxProfit_(prices []int) int {
	ans := 0
	for i := 1; i < len(prices)-1; i++ {

		minPrice := prices[0]
		for j := 0; j <= i; j++ {
			minPrice = min(minPrice, prices[j])
		}

		maxPrice := prices[i+1]
		for k := i + 1; k < len(prices); k++ {
			maxPrice = max(maxPrice, prices[k])
		}
		ans = maxInt(ans, maxPrice-minPrice)
	}
	return ans
}

// 优化第二步 j往前找到最小的点 j往后 找到最大的节点
func maxProfit3(prices []int) int {
	// 今天卖出 核心点 是今天卖出 然后顺便记录每天卖出点前的最小值
	minPrice := math.MaxInt
	ans := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > ans {
			ans = prices[i] - minPrice
		}
	}
	return ans
}

func maxProfit4(prices []int) int {
	minPrice := prices[0]
	mp := 0
	for _, price := range prices {
		mp = max(mp, price-minPrice)
		minPrice = min(minPrice, price)
	}
	return mp
}

func maxProfit_dp(prices []int) int {
	dp := make([]int, len(prices)) // 第 i天 获取的利润最大
	dp[0] = 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			prices[i] = minPrice
		}

		dp[i] = max(dp[i-1], prices[i]-minPrice)
	}
	return dp[len(dp)-1]

}

func maxProfit_stack(prices []int) int {
	ans := 0
	var stack []int
	for i := 0; i < len(prices); i++ {
		for len(stack) > 0 && prices[stack[len(stack)-1]] > prices[i] {
			bottom := stack[0]
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = max(ans, top-bottom)
		}
		stack = append(stack, prices[i])
	}
	return ans
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
