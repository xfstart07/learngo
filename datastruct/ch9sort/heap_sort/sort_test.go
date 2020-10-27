// Author: xufei
// Date: 2019-08-05

package heap_sort

import "testing"

func TestHeapSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "堆排序",
			args: args{
				list: []int{9, 1, 5, 8, 3, 7, 4, 6, 2},
			},
		},
		{
			name: "堆排序",
			args: args{
				list: []int{50, 10, 90, 30, 70, 40, 80, 60, 20},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.list)
			t.Log(tt.args.list)
		})
	}
}
