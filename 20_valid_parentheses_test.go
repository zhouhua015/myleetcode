package leetcode

import "testing"

var opposites = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else if c == ')' || c == '}' || c == ']' {
			if len(stack) < 1 {
				return false
			}
			if stack[len(stack)-1] != opposites[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func TestValidParentheses(t *testing.T) {
	cases := []struct {
		s     string
		valid bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for i, c := range cases {
		r := isValid(c.s)
		if r != c.valid {
			t.Errorf("case #%d, want %t got %t\n", i+1, c.valid, r)
		}
	}

}
