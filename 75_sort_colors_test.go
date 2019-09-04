package leetcode

func sortColorsCounting(nums []int) {
	counts := make([]int, 3)
	for i := 0; i < len(nums); i++ {
		counts[nums[i]]++
	}

	x := 0
	for i := 0; i < len(nums); i++ {
		if counts[x] > 0 {
			counts[x]--
			nums[i] = x
		} else {
			x++
			i--
		}
	}
}

func sortColorsInnerLoop(nums []int) {
	l, r := 0, len(nums)-1
	for i := 0; i <= r; i++ {
		for (nums[i] == 0 && i != l) || (nums[i] == 2 && i != r) {
			if nums[i] == 0 {
				nums[i], nums[l] = nums[l], nums[i]
				l++
			} else {
				nums[i], nums[r] = nums[r], nums[i]
				r--
			}
		}
	}
}

func sortColors(nums []int) {
	for i, j, k := 0, 0, len(nums)-1; i <= k; i++ {
		if nums[i] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		} else if nums[i] == 2 {
			nums[i], nums[k] = nums[k], nums[i]
			k--
			i--
		}
	}
}
