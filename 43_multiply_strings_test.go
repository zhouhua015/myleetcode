package leetcode

import "testing"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	const zero = '0'
	ret := make([]byte, len(num1)+len(num2), len(num1)+len(num2))
	for i := 0; i < len(ret); i++ {
		ret[i] = zero
	}

	for i := len(num1) - 1; i >= 0; i-- {
		carry := 0
		for j := len(num2) - 1; j >= 0; j-- {
			r := int(num1[i]-zero)*int(num2[j]-zero) + carry + int(ret[i+j+1]-zero)
			carry = r / 10
			ret[i+j+1] = byte(r%10) + zero
		}
		if carry > 0 {
			ret[i] = byte(carry) + zero
		} else if i == 0 {
			ret = ret[1:]
		}
	}

	return string(ret)
}

func TestMultiply(t *testing.T) {
	cases := []struct {
		num1, num2 string
		expected   string
	}{
		{"2", "3", "6"},
		{"123", "456", "56088"},
		{"9", "9", "81"},
		{"90", "90", "8100"},
		{"792", "687", "544104"},
	}

	for k, v := range cases {
		r := multiply(v.num1, v.num2)
		if r != v.expected {
			t.Errorf("case #%d, want '%s' got '%s'\n", k+1, v.expected, r)
		}
	}
}
