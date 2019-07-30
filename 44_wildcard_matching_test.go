package leetcode

import "testing"

func isMatchSlice(s, p []byte) bool {
	if len(s) == 0 || len(p) == 0 {
		if len(p) > 0 {
			for i := 0; i < len(p); i++ {
				if p[i] != '*' {
					return false
				}
			}
			return true
		}

		return len(s) == 0 && len(p) == 0
	}

	if p[0] == '*' {
		return isMatchSlice(s, p[1:]) || isMatchSlice(s[1:], p[1:]) || isMatchSlice(s[1:], p)
	}
	return (s[0] == p[0] || p[0] == '?') && isMatchSlice(s[1:], p[1:])
}

func isWildcardMatchNaive(s string, p string) bool {
	return isMatchSlice([]byte(s), []byte(p))
}

func isWildcardMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			if p[j] == '*' {
				dp[i][j] = dp[i][j+1] || (i < len(s) && (dp[i+1][j+1] || dp[i+1][j]))
			} else {
				first := i < len(s) && (s[i] == p[j] || p[j] == '?')
				dp[i][j] = first && dp[i+1][j+1]
			}
		}
	}

	return dp[0][0]
}

func TestWildcardMatch(t *testing.T) {
	testcases := []struct {
		s, p string
		m    bool
	}{
		{"", "*", true},
		{"aa", "", false},
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", "?*", true},
		{"aa", "*", true},
		{"aacd", "?*", true},
		{"aacd", "a*", true},
		{"aacd", "a?cd", true},
		{"aacd", "*a?cd", true},
		{"aacd", "*cd", true},
		{"aacd", "a*c*d", true},
	}

	for i := range testcases {
		if isWildcardMatch(testcases[i].s, testcases[i].p) != testcases[i].m {
			t.Errorf("case #%d, want %t got %t\n", i+1, testcases[i].m, !testcases[i].m)
		}
	}
}
