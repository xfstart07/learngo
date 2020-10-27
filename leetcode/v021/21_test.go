// Author: xufei
// Date: 2020/3/27

package v021

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	l1 := newNode([]int{1, 2, 4})
	l2 := newNode([]int{1, 3, 4})

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "merge ListNode",
			args: args{
				l1: l1,
				l2: l2,
			},
			want: &ListNode{Val: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.l1, tt.args.l2)
			if !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}

			for ; got != nil; got = got.Next {
				fmt.Printf("%d ", got.Val)
			}
		})
	}
}
