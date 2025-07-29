package simple

func massage(nums []int) int {
	dp := make([]int, len(nums)) //	 表示选择 nums[i]时段的最优集合
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]

}
