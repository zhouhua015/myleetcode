package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	pseudo := &ListNode{Next: head}
	cur := pseudo
	var (
		i    int
		prev *ListNode
	)

	for ; i < m && cur != nil; i++ {
		prev, cur = cur, cur.Next
	}
	if cur == nil {
		return nil
	}

	rb, next := cur, cur.Next
	for ; i < n && next != nil; i++ {
		next.Next, cur, next = cur, next, next.Next
	}
	prev.Next, rb.Next = cur, next
	return pseudo.Next
}
