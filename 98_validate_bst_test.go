package leetcode

func validateBST(n *TreeNode, low, high *int) bool {
	if n == nil {
		return true
	}

	if (low != nil && n.Val <= *low) || (high != nil && n.Val >= *high) {
		return false
	}

	ans := true
	if n.Left != nil {
		ans = validateBST(n.Left, low, &n.Val)
	}

	if ans && n.Right != nil {
		ans = validateBST(n.Right, &n.Val, high)
	}
	return ans
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return validateBST(root, nil, nil)
}
