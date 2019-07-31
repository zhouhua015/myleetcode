package leetcode

import "testing"

func permuteBacktrace(list [][]int, cur, nums []int, visited map[int]bool) [][]int {
	if len(visited) == len(nums) {
		tmp := make([]int, len(cur), len(cur))
		copy(tmp, cur)
		return append(list, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		}
		visited[nums[i]] = true
		cur = append(cur, nums[i])
		list = permuteBacktrace(list, cur, nums, visited)
		cur = cur[:len(cur)-1]
		delete(visited, nums[i])
	}
	return list
}

func permute(nums []int) [][]int {
	visited := make(map[int]bool)
	return permuteBacktrace(nil, nil, nums, visited)
}

func TestPermute(t *testing.T) {
	cases := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, nil},
		{[]int{1}, nil},
	}

	for k, v := range cases {
		r := permute(v.nums)
		t.Logf("case #%d, %v\n", k+1, r)
	}
}
