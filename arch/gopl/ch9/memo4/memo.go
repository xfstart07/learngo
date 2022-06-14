// Author: Xu Fei
// Date: 2018/8/24
package memo

import "sync"

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//!+
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
		// <- 可以多次调用吗？
		// 允许多次调用，在没有写入的情况返回零值
	}
	return e.res.value, e.res.err
}
