package leetcode

import (
	"sort"
	"testing"
)

func isAnagram(t1, t2 map[byte]int) bool {
	if len(t1) != len(t2) {
		return false
	}

	for k, v := range t1 {
		if v != t2[k] {
			return false
		}
	}

	return true
}

func groupAnagrams(strs []string) [][]string {
	var tables []map[byte]int
	for i := 0; i < len(strs); i++ {
		tmp := make(map[byte]int)
		for j := 0; j < len(strs[i]); j++ {
			tmp[strs[i][j]]++
		}
		tables = append(tables, tmp)
	}

	visited := make(map[int]bool)
	var groups [][]string
	for i := 0; i < len(strs); i++ {
		base := strs[i]
		if visited[i] {
			continue
		}

		var group []string
		group = append(group, base)

		for j := i + 1; j < len(strs); j++ {
			if visited[j] {
				continue
			}

			if isAnagram(tables[i], tables[j]) {
				group = append(group, strs[j])
				visited[j] = true
			}
		}
		groups = append(groups, group)
	}

	return groups
}

func groupAnagramsSort(strs []string) [][]string {
	table := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		x := []rune(strs[i])
		sort.Slice(x, func(i, j int) bool {
			return x[i] < x[j]
		})
		table[string(x)] = append(table[string(x)], strs[i])
	}

	var groups [][]string
	for _, v := range table {
		groups = append(groups, v)
	}
	return groups
}

func TestGroupAnagrams(t *testing.T) {
	cases := []struct {
		strs []string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
	}

	for k, v := range cases {
		r := groupAnagrams(v.strs)
		t.Logf("case #%d, %v\n", k+1, r)
	}

}
