// Author: xufei
// Date: 2020/3/25

package v235

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	root := NewNode([]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}, 1)
	root2 := NewNode([]int{2, 1, 4}, 1)
	root4 := NewNode([]int{2, 1}, 1)
	root5 := NewNode([]int{3, 1, 4, -1, 2}, 1)
	root6 := NewNode([]int{5, 3, 6, 2, 4, -1, -1}, 1)

	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "最近公共祖先",
			args: args{
				root: root,
				p:    &TreeNode{Val: 2},
				q:    &TreeNode{Val: 8},
			},
			want: &TreeNode{Val: 6},
		},
		{
			name: "Test1",
			args: args{
				root: root,
				p:    &TreeNode{Val: 2},
				q:    &TreeNode{Val: 4},
			},
			want: &TreeNode{Val: 2},
		},
		{
			name: "Test2",
			args: args{
				root: root2,
				p:    &TreeNode{Val: 2},
				q:    &TreeNode{Val: 4},
			},
			want: &TreeNode{Val: 2},
		},
		{
			name: "Test4",
			args: args{
				root: root4,
				p:    &TreeNode{Val: 1},
				q:    &TreeNode{Val: 2},
			},
			want: &TreeNode{Val: 2},
		},
		{
			name: "Test5",
			args: args{
				root: root5,
				p:    &TreeNode{Val: 2},
				q:    &TreeNode{Val: 3},
			},
			want: &TreeNode{Val: 3},
		},
		{
			name: "Test6",
			args: args{
				root: root6,
				p:    &TreeNode{Val: 1},
				q:    &TreeNode{Val: 4},
			},
			want: &TreeNode{Val: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
