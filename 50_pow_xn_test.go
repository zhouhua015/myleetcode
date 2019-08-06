package leetcode

import (
	"math"
	"testing"
)

func myPower(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		p := myPower(x, n/2)
		return p * p
	}
	return myPower(x, n-1) * x
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		return 1 / myPower(x, n)
	}
	return myPower(x, n)
}

func TestPowxn(t *testing.T) {
	cases := []struct {
		x        float64
		n        int
		expected float64
	}{
		{2.00000, 10, 1024.00000},
		{2.10000, 3, 9.26100},
		{2.00000, -2, 0.25000},
		{0.00001, 2147483647, 0.00000},
	}

	for k, v := range cases {
		r := myPow(v.x, v.n)
		if math.Abs(r-v.expected) > 0.000001 {
			t.Errorf("case #%d, want %.5f got %.5f\n", k+1, v.expected, r)
		}
	}
}
