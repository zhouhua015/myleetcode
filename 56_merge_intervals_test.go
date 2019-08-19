package leetcode

import (
	"sort"
	"testing"
)

func min56(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max56(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func mergeR(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 1 {
		return intervals
	}

	i1, i2 := mergeR(intervals[:n/2]), mergeR(intervals[n/2:])
	i, j := len(i1)-1, 0
	end1, start2 := i1[i][1], i2[j][0]

	var result [][]int
	if start2 > end1 {
		result = append(result, i1...)
		result = append(result, i2...)
		return result
	}

	result = append(result, i1[:i]...)
	result = append(result, []int{min56(i1[i][0], i2[j][0]), max56(i1[i][1], i2[j][1])})
	result = append(result, i2[j+1:]...)
	return mergeR(result)
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	return mergeR(intervals)
}

func TestMergeIntervals(t *testing.T) {
	cases := []struct {
		intervals [][]int
		expected  [][]int
	}{
		{
			[][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}},
			[][]int{[]int{1, 6}, []int{8, 10}, []int{15, 18}},
		},
		{
			[][]int{[]int{1, 4}, []int{4, 5}},
			[][]int{[]int{1, 5}},
		},
		{
			[][]int{[]int{1, 4}, []int{5, 8}, []int{2, 10}, []int{6, 15}},
			[][]int{[]int{1, 15}},
		},
		{
			[][]int{[]int{1, 4}},
			[][]int{[]int{1, 4}},
		},
		{
			[][]int{},
			[][]int{},
		},
		{
			[][]int{[]int{2, 3}, []int{4, 5}, []int{6, 7}, []int{8, 9}, []int{1, 10}},
			[][]int{[]int{1, 10}},
		},
	}

	for k, v := range cases {
		r := merge(v.intervals)
		if len(r) != len(v.expected) {
			t.Errorf("case #%d, length mismatching, want %v got %v\n", k+1, v.expected, r)
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
