package reverseList

import (
	"reflect"
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

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
