package main

import (
	"learngo/books/goaction/ch7/runner"
	"log"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")

	r := runner.New(timeout)

	// r.Add(createTask(), createTask(), createTask())
	r.Add(createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Fatalln("Terminating due to timeout.")
		case runner.ErrInterrupt:
			log.Fatalln("Terminating due to interrupt.")
		}
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
