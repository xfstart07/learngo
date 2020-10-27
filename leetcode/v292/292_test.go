// Author: xufei
// Date: 2020/3/24

package v292

import "testing"

func Test_canWinNim(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nim game",
			args: args{n: 4},
			want: false,
		},
		{
			name: "nim game",
			args: args{n: 5},
			want: true,
		},
		{
			name: "nim game",
			args: args{n: 12},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canWinNim(tt.args.n); got != tt.want {
				t.Errorf("canWinNim() = %v, want %v", got, tt.want)
			}
		})
	}
}
