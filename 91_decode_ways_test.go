package leetcode

import "testing"

func numDecodings(s string) int {
	if len(s) == 0 {
		return 1
	}
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	if s[:2] > "26" {
		return numDecodings(s[1:])
	}
	return numDecodings(s[1:]) + numDecodings(s[2:])
}

func numDecodingsDP(s string) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	if s[len(s)-1] != '0' {
		dp[len(s)-1] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
			continue
		}

		if s[i:i+2] > "26" {
			dp[i] = dp[i+1]
		} else {
			dp[i] = dp[i+1] + dp[i+2]
		}
	}
	return dp[0]
}

func TestDecodeWays(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{"12", 2},
		{"226", 3},
		{"012", 0},
		{"101", 1},
		{"100", 0},
	}

	for k, v := range cases {
		// ans := numDecodings(v.s)
		ans := numDecodingsDP(v.s)
		if ans != v.expected {
			t.Fatalf("case #%d, want %d got %d\n", k+1, v.expected, ans)
		}
	}
}
