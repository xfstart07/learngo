// Author: xufei
// Date: 2018/12/2

package engine

import "log"

type ConcurrentEngine struct {
	Scheduler       Scheduler
	WorkerCount     int
	ItemSaver       chan Item
	WorkerProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	Submit(Request)
	WorkerReady(chan Request)
	ConfigureWorkerChan(chan Request)
	Run()
}

func (e ConcurrentEngine) Run(requests ...Request) {
	e.Scheduler.Run()

	out := make(chan ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler, out)
	}

	// 发送 req 给 scheduler
	for _, r := range requests {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out

		for _, item := range result.Items {
			e.ItemSaver <- item
		}

		for _, request := range result.Requests {
			// 这里要全部提交完，上面的 out 才能接受结果，所有在这里解锁释放出来，
			// 保证一定能提交成功，在 submit 中起 goroutine 去等待发送
			e.Scheduler.Submit(request)
		}
	}

}

func (e ConcurrentEngine) createWorker(s Scheduler, out chan ParseResult) {
	in := make(chan Request)
	go func() {
		for {
			// 告诉 schedule Worker 已经准备好
			s.WorkerReady(in)
			request := <-in

			parseResult, err := e.WorkerProcessor(request)
			if err != nil {
				log.Printf("processor: %v\n", err)
				continue
			}
			out <- parseResult
		}
	}()
}
