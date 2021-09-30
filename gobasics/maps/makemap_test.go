// Author: xufei
// Date: 2019-06-26

package maps

import "testing"

func TestMakeLenMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "make map",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeLenMap()
		})
	}
}
