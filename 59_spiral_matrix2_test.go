package leetcode

import "testing"

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	r1, r2 := 0, n-1
	c1, c2 := 0, n-1

	x := 1
	for r1 <= r2 && c1 <= c2 {
		for i := c1; i <= c2; i++ {
			ans[r1][i] = x
			x++
		}
		for i := r1 + 1; i <= r2; i++ {
			ans[i][c2] = x
			x++
		}
		if r1 < r2 && c1 < c2 {
			for i := c2 - 1; i > c1; i-- {
				ans[r2][i] = x
				x++
			}

			for i := r2; i > r1; i-- {
				ans[i][c1] = x
				x++
			}
		}

		r1++
		r2--
		c1++
		c2--
	}

	return ans
}

func TestSpiralMatrix2(t *testing.T) {
	cases := []struct {
		n        int
		expected [][]int
	}{
		{1, [][]int{[]int{1}}},
		{
			2,
			[][]int{
				[]int{1, 2},
				[]int{4, 3},
			},
		},
		{
			3,
			[][]int{
				[]int{1, 2, 3},
				[]int{8, 9, 4},
				[]int{7, 6, 5},
			},
		},
		{
			4,
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{12, 13, 14, 5},
				[]int{11, 16, 15, 6},
				[]int{10, 9, 8, 7},
			},
		},
		{
			5,
			[][]int{
				[]int{1, 2, 3, 4, 5},
				[]int{16, 17, 18, 19, 6},
				[]int{15, 24, 25, 20, 7},
				[]int{14, 23, 22, 21, 8},
				[]int{13, 12, 11, 10, 9},
			},
		},
	}

	for k, v := range cases {
		r := generateMatrix(v.n)
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

			for j := 0; j < len(r[i]); j++ {
				if r[i][j] != v.expected[i][j] {
					t.Errorf("case #%d, %dth element mismatch, want %v got %v\n",
						k+1, i+1, v.expected[i], r[i])
				}
			}

		}

	}
}
