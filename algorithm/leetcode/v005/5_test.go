package v005

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "max palindrome",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "max palindrome",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "max palindrome",
			args: args{
				s: "ac",
			},
			want: "a",
		},
		{
			name: "max palindrome",
			args: args{
				s: "abc",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
