package leetcode

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cur, length, max := nums[0], 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > cur {
			length++
		} else {
			length = 1
		}

		if length > max {
			max = length
		}
		cur = nums[i]
	}
	return max
}
