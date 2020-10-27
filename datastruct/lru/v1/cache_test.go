// Author: xufei
// Date: 2019-08-07

package v1

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "LRU 算法",
			args: args{
				capacity: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := Constructor(tt.args.capacity)

			cache.Put(2, 1)
			t.Log(cache.Get(2))
			cache.Put(3, 2)
			t.Log(cache.Get(2))
			t.Log(cache.Get(3))

			cache.print()
		})
	}
}
