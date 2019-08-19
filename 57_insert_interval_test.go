package leetcode

import (
	"testing"
)

func min57(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max57(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func overlapped(i1, i2 []int) bool {
	if i1[1] < i2[0] || i2[1] < i1[0] {
		return false
	}

	return true
}

func insert57(intervals [][]int, newInterval []int) [][]int {
	l, r := 0, len(intervals)
	for l < r {
		m := (l + r) / 2

		if newInterval[0] <= intervals[m][0] {
			r = m
		} else {
			l = m + 1
		}
	}
	leftmost := l

	l, r = 0, len(intervals)
	for l < r {
		m := (l + r) / 2

		if newInterval[1] < intervals[m][1] {
			r = m
		} else {
			l = m + 1
		}
	}
	rightmost := l

	var result [][]int
	if leftmost > 0 &&
		overlapped(intervals[leftmost-1], newInterval) &&
		rightmost < len(intervals) &&
		overlapped(newInterval, intervals[rightmost]) {
		result = append(result, intervals[:leftmost-1]...)
		result = append(result, []int{
			intervals[leftmost-1][0],
			intervals[rightmost][1],
		})
		result = append(result, intervals[rightmost+1:]...)
	} else if leftmost > 0 && overlapped(intervals[leftmost-1], newInterval) {
		result = append(result, intervals[:leftmost-1]...)
		result = append(result, []int{intervals[leftmost-1][0], newInterval[1]})
		result = append(result, intervals[rightmost:]...)
	} else if rightmost < len(intervals) && overlapped(newInterval, intervals[rightmost]) {
		result = append(result, intervals[:leftmost]...)
		result = append(result, []int{newInterval[0], intervals[rightmost][1]})
		result = append(result, intervals[rightmost+1:]...)
	} else {
		result = append(result, intervals[:leftmost]...)
		result = append(result, newInterval)
		result = append(result, intervals[rightmost:]...)
	}

	return result
}

func TestInsertInterval(t *testing.T) {
	cases := []struct {
		newInterval []int
		intervals   [][]int
		expected    [][]int
	}{
		{
			[]int{2, 5},
			[][]int{[]int{1, 3}, []int{6, 9}},
			[][]int{[]int{1, 5}, []int{6, 9}},
		},
		{
			[]int{4, 8},
			[][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}},
			[][]int{[]int{1, 2}, []int{3, 10}, []int{12, 16}},
		},
		{
			[]int{1, 16},
			[][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}},
			[][]int{[]int{1, 16}},
		},
		{
			[]int{3, 4},
			[][]int{[]int{1, 2}},
			[][]int{[]int{1, 2}, []int{3, 4}},
		},
		{
			[]int{1, 2},
			[][]int{[]int{3, 4}},
			[][]int{[]int{1, 2}, []int{3, 4}},
		},
	}

	for k, v := range cases {
		r := insert57(v.intervals, v.newInterval)
		if len(r) != len(v.expected) {
			t.Errorf("case #%d, length mismatching, want %v got %v\n", k+1, v.expected, r)
			continue
		}

		for i := 0; i < len(r); i++ {
			if len(r[i]) != len(v.expected[i]) {
				t.Errorf("case #%d, %dth element length mismatching, want %v got %v\n",
					k+1, i+1, v.expected, r)
				continue
			}

			if r[i][0] != v.expected[i][0] || r[i][1] != v.expected[i][1] {
				t.Errorf("case #%d, %dth element mismatch, want %v got %v\n",
					k+1, i+1, v.expected[i], r[i])
			}
		}
	}
}
