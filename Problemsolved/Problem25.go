package Problemsolved

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func reverseListNode(q *ListNode, k int) *ListNode {
	// head, tail track for the first and last (each cycle)
	var head, tail *ListNode = nil, nil
	crt := q
	var prev *ListNode = nil

	count, time := 0, 0
	for crt != nil {
		count++
		crt = crt.Next
	}

	crt = q

	for crt != nil && time < count/k {
		// -- swap part -- //
		prev = nil
		i := 0
		for crt != nil && i < k {
			nextTmp := crt.Next // store next
			crt.Next = prev     // reverse pointer around crt
			prev = crt          // move prev forward
			crt = nextTmp       // move crt forward

			i++ // each time swap 1
		}
		// p->2->1->0 and c->3->4->5->0
		// prev is the first of reverse part, crt is the first of origin remain part

		// -- link part -- //
		// link head to prev (if not head) then sweep tail to the last (for link with prev)
		if i == k { // which mean multiple of k
			if head == nil {
				head = prev // new head
				tail = prev // initial tail to sweep later
			} else {
				tail.Next = prev // already have head so head->tail->prev
			}
			for tail.Next != nil { // sweep until reach the last node
				tail = tail.Next
			}
		}
		time++
	}
	tail.Next = crt // add the remainder

	if head == nil {
		return q
	} else {
		return head
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	return reverseListNode(head, k)
}
