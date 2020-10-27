// Author: xufei
// Date: 2019-08-05

package straight_insertion_sort

import "testing"

func TestInsertSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "直接插入排序",
			args: args{
				list: []int{9, 1, 5, 8, 3, 7, 4, 6, 2},
			},
		},
		{
			name: "直接插入排序",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertSort(tt.args.list)
			t.Log(tt.args.list)
		})
	}
}
