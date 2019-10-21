package leetcode

import "testing"

func removeDuplicates80(nums []int) int {
	var i, c, n int
	for j := 0; j < len(nums); j++ {
		if n == nums[j] {
			if c < 2 {
				nums[i] = nums[j]
				i++
			}
			c++
		} else {
			c, n = 1, nums[j]
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func TestRemoveDups(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 1, 2, 2, 3}, []int{1, 1, 2, 2, 3}},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, []int{0, 0, 1, 1, 2, 3, 3}},
		{[]int{0, 0, 1, 1, 1, 1, 2, 2, 2, 3, 3}, []int{0, 0, 1, 1, 2, 2, 3, 3}},
	}

	for k, v := range cases {
		ans := removeDuplicates80(v.nums)
		if ans != len(v.expected) {
			t.Errorf("case #%d, want %v got %v\n", k+1, v.expected, v.nums[:ans])
			t.FailNow()
		}

		for i, ans := range v.expected {
			if v.nums[i] != ans {
				t.Errorf("case #%d, want %v got %v\n", k+1, v.expected, v.nums[:ans])
				t.FailNow()
			}
		}
	}

}
