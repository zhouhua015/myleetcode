package leetcode

import "testing"

func lengthOfLastWord(s string) int {
	pos := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && pos < 0 {
			pos = i
		} else if s[i] == ' ' && pos >= 0 {
			return pos - i
		}
	}

	return pos + 1
}

func TestLengthOfLastWord(t *testing.T) {
	cases := []struct {
		s string
		e int
	}{
		{"hello world", 5},
		{"hello world   ", 5},
		{"helloworld   ", 10},
		{"             ", 0},
		{"hello", 5},
		{" hello ", 5},
		{"a", 1},
	}

	for k, v := range cases {
		r := lengthOfLastWord(v.s)
		if r != v.e {
			t.Errorf("case #%d, want %d got %d\n", k+1, v.e, r)
		}
	}
}
