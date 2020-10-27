// Author: xufei
// Date: 2020/7/30

package logger

import (
	"context"
	"testing"
)

func BenchmarkInfof(t *testing.B) {
	ctx := NewTraceIDContext(context.Background(), "1234567890")

	for i := 0; i < t.N; i++ {
		Infof(ctx, "带ctx 的Logger")
	}

	//BenchmarkInfof-8   	   45920	     25044 ns/op
}
