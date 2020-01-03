package leetcode

import (
	"sort"
	"testing"
)

func lengthOfLISCache(nums []int, start int, cache []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	if cache[start] > 0 {
		return cache[start]
	}

	ans := 1
	for i := start + 1; i < len(nums); i++ {
		if nums[i] > nums[start] {
			length := lengthOfLISCache(nums, i, cache) + 1
			if length > ans {
				ans = length
			}
		}
	}
	cache[start] = ans
	return ans
}

func lengthOfLISNaive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	cache := make([]int, len(nums))
	var ans int
	for i := 0; i < len(nums); i++ {
		l := lengthOfLISCache(nums, i, cache)
		if l > ans {
			ans = l
		}
	}
	return ans
}

func lengthOfLISDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1

	ans := 1
	for i := 1; i < len(nums); i++ {
		var max int
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] > max {
				max = dp[j]
			}
		}
		dp[i] = max + 1
		if ans < max+1 {
			ans = max + 1
		}
	}
	return ans
}

func lengthOfLIS(nums []int) int {
	var lis []int
	for i := 0; i < len(nums); i++ {
		pos := sort.SearchInts(lis, nums[i])
		if pos == len(lis) {
			lis = append(lis, nums[i])
		} else {
			lis[pos] = nums[i]
		}
	}
	return len(lis)
}

func TestLengthOFLIS(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{10, 9, 2, 5, 3, 4}, 3},
		{[]int{4, 10, 4, 3, 8, 9}, 3},
		{[]int{4, 10, 2, 1}, 2},
		{[]int{2, 2}, 1},
	}

	for k, v := range cases {
		ans := lengthOfLIS(v.nums)
		if ans != v.expected {
			t.Fatalf("case #%d, want %d got %d\n", k+1, v.expected, ans)
		}
	}
}
