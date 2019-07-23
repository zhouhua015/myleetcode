package leetcode

import "testing"

const (
	dot        = '.'
	one        = '1'
	nine       = '9'
	unsolvable = 82
)

func next(i, j int) (int, int) {
	j++
	if j >= 9 {
		j = 0
		i++
	}

	return i, j
}

func isValid37(v byte, i, j int, row, col, sub [][]bool) bool {
	return !(row[i][v-one] || col[j][v-one] || sub[i/3*3+j/3][v-one])
}

func solveSudokuAt(board [][]byte, i, j int, row, col, sub [][]bool) int {
	if i == 9 {
		return 0 // solved
	}

	for board[i][j] != dot {
		i, j = next(i, j)
		if i == 9 {
			return 0 // solved
		}
	}

	var v byte = one
	for ; v <= nine; v++ {
		if !isValid37(v, i, j, row, col, sub) {
			continue
		}

		row[i][v-one], col[j][v-one], sub[i/3*3+j/3][v-one] = true, true, true
		board[i][j] = v
		ni, nj := next(i, j)
		if solveSudokuAt(board, ni, nj, row, col, sub) == 0 {
			return 0
		}
		row[i][v-one], col[j][v-one], sub[i/3*3+j/3][v-one] = false, false, false
	}

	board[i][j] = dot
	return unsolvable
}

func solveSudoku(board [][]byte) {
	row := make([][]bool, 9, 9)
	col := make([][]bool, 9, 9)
	sub := make([][]bool, 9, 9)
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 9, 9)
		col[i] = make([]bool, 9, 9)
		sub[i] = make([]bool, 9, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == dot {
				continue
			}

			row[i][board[i][j]-one] = true
			col[j][board[i][j]-one] = true
			sub[i/3*3+j/3][board[i][j]-one] = true
		}
	}
	solveSudokuAt(board, 0, 0, row, col, sub)
}

func TestSudokuSolver(t *testing.T) {
	cases := []struct {
		board [][]byte
	}{
		{
			[][]byte{
				[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
		},
		{
			[][]byte{
				[]byte{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
				[]byte{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
				[]byte{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
				[]byte{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
				[]byte{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
				[]byte{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
				[]byte{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
				[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
				[]byte{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
			},
		},
	}

	for k, v := range cases {
		solveSudoku(v.board)
		t.Logf("case #%d, \t\n%v\n", k+1, v.board)
	}
}
