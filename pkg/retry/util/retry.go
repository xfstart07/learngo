// Author: Xu Fei
// Date: 2018/8/27
package util

import (
	"time"
	"fmt"
	"log"
)

func Retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			log.Println("stop")
			return s.error
		}

		if attempts--; attempts > 0 {
			fmt.Printf("Warn: retry func error %s. attemps #%d after %s.", err.Error(), attempts, sleep)
			time.Sleep(sleep)
			return Retry(attempts, sleep, fn)
		}

		return err
	}

	return nil
}

type stop struct {
	error
}

func NoRetryError(err error) stop {
	return stop{err}
}
