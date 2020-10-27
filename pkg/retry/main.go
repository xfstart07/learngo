// Author: Xu Fei
// Date: 2018/8/28
package main

import (
	"fmt"
	"learngo/pkg/retry/util"
	"time"
)

func main() {
	cnt := 0
	err := fmt.Errorf("test error every time")
	a := func() error {
		if cnt == 0 {
			cnt++
			return err
		} else {
			cnt++
			return util.NoRetryError(fmt.Errorf("stoperr"))
		}
	}
	errFn := util.Retry(3, 1*time.Millisecond, a)
	fmt.Println(errFn)
}
