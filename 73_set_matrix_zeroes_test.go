package leetcode

func setZeroesMN(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rows := make([]int, m)
	cols := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = 1
				cols[j] = 1
			}
		}
	}

	for i := 0; i < len(rows); i++ {
		if rows[i] == 0 {
			continue
		}

		for j := 0; j < n; j++ {
			matrix[i][j] = 0
		}
	}

	for i := 0; i < len(cols); i++ {
		if cols[i] == 0 {
			continue
		}

		for j := 0; j < m; j++ {
			matrix[j][i] = 0
		}
	}
}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	var firstCol bool
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstCol = true
		}

		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if firstCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
