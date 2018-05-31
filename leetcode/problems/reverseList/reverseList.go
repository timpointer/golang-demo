package reverseList

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		tempNext := curr.Next
		curr.Next = pre
		pre = curr
		curr = tempNext
	}
	return pre
}
