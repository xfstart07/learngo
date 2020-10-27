// Author: xufei
// Date: 2020/4/1

package v142

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	head, pos := newNode([]int{3, 2, 0, -4}, 1)

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "detect Cycle",
			args: args{head: head},
			want: pos,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
