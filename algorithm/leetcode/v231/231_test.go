// Author: xufei
// Date: 2020/3/31

package v231

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2的幂",
			args: args{n: 1},
			want: true,
		},
		{
			name: "2的幂1",
			args: args{n: 16},
			want: true,
		},
		{
			name: "2的幂2",
			args: args{n: 218},
			want: false,
		},
		{
			name: "2的幂3",
			args: args{n: 33554432},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
