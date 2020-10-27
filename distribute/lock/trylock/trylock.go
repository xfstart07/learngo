// Author: xufei
// Date: 2019-12-24

package main

import "sync"

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock

	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}

	return l
}

func (l Lock) Lock() bool {
	flag := false
	select {
	case <-l.c:
		flag = true
	default:
	}

	return flag
}

func (l Lock) UnLock() {
	l.c <- struct{}{}
}

var counter int

func main() {
	var l = NewLock()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() { // 不会阻塞
				println("lock fail")
				return
			}

			counter++
			println("current counter: ", counter)
			l.UnLock()
		}()
	}

	wg.Wait()
}
