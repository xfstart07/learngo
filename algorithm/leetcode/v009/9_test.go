// Author: xufei
// Date: 2020/3/30

package v009

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isPalindrome",
			args: args{x: 121},
			want: true,
		},
		{
			name: "isPalindrome1",
			args: args{x: -121},
			want: false,
		},
		{
			name: "isPalindrome2",
			args: args{x: 10},
			want: false,
		},
		{
			name: "isPalindrome3",
			args: args{x: 0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
