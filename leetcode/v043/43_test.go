package v043

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "big multiply",
			args: args{
				num1: "0",
				num2: "1",
			},
			want: "0",
		},
		{
			name: "big multiply",
			args: args{
				num1: "123",
				num2: "45622",
			},
			want: "5611506",
		},
		{
			name: "big multiply",
			args: args{
				num1: "111",
				num2: "111",
			},
			want: "12321",
		},
		{
			name: "big multiply",
			args: args{
				num1: "408",
				num2: "5",
			},
			want: "2040",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
