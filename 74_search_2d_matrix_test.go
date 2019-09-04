package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	R, C := len(matrix), len(matrix[0])
	l, r := 0, R*C

	for l < r {
		m := (l + r) / 2

		v := matrix[m/C][m%C]
		if v > target {
			l = m + 1
		} else if v < target {
			r = m
		} else {
			return true
		}
	}

	return false
}
