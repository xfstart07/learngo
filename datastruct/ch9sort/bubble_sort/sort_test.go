// Author: xufei
// Date: 2019-08-05

package bubble_sort

import (
	"testing"
)

func TestBubbleSort0(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "冒泡排序",
			args: args{
				list: []int{9, 1, 5, 8, 3, 7, 4, 6, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort0(tt.args.list)
			t.Log(tt.args.list)
		})
	}
}

func TestBubbleSort1(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "冒泡排序",
			args: args{
				list: []int{9, 1, 5, 8, 3, 7, 4, 6, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort1(tt.args.list)
			t.Log(tt.args.list)
		})
	}
}

func TestBubbleSort2(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "冒泡排序优化",
			args: args{
				list: []int{9, 1, 5, 8, 3, 7, 4, 6, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort2(tt.args.list)
			t.Log(tt.args.list)
		})
	}
}
