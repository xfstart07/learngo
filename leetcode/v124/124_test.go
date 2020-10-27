package v124

import "testing"

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "maxPathSum01",
			args: args{root: NewNode([]int{1, 2, 3}, 1)},
			want: 6,
		},
		{
			name: "maxPathSum02",
			args: args{root: NewNode([]int{-10, 9, 20, 0, 0, 15, 7}, 1)},
			want: 42,
		},
		{
			name: "maxPathSum03",
			args: args{root: NewNode([]int{1, 2}, 1)},
			want: 3,
		},
		{
			name: "maxPathSum03",
			args: args{root: NewNode([]int{-2, -1}, 1)},
			want: -1,
		},
		{
			name: "maxPathSum03",
			args: args{root: NewNode([]int{1}, 1)},
			want: 1,
		},
		{
			name: "maxPathSum04",
			args: args{root: NewNode([]int{1, -2, -3, 1, 3, -2, 0, -1}, 1)},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
