package leetcode

import "testing"

func insert(val int, nums []int, l, h int) (int, []int) {
	for l < h {
		m := (l + h) / 2
		if nums[m] < val {
			l = m + 1
		} else {
			h = m
		}
	}

	var newnums []int
	newnums = append(newnums, nums[0:l]...)
	newnums = append(newnums, val)
	newnums = append(newnums, nums[l:]...)
	return l, newnums
}

func myFindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := (len(nums1) + len(nums2)) / 2

	short, long := nums1, nums2
	if len(nums1) > len(nums2) {
		short, long = nums2, nums1
	}

	slow, shigh, llow, lhigh := 0, len(short)-1, 0, len(long)
	var middle int
	for slow <= shigh {
		middle = (slow + shigh) / 2
		var pos int
		pos, long = insert(short[middle], long, llow, lhigh)
		if pos+middle-slow < m {
			m -= middle - slow
			slow = middle + 1
			llow = pos + 1
			lhigh = len(long)
		} else if pos+middle-slow > m {
			shigh = middle - 1
			lhigh = pos
		} else {
			m -= middle - slow
			break
		}
	}

	if (len(nums1)+len(nums2))%2 != 0 {
		return float64(long[m])
	}

	n := long[m-1]
	if middle-1 >= 0 && short[middle-1] > n {
		n = short[middle-1]
	}
	return float64(n+long[m]) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	short, long := nums1, nums2
	if len(nums1) > len(nums2) {
		short, long = nums2, nums1
	}

	m, n := len(short), len(long)
	l, h := 0, len(short)
	for l <= h {
		i := (l + h) / 2
		j := (m+n+1)/2 - i
		switch {
		case i > l && short[i-1] > long[j]:
			h = i - 1
		case i < h && long[j-1] > short[i]:
			l = i + 1
		default:
			var median1 int
			if i == 0 {
				median1 = long[j-1]
			} else if j == 0 {
				median1 = short[i-1]
			} else {
				if median1 = long[j-1]; short[i-1] > median1 {
					median1 = short[i-1]
				}
			}

			if (m+n)%2 != 0 {
				return float64(median1)
			}

			var median2 int
			if i == m {
				median2 = long[j]
			} else if j == n {
				median2 = short[i]
			} else {
				if median2 = long[j]; short[i] < median2 {
					median2 = short[i]
				}
			}

			return float64(median1+median2) / 2
		}
	}
	return -1.0
}

func TestMedianSortedArrays(t *testing.T) {
	testcases := []struct {
		nums1, nums2 []int
		expected     float64
	}{
		{[]int{3, 4, 5}, []int{1, 2, 6}, 3.5},
		{[]int{1, 2}, []int{-1, 3}, 1.5},
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{nums2: []int{3, 4}, expected: 3.5},
		{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, 4.5},
		{[]int{2, 3, 8, 15, 24}, []int{7, 11, 21, 35, 44, 45}, 15},
	}

	f := func(nums1, nums2 []int, expected float64) {
		if actual := findMedianSortedArrays(nums1, nums2); actual != expected {
			t.Errorf("want %.1f got %.1f, test case: {nums1: %v, nums2: %v}\n",
				expected, actual, nums1, nums2)
		} else {
			t.Logf("passed test case: {nums1: %v, nums2: %v}\n", nums1, nums2)
		}
	}

	for _, c := range testcases {
		f(c.nums1, c.nums2, c.expected)
	}
}
