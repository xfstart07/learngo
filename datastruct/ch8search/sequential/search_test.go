// Author: xufei
// Date: 2019-07-31

package sequential

import "testing"

func TestSequential_Search2(t *testing.T) {
	type args struct {
		a   []int
		n   int
		key int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "优化查找",
			args: args{
				a:   []int{1, 2, 4, 3, 6, 5},
				n:   5,
				key: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sequential_Search2(tt.args.a, tt.args.n, tt.args.key); got != tt.want {
				t.Errorf("Sequential_Search2() = %v, want %v", got, tt.want)
			}
		})
	}
}
