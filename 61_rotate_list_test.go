package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	last, p := head, head
	n := 0
	for ; p != nil; p = p.Next {
		last = p
		n++
	}

	k = k % n
	if k == 0 {
		return head
	}

	p = head
	for i := n - k - 1; i > 0; i-- {
		p = p.Next
	}
	last.Next, head, p.Next = head, p.Next, nil
	return head
}
