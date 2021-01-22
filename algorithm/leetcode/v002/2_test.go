// Author: xufei
// Date: 2019-06-17

package p002

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{}
	la := []int{3, 4, 2}
	for _, value := range la {
		node := new(ListNode)
		node.Val = value
		node.Next = l1

		l1 = node
	}

	l2 := new(ListNode)
	la = []int{4, 6, 5}
	for _, value := range la {
		node := new(ListNode)
		node.Val = value
		node.Next = l2

		l2 = node
	}

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "problem002",
			args: args{
				l1: l1,
				l2: l2,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); got.Val != tt.want {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
