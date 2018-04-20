package reverseList

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	next := head.Next
	if next != nil {
		nnext := next.Next
		if nnext == nil {
			next.Next = head
			return head
		} else {
			reverseList(next)
		}
	} else {
		return head
	}
	return nil
}
