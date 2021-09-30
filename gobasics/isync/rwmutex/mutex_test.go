// Author: xufei
// Date: 2019-08-12

package rwmutex

import "testing"

func TestUse(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "读写锁",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Use()
		})
	}
}
