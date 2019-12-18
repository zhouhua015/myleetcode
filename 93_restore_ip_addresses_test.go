package leetcode

import (
	"strings"
	"testing"
)

func restoreIP(all *[]string, cur []string, remain string) {
	if len(cur) == 4 && len(remain) == 0 {
		*all = append(*all, strings.Join(cur, "."))
		return
	} else if len(remain) == 0 || len(cur) > 4 {
		return
	}

	for i := 1; i < 4 && i <= len(remain); i++ {
		tmp := remain[0:i]
		if i != 1 && tmp[0] == '0' {
			continue
		} else if i == 3 && tmp > "255" {
			continue
		}

		cur = append(cur, tmp)
		restoreIP(all, cur, remain[i:])
		cur = cur[:len(cur)-1]
	}
}

func restoreIpAddresses(s string) []string {
	var all []string
	restoreIP(&all, nil, s)
	return all
}

func TestRestoreIPAddresses(t *testing.T) {
	cases := []struct {
		s        string
		expected []string
	}{
		{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
		{"10101010", []string{"10.10.10.10", "10.10.101.0", "10.101.0.10", "101.0.10.10", "101.0.101.0"}},
	}

	for _, v := range cases {
		ans := restoreIpAddresses(v.s)
		t.Log(ans)
	}
}
