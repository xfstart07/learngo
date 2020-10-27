// Author: xufei
// Date: 2020/3/24

package v557

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reverse string",
			args: args{s: "Let's take LeetCode contest"},
			want: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			name: "reverse null",
			args: args{s: ""},
			want: "",
		},
		{
			name: "reverse ",
			args: args{s: " "},
			want: " ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
