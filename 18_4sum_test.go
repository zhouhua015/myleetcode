package leetcode

import (
	"sort"
	"testing"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	var quadruplets [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			v := target - nums[i] - nums[j]
			l, r := j+1, len(nums)-1
			for l < r {
				s := nums[l] + nums[r]
				if s == v {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--

					for l < len(nums)-1 && nums[l] == nums[l-1] {
						l++
					}
					for r > l && nums[r] == nums[r+1] {
						r--
					}
				} else if s < v {
					l++
				} else {
					r--
				}
			}
		}
	}

	return quadruplets
}

func Test4Sum(t *testing.T) {
	testcases := []struct {
		target   int
		nums     []int
		expected [][]int
	}{
		{0, []int{1, 0, -1, 0, -2, 2}, [][]int{[]int{-1, 0, 0, 1}, []int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}}},
	}

	for _, c := range testcases {
		r := fourSum(c.nums, c.target)
		t.Log(r)
	}

}
