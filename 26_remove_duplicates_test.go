package leetcode

import "testing"

func removeDuplicatesNaive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	var i, j int
	for i = 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		j++
		nums[j] = nums[i]
	}

	return j + 1
}

func findNext(nums []int, index int) int {
	v := nums[index]
	l, r := index+1, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] <= v {
			l = m + 1
		} else {
			r = m
		}
	}

	if nums[r] != v {
		return r
	}
	return len(nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	var pos int
	for i := 0; i < len(nums); {
		i = findNext(nums, i)
		if i >= len(nums) {
			break
		}
		pos++
		nums[pos] = nums[i]
	}

	return pos + 1
}

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for i, c := range cases {
		n := removeDuplicates(c.nums)
		if n != c.expected {
			t.Errorf("case #%d, want %d got %d\n", i+1, c.expected, n)
		}

		for j := 1; j < n; j++ {
			if c.nums[j] == c.nums[j-1] {
				t.Errorf("case #%d, duplicates found at index %d\n", i+1, j)
			}
		}
	}
}
