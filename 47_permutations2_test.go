package leetcode

import "testing"

func permuteUniqueBacktrace(list [][]int, cur, nums []int, visited map[int]bool) [][]int {
	if len(visited) == len(nums) {
		tmp := make([]int, len(cur), len(cur))
		copy(tmp, cur)
		return append(list, tmp)
	}

	used := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if visited[i] || used[nums[i]] {
			continue
		}

		used[nums[i]] = true
		visited[i] = true
		cur = append(cur, nums[i])
		list = permuteUniqueBacktrace(list, cur, nums, visited)
		cur = cur[:len(cur)-1]
		delete(visited, i)
	}
	return list
}

func permuteUnique(nums []int) [][]int {
	visited := make(map[int]bool)
	return permuteUniqueBacktrace(nil, nil, nums, visited)
}

func TestPermute2(t *testing.T) {
	cases := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 1, 2}, nil},
		{[]int{3, 3, 0, 3}, nil},
		{[]int{1}, nil},
	}

	for k, v := range cases {
		r := permuteUnique(v.nums)
		t.Logf("case #%d, %v\n", k+1, r)
	}
}
