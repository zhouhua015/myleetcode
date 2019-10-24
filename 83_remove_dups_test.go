package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates83(head *ListNode) *ListNode {
	var ans *ListNode
	for cur := head; cur != nil; cur = cur.Next {
		next := cur.Next
		if next != nil && next.Val == cur.Val {
			continue
		}

		if ans == nil {
			head, ans = cur, cur
		} else {
			ans.Next = cur
			ans = cur
		}
	}
	return head
}
