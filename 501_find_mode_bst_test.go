package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findModeR(node *TreeNode, mode *[]int, freq, last, repeated *int) {
	if node.Left != nil {
		findModeR(node.Left, mode, freq, last, repeated)
	}

	index := len(*mode) - 1
	if index < 0 {
		*mode = append(*mode, node.Val)
		*freq = 1
	} else if (*mode)[index] == node.Val {
		*mode = []int{node.Val}
		*freq++
	} else if *freq == 1 {
		*mode = append(*mode, node.Val)
	} else if *last == node.Val {
		*repeated++
		if *repeated == *freq {
			*mode = append(*mode, node.Val)
		}
	} else {
		*last = node.Val
		*repeated = 1
	}

	if node.Right != nil {
		findModeR(node.Right, mode, freq, last, repeated)
	}
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		mode     []int
		freq     int
		last     int
		repeated int
	)
	findModeR(root, &mode, &freq, &last, &repeated)
	return mode
}
