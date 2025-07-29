package simple

// 普通做法
func countBits(num int) []int {
	res := make([]int, num+1)

	for i := 0; i <= num; i++ {
		res[i] = count(i)
	}
	return res

}

func count(num int) int {
	var cnt int
	for num > 0 {
		num &= num - 1
		cnt++
	}
	return cnt
}

func countBits2(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	for i := 1; i <= num; i++ {
		if res[i]%2 == 0 {
			res[i] = res[i>>1]
		} else {
			res[i] = res[i-1] + 1
		}
	}
	return res
}
