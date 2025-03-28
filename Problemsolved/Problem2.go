package Problemsolved

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	du := 0
	l0 := &ListNode{}
	current := l0
	for l1 != nil || l2 != nil || du == 1 {
		val := du
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		//fmt.Println(l1, l2)

		fmt.Println(val, current)
		du = val / 10
		current.Next = &ListNode{Val: val % 10}
		current = current.Next
	}
	return l0.Next
}

// support
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0], Next: nil}
	curr := head

	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i], Next: nil}
		curr = curr.Next
	}

	return head
}

func printLinkedList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d", curr.Val)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	fmt.Println()
}
