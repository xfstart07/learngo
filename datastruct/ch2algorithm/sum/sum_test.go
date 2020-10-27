// Author: xufei
// Date: 2019-07-24

package sum

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "100 sum",
			args: args{
				n: 100,
			},
			want: 5050,
		},
		{
			name: "1000 sum",
			args: args{
				n: 1000,
			},
			want: 500500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.n); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
