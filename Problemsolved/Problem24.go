package Problemsolved

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	a := head

	if head.Next != nil {
		tmp := head.Next     // b
		head.Next = tmp.Next // a.Next = c
		tmp.Next = head      // b.Next = a
		// a(root)->b->a->c
		a = tmp
	}

	for head.Next != nil && head.Next.Next != nil {
		tmp1 := head.Next      // a
		tmp2 := head.Next.Next // b

		tmp1.Next = tmp2.Next // a.Next = c
		tmp2.Next = tmp1      // b.Next = a
		// head->b->a->c
		head.Next = tmp2

		head = tmp1 // new address
	}
	return a
}
