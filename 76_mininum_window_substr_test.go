package leetcode

import "testing"

func findMinWindow(s string, c map[rune]int, n int) string {
	hit := make(map[rune]int)
	index := make([]int, 0, len(s))
	var meet, start int
	var min string
	for i, r := range s {
		if _, ok := c[r]; !ok {
			continue
		}

		index = append(index, i)
		hit[r]++
		if hit[r] <= c[r] {
			meet++
		}

		if meet == n {
			idx := index[start]
			start++
			if len(min) == 0 || len(min) > i-idx+1 {
				min = string(s[idx : i+1])
			}

			ch := rune(s[idx])
			hit[ch]--
			for hit[ch] >= c[ch] {
				idx = index[start]
				start++

				if len(min) > i-idx+1 {
					min = string(s[idx : i+1])
				}
				ch = rune(s[idx])
				hit[ch]--
			}
			meet--
		}
	}
	return min
}

func minWindow(s string, t string) string {
	count := make(map[rune]int)
	for _, r := range t {
		count[r]++
	}

	return findMinWindow(s, count, len(t))
}

func TestMinWindow(t *testing.T) {
	cases := []struct {
		s, t     string
		expected string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"AAABBCCCCBAAC", "ABC", "CBA"},
		{"a", "a", "a"},
	}

	for k, v := range cases {
		r := minWindow(v.s, v.t)
		if r != v.expected {
			t.Errorf("case #%d, want '%s' got '%s'\n", k+1, v.expected, r)
		}
	}
}
