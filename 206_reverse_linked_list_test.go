package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList206(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur, next := head, head.Next
	cur.Next = nil
	for next != nil {
		next.Next, cur, next = cur, next, next.Next
	}
	return cur
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tail := head.Next
	r := reverseListRecursive(tail)
	if r == nil {
		return head
	}
	head.Next = nil
	tail.Next = head
	return r
}
