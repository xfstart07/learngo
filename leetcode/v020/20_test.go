// Author: xufei
// Date: 2020/3/31

package v020

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isValid",
			args: args{s: "()"},
			want: true,
		},
		{
			name: "isValid2",
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			name: "isValid3",
			args: args{s: "(]"},
			want: false,
		},
		{
			name: "isValid4",
			args: args{s: "(])]"},
			want: false,
		},
		{
			name: "isValid5",
			args: args{s: "{[]}"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
