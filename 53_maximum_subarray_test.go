package leetcode

import "testing"

func maxSubArrayNaive(nums []int) int {
	var sum, max int
	for i := 0; i < len(nums); i++ {
		sum = 0
		if i == 0 {
			max = nums[i]
		}
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}

	return max
}

func maxSubArrayOn(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}

	sum, max := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}

		if sum > max {
			max = sum
		}
	}
	return max
}

func maxSubArray(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}

	for i := 1; i < n; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
	}

	max := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func TestMaxSubarray(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{-2, 1}, 1},
		{[]int{-2, -1, -3, -4, -1}, -1},
	}

	for k, v := range cases {
		r := maxSubArray(v.nums)
		if r != v.expected {
			t.Errorf("case #%d, want %d got %d\n", k+1, v.expected, r)
		}
	}
}
