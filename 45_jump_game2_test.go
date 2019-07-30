package leetcode

import "testing"

func jump2(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	if nums[0]+1 >= len(nums) {
		return 1
	}

	var max, ti int
	for i := nums[0]; i > 0; i-- {
		if nums[i] == 0 {
			continue
		}
		if nums[i]+i > max {
			ti = i
			max = nums[i] + i
		}
	}

	return 1 + jump2(nums[ti:])
}

func TestJumpGame2(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{4, 3, 1, 1, 4}, 1},
		{[]int{1, 1, 1, 1}, 3},
		{[]int{0}, 0},
		{[]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}, 3},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}, 2},
	}

	for k, v := range cases {
		r := jump2(v.nums)
		if r != v.expected {
			t.Errorf("case #%d, want %d got %d\n", k+1, v.expected, r)
		}
	}
}
