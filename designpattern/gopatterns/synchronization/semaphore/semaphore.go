// Author: xufei
// Date: 2019-05-15

package semaphore

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrNoTickets      = errors.New("semaphore: could not aquire semaphore")
	ErrIllegalRelease = errors.New("semaphore: can't release the semaphore without acquiring it first")
)

type Interface interface {
	Acquire() error
	Release() error
}

type implementation struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *implementation) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		fmt.Println("获得锁")
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *implementation) Release() error {
	select {
	case _ = <-s.sem:
		fmt.Println("释放锁")
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
}

func New(tickets int, timeout time.Duration) Interface {
	return &implementation{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}
