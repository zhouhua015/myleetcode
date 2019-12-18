package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack94 []*TreeNode

func (s *stack94) Push(v *TreeNode) {
	*s = append(*s, v)
}

func (s *stack94) Pop() (x *TreeNode) {
	x, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return x
}

func (s *stack94) Empty() bool {
	return len(*s) == 0
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var s stack94
	for cur := root; cur != nil; cur = cur.Left {
		s.Push(cur)
	}

	var ans []int
	for !s.Empty() {
		cur := s.Pop()
		ans = append(ans, cur.Val)

		if cur.Right != nil {
			for cur := cur.Right; cur != nil; cur = cur.Left {
				s.Push(cur)
			}
		}
	}
	return ans
}

func morrisInorderTraversal(root *TreeNode) (ans []int) {
	for cur := root; cur != nil; {
		if cur.Left == nil {
			ans = append(ans, cur.Val)
			cur = cur.Right
		} else {
			mostright := cur.Left
			for mostright.Right != nil {
				mostright = mostright.Right
			}
			mostright.Right = cur
			cur, cur.Left = cur.Left, nil
		}
	}
	return ans
}
