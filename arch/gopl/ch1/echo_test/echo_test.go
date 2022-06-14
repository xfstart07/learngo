// Author: Xu Fei
// Date: 2018/8/17
package echo_test

import (
	"testing"
	"strings"
)

func BenchmarkEcho1(b *testing.B) {
	var s, sep string
	strs := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		strs[i] = "string"
	}

	sep  = " "
	for i := 0; i < b.N; i++ {
		s += sep + strs[i]
	}
}
//200000	    159500 ns/op
//PASS

func BenchmarkEcho2(b *testing.B) {
	strs := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		strs[i] = "string"
	}

	strings.Join(strs, " ")
}
//30000000	        36.0 ns/op
//PASS