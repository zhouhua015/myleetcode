package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 * 		Val int
 *      Left *TreeNode
 *      Right *TreeNode
 * }
 */

func sumR2L(node *TreeNode, cur int) int {
	var sum int
	if node.Left != nil {
		sum = sumR2L(node.Left, cur<<1+node.Val)
	}
	if node.Right != nil {
		sum += sumR2L(node.Right, cur<<1+node.Val)
	}

	if sum == 0 {
		return cur<<1 + node.Val
	}
	return sum
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sumR2L(root, 0)
}
