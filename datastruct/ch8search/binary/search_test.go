// Author: xufei
// Date: 2019-07-31

package binary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	type args struct {
		arr []int
		key int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "折半查找",
			args: args{
				arr: []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99},
				key: 62,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.arr, tt.args.key); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterpolation(t *testing.T) {
	type args struct {
		arr []int
		key int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "折半查找",
			args: args{
				arr: []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99},
				key: 62,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Interpolation(tt.args.arr, tt.args.key); got != tt.want {
				t.Errorf("Interpolation() = %v, want %v", got, tt.want)
			}
		})
	}
}
