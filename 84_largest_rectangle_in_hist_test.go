package leetcode

import (
	"testing"
)

func largestRectangleAreaBrutalforce(heights []int) int {
	var ans int

	for i := 0; i < len(heights); i++ {
		min := heights[i]
		for j := i; j < len(heights); j++ {
			if min > heights[j] {
				min = heights[j]
			}

			if min*(j-i+1) > ans {
				ans = min * (j - i + 1)
			}
		}
	}

	return ans
}

func largestRectangleAreaDivideAndConquer(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	var index, value int
	for i, v := range heights {
		if v < value || i == 0 {
			value = v
			index = i
		}
	}

	ans := len(heights) * value
	if t := largestRectangleAreaDivideAndConquer(heights[:index]); ans < t {
		ans = t
	}

	if t := largestRectangleAreaDivideAndConquer(heights[index+1:]); ans < t {
		ans = t
	}

	return ans
}

type stack84 []int

func (s *stack84) push(n int) {
	*s = append(*s, n)
}

func (s *stack84) pop() (p int) {
	x := len(*s) - 1
	p, *s = (*s)[x], (*s)[:x]

	return p
}

func (s *stack84) peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack84) empty() bool {
	return len(*s) == 0
}

func largestRectangleAreaStack(heights []int) int {
	var ans int

	var s stack84

	for i, v := range heights {
		for !s.empty() && heights[s.peek()] > v {
			var tmp int

			top := s.pop()

			if s.empty() {
				tmp = i
			} else {
				tmp = i - s.peek() - 1
			}

			if ans < heights[top]*tmp {
				ans = heights[top] * tmp
			}
		}

		s.push(i)
	}

	for !s.empty() {
		top := s.pop()

		var tmp int

		if s.empty() {
			tmp = len(heights)
		} else {
			tmp = len(heights) - s.peek() - 1
		}

		if ans < heights[top]*tmp {
			ans = heights[top] * tmp
		}
	}

	return ans
}

func largestRectangleArea(heights []int) int {
	return largestRectangleAreaStack(heights)
}

func TestLargestRectArea(t *testing.T) {
	cases := []struct {
		heights  []int
		expected int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{0, 9}, 9},
		{[]int{2, 0, 1, 0, 1, 0}, 2},
		{[]int{1, 2, 2}, 4},
		{[]int{}, 0},
	}

	for k, v := range cases {
		ans := largestRectangleArea(v.heights)
		if ans != v.expected {
			t.Errorf("case #%d, want %d got %d\n", k+1, v.expected, ans)
			t.FailNow()
		}
	}
}
