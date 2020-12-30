// Author: xufei
// Date: 2020/11/19

package xgo

import (
	"fmt"
	"log"
	"runtime"
)

// Go goroutine
func Go(fn func()) {
	go try2(fn, nil)
}

func try2(fn func(), cleaner func()) (ret error) {
	if cleaner != nil {
		defer cleaner()
	}

	defer func() {
		_, file, line, _ := runtime.Caller(5)
		if err := recover(); err != nil {
			log.Printf("recover: err: %v, line: %s:%d", err, file, line)

			if _, ok := err.(error); ok {
				ret = err.(error)
			} else {
				ret = fmt.Errorf("%+v", err)
			}
		}
	}()

	fn()
	return nil
}
