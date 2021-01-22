// Author: xufei
// Date: 2020/3/31

package v230

import "testing"

func Test_kthSmallest(t *testing.T) {
	root := newTree([]int{3, 1, 4, -1, 2}, 1)

	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find k",
			args: args{
				root: root,
				k:    1,
			},
			want: 1,
		},
		{
			name: "find k",
			args: args{
				root: newTree([]int{5, 3, 6, 2, 4, -1, -1, 1}, 1),
				k:    3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
