package leetcode

import (
	"strings"
	"testing"
)

func intToRoman(num int) string {
	if num > 3999 || num < 1 {
		return ""
	}

	switch {
	case num >= 1000:
		return strings.Repeat("M", num/1000) + intToRoman(num%1000)
	case num >= 900:
		return "CM" + intToRoman(num%900)
	case num >= 500:
		return "D" + intToRoman(num%500)
	case num >= 400:
		return "CD" + intToRoman(num%400)
	case num >= 100:
		return strings.Repeat("C", num/100) + intToRoman(num%100)
	case num >= 90:
		return "XC" + intToRoman(num%90)
	case num >= 50:
		return "L" + intToRoman(num%50)
	case num >= 40:
		return "XL" + intToRoman(num%40)
	case num >= 10:
		return strings.Repeat("X", num/10) + intToRoman(num%10)
	case num >= 9:
		return "IX" + intToRoman(num%9)
	case num >= 5:
		return "V" + intToRoman(num%5)
	case num >= 4:
		return "IV" + intToRoman(num%4)
	default:
		return strings.Repeat("I", num)
	}
}

func TestInteger2Roman(t *testing.T) {
	testcases := []struct {
		num      int
		expected string
	}{
		{1994, "MCMXCIV"},
		{58, "LVIII"},
		{9, "IX"},
		{4, "IV"},
		{3, "III"},
		{3997, "MMMCMXCVII"},
	}

	for i, c := range testcases {
		r := intToRoman(c.num)
		if r != c.expected {
			t.Errorf("case #%d, want '%s' got '%s'\n", i+1, c.expected, r)
		}
	}
}
