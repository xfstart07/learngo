// Author: xufei
// Date: 2019-08-05

package shell_sort

import "testing"

func TestShellSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "希尔排序",
			args: args{
				list: []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShellSort(tt.args.list)
			t.Log(tt.args.list)
		})
	}
}
