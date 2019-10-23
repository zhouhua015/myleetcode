package leetcode

import "testing"

func search81(nums []int, target int) bool {
	l, h := 0, len(nums)-1
	for l <= h {
		m := l + (h-l)/2
		if nums[m] > target {
			if nums[m] > nums[l] && nums[l] > target {
				l = m + 1
			} else if nums[m] == nums[l] {
				l++
			} else {
				h = m - 1
			}
		} else if nums[m] < target {
			if nums[h] > nums[m] && nums[h] < target {
				h = m - 1
			} else if nums[h] == nums[m] {
				h--
			} else {
				l = m + 1
			}
		} else {
			return true
		}
	}

	return false
}

func TestSearchInRotatedArray81(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected bool
	}{
		{[]int{5, 1, 3}, 3, true},
		{[]int{1, 2, 3}, 3, true},
		{[]int{3, 1, 1}, 3, true},
		{[]int{1, 1, 3}, 3, true},
		{[]int{3, 1, 3, 3, 3}, 1, true},
		{[]int{1, 3, 1, 1}, 3, true},
		{[]int{1, 1, 3, 1, 1}, 1, true},
		{[]int{1, 1, 1, 3, 1}, 3, true},
		{[]int{2, 5, 6, 7, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 7, 0, 1, 2}, 3, false},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{5, 3}, 3, true},
		{[]int{5, 3}, 5, true},
		{[]int{5, 5, 3, 3, 3}, 3, true},
		{[]int{5, 5, 3, 3, 3}, 0, false},
	}

	for k, v := range cases {
		ans := search81(v.nums, v.target)
		if ans != v.expected {
			t.Errorf("case #%d, want %t got %t\n", k+1, v.expected, ans)
			t.FailNow()
		}
	}
}
