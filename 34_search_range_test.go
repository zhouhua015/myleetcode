package leetcode

import "testing"

func searchRange(nums []int, target int) []int {
	var rng = []int{-1, -1}
	if len(nums) == 0 {
		return rng
	}

	l, h := 0, len(nums)
	for l < h {
		m := (l + h) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			h = m
		} else {
			rng[0] = m
			h = m
		}
	}

	l, h = 0, len(nums)
	for l < h {
		m := (l + h) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			h = m
		} else {
			rng[1] = m
			l = m + 1
		}
	}

	return rng
}

func TestSearchRange(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 7, []int{1, 2}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{7, 7}, 7, []int{0, 1}},
		{[]int{7}, 7, []int{0, 0}},
	}

	for k, v := range cases {
		x := searchRange(v.nums, v.target)
		if x[0] != v.expected[0] || x[1] != v.expected[1] {
			t.Errorf("case #%d, want %v got %v\n", k+1, v, x)
		}
	}
}
