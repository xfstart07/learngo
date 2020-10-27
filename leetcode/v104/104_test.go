// Author: xufei
// Date: 2020/3/24

package v104

import "testing"

func Test_maxDepth(t *testing.T) {
	root := NewNode([]int{3, 9, 20, 0, 0, 15, 7}, 1)

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "MaxDeep",
			args: args{root: root},
			want: 3,
		},
		{
			name: "MaxDeep",
			args: args{root: nil},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
