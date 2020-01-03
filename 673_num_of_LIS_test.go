package leetcode

import (
	"testing"
)

type segTreeValue struct {
	length, count int
}

func newSegTreeValue() segTreeValue {
	return segTreeValue{0, 1}
}

type segmentTree struct {
	value       segTreeValue
	lo, hi      int
	left, right *segmentTree
}

func newSegmentTree(lo, hi int) *segmentTree {
	return &segmentTree{value: newSegTreeValue(), lo: lo, hi: hi}
}

func (t *segmentTree) mid() int {
	return t.lo + (t.hi-t.lo)/2
}

func (t *segmentTree) leftChild() *segmentTree {
	if t.left == nil {
		t.left = newSegmentTree(t.lo, t.lo+(t.hi-t.lo)/2)
		t.left.value = newSegTreeValue()
	}
	return t.left
}

func (t *segmentTree) rightChild() *segmentTree {
	if t.right == nil {
		t.right = newSegmentTree(t.lo+(t.hi-t.lo)/2+1, t.hi)
		t.right.value = newSegTreeValue()
	}
	return t.right
}

func (t *segmentTree) merge(v1, v2 segTreeValue) segTreeValue {
	if v1.length == v2.length {
		if v1.length == 0 {
			return newSegTreeValue()
		}
		return segTreeValue{v1.length, v1.count + v2.count}
	} else if v1.length > v2.length {
		return v1
	}
	return v2
}

func (t *segmentTree) query(key int) segTreeValue {
	if key >= t.hi {
		return t.value
	} else if key < t.lo {
		return newSegTreeValue()
	}
	return t.merge(t.leftChild().query(key), t.rightChild().query(key))
}

func (t *segmentTree) insert(key int, value segTreeValue) {
	if t.lo == t.hi {
		t.value = t.merge(value, t.value)
		return
	}

	if key <= t.mid() {
		t.leftChild().insert(key, value)
	} else {
		t.rightChild().insert(key, value)
	}
	t.value = t.merge(t.leftChild().value, t.rightChild().value)
}

func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	tree := newSegmentTree(min, max)
	for _, n := range nums {
		v := tree.query(n - 1)
		tree.insert(n, segTreeValue{v.length + 1, v.count})
	}
	return tree.value.count
}

func findNumberOfLISDP(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	dp, count := make([]int, len(nums)), make([]int, len(nums))
	dp[0], count[0] = 1, 1

	var length int
	for i := 1; i < len(nums); i++ {
		max, cnt := 0, 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j] > max {
					max, cnt = dp[j], count[j]
				} else if dp[j] == max {
					cnt += count[j]
				}
			}
		}

		dp[i], count[i] = max+1, cnt
		if dp[i] > length {
			length = dp[i]
		}
	}

	var ans int
	for i := 0; i < len(count); i++ {
		if dp[i] == length {
			ans += count[i]
		}
	}
	return ans
}

func TestNumberofLIS(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3, 5, 4, 7}, 2},
		{[]int{2, 2, 2, 2, 2}, 5},
		{[]int{1, 2, 4, 3, 5, 4, 7, 2}, 3},
		{[]int{2, 1}, 2},
	}

	for k, v := range cases {
		ans := findNumberOfLIS(v.nums)
		if ans != v.expected {
			t.Errorf("case #%d, want %d got %d\n", k+1, v.expected, ans)
		}
	}
}
