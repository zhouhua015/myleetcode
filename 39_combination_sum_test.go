package leetcode

import (
	"testing"
)

type combstack []int

func (cs *combstack) Push(v int) {
	*cs = append(*cs, v)
}

func (cs *combstack) Pop() {
	*cs = (*cs)[:len(*cs)-1]
}

func (cs *combstack) Peek() int {
	if len(*cs) > 0 {
		return (*cs)[len(*cs)-1]
	}
	panic("empty stack")
}

func (cs *combstack) Size() int {
	return len(*cs)
}

func combSum(sln [][]int, candidates []int, index int, comb *combstack, left int) [][]int {
	if left == 0 {
		c := make([]int, comb.Size(), comb.Size())
		copy(c, *comb)

		sln = append(sln, c)
		return sln
	}

	for i := index; i < len(candidates); i++ {
		if candidates[i] <= left {
			v := candidates[i]
			comb.Push(v)
			sln = combSum(sln, candidates, i, comb, left-v)
			comb.Pop()
		}
	}

	return sln
}

func combinationSum(candidates []int, target int) [][]int {
	comb := make(combstack, 0)
	return combSum(nil, candidates, 0, &comb, target)
}

func TestCombinationSum(t *testing.T) {
	cases := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{[]int{2, 2, 3}, []int{7}}},
		{[]int{2, 3, 5}, 8, [][]int{[]int{2, 2, 2, 2}, []int{2, 3, 3}, []int{3, 5}}},
	}

	for k, v := range cases {
		r := combinationSum(v.candidates, v.target)
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
