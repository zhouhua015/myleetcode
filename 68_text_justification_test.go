package leetcode

import (
	"strings"
	"testing"
)

func justifiedLine(words []string, extraSpaces int) string {
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", extraSpaces)
	}

	paddings := make([]int, len(words)-1)
	limit := len(paddings)
	spaces := extraSpaces
	if spaces < limit {
		limit = spaces
	}
	for spaces > 0 {
		each := spaces / limit
		for i := 0; i < limit; i++ {
			paddings[i] += each
		}
		spaces -= each * limit
		limit--
		if spaces < limit {
			limit = spaces
		}
	}

	var line string
	for i := 0; i < len(words); i++ {
		line += words[i]
		if i != len(words)-1 {
			line += strings.Repeat(" ", paddings[i]+1)
		}
	}
	return line
}

func fullJustify(words []string, maxWidth int) []string {
	var ans []string
	i := 0
	for i < len(words) {
		var line []string
		width := 0
		for ; width < maxWidth && i < len(words); i++ {
			if width+len(words[i]) > maxWidth {
				break
			}

			line = append(line, words[i])
			width += len(words[i]) + 1
		}

		padding := maxWidth - width + 1
		if i == len(words) {
			padding = 0
		}
		ans = append(ans, justifiedLine(line, padding))
	}
	last := len(ans[len(ans)-1])
	if last < maxWidth {
		ans[len(ans)-1] += strings.Repeat(" ", maxWidth-last)
	}

	return ans
}

func TestTextJustification(t *testing.T) {
	cases := []struct {
		words    []string
		maxWidth int
		expected []string
	}{
		{
			[]string{"This", "is", "an", "example", "of", "text", "justification."},
			16,
			[]string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
			16,
			[]string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			[]string{
				"Science", "is", "what", "we", "understand", "well",
				"enough", "to", "explain", "to", "a", "computer.",
				"Art", "is", "everything", "else", "we", "do",
			},
			20,
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}

	for k, v := range cases {
		ans := fullJustify(v.words, v.maxWidth)
		if len(ans) != len(v.expected) {
			t.Errorf("case #%d, length mismatched, want %v got %v\n", k+1, v.expected, ans)
			continue
		}

		for i := 0; i < len(ans); i++ {
			if ans[i] != v.expected[i] {
				t.Errorf("case #%d, %dth line mismatched, want '%s' got '%s'\n",
					k+1, i+1, v.expected[i], ans[i])
				continue
			}
		}

	}
}
