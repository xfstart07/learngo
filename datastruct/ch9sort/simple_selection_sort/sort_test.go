// Author: xufei
// Date: 2019-08-05

package simple_selection_sort

import "testing"

func TestSelectSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "简单选择排序",
			args: args{
				list: []int{9, 1, 5, 8, 3, 7, 4, 6, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectSort(tt.args.list)
			t.Log(tt.args.list)
		})
	}
}
