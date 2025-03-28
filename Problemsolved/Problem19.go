package Problemsolved

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	q := head
	//var c byte = 1
	c := 1
	for q.Next != nil {
		q = q.Next
		c++
	}
	// suicide case
	if c == n {
		return head.Next
	}
	q = head
	//var i byte = 1
	for i := 1; i < c; i++ {
		if i == c-n {
			q.Next = q.Next.Next
			break
		} else {
			q = q.Next
		}
	}
	return head
}