package leetcode

type ListNode24 struct {
	Val  int
	Next *ListNode24
}

func swapPairs(head *ListNode24) *ListNode24 {
	if head == nil || head.Next == nil {
		return head
	}

	cur, next, head := head, head.Next, head.Next
	for next != nil {
		nextcur := next.Next
		next.Next = cur

		if nextcur == nil {
			next = nil
			cur.Next = nil
		} else {
			next = nextcur.Next
			if next == nil {
				cur.Next = nextcur
			} else {
				cur.Next = next
			}
		}
		cur = nextcur
	}

	return head
}
