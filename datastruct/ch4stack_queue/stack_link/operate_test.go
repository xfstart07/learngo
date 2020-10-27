// Author: xufei
// Date: 2019-07-26

package stack_link

import (
	"testing"
)

func TestPush(t *testing.T) {
	type args struct {
		link *LinkStack
		e    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "入栈",
			args: args{
				link: InitStack(),
				e:    100,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Push(tt.args.link, tt.args.e); got != tt.want {
				t.Errorf("Push() = %v, want %v", got, tt.want)
			} else {
				for node := tt.args.link.top; node != nil; node = node.next {
					t.Log(node.data)
				}
			}
		})
	}
}

func TestLinkStack_Push(t *testing.T) {
	type args struct {
		link *LinkStack
		e    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "入栈",
			args: args{
				link: InitStack(),
				e:    100,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.link.Push(tt.args.e); got != tt.want {
				t.Errorf("LinkStack.Push() = %v, want %v", got, tt.want)
			} else {
				for node := tt.args.link.top; node != nil; node = node.next {
					t.Log(node.data)
				}
			}
		})
	}
}
