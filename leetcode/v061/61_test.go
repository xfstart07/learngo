// Author: xufei
// Date: 2020/4/1

package v061

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "rotate",
			args: args{
				head: newNode([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
			want: &ListNode{Val: 4},
		},
		{
			name: "rotate2",
			args: args{
				head: newNode([]int{0, 1, 2}),
				k:    4,
			},
			want: &ListNode{Val: 2},
		},
		{
			name: "rotate3",
			args: args{
				head: newNode([]int{1, 2, 3, 4, 5}),
				k:    10,
			},
			want: &ListNode{Val: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateRight(tt.args.head, tt.args.k)
			if !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
			for n := got; n != nil; n = n.Next {
				fmt.Printf("%v, ", n.Val)
			}
		})
	}
}
