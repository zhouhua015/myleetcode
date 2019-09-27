package leetcode

import "testing"

func combs(ans [][]int, cur []int, n, k, b int) [][]int {
	if k == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		ans = append(ans, tmp)
		return ans
	}

	count := len(cur)
	for i := b; i <= n-k+1; i++ {
		cur = append(cur, i)
		ans = combs(ans, cur, n, k-1, i+1)
		cur = cur[:count]
	}
	return ans
}

func combine(n int, k int) [][]int {
	return combs(nil, nil, n, k, 1)
}

func TestCombine(t *testing.T) {
	cases := []struct {
		n, k int
	}{
		{4, 2},
	}
	for _, v := range cases {
		r := combine(v.n, v.k)
		t.Log(r)
	}
}
