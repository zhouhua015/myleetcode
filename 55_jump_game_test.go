package leetcode

import "testing"

func canJump(nums []int) bool {
	if nums[0] >= len(nums)-1 {
		return true
	}

	var next int
	for j := 1; j <= nums[0]; j++ {
		if nums[j] == 0 {
			continue
		}

		if next == 0 || j+nums[j] >= next+nums[next] {
			next = j
		}
	}
	if next == 0 {
		return false
	}
	return canJump(nums[next:])
}

func canJumpLoop(nums []int) bool {
	var i int
	for {
		if i+nums[i] >= len(nums)-1 {
			return true
		}

		var next int
		for j := i + 1; j <= i+nums[i]; j++ {
			if nums[j] == 0 {
				continue
			}

			if next == 0 || j+nums[j] >= next+nums[next] {
				next = j
			}
		}
		if next == 0 {
			return false
		}
		i = next
	}
}

func TestJumpGame(t *testing.T) {
	cases := []struct {
		nums    []int
		canjump bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
		{[]int{1, 0}, true},
		{[]int{1, 1, 0}, true},
		{[]int{1, 1, 1, 0}, true},
	}

	for k, v := range cases {
		r := canJump(v.nums)
		if r != v.canjump {
			t.Errorf("case #%d,  want %t got %t\n", k+1, v.canjump, r)
		}
	}
}
