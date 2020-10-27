// Author: xufei
// Date: 2019-07-31

package fibonacci

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		arr []int
		n   int
		key int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "斐波那契查找",
			args: args{
				arr: []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99},
				n:   9,
				key: 59,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.arr, tt.args.n, tt.args.key); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
