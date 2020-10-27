// Author: xufei
// Date: 2020/3/25

package v236

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestorV1(t *testing.T) {
	root := NewNode([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, 1)
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
				p:    &TreeNode{Val: 5},
				q:    &TreeNode{Val: 1},
			},
			want: &TreeNode{Val: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestorV1(tt.args.root, tt.args.p, tt.args.q); (got == nil && got != tt.want) ||
				(got != nil && !reflect.DeepEqual(got.Val, tt.want.Val)) {
				t.Errorf("lowestCommonAncestorV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
