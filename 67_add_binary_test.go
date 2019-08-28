package leetcode

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	if i < 0 {
		return b
	} else if j < 0 {
		return a
	}

	var ans []byte
	if i > j {
		ans = make([]byte, i+1)
	} else {
		ans = make([]byte, j+1)
	}
	m := len(ans) - 1

	var carry bool
	for i >= 0 && j >= 0 {
		if a[i] == b[j] {
			if carry {
				ans[m] = '1'
			} else {
				ans[m] = '0'
			}
			carry = a[i] == '1'
		} else {
			if carry {
				ans[m] = '0'
			} else {
				ans[m] = '1'
			}
		}

		i--
		j--
		m--
	}

	residue := func(k int, s string) {
		for ; k >= 0; k-- {
			if carry {
				if s[k] == '0' {
					ans[m] = '1'
					carry = false
				} else {
					ans[m] = '0'
				}
			} else {
				ans[m] = s[k]
			}
			m--
		}
	}

	residue(i, a)
	residue(j, b)

	if carry {
		ans[m] = '1'
		return string(ans)
	}
	return string(ans[1:])
}
