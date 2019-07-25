package leetcode

import "testing"

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	var i int
	v := nums[i]
	for i < len(nums) {
		if v <= 0 || v > len(nums) || nums[v-1] == v {
			if v <= 0 || v > len(nums) {
				nums[i] = v
			}

			i++
			if i >= len(nums) {
				break
			}

			v = nums[i]
			continue
		}

		v, nums[v-1] = nums[v-1], v
	}

	for i = 0; i < len(nums); i++ {
		if nums[i] <= 0 || nums[i] > len(nums) || nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

func TestFirstMissingPositive(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{0, 1}, 2},
		{[]int{2, 1}, 3},
		{[]int{1, 2}, 3},
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 10, 11, 13}, 1},
		{[]int{1, 2, 3, 4, 6}, 5},
	}

	for k, v := range cases {
		r := firstMissingPositive(v.nums)
		if r != v.expected {
			t.Errorf("case #%d, want %d got %d\n", k+1, v.expected, r)
		}
	}
}
