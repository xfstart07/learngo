// Author: xufei
// Date: 2019-12-23

package snowflake

import (
	"testing"
)

func TestNode_NewID(t *testing.T) {
	tests := []struct {
		name string
		node int64
	}{
		{
			name: "new ID",
			node: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, err := NewNode(tt.node)
			if err != nil {
				t.Error(err)
			}
			t.Log(NewID())
			t.Log(NewID())
			t.Log(NewID())
		})
	}
}
