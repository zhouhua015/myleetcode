package leetcode

func findSubset(ans [][]int, cur []int, nums []int, k int) [][]int {
	if k == 0 {
		temp := make([]int, len(cur))
		copy(temp, cur)
		ans = append(ans, temp)
		return ans
	}

	for i := 0; i < len(nums); i++ {
		cur = append(cur, nums[i])
		ans = findSubset(ans, cur, nums[i+1:], k-1)
		cur = cur[:len(cur)-1]
	}
	return ans
}

func subsets(nums []int) [][]int {
	var ans [][]int
	for i := 0; i <= len(nums); i++ {
		tmp := findSubset(nil, nil, nums, i)
		ans = append(ans, tmp...)
	}
	return ans
}
