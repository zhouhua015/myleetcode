package leetcode

func abs52(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func attacked52(x, y int, cols []int) bool {
	for i := 0; i < len(cols); i++ {
		if x == i || y == cols[i] {
			return true
		}

		if abs52(x-i) == abs52(y-cols[i]) {
			return true
		}
	}

	return false
}

func nqueens52(sln int, cols []int, N int) int {
	i := len(cols)
	if i == N {
		sln++
		return sln
	}

	for j := 0; j < N; j++ {
		if attacked52(i, j, cols) {
			continue
		}

		cols = append(cols, j)
		sln = nqueens52(sln, cols, N)
		cols = cols[:i]
	}
	return sln
}

func totalNQueens(n int) int {
	return nqueens52(0, nil, n)
}
