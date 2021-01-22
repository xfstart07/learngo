// Author: xufei
// Date: 2020/4/1

package v148

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "sortList",
			args: args{head: newNode([]int{4, 2, 1, 3})},
			want: newNode([]int{1, 2, 3, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
