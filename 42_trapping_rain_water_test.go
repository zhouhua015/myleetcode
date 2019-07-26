package leetcode

import "testing"

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	var i, sum int
	for i < len(height) {
		j := i + 1
		for ; j < len(height); j++ {
			if height[j] >= height[i] {
				break
			}
		}
		if j == len(height) {
			break
		}

		b := height[i]
		for i = i + 1; i < j; i++ {
			sum += b - height[i]
		}
		i = j
	}

	end := i
	i = len(height) - 1
	for i > end {
		j := i - 1
		for ; j >= end; j-- {
			if height[j] >= height[i] {
				break
			}
		}

		b := height[i]
		for i = i - 1; i > j; i-- {
			sum += b - height[i]
		}
		i = j
	}

	return sum
}

func trapDP(height []int) int {
	if len(height) < 3 {
		return 0
	}

	lmax := make([]int, len(height), len(height))
	lmax[0] = height[0]
	for i := 1; i < len(height); i++ {
		lmax[i] = lmax[i-1]
		if lmax[i] < height[i] {
			lmax[i] = height[i]
		}
	}

	rmax := make([]int, len(height), len(height))
	rmax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rmax[i] = rmax[i+1]
		if rmax[i] < height[i] {
			rmax[i] = height[i]
		}
	}

	var sum int
	for i := 1; i < len(height); i++ {
		var min = lmax[i]
		if min > rmax[i] {
			min = rmax[i]
		}
		sum += min - height[i]
	}

	return sum
}

type stack42 []int

func (s *stack42) Push(v int) {
	*s = append(*s, v)
}

func (s *stack42) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *stack42) Peek() int {
	if len(*s) > 0 {
		return (*s)[len(*s)-1]
	}
	panic("empty stack")
}

func (s *stack42) Size() int {
	return len(*s)
}

func trapStack(height []int) int {
	s := make(stack42, 0)
	ans := 0
	for i := 0; i < len(height); i++ {
		for s.Size() > 0 && height[i] > height[s.Peek()] {
			top := s.Peek()
			s.Pop()
			if s.Size() == 0 {
				break
			}

			distance := i - s.Peek() - 1
			min := height[s.Peek()]
			if min > height[i] {
				min = height[i]
			}
			boundedHeight := min - height[top]
			ans += distance * boundedHeight
		}

		s.Push(i)
	}

	return ans
}

func trap2Points(height []int) int {
	if len(height) < 3 {
		return 0
	}

	l, r := 0, len(height)-1
	var lmax, rmax, ans int
	for l < r {
		if height[l] < height[r] {
			if height[l] >= lmax {
				lmax = height[l]
			} else {
				ans += lmax - height[l]
			}
			l++
		} else {
			if height[r] >= rmax {
				rmax = height[r]
			} else {
				ans += rmax - height[r]
			}
			r--
		}
	}

	return ans
}

func TestTrapRainWater(t *testing.T) {
	cases := []struct {
		height   []int
		expected int
	}{
		{[]int{4, 2, 3}, 1},
		{[]int{3, 0, 0, 0, 1}, 3},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 0},
		{[]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 0},
		{[]int{1, 3, 2, 1, 0, 0, 0, 0, 0, 0}, 0},
	}

	for k, v := range cases {
		r := trap(v.height)
		if r != v.expected {
			t.Errorf("case #%d, want %d got %d\n", k+1, v.expected, r)
		}
	}
}
