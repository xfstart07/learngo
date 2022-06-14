// Author: Xu Fei
// Date: 2018/8/24
package memo_test

import (
	memo "learngo/books/gopl/ch9/memo4"
	"learngo/books/gopl/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
