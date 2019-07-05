package leetcode

type ListNode19 struct {
	Val  int
	Next *ListNode19
}

func removeNthFromEnd(head *ListNode19, n int) *ListNode19 {
	if head == nil {
		return nil
	}

	b, e := head, head
	for n > 0 && e != nil {
		e = e.Next
		n--
	}

	if e == nil && n > 0 {
		return nil
	}

	var p *ListNode19
	for e != nil {
		e, p, b = e.Next, b, b.Next
	}

	if p == nil {
		head = head.Next
		return head
	}

	p.Next, b.Next = b.Next, nil
	return head
}
