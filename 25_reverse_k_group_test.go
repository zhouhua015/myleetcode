package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type ListNode25 struct {
	Val  int
	Next *ListNode25
}

func reverseList(head *ListNode25, n int) *ListNode25 {
	if head == nil {
		return head
	}

	cur, next := head, head.Next
	for i := 0; i < n; i++ {
		tmp := next.Next
		next.Next, cur, next = cur, next, tmp
	}
	head.Next = nil

	return cur
}

func reverseKGroup(head *ListNode25, k int) *ListNode25 {
	if k < 2 {
		return head
	}

	var pre *ListNode25
	cur := head
	for cur != nil {
		hdr := cur
		var i int
		for i = 0; i < k && cur != nil; i++ {
			cur = cur.Next
		}

		if i < k {
			if pre != nil {
				pre.Next = hdr
			}
			return head
		}

		rhdr := reverseList(hdr, k-1)
		if pre != nil {
			pre.Next = rhdr
		} else {
			head = rhdr
		}
		pre = hdr
	}

	return head
}

func constructListNodes(nodes string) *ListNode25 {
	list := strings.Split(strings.Trim(nodes, "[]"), ",")
	var head, cur *ListNode25
	for i := 0; i < len(list); i++ {
		v, err := strconv.Atoi(list[i])
		if err != nil {
			return nil
		}

		p := &ListNode25{v, nil}
		if i == 0 {
			head, cur = p, p
			continue
		}
		cur.Next = p
		cur = p
	}
	return head
}

func toString(head *ListNode25) string {
	var str string
	for head != nil {
		str += fmt.Sprintf("%d,", head.Val)
		head = head.Next
	}

	if len(str) > 1 {
		str = str[:len(str)-1]
	}

	return "[" + str + "]"
}

func TestReverseKGroup(t *testing.T) {
	cases := []struct {
		list     string
		k        int
		expected string
	}{
		{"[1,2,3,4,5]", 5, "[5,4,3,2,1]"},
		{"[1,2,3,4]", 2, "[2,1,4,3]"},
	}
	for i, c := range cases {
		head := constructListNodes(c.list)
		if head == nil {
			t.Fatal("Invalid lists")
		}

		str := toString(reverseKGroup(head, c.k))
		if str != c.expected {
			t.Errorf("case #%d, want %s got %s\n", i+1, c.expected, str)
		}
	}
}
