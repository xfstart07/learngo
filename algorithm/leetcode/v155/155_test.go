// Author: xufei
// Date: 2020/3/30

package v155

import "testing"

func TestMinStack_GetMin(t *testing.T) {
	tests := []struct {
		name string
		s    *MinStack
		want int
	}{
		{
			name: "MinStack0",
			s: func() *MinStack {
				ms := Constructor()
				ms.Push(-2)
				ms.Push(0)
				ms.Push(-3)
				ms.Pop()

				return &ms
			}(),
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetMin(); got != tt.want {
				t.Errorf("GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
