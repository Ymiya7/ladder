package leet

import "sort"

// getKth 将整数按权重排序
func getKth(lo int, hi int, k int) int {
	nums := []int{}
	for i := lo; i <= hi; i++ {
		nums = append(nums, i)
	}

	sort.Slice(nums, func(i, j int) bool {
		if dealNum(nums[i]) != dealNum(nums[j]) {
			return dealNum(nums[i]) < dealNum(nums[j])
		}
		return nums[i] < nums[j]
	})
	return nums[k-1]
}

// getNum 获取某个数的权重
func getNum(num int) int {
	if num == 1 {
		return 0
	}
	if num%2 == 0 {
		return 1 + getNum(num/2)
	}
	return 1 + getNum(num*3+1)
}

var memo = map[int]int{}

// dealNum 记忆化搜索
func dealNum(num int) int {
	if val, ok := memo[num]; ok {
		return val
	}
	if num == 1 {
		memo[num] = 0
		return 0
	}
	if num%2 == 0 {
		memo[num] = getNum(num/2) + 1
	} else {
		memo[num] = getNum(num*3+1) + 1
	}
	return memo[num]
}
