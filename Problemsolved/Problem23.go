package Problemsolved

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * Definition for mergeTwoLists
 * func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
 *     ...
 * }
 */

func recu(left int, right int, lists []*ListNode) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)/2
	l1 := recu(left, mid, lists)
	l2 := recu(mid+1, right, lists)
	return mergeTwoLists(l1, l2)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	return recu(0, len(lists)-1, lists)
}
