// Author: xufei
// Date: 2019-05-15

package main

import (
	"learngo/designpattern/gopatterns/synchronization/semaphore"
	"os"
	"time"
)

func main() {
	Timeouts()
	NonBlocking()
}

func Timeouts() {
	tickets, timeout := 1, 3*time.Second
	s := semaphore.New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		panic(err)
	}

	// doSomething

	if err := s.Release(); err != nil {
		panic(err)
	}
}

func NonBlocking() {
	tickets, timeout := 0, time.Duration(0)
	s := semaphore.New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		if err != semaphore.ErrNoTickets {
			panic(err)
		}

		os.Exit(1)
	}

}
