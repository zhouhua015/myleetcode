package leetcode

import "testing"

// Time Limit Exceeded
func longestValidParentheses(s string) int {
	dp := make([][]bool, len(s), len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s), len(s))
	}
	for i := 0; i+1 < len(s); i++ {
		dp[i][i+1] = s[i] == '(' && s[i+1] == ')'
	}

	for sz := 4; sz <= len(s); sz += 2 {
		for i := 0; i+sz <= len(s); i++ {
			if s[i] == '(' && s[i+sz-1] == ')' && dp[i+1][i+sz-2] {
				dp[i][i+sz-1] = true
				continue
			}

			for j := i + 1; j < i+sz-2; j += 2 {
				if dp[i][j] && dp[j+1][i+sz-1] {
					dp[i][i+sz-1] = true
					break
				}
			}
		}
	}

	var max int
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= 0; j-- {
			if dp[i][j] && max < j-i+1 {
				max = j - i + 1
			}
		}
	}
	return max
}

func longestValidParenthesesDP(s string) int {
	dp := make([]int, len(s), len(s))
	var max int
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = 2
				if i-2 >= 0 {
					dp[i] += dp[i-2]
				}
			} else if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
				dp[i] = dp[i-1] + 2
				if i-dp[i-1]-2 > 0 {
					dp[i] += dp[i-dp[i-1]-2]
				}
			}

			if max < dp[i] {
				max = dp[i]
			}
		}
	}
	return max
}

func longestValidParenthesesStack(s string) int {
	stack := make([]int, 0)
	stack = append(stack, -1)
	var max int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			sz := len(stack)
			if sz > 0 {
				stack = stack[:sz-1]
				sz = len(stack)
			}

			if sz == 0 {
				stack = append(stack, i)
			} else if max < i-stack[sz-1] {
				max = i - stack[sz-1]
			}
		}
	}

	return max
}

func longestValidParentheses2Rounds(s string) int {
	var c, o, max int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			o++
		case ')':
			c++
		default:
			return -1
		}

		if c > o {
			o, c = 0, 0
		} else if c == o && max < 2*c {
			max = 2 * c
		}
	}

	c, o = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '(':
			o++
		case ')':
			c++
		default:
			return -1
		}

		if o > c {
			o, c = 0, 0
		} else if c == o && max < 2*o {
			max = 2 * o
		}
	}

	return max
}

func TestLongestValidParentheses(t *testing.T) {
	testcases := []struct {
		input    string
		expected int
	}{
		{"(", 0},
		{")", 0},
		{"((", 0},
		{"(((", 0},
		{"()", 2},
		{"(()", 2},
		{"()(()", 2},
		{")()())", 4},
		{"()()())", 6},
		{"()()()()", 8},
		{"(())", 4},
		{"((())", 4},
		{"(()))", 4},
		{"(())()", 6},
		{"(())())", 6},
		{"((()))", 6},
		{"((()))())()", 8},
	}

	for i, c := range testcases {
		n := longestValidParentheses(c.input)
		if n != c.expected {
			t.Errorf("case #%d, want %d got %d\n", i+1, c.expected, n)
		}
	}
}
