package leetcode

import (
	"sort"
	"testing"
)

type combstack2 []int

func (cs *combstack2) Push(v int) {
	*cs = append(*cs, v)
}

func (cs *combstack2) Pop() {
	*cs = (*cs)[:len(*cs)-1]
}

func (cs *combstack2) Peek() int {
	if len(*cs) > 0 {
		return (*cs)[len(*cs)-1]
	}
	panic("empty stack")
}

func (cs *combstack2) Size() int {
	return len(*cs)
}

func combSum2(sln [][]int, candidates []int, index int, comb *combstack2, left int) [][]int {
	if left == 0 {
		c := make([]int, comb.Size(), comb.Size())
		copy(c, *comb)
		sln = append(sln, c)
		return sln
	}

	var pre int
	for i := index; i < len(candidates); i++ {
		if candidates[i] <= left {
			v := candidates[i]
			if i > index && pre == v {
				continue
			}
			pre = v
			comb.Push(v)
			sln = combSum2(sln, candidates, i+1, comb, left-v)
			comb.Pop()
		}
	}

	return sln
}

func combinationSum2(candidates []int, target int) [][]int {
	comb := make(combstack2, 0)
	sort.Ints(candidates)
	return combSum2(nil, candidates, 0, &comb, target)
}

func TestCombinationSum2(t *testing.T) {
	cases := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{[]int{1, 7}, []int{1, 2, 5}, []int{2, 6}, []int{1, 1, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{[]int{1, 2, 2}, []int{5}}},
	}

	for k, v := range cases {
		r := combinationSum2(v.candidates, v.target)
		t.Logf("expected: %v, r: %v\n", v.expected, r)
		if len(r) != len(v.expected) {
			t.Errorf("case #%d, result mismatch, want %v got %v\n", k+1, v.expected, r)
			continue
		}

		for i := 0; i < len(v.expected); i++ {
			if len(r[i]) != len(v.expected[i]) {
				t.Errorf("case #%d, %dth result mismatch, want %v got %v\n", k+1, i+1, v.expected[i], r[i])
				continue
			}

			for j := 0; j < len(v.expected[i]); j++ {
				if r[i][j] != v.expected[i][j] {
					t.Errorf("case #%d, %dth result %dth element mismatch, want %v got %v\n", k+1, i+1, j+1, v.expected[i], r[i])
					break
				}
			}
		}
	}
}
