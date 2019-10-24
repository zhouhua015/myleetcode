package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var ans *ListNode
	for cur := head; cur != nil; {
		next := cur.Next
		if next != nil && next.Val == cur.Val {
			for next != nil && next.Val == cur.Val {
				cur, next = next, next.Next
			}
			cur = next
			continue
		}

		cur.Next = nil
		if ans == nil {
			head, ans = cur, cur
		} else {
			ans.Next = cur
			ans = cur
		}
		cur = next
	}
	if ans == nil {
		return nil
	}

	return head
}
