// Author: xufei
// Date: 2019-07-26

package sqstack

import (
	"testing"
)

func TestPush(t *testing.T) {
	type args struct {
		stack *SqStack
		e     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "入栈",
			args: args{
				stack: InitStack(),
				e:     100,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Push(tt.args.stack, tt.args.e); got != tt.want {
				t.Errorf("Push() = %v, want %v", got, tt.want)
			} else {
				t.Log(tt.args.stack.data)
				t.Log(tt.args.stack.top)
			}
		})
	}
}

func TestPop(t *testing.T) {
	type args struct {
		stack *SqStack
		e     *int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "出栈",
			args: args{
				stack: InitStack(),
				e:     new(int),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Push(tt.args.stack, 100)
			Push(tt.args.stack, 101)
			if got := Pop(tt.args.stack, tt.args.e); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			} else {
				t.Log("出栈值", *tt.args.e)
			}
		})
	}
}
