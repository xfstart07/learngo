// Author: xufei
// Date: 2018/12/3

package scheduler

import "learngo/crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
	panic("implement me")
}

func (s *SimpleScheduler) Run() {
	panic("implement me")
}

func (s *SimpleScheduler) ConfigureWorkerChan(workerChan chan engine.Request) {
	s.WorkerChan = workerChan
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	//  由 Go 的 goroutine 来调度 chan 接收
	go func() { s.WorkerChan <- request }()
}
