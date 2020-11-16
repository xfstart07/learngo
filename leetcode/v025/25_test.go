// Author: xufei
// Date: 2020/11/12

package v025

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "node list",
			args: args{
				head: newNode([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
		},
		{
			name: "node list",
			args: args{
				head: newNode([]int{1, 2, 3, 4, 5}),
				k:    3,
			},
		},
		{
			name: "node list",
			args: args{
				head: newNode([]int{1, 2, 3, 4, 5}),
				k:    6,
			},
		},
		{
			name: "node list",
			args: args{
				head: newNode([]int{1, 2, 3, 4, 5}),
				k:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup1(tt.args.head, tt.args.k)
			for n := got; n != nil; n = n.Next {
				t.Logf("%d ", n.Val)
			}
		})
	}
}

func Test_reverseKGroup1(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "node list",
			args: args{
				head: newNode([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(tt.args.head, tt.args.k)
			for n := got; n != nil; n = n.Next {
				t.Logf("%d ", n.Val)
			}
		})
	}
}
