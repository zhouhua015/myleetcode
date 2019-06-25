package leetcode

import (
	"testing"
)

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			first := (i < len(s) && (s[i] == p[j] || p[j] == '.'))
			if j+1 < len(p) && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || (first && dp[i+1][j])
			} else {
				dp[i][j] = first && dp[i+1][j+1]
			}
		}
	}

	return dp[0][0]
}

func TestRegExpMatch(t *testing.T) {
	testcases := []struct {
		s, p string
		m    bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"aab", "a*bc*", true},
		{"aab", "a*c*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"", "", true},
		{"xxxx", "", false},
		{"", ".", false},
		{"", ".*", true},
		{"", ".*a*", true},
		{"", ".*a", false},
		{"", "a*", true},
		{"abcd", "d*", false},
		{"ab", ".*c", false},
		{"aaa", "aaaa", false},
		{"aaa", "a*b*", true},
		{"aaa", "a*b*c", false},
		{"aaa", "a*a", true},
	}

	for i := range testcases {
		if isMatch(testcases[i].s, testcases[i].p) != testcases[i].m {
			t.Errorf("case #%d, want %t got %t\n", i+1, testcases[i].m, !testcases[i].m)
		}
	}
}
