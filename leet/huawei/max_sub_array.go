package huawei

func maxSubArrayLen(nums []int, k int) int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	sumMap := make(map[int]int)
	sumMap[prefix[0]] = 0
	maxLen := 0
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
		if prefix[i] == k {
			maxLen = i + 1
		}
		if _, ok := sumMap[prefix[i]-k]; ok {
			maxLen = max(i-sumMap[prefix[i]-k], maxLen)
		}
		if _, ok := sumMap[prefix[i]]; !ok {
			sumMap[prefix[i]] = i
		}
	}

	return maxLen
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
