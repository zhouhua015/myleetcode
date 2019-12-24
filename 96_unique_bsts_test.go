package leetcode

func nTrees(low, high int, cache [][]int) int {
	if low == high {
		return 0
	}
	if cache[low][high] > 0 {
		return cache[low][high]
	}

	var n int
	for i := low; i < high; i++ {
		nl := nTrees(low, i, cache)
		nr := nTrees(i+1, high, cache)
		if nl == 0 && nr == 0 {
			n++
		} else if nl == 0 {
			n += nr
		} else if nr == 0 {
			n += nl
		} else {
			n += nl * nr
		}
	}
	cache[low][high] = n
	return n
}

func numTrees(n int) int {
	// dp := make([][]int, n+1)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([]int, n+1)
	// }
	// return nTrees(0, n, dp)

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}
