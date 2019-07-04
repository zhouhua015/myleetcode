package leetcode

import (
	"sort"
	"testing"
)

var letters = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return nil
	}

	if len(digits) == 1 {
		return letters[digits[0]]
	}

	dp := make([][]string, len(digits))
	i := len(digits) - 1
	dp[i] = letters[digits[i]]

	i--
	for ; i >= 0; i-- {
		l := letters[digits[i]]
		for j := 0; j < len(l); j++ {
			for k := 0; k < len(dp[i+1]); k++ {
				dp[i] = append(dp[i], l[j]+dp[i+1][k])
			}

		}
	}
	return dp[0]
}

func TestLetterCombinations(t *testing.T) {
	testcases := []struct {
		digits   string
		expected []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}

	for i, c := range testcases {
		sort.Slice(c.expected, func(i, j int) bool { return c.expected[i] < c.expected[j] })

		comb := letterCombinations(c.digits)
		sort.Slice(comb, func(i, j int) bool { return comb[i] < comb[j] })

		if len(c.expected) != len(comb) {
			t.Errorf("case #%d, count mismatched, want %v got %v\n", i+1, c.expected, comb)
		}

		for j := 0; j < len(c.expected); j++ {
			if c.expected[j] != comb[j] {
				t.Errorf("case #%d, %dth element, want '%s' got '%s'\n", i+1, j+1, c.expected[j], comb[j])
			}
		}
	}
}
