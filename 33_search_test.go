package leetcode

import "testing"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, h := 0, len(nums)
	for l < h {
		m := (l + h) / 2

		if nums[m] > target {
			if nums[h-1] >= nums[m] || target >= nums[l] {
				h = m
			} else {
				l = m + 1
			}
		} else if nums[m] < target {
			if nums[m] >= nums[l] || target <= nums[h-1] {
				l = m + 1
			} else {
				h = m
			}
		} else {
			return m
		}
	}

	return -1
}

func TestSearchInRotatedSortedArray(t *testing.T) {
	testcases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
		{[]int{5}, 3, -1},
		{[]int{5}, 5, 0},
		{[]int{3, 5}, 3, 0},
		{[]int{3, 5}, 5, 1},
		{[]int{5, 3}, 5, 0},
		{[]int{5, 3}, 3, 1},
		{[]int{5, 3}, 2, -1},
		{[]int{7, 8, 1, 2, 3, 4, 5, 6}, 2, 3},
	}

	for i, c := range testcases {
		n := search(c.nums, c.target)
		if n != c.expected {
			t.Errorf("case #%d, want %d got %d\n", i+1, c.expected, n)
		}
	}
}
