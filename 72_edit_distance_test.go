package leetcode

import (
	"testing"
)

func min3ways(a, b, c int) int {
	d := a
	if d > b {
		d = b
	}

	if d > c {
		d = c
	}

	return d
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := m; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			if i == m {
				dp[i][j] = n - j
			} else if j == n {
				dp[i][j] = m - i
			} else if word1[i] == word2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = 1 + min3ways(dp[i+1][j], dp[i][j+1], dp[i+1][j+1])
			}
		}
	}

	return dp[0][0]
}

func TestEditDistance(t *testing.T) {
	cases := []struct {
		word1 string
		word2 string
		dist  int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"sss", "s", 2},
		{"", "s", 1},
		{"a", "", 1},
		{"a", "s", 1},
	}

	for k, v := range cases {
		ans := minDistance(v.word1, v.word2)
		if ans != v.dist {
			t.Errorf("case #%d, want %d got %d\n", k+1, v.dist, ans)
		}
	}
}
