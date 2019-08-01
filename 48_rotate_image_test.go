package leetcode

import "testing"

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			t00 := &matrix[i][j]
			t0n := &matrix[j][n-i-1]
			tnn := &matrix[n-i-1][n-j-1]
			tn0 := &matrix[n-j-1][i]

			tmp := *t00
			*t00, *tn0, *tnn, *t0n = *tn0, *tnn, *t0n, tmp
		}
	}
}

func TestRotateImage(t *testing.T) {
	cases := []struct {
		matrix [][]int
	}{
		{
			[][]int{
				[]int{5, 1, 9, 11},
				[]int{2, 4, 8, 10},
				[]int{13, 3, 6, 7},
				[]int{15, 14, 12, 16},
			},
		},
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
		},
		{
			[][]int{
				[]int{2, 29, 20, 26, 16, 28},
				[]int{12, 27, 9, 25, 13, 21},
				[]int{32, 33, 32, 2, 28, 14},
				[]int{13, 14, 32, 27, 22, 26},
				[]int{33, 1, 20, 7, 21, 7},
				[]int{4, 24, 1, 6, 32, 34},
			},
		},
		{
			[][]int{
				[]int{1, 2, 29, 20, 26, 16, 28},
				[]int{4, 12, 27, 9, 25, 13, 21},
				[]int{7, 32, 33, 32, 2, 28, 14},
				[]int{11, 13, 14, 32, 27, 22, 26},
				[]int{15, 33, 1, 20, 7, 21, 7},
				[]int{19, 4, 24, 1, 6, 32, 34},
				[]int{21, 18, 8, 3, 14, 23, 27},
			},
		},
	}

	for k, v := range cases {
		rotate(v.matrix)
		t.Logf("case #%d:\n", k+1)
		for i := 0; i < len(v.matrix); i++ {
			t.Logf("\t%v", v.matrix[i])
		}
	}
}
