package leetcode

import "testing"

func longestPalindrome(s string) string {
	bytes := []byte(s)
	var longest string
	palindromes := make(map[string]bool)
	for size := 1; size <= len(bytes); size++ {
		for i := 0; i <= len(bytes)-size; i++ {
			substr := bytes[i : i+size]
			if size == 1 {
				palindromes[string(substr)] = true
				longest = string(substr)
			} else if size == 2 && substr[0] == substr[1] {
				palindromes[string(substr)] = true
				longest = string(substr)
			} else {
				if substr[0] != substr[size-1] {
					continue
				}

				if palindromes[string(substr[1:size-1])] {
					palindromes[string(substr)] = true
					longest = string(substr)
				}
			}
		}
	}

	return longest
}

func TestLongestPalindrome(t *testing.T) {
	testcases := []struct {
		str      string
		expected map[string]bool
	}{
		{"babad", map[string]bool{"bab": true, "aba": true}},
		{"cbbd", map[string]bool{"bb": true}},
	}

	f := func(s string, expected map[string]bool) {
		expects := make([]string, 0, len(expected))
		for k, _ := range expected {
			expects = append(expects, k)
		}

		r := longestPalindrome(s)
		if len(r) == 0 {
			t.Fatalf("wrong answer, length should always > 0, '%s'\n", s)
		}

		if expected[r] {
			t.Logf("passed test case '%s', got '%s'\n", s, r)
		} else {
			t.Errorf("want %v got %s, test case '%s'\n", expects, r, s)
		}
	}

	for i := 0; i < len(testcases); i++ {
		f(testcases[i].str, testcases[i].expected)
	}
}
