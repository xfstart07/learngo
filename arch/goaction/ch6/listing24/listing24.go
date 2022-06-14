package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task, %d", post)
	}

	close(tasks) // 发送完之后，关闭通道

	wg.Wait()

}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks // 会在这里堵塞，等待接收信息
		if !ok {
			fmt.Printf("Worker: %d: Shutting Down\n", worker)
			return
		}

		fmt.Println("Worker Started ", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)

		fmt.Println("Worker Completed ", worker)
	}
}
