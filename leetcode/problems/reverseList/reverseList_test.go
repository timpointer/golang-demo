package reverseList

import (
	"log"
	"testing"
)

func Test_reverseList(t *testing.T) {
	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	got := reverseList(node1)
	log.Println("val %v", got)
	for {
		next := got.Next
		if next == nil {
			break
		}
		log.Println("val %v", got.Val)
		got = next
	}
}
