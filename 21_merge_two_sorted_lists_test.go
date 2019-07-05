package leetcode

type ListNode21 struct {
	Val  int
	Next *ListNode21
}

func mergeTwoLists(l1, l2 *ListNode21) *ListNode21 {
	p1, p2 := l1, l2

	var r, p *ListNode21
	for p1 != nil && p2 != nil {
		var c *ListNode21
		if p1.Val <= p2.Val {
			c, p1 = p1, p1.Next
		} else {
			c, p2 = p2, p2.Next
		}

		if p == nil {
			p, r = c, c
		} else {
			p.Next, p = c, c
		}
	}

	if p1 != nil {
		if p == nil {
			p, r = p1, p1
		} else {
			p.Next = p1
		}
	} else if p2 != nil {
		if p == nil {
			p, r = p2, p2
		} else {
			p.Next = p2
		}
	}

	return r
}
