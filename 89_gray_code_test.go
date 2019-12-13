package leetcode

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	codes := grayCode(n - 1)
	l := 2 * len(codes)
	ans := make([]int, l)
	for i, c := range codes {
		ans[i] = c << 1
		ans[l-i-1] = c<<1 + 1
	}
	return ans
}
