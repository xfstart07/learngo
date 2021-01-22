// Author: xufei
// Date: 2020/11/9

package v001

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "长度",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "长度",
			args: args{
				s: "bbbbbb",
			},
			want: 1,
		},
		{
			name: "长度",
			args: args{
				s: "dvdf",
			},
			want: 3,
		},
		{
			name: "长度",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
