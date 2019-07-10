package leetcode

import (
	"testing"
)

func divide(dividend int, divisor int) int {
	var neg bool
	if dividend > 0 && divisor < 0 {
		divisor = -divisor
		neg = true
	} else if dividend < 0 && divisor > 0 {
		dividend = -dividend
		neg = true
	} else if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	}

	var q int
	if divisor == 1 {
		q = dividend
	} else {
		s := divisor
		for s <= dividend {
			s += divisor
			q++
		}
	}

	if neg {
		q = -q
	}

	if q > 1<<31-1 {
		return 1<<31 - 1
	} else if q < -1<<31 {
		return -1 << 31
	}
	return q
}

func TestDivide(t *testing.T) {
	cases := []struct {
		dd  int
		dr  int
		exp int
	}{
		{1, 1, 1},
		{1, 3, 0},
		{10, 3, 3},
		{7, -3, -2},
	}

	for i, c := range cases {
		v := divide(c.dd, c.dr)
		if v != c.exp {
			t.Errorf("case #%d, want %d got %d\n", i+1, c.exp, v)
		}
	}
}
