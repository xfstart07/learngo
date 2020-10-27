package v008

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find digital",
			args: args{
				str: "   -42",
			},
			want: -42,
		},
		{
			name: "find digital",
			args: args{
				str: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "find digital",
			args: args{
				str: "words and 987",
			},
			want: 0,
		},
		{
			name: "find digital",
			args: args{
				str: " -91283472332 ",
			},
			want: -2147483648,
		},
		{
			name: "find digital",
			args: args{
				str: " 91283472332 ",
			},
			want: 2147483647,
		},
		{
			name: "find digital",
			args: args{
				str: "20000000000000000000 ",
			},
			want: 2147483647,
		},
		{
			name: "find digital",
			args: args{
				str: "-20000000000000000000 ",
			},
			want: -2147483648,
		},
		{
			name: "find digital",
			args: args{
				str: " 0000000000012345678 ",
			},
			want: 12345678,
		},
		{
			name: "find digital",
			args: args{
				str: "-000000000000001",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
