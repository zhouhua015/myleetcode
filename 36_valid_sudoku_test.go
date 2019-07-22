package leetcode

import "testing"

func isValidSudoku(board [][]byte) bool {
	rows := make([]bool, 9, 9)
	cols := make([]bool, 9, 9)
	subs := make([]bool, 9, 9)

	isempty := func(i, j int) bool {
		return board[i][j] == '.'
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if !isempty(i, j) {
				rows[i] = true
				cols[j] = true
				subs[i/3*3+j/3] = true
			}
		}
	}

	for i := 0; i < 9; i++ {
		if !rows[i] {
			continue
		}

		digits := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if isempty(i, j) {
				continue
			}
			if digits[board[i][j]] {
				return false
			}
			digits[board[i][j]] = true
		}
	}

	for i := 0; i < 9; i++ {
		if !cols[i] {
			continue
		}

		digits := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if isempty(j, i) {
				continue
			}
			if digits[board[j][i]] {
				return false
			}
			digits[board[j][i]] = true
		}
	}

	for i := 0; i < 9; i++ {
		if !subs[i] {
			continue
		}

		digits := make(map[byte]bool)
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				r, c := i/3*3+j, i%3*3+k
				if isempty(r, c) {
					continue
				}
				if digits[board[r][c]] {
					return false
				}
				digits[board[r][c]] = true
			}
		}
	}

	return true
}

func TestValidSudoku(t *testing.T) {
	cases := []struct {
		board    [][]byte
		expected bool
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
			true,
		},
		{
			[][]byte{
				[]byte{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			false,
		},
	}

	for k, v := range cases {
		r := isValidSudoku(v.board)
		if r != v.expected {
			t.Errorf("case #%d, want %t got %t\n", k+1, v.expected, r)
		}
	}
}
