package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodesNaive(head *ListNode) []int {
	var ans []int
	for ; head != nil; head = head.Next {
		ans = append(ans, 0)
		i := len(ans) - 1
		for cur := head.Next; cur != nil; cur = cur.Next {
			if cur.Val > head.Val {
				ans[i] = cur.Val
				break
			}
		}
	}
	return ans
}

func nextLargerNodesOptimized(head *ListNode) []int {
	var vals []int
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}

	var max int
	ans := make([]int, len(vals))
	for i := len(vals) - 1; i >= 0; i-- {
		if i == len(vals)-1 || vals[i] >= max {
			max, ans[i] = vals[i], 0
		} else {
			for j := i + 1; j < len(vals); j++ {
				if vals[j] > vals[i] {
					ans[i] = vals[j]
					break
				}
			}
		}
	}

	return ans
}

type stack1019 []int

func (s *stack1019) push(n int) {
	*s = append(*s, n)
}

func (s *stack1019) pop() (n int) {
	x := len(*s) - 1
	n, *s = (*s)[x], (*s)[:x]
	return n
}

func (s *stack1019) empty() bool {
	return len(*s) == 0
}

func (s *stack1019) peek() int {
	return (*s)[len(*s)-1]
}

func nextLargerNodes(head *ListNode) []int {
	var vals []int
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}

	var s stack1019
	ans := make([]int, len(vals))
	for i := len(vals) - 1; i >= 0; i-- {
		for !s.empty() && s.peek() <= vals[i] {
			s.pop()
		}

		if s.empty() {
			ans[i] = 0
		} else {
			ans[i] = s.peek()
		}
		s.push(vals[i])
	}

	return ans
}

func nextLargerNodesReverseAndStack(head *ListNode) []int {
	var cur *ListNode
	for cur, head.Next = head.Next, nil; cur != nil; {
		tmp := cur.Next
		cur.Next = head
		head = cur
		cur = tmp
	}

	var ans []int
	var s stack1019
	for ; head != nil; head = head.Next {
		for !s.empty() {
			v := s.pop()
			if v > head.Val {
				ans = append(ans, v)
				s.push(v)
				s.push(head.Val)
				break
			}
		}

		if s.empty() {
			ans = append(ans, 0)
			s.push(head.Val)
		}
	}

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}

	return ans
}
