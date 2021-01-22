// Author: xufei
// Date: 2020/3/31

package v007

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "reverse int",
			args: args{x: 123},
			want: 321,
		},
		{
			name: "reverse int 1",
			args: args{x: -123},
			want: -321,
		},
		{
			name: "reverse int 2",
			args: args{x: 120},
			want: 21,
		},
		{
			name: "reverse int 3",
			args: args{x: 1},
			want: 1,
		},
		{
			name: "reverse int 3",
			args: args{x: 7563847412},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
