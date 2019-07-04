package leetcode

import (
	"math"
	"sort"
	"testing"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	sum := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			n := nums[l] + nums[r] + nums[i]
			if n > target {
				r--
				for r > i && nums[r] == nums[r+1] {
					r--
				}
			} else if n < target {
				l++
				for l < len(nums) && nums[l] == nums[l-1] {
					l++
				}
			} else {
				return n
			}

			if abs(sum-target) > abs(n-target) {
				sum = n
			}
		}
	}

	return sum
}

func Test3SumClosest(t *testing.T) {
	testcases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{-1, 2, 1, -4}, 3, 2},
		{[]int{-1, 2, 1, -4}, -2, -3},
	}

	for i, c := range testcases {
		sum := threeSumClosest(c.nums, c.target)
		if sum != c.expected {
			t.Errorf("case #%d, want %d got %d\n", i+1, c.expected, sum)
		}
	}
}
