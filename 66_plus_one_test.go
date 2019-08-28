package leetcode

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}

	i := len(digits) - 1
	digits[i]++
	if digits[i] < 10 {
		return digits
	}

	carry := 1
	digits[i] = 0
	i--
	for ; i >= 0; i-- {
		digits[i] += carry
		carry = 0

		if digits[i] >= 10 {
			digits[i] = 0
			carry = 1
			continue
		}
		return digits
	}

	if carry > 0 {
		return append([]int{carry}, digits...)
	}
	return digits
}
