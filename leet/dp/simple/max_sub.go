package simple

func maxSubArray(nums []int) int {

	dp := make([]int, len(nums)) // 表示已数字nums[i]结尾的最大子数组的和
	dp[0] = nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}
