package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	var sorted *ListNode
	for head != nil {
		var pre, cur *ListNode
		for cur = sorted; cur != nil; cur, pre = cur.Next, cur {
			if cur.Val >= head.Val {
				break
			}
		}
		next := head.Next
		if pre == nil {
			tmp := sorted
			sorted = head
			sorted.Next = tmp
		} else {
			head.Next = pre.Next
			pre.Next = head
		}

		head = next
	}

	return sorted
}

func insertionSortListOptimized(head *ListNode) *ListNode {
	var sorted, last *ListNode
	for head != nil {
		var begin, end *ListNode
		if last == nil || last.Val < head.Val {
			begin = last
			end = nil
		} else {
			begin = sorted
			end = last
		}

		var pre, cur *ListNode
		for cur = begin; cur != end; cur, pre = cur.Next, cur {
			if cur.Val >= head.Val {
				break
			}
		}
		next := head.Next
		if pre == nil {
			tmp := sorted
			sorted = head
			sorted.Next = tmp
		} else {
			head.Next = pre.Next
			pre.Next = head
		}

		last, head = head, next
	}

	return sorted
}
