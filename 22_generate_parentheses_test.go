package leetcode

import "testing"

func generateParenthesisRecursive(n int) []string {
	if n == 1 {
		return []string{"()"}
	}

	p := generateParenthesisRecursive(n - 1)
	var parentheses []string
	visited := make(map[string]bool)
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[i]); j++ {
			ns := p[i][:j] + "()" + p[i][j:]
			if !visited[ns] {
				parentheses = append(parentheses, ns)
				visited[ns] = true
			}
		}
	}
	return parentheses
}

func backtrace(list []string, cur string, open, close, n int) []string {
	if len(cur) == n*2 {
		list = append(list, cur)
		return list
	}

	if open < n {
		list = backtrace(list, cur+"(", open+1, close, n)
	}
	if close < open {
		list = backtrace(list, cur+")", open, close+1, n)
	}

	return list
}

func generateParenthesisBacktrace(n int) []string {
	return backtrace(nil, "", 0, 0, n)
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	var parentheses []string
	for i := 0; i < n; i++ {
		left := generateParenthesis(i)
		for j := 0; j < len(left); j++ {
			right := generateParenthesis(n - 1 - i)
			for k := 0; k < len(right); k++ {
				parentheses = append(parentheses, "("+left[j]+")"+right[k])
			}
		}
	}
	return parentheses
}

func TestGenParentheses(t *testing.T) {
	cases := []struct {
		n           int
		parentheses []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for _, c := range cases {
		p := generateParenthesis(c.n)
		t.Log(p)
	}
}
