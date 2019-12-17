package leetcode

import (
	"sort"
	"testing"
)

func subsetsr(ans [][]int, cur []int, nums []int, c int) [][]int {
	if c == 0 {
		return append(ans, append([]int(nil), cur...))
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		cur = append(cur, nums[i])
		ans = subsetsr(ans, cur, nums[i+1:], c-1)
		cur = cur[:len(cur)-1]

	}
	return ans
}

func subsetsWithDupLoopAndRecursive(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i <= len(nums); i++ {
		tmp := subsetsr(nil, nil, nums, i)
		ans = append(ans, tmp...)
	}
	return ans
}

func subsetBacktrace(all [][]int, cur []int, nums []int) [][]int {
	all = append(all, append([]int(nil), cur...))
	for j := 0; j < len(nums); {
		cur = append(cur, nums[j])
		all = subsetBacktrace(all, cur, nums[j+1:])
		cur = cur[:len(cur)-1]

		j++
		for j < len(nums) && nums[j] == nums[j-1] {
			j++
		}
	}
	return all
}

func subsetsWithDupBacktrace(nums []int) [][]int {
	sort.Ints(nums)
	return subsetBacktrace(nil, nil, nums)
}

func backStack(nums []int, start int, list []int, allList *[][]int) {
	*allList = append(*allList, list)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		list := append([]int{nums[i]}, list...)
		backStack(nums, i+1, list, allList)
	}
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var all [][]int
	backStack(nums, 0, []int{}, &all)
	return all
}

func TestSubsets2Dup(t *testing.T) {
	cases := []struct {
		nums []int
	}{
		{[]int{1, 2, 2}},
	}

	for _, v := range cases {
		ans := subsetsWithDup(v.nums)
		t.Log(ans)
	}

}
