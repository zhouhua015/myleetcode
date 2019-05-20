package leetcode

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	sum, carry := new(ListNode), 0
	n1, n2 := l1, l2
	n := sum
	for {
		v := carry
		if n1 != nil {
			v += n1.Val
		}
		if n2 != nil {
			v += n2.Val
		}
		n.Val = v % 10

		if carry = 0; v >= 10 {
			carry = 1
		}

		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}

		if n1 == nil && n2 == nil {
			break
		}

		n.Next = new(ListNode)
		n = n.Next
	}

	if carry > 0 {
		n.Next = &ListNode{Val: carry}
	}
	return sum
}

func TestAddTwoNumbers(t *testing.T) {
	testcases := []struct {
		l1  *ListNode
		l2  *ListNode
		sum *ListNode
	}{
		{
			l1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 3},
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  6,
					Next: &ListNode{Val: 4},
				},
			},
			sum: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val:  0,
					Next: &ListNode{Val: 8},
				},
			},
		},
		{
			l1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 3},
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
				},
			},
			sum: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val:  0,
					Next: &ListNode{Val: 4},
				},
			},
		},
		{
			l1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  6,
					Next: &ListNode{Val: 4},
				},
			},
			sum: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val:  0,
					Next: &ListNode{Val: 5},
				},
			},
		},
		{
			l1: &ListNode{
				Val: 2,
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  6,
					Next: &ListNode{Val: 4},
				},
			},
			sum: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val:  6,
					Next: &ListNode{Val: 4},
				},
			},
		},
	}

	f := func(l1, l2, sum *ListNode) {
		actual := addTwoNumbers(l1, l2)
		if actual == nil {
			t.Fatal("Got nil result\n")
		}

		n1, n2 := actual, sum
		for i := 0; n1 != nil && n2 != nil; i++ {
			if n1.Val != n2.Val {
				t.Errorf("index: %d, want %d, got %d\n", i, n2.Val, n1.Val)
			}
			if (n1.Next == nil) != (n2.Next == nil) {
				t.Errorf("index: %d, length mismatched found\n", i)
			}
			n1, n2 = n1.Next, n2.Next
		}
	}

	for _, c := range testcases {
		f(c.l1, c.l2, c.sum)
	}
}
