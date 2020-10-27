// Author: xufei
// Date: 2020/3/24

package v206

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "reverse link",
			args: args{head: newNode([]int{1, 2, 3, 4, 5})},
			want: newNode([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "reverse link",
			args: args{head: newNode([]int{9, 8, 7, 6, 5, 4, 3})},
			want: newNode([]int{3, 4, 5, 6, 7, 8, 9}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList2(tt.args.head); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
