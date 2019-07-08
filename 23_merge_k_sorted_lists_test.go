package leetcode

type ListNode23 struct {
	Val  int
	Next *ListNode23
}

func mergeKLists(lists []*ListNode23) *ListNode23 {
	if len(lists) == 0 {
		return nil
	}

	var head, pre *ListNode23
	run := make([]*ListNode23, len(lists))
	for {
		var index int
		for i := 0; i < len(run); i++ {
			if run[i] == nil {
				run[i] = lists[i]
				if lists[i] != nil {
					lists[i] = lists[i].Next
				}
			}
			if run[i] == nil {
				continue
			}
			if run[index] == nil || run[i].Val < run[index].Val {
				index = i
			}
		}

		if head == nil {
			head = run[index]
		} else {
			pre.Next = run[index]
		}
		pre = run[index]
		run[index] = nil

		if pre == nil {
			break
		}
	}
	return head
}
