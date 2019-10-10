package leetcode

import "testing"

func match(board [][]byte, word []byte, pos, i, j int) bool {
	if pos >= len(word) {
		return true
	}

	m, n := len(board), len(board[0])
	if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != word[pos] {
		return false
	}

	board[i][j] = '_'
	if match(board, word, pos+1, i-1, j) ||
		match(board, word, pos+1, i+1, j) ||
		match(board, word, pos+1, i, j-1) ||
		match(board, word, pos+1, i, j+1) {
		return true
	}
	board[i][j] = word[pos]
	return false
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if match(board, []byte(word), 0, i, j) {
				return true
			}
		}
	}
	return false
}

// func exist(board [][]byte, word string) bool {
// 	m, n := len(board), len(board[0])
// 	visited := make(map[int]bool)
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if board[i][j] != word[0] {
// 				continue
// 			}
//
// 			visited[i*n+j] = true
// 			if walk(board, visited, word[1:], i, j) {
// 				return true
// 			}
// 			visited[i*n+j] = false
// 		}
// 	}
// 	return false
// }

func TestWordSearch(t *testing.T) {
	cases := []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{
			[][]byte{
				[]byte{'A', 'B', 'C', 'E'},
				[]byte{'S', 'F', 'C', 'S'},
				[]byte{'A', 'D', 'E', 'E'},
			},
			"ABCCED",
			true,
		},
		{
			[][]byte{
				[]byte{'A', 'B', 'C', 'E'},
				[]byte{'S', 'F', 'C', 'S'},
				[]byte{'A', 'D', 'E', 'E'},
			},
			"SEE",
			true,
		},
		{
			[][]byte{
				[]byte{'A', 'B', 'C', 'E'},
				[]byte{'S', 'F', 'C', 'S'},
				[]byte{'A', 'D', 'E', 'E'},
			},
			"ABCB",
			false,
		},
	}

	for k, v := range cases {
		ans := exist(v.board, v.word)
		if ans != v.expected {
			t.Errorf("case #%d, want %t got %t\n", k+1, v.expected, ans)
		}
	}
}
