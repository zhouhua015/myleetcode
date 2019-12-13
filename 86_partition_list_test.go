package leetcode

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	lh, gh := &ListNode{}, &ListNode{}
	lt, ge := lh, gh
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			lt.Next = cur
			lt = lt.Next
		} else {
			ge.Next = cur
			ge = ge.Next
		}
	}
	ge.Next = nil
	lt.Next = gh.Next
	return lh.Next
}

func TestPartitionList(t *testing.T) {
	cases := []struct {
		head     *ListNode
		x        int
		expected *ListNode
	}{
		{
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 2,
								},
							},
						},
					},
				},
			},
			3,
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
		},
	}

	for k, v := range cases {
		ans := partition(v.head, v.x)
		for p, cur, n := v.expected, ans, 0; p != nil; p, cur, n = p.Next, cur.Next, n+1 {
			if cur == nil {
				t.Fatalf("case #%d, %dth node, unexpected end.", k+1, n+1)
			}
			if cur.Val != p.Val {
				t.Fatalf("case #%d, %dth node, want %d got %d.", k+1, n+1, p.Val, cur.Val)
			}
		}
	}

}
