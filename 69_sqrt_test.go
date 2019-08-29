package leetcode

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	l, r := 0, x
	for l < r {
		m := (l + r) / 2
		if m*m > x {
			r = m
		} else if m*m < x {
			l = m + 1
		} else {
			return m
		}
	}
	return r - 1
}
