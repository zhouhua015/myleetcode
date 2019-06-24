package leetcode

import (
	"testing"
)

func power(x, y int) int {
	if y <= 0 {
		return 1
	}

	n := x
	for y-1 > 0 {
		n *= x
		y--
	}
	return n
}

func isPalindromeMem(x int) bool {
	if x < 0 {
		return false
	}

	var digits []int
	n := x
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	for i, j := 0, len(digits)-1; i < len(digits) && j >= 0 && i <= j; {
		if digits[i] != digits[j] {
			return false
		}

		i++
		j--
	}
	return true
}

func isPalindrome(x int) bool {
	// negative number and 10s are surely not palindrome
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	var reverted int
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}

	return x == reverted || reverted/10 == x
}

func TestPalindromeNumber(t *testing.T) {
	cases := []struct {
		num          int
		isPalindrome bool
	}{
		{1, true},
		{11, true},
		{12, false},
		{121, true},
		{-121, false},
		{10, false},
		{101, true},
		{1001, true},
		{1111, true},
		{1221, true},
		{1122, false},
	}

	for i := 0; i < len(cases); i++ {
		v := isPalindrome(cases[i].num)
		if cases[i].isPalindrome != v {
			t.Errorf("case #%d, want %t got %t\n", i+1, cases[i].isPalindrome, v)
		}
	}
}
