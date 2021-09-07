package mapreduce

import (
	"strings"
	"testing"
)

// Map/Reduce/Filter只是一种控制逻辑，真正的业务逻辑是在传给他们的数据和那个函数来定义的。

func TestMapStrToStr(t *testing.T) {
	var list = []string{"Leon", "Study", "Go"}

	ret1 := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})

	t.Logf("MapStrToStr: %+v", ret1)

	ret2 := MapStrToInt(list, func(s string) int {
		return len(s)
	})

	t.Logf("MapStrToInt: %+v", ret2)
}

func TestReduce(t *testing.T) {
	var list = []string{"Leon", "Study", "Go"}
	ret := Reduce(list, func(s string) int {
		return len(s)
	})
	t.Logf("Reduce: %+v", ret)
}

func TestFilter(t *testing.T) {
	var list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ret := Fileter(list, func(n int) bool {
		return n%2 == 1
	})
	t.Logf("Filter: %+v", ret)
}