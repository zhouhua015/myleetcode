package leetcode

import "testing"

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m < 1 {
		return nil
	}
	n := len(matrix[0])

	var spiral []int
	var i, j, round int
	for len(spiral) < m*n {
		spiral = append(spiral, matrix[i][j])
		switch {
		case j == n-round-1:
			if i == m-round-1 {
				j--
			} else {
				i++
			}
		case i == round:
			if j == n-round-1 {
				i++
			} else {
				j++
			}
		case i == m-round-1:
			if j == round {
				i--
			} else {
				j--
			}
		case j == round:
			i--
			if i == round {
				round++
				i, j = round, round
			}
		}
	}

	return spiral
}

func TestSpiralOrder(t *testing.T) {
	cases := []struct {
		matrix   [][]int
		expected []int
	}{
		{
			[][]int{
				[]int{1, 2, 3},
			},
			[]int{1, 2, 3},
		},
		{
			[][]int{
				[]int{1},
				[]int{2},
				[]int{3},
			},
			[]int{1, 2, 3},
		},
		{
			[][]int{
				[]int{1, 2},
				[]int{4, 3},
			},
			[]int{1, 2, 3, 4},
		},
		{
			[][]int{
				[]int{1, 2},
				[]int{3, 4},
				[]int{5, 6},
			},
			[]int{1, 2, 4, 6, 5, 3},
		},
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{
				[]int{2, 3, 4},
				[]int{5, 6, 7},
				[]int{8, 9, 10},
				[]int{11, 12, 13},
				[]int{14, 15, 16},
			},
			[]int{2, 3, 4, 7, 10, 13, 16, 15, 14, 11, 8, 5, 6, 9, 12},
		},
		{
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
				[]int{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			[][]int{
				[]int{1, 2, 3, 4, 5},
				[]int{6, 7, 8, 9, 10},
				[]int{11, 12, 13, 14, 15},
				[]int{16, 17, 18, 19, 20},
			},
			[]int{1, 2, 3, 4, 5, 10, 15, 20, 19, 18, 17, 16, 11, 6, 7, 8, 9, 14, 13, 12},
		},
	}

	for k, v := range cases {
		r := spiralOrder(v.matrix)
		if len(r) != len(v.expected) {
			t.Errorf("case #%d, length mismatched, want %d got %d\n", k+1, len(v.expected), len(r))
		}

		for i := 0; i < len(r); i++ {
			if r[i] != v.expected[i] {
				t.Errorf("case #%d, %dth element, want %d got %d\n", k+1, i+1, v.expected[i], r[i])
			}
		}
	}
}
