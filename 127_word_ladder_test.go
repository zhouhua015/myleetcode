package leetcode

import (
	"testing"
)

func compileWordList(wordList []string) map[string][]string {
	result := make(map[string][]string)
	const c = byte('*')
	for _, word := range wordList {
		bytes := []byte(word)
		for i, b := range bytes {
			bytes[i] = c
			result[string(bytes)] = append(result[string(bytes)], word)
			bytes[i] = b
		}
	}

	return result
}

type node struct {
	word  string
	depth int
}

func findShortestPath(begins []node, endWord string, visited map[string]bool, compiledList map[string][]string) ([]node, int) {
	var result []node
	const c = byte('*')
	for i := 0; i < len(begins); i++ {
		bytes := []byte(begins[i].word)
		for j, b := range bytes {
			bytes[j] = c
			if list, ok := compiledList[string(bytes)]; ok {
				for k := 0; k < len(list); k++ {
					if list[k] == endWord {
						return nil, begins[i].depth + 1
					}

					if visited[list[k]] {
						continue
					}
					result = append(result, node{list[k], begins[i].depth + 1})
					visited[list[k]] = true
				}
			}
			bytes[j] = b
		}
	}

	if len(result) == 0 {
		return nil, 0
	}
	return findShortestPath(result, endWord, visited, compiledList)
}

func ladderLength(beginWord, endWord string, wordList []string) int {
	compiledList := compileWordList(wordList)
	visited := make(map[string]bool)
	visited[beginWord] = true
	_, len := findShortestPath([]node{node{beginWord, 1}}, endWord, visited, compiledList)
	return len
}

func getAdjacents(words []node, visited map[string]int, compiledList map[string][]string) []node {
	var result []node
	const c = byte('*')
	for i := 0; i < len(words); i++ {
		bytes := []byte(words[i].word)
		for j, b := range bytes {
			bytes[j] = c
			if list, ok := compiledList[string(bytes)]; ok {
				for k := 0; k < len(list); k++ {
					if _, ok := visited[list[k]]; ok {
						continue
					}
					result = append(result, node{list[k], words[i].depth + 1})
					visited[list[k]] = words[i].depth + 1
				}
			}
			bytes[j] = b
		}
	}
	return result
}

func findShortestPathBi(begins, ends []node, bgVisited, edVisited map[string]int, compiledList map[string][]string) int {
	newBegins := getAdjacents(begins, bgVisited, compiledList)
	newEnds := getAdjacents(ends, edVisited, compiledList)

	depth := 0
	for w, bl := range bgVisited {
		if el, ok := edVisited[w]; ok {
			if depth == 0 || depth > bl+el-1 {
				depth = bl + el - 1
			}
		}
	}
	if depth > 0 {
		return depth
	}

	if len(newBegins) == 0 || len(newEnds) == 0 {
		return 0
	}
	return findShortestPathBi(newBegins, newEnds, bgVisited, edVisited, compiledList)
}

func ladderLengthBi(beginWord, endWord string, wordList []string) int {
	i := 0
	for ; i < len(wordList); i++ {
		if wordList[i] == endWord {
			break
		}
	}
	if i == len(wordList) {
		return 0
	}

	compiledList := compileWordList(wordList)
	bgVisited := make(map[string]int)
	bgVisited[beginWord] = 1
	edVisited := make(map[string]int)
	edVisited[endWord] = 1
	return findShortestPathBi([]node{node{beginWord, 1}}, []node{node{endWord, 1}}, bgVisited, edVisited, compiledList)
}

// TestLadderLength verifies ladder length problem
func TestLadderLength(t *testing.T) {
	testcases := []struct {
		bw       string
		ew       string
		words    []string
		expected int
	}{
		{"hog", "cog", []string{"cog"}, 2},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
		{"hot", "dog", []string{"hot", "cog", "dog", "tot", "hog", "hop", "pot", "dot"}, 3},
		{"ta", "if", []string{"ts", "sc", "ph", "ca", "jr", "hf", "to", "if", "ha", "is", "io", "cf", "ta"}, 4},
		{"qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}, 5},
	}

	f := func(b, e string, words []string, length int, how func(string, string, []string) int) {
		if len := how(b, e, words); len != length {
			t.Errorf("want %d got %d, test case {%s, %s, %v}\n", length, len, b, e, words)
		} else {
			t.Logf("passed test case {%s, %s, %v}\n", b, e, words)
		}
	}

	for _, c := range testcases {
		f(c.bw, c.ew, c.words, c.expected, ladderLength)
		f(c.bw, c.ew, c.words, c.expected, ladderLengthBi)
	}
}
