// Author: xufei
// Date: 2020/3/30

package v160

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	headA := newNode([]int{4, 1})
	headB := newNode([]int{5, 0, 1})
	posA := newNode([]int{8, 4, 5})
	headA.Next = posA
	headB.Next = posA

	posLA := newNode([]int{2, 4})
	headLA := newNode([]int{0, 9, 1})
	headLB := newNode([]int{3})
	headLA.Next = posLA
	headLB.Next = posLA

	headOA := newNode([]int{2, 6, 4})
	headOB := newNode([]int{1, 5})

	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "getIntersectionNode",
			args: args{
				headA: headA,
				headB: headB,
			},
			want: posA,
		},
		{
			name: "getIntersectionNode1",
			args: args{
				headA: headLA,
				headB: headLB,
			},
			want: posLA,
		},
		{
			name: "getIntersectionNode2",
			args: args{
				headA: headOA,
				headB: headOB,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
