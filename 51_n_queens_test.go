package leetcode

import (
	"strings"
	"testing"
)

func abs51(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func attacked(x, y int, cols []int) bool {
	for i := 0; i < len(cols); i++ {
		if x == i || y == cols[i] {
			return true
		}

		if abs51(x-i) == abs51(y-cols[i]) {
			return true
		}
	}

	return false
}

func nqueens(sln [][]string, cur [][]rune, cols []int) [][]string {
	i, N := len(cols), len(cur)
	if i == N {
		tmp := make([]string, len(cur))
		for j := 0; j < len(cur); j++ {
			tmp[j] = string(cur[j])
		}
		sln = append(sln, tmp)
		return sln
	}

	for j := 0; j < N; j++ {
		if attacked(i, j, cols) {
			continue
		}

		cur[i][j] = 'Q'
		cols = append(cols, j)
		sln = nqueens(sln, cur, cols)
		cur[i][j] = '.'
		cols = cols[:i]
	}
	return sln
}

func solveNQueens(n int) [][]string {
	cur := make([][]rune, n, n)
	for i := 0; i < n; i++ {
		cur[i] = []rune(strings.Repeat(".", n))
	}
	return nqueens(nil, cur, nil)
}

func TestNQueens(t *testing.T) {
	cases := []int{4, 8}
	for k, v := range cases {
		r := solveNQueens(v)
		t.Logf("case #%d, solutions: %d\n", k+1, len(r))
		for i := 0; i < len(r); i++ {
			t.Logf("\t%v\n", r[i])
		}
	}
}
